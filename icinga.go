package main

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type icingaConf struct {
	Server string
	Port   int
	Pass   string
	User   string
	Proto  string
}

type icinga struct {
	baseUrl string
	user    string
	pass    string
}

func newIcinga(conf icingaConf) icinga {
	baseUrl := fmt.Sprintf("%s://%s:%d/v1", conf.Proto, conf.Server, conf.Port)
	i := icinga{
		baseUrl: baseUrl,
		user:    conf.User,
		pass:    conf.Pass,
	}
	return i
}

func (i icinga) ServiceStatus(host, service string) (bool, error) {
	// proper encoding for the host string
	hostUrl := &url.URL{Path: host}
	host = hostUrl.String()
	// proper encoding for the service string
	serviceUrl := &url.URL{Path: service}
	service = serviceUrl.String()
	// build url
	url := fmt.Sprintf("%s/objects/services?service=%s!%s", i.baseUrl, host, service)
	// query api
	// TODO: read rootca from file
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return true, err
	}
	req.SetBasicAuth(i.user, i.pass)
	resp, err := client.Do(req)
	if err != nil {
		return true, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return true, errors.New("HTTP error " + resp.Status)
	}
	// parse response body
	var results serviceStatusResults
	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &results); err != nil {
		return true, err
	}
	return results.Status()
}

type serviceStatusResults struct {
	Results []serviceStatusResult `json:"results"`
}

func (r serviceStatusResults) Status() (bool, error) {
	if len(r.Results) != 1 {
		return true, errors.New("not exactly one result found")
	}
	// only fail if critical and active on icinga
	// TODO: Ignore when downtime scheduled
	if r.Results[0].Attrs.LastCheckResult.State == 2 && r.Results[0].Attrs.LastCheckResult.Active {
		return false, nil
	}
	return true, nil
}

type serviceStatusResult struct {
	Attrs struct {
		LastCheckResult struct {
			State  float64 `json:"state"`
			Active bool    `json:"active"`
		} `json:"last_check_result"`
	} `json:"attrs"`
}
