// Code generated by "esc -o static.go -pkg dashboard -prefix static static"; DO NOT EDIT.

package dashboard

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/index.html": {
		local:   "static/index.html",
		size:    3610,
		modtime: 1513356749,
		compressed: `
H4sIAAAAAAAC/5xX247bNhO+36dg+EM5AJG0m/xA26ykIqkbtEDTLJLtRVAUi5FESxNTJEuOvTWKvntB
HWxZljdOb2xpzvPNgVTyaPH+h9tPNz+ymhqZXST+j0lQVcqF4p4goMwuGGMsaQQBK2qwTlDKf7t9G37L
xywFjUj5BsW90ZY4K7QioSjl91hSnZZig4UI25fnDBUSggxdAVKkV9HlYIqQpMje3Lx7/ytbgKtzDbZM
4o7ciUhUK1ZbsUx5TWTcqzheakUuqrSupACDLip0ExfOfb+EBuU2/aBzTfrV/y8vn39zecmZFTLljrZS
uFoI4mPLUx6jrREpJ/EXeZO8dw3OCXJxoZtGq8gz/oORVqJXTuIO7CTX5Ta7SErcsEKCcyn3UAIqYQcf
I55X8oxjWohqrzKvFjpREGrFpFjSSHIqLXWlJ+xWBJuqF2GtTIhNxZmzxQ6f3Hh4jKo4A0kpxwJVBY0u
VnP2nAGVPVa5M9eL1x9/evP+9YdFErfUw9jiEjejxCavp/O0WNU+0YsDfEurTanv1Qy8+ZpIK9b9hQPY
H4UUBbFbbIQFVYlRAHN2w34WwnsLxsyWcSo6iLS5/wKO2MtLtoCtG8Mx5rJ3WlF9zL2t0bFPAixzmi3B
7iV2QfcPExRbE4/CkPEuzi4uUSJBLn0v27XgQ/iozJpY+zuAxLBM+RIl+ecSCEIjoRC1lqWwKX/bcthS
W/Zm7VAJ59iN1YVwLooinnXRsDCcLXP32BP2YHqX3tWJppcaSvSt6OWGl4PJcaYfmgMHI+eDy8QVFg2N
59ovz86035J0t7TQiDtqjJypd8tkSZA6AhJBxlm7C/pt+cpzSnRGwvZG2EIoCrLguseRtJaEJn2SBGn/
HGRPeow76ZQnQWoGzV7PEdiO0z4F2ZDmRRJ36XwpMVQkrF0bP0vuILX6xQ7Cdd5ua579PJZO4vrFsON9
A40qRPtTZk+zWUJ15icsialuXxZrC97WjvBaKU0jUkx23C0Tuwl1m/XAT9B24NMNWIYMFTvI8Bn7mwXH
gR3vLSqzJEgPdH/HPyLCRjiCxgQ+tvJ8xbLPNMjc1ynCDpE5l4cA9fn/M04xiUcgJXFfqXPbY2XwqN89
PQnSlUEsfcftXe1bZmXQ3x7YvnU6DU8MMva4KcHV191CGwzmxtu725u+gw2ghBwl0pZnURT1qy5gPUeK
fRNO5/FQ+eRB2E42Px1Dxz/3sPrK6cvNqW1iQAk5CevEDmxFJ6fMnFxu5g7oycna7s8jqa66V1kbyhk1
PKtwLSWJ66uZqI4Prq8o79ll7mt7ytvM6TlDGh/5ggClm7oZyA8ntDJ4pNnSsjPQOFgbUyuHzOzLeT1w
+xqaTUowzt8DtCokFquUk64qKRZdrk8f/+/ld9e7EPzbs+tpb+K0HIPVO7dtci13d5ElOLaEEFQlRdhd
65IYHxrAuQkcX2I//7kWdhu+jF5EV1GDKvrcAjPozKn4LxK4F043IgQpz9Ag0RgJhKo6Q7j/7HhAsIFT
cWb+4+vps+tx0nH/0RF3H4IX/wYAAP//fpLgUxoOAAA=
`,
	},

	"/main.js": {
		local:   "static/main.js",
		size:    6645,
		modtime: 1513356783,
		compressed: `
H4sIAAAAAAAC/+xZX2/bOBJ/jj/FlAnOVO0oaYu7B7vKAdfmgG6x3WDTfQoCg7ZGNmuaFETKSdB6P/uC
pGRLsuz8wS6Sh/qhoDnD4fz5zfAXN8nlxHAlgUtuaADfOwAAMx7jZcqlxIwGQ7d1RGM1yRcoTRAqSaE7
x7s87fahe5hwYTDr9qE0trZjP0uWQcwMgwiOaPfQLcOUSRTdwrT9TJTUSmAo1JQeUTPjOggN3hoaBDVL
GUIEEm/gd5ye36YN3T4QTipW7WWhnqkbGoTeS9rqY2ldsgU6P73RhMuYknCc2n1SXjKsHcvQ5JmEVxmG
BrWhVrWisgpCm8zy1CoYdtxCKBb/70Lb/VWns67Certw7iicovnl8rcvlJywlJ+MU31CKom2AVYD4QlQ
n+wITuHHDyi/yFyI6vdcxphwiXEzDUeUvI/5EiaCaR11pTI84RPmvPNVOzvPMpXBzYwLBMG14XIK41xz
iVpDmqkJao06DMP3JzFfnpEyBSFLU5TxV0WJgwEJwoTF+ElSooW6qRZuBSg0tlRoonJpIILTehUSlQG1
8jneAZfQzMsGZrk0vd5wS2DPWme4hAiov+UM3gXwX/jP6SkM4M2/T+G1P1+UsHl8nEIE32Gc8nhg/eiD
B87AeXM1x7trWLXf7BILEZhFKigZpyO7INZAsH3giDr1+7Pq42mx4FF2vkRpNJ3j3Q6Vzxef2sSrTn21
asVwYd1mYy+Yoecy9tyY/qByEYNUBhI0kxmgc9/hqglt6JZOQw9I95/BeYrZBKVhU9QWVQ3cOLRmbOGE
V9fbwsL9beFDG2V9A0Qb+LaDVxtmcnvXFM2lW1+ajMspdedDLw6GnYODgwOrz6RUxic/gkKpsvcqAkJs
321JBkCOyY4G8umCCIpjcZ65Q6NC0tKzFleVNF95R6+ryNoh3gm08tN6rjmzfAfZrLSrt+32YB1PWxqM
UsJwO4eIr8QALFaLCvWAwL8Ob9nw0rDMeFGisgUzX/kCtWGLdFOzzASbA+cy3qeOMnbKTtf9Y3U3ldsx
bt24tL7hoHCxDzHXqWB3Fz7KQRluv1xUdrSPouJxv0zAoFy0zFvfN2Ga6xlN7HNsK2CxUIWrrfIbW92D
ErXovDVl5IO9metXgh9U1n0ocTlo4LTtZbAf38jeWyyaaNUYx53t8CAClDrP8Fcu+YKJj7W0sil6d3XQ
6Iv1eOB2OBQ6u8YDRIXGFd8xGiYzlpn1u+a+jdyR8oFLWt83p9g2TTeDl4ycUmW0JkwbsvVUdZodXyRU
oJyaGZzB6a7wCs76Hbg0mGV5asukB+Vo3fGSV5XXgdcslKG7ydsWfVX7viTULD8+GdURQ9ScXNvp2zb8
1sK9o88mQOW24F/yxRgzun1BEBr1f36LMX0XtEZfjQ/IiC0ZF2zMBTd3JQNXuXkCJXFs5qF0xN59Mk+5
fpFk24bSSk1QQzWCv5+aPIY+zFNeIcSe5M1TvmbH85Q/iR7PU1720DzljyPINWy5+u7vmQoVrhDaPjyN
FW9Zsdl4FBzdjjv2ggnz54tPQAo3HYdQyTaJfgYO/XAK/Yws+Se7xZ1/YT8jvX3RFPYnHQQyWk9G9+2J
7PDlE6J6mG30iPxpdR5PkpqDyeP5S74oI9M33Exm9f2DgwnTCKcD2+rFb6E2rmEpeVOTyEIUY8JyYWqy
XM6lupFWvqo59jDoFk7alC+4rKvZ8RS+9em4D9N7sOz+Sgzr/Q3vt68LdrGmb/e30XJz9beWNrIuLJsu
nAF9C69b/Ng1wdcXNC1FsGX8eNvucJ9R3mJ0K2m9hxodZ8jmbS/LXmR3Kr/Me7fqUDdqOhX4EQ3jov5H
gQNGbF+YQ9Kzgh4ZxV6PFMPRluCI8jgIuaZksOSajwWSWrK9XAse4x/p+sf/JlGpaH1UN5KW47cYTE0/
JkoIlmrMRvpuMVaClP8zY634kD5Ywman3TGTU4HHeQrrdWyby/X8XwEAAP///2/EzPUZAAA=
`,
	},

	"/style.css": {
		local:   "static/style.css",
		size:    850,
		modtime: 1513351262,
		compressed: `
H4sIAAAAAAAC/4SS4Y6bMBCE//MUK1WVWqkgyEG58z3N2rtgC5+NnE1yUdV3rzhIVa6E/N2Zj9llnBXG
YhL4lQEAXByJVVCV5dfXj4Fl11tRUJfj+zwxnjEp0FHsPNBohj7FU6DcRB+Tgi/M/Jr9zrKCWND54/J1
csfR41VBiGEx6JNIDDfDG6behVziqODwkTh5htEdocAzOo/aeSdXWC1927H6S8xq0SV84/9Mt9M6H1EU
eO5kRcVhQTYOa8mY9ucnu7LxzOk+1NT4UpsVFPZCuHtu2+az/1EKlU39xCvqFIYQL2En6Uk33WGLeZSG
1JZmJgW15633M2n2Bwgt4ohELvQKnpeexC5KF4Pkl6UhHT3NDQm/S47e9eGfmsQy0v3FKtTmZf7ZoiNd
QZIKYnNjnadvkej7zlEdVQYndgPlM4cdlg7ccDWxfwIAAP//yAkXK1IDAAA=
`,
	},

	"/": {
		isDir: true,
		local: "static",
	},
}
