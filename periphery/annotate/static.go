// Code generated by "esc -o static.go -pkg annotate -prefix static static"; DO NOT EDIT.

package annotate

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
		size:    2807,
		modtime: 1520512558,
		compressed: `
H4sIAAAAAAAC/5xWX2/jRBB/Tj7FsiinVjrbPQ4JaG2jcpQ32gLlASFUje2JPdf1rs87SS9CfHe0/tPY
jlMiXpLdnd/82Zn5jTf84se7Dw9/3N+IgksVL0P3JxToPJKopTtAyOKlEEKEJTKItIDaIkfy94efvG/l
UKShxEhuCZ8rU7MUqdGMmiP5TBkXUYZbStFrNm8FaWIC5dkUFEbv/IveFBMrjH+4//nuVlxrbRgYw6A9
bRGK9JOoUUWSUqOlKGpcRxKsRbbBGrbu1KfUSMG7CiNJJeQYfPYa9NBGq1gwV/YyCNZGs/VzY3KFUJH1
U1MGqbXfr6EktYt+NYlhc/n1xcXbby4uZBuB5Z1CWyCyPIhuIOsiYfzMzuQk5tSUpdG+E/wPIw2iUw6D
tl5hYrJdvAwz2opUgbWRdNUA0lj3PgYyp+QEh2ce6b3KvJpnMWUyWihc8wA5RSuTm4m4gVCZdxDRYDwq
cylsnb7kJ6lceiqdSwGKIzlujjmbtgLd+00gy1EKylwONppl7Pt+GDjIEc34jU5sdSWub2/vHq4fbubA
YZDRdpCWyfZ4lmrKi/k0tXTBjBgS5Qpdb1D2VkhXGxbNr9eVq7nSpw2lT4/Q5oIcHTJg8CoFKRZGZVhH
8heH6fNFRvu+jI/H3y67g32rOGfO9JFeUAYychVyuH4zaihbdb00cjBw3rsMbVpTxcN2d2OpNY1b1PzI
ZaVm+rgCjR0uXEWUrWIZLxeLA4jX5XpwlxGI9NqIxtNrzT1SYSrRMpTVTDsuwlX0Il9NjYwLcbgfOoHc
nhLXSqxNLc62UAuGXJAWfC7+Fqt4uTjGToa8MzOGrCKGfBWHSdys/2TI/1rFYZAMQpwhQB/HPyOnA9Bi
GsD4WsmG2Wh7wJbFnt5Gp4rSp44C/TQ4e/Pl+++uuuo36/MXDrVGRYOX8YgVOMfxxcSRLczzTUZs6v/2
IuOp5ZOysCeyh42nUTcPeP7Yi5vxH8mMbKVgdym00Xh1kLHXXHhspl27OHEqDWy1A6rGfKPgaNRz0+m+
NlvKUOxRwnFeFFijPzOq5lhywJTDSyaG2ZSz7BxXuaAMT61y30sfQKeojn1Uph7g5F6V8W+os/nePD43
+vXiZbC2E3UwW4ef14+fNljvvPf+V/47vyTtf7Qu6b3OnIp7K8EzWlOiB0qdoMFYVgqYdH4CuHsQvQIs
4VicsXtZnp1fDS8ddM+hoH3lLv8NAAD//7ho6+L3CgAA
`,
	},

	"/main.js": {
		local:   "static/main.js",
		size:    2950,
		modtime: 1522847014,
		compressed: `
H4sIAAAAAAAC/+xVXW/bNhR9z6+4YAyYmlXa2bA9OFWGAulDh6LplvRpGAJWvLK4SJeaSNnxWv/3gZTk
yY48ZMNQYMD0Yov34xyS5x5lDaVOG4LCSPV6jeQsj+DTGQDARKzQ/XB7846zuaz0fH0x/1jNMSQxmMFG
kzIbUZhU+hbCoqzTPIa+J1fSyb6Zf3Kt8LbSRFjz6HK/rDMIqZAksIDPn6F7SYCaojhYaEhhpgnVsG3g
ytlLpdeQFtLaZErG6Uy3vKCShMX06p2BhiSRcdKhgnYfkJmGlBDi5Vzp9RWLhCfJIyGrCkndGc7OPTqL
RCYVviHufzQN6Hfw56lpyPkGriw4W7BByg6wsHhEeS1rCDWQwOKwXWZq4D7+gFvQBMcH2T+hfDa7fBLw
tS1RSIC3KFfwTQTfw3eLBSzh4tsFfNXWj1ebyp+dhWQE1j9KbpfAqCmx1imLYYuyPlgoDbl8CczmpnYs
flaT3DS+ydcvlF5p55toahweLFlMDanB0pPOu/EdKekQEiDcwLV0GCT38wNufxFOlxidLrp1taYVJBDe
hDNvTSoLvNNlF+JM4Ys0Z3F/atHl2Wi3oDl/pFot4U98rWLwHKyTZbUcYMbghnlOruyJzQWNQwKurArO
As69/8/iFnSM0YSHqn8o+d1fD0B4GY5A+LeLLs92Z2d71/mt0enDq24mud6Ptd9SidbKlb+yCZ+eh8z7
bny1oWkkHD663ke8h3QFokBa5W7vJvs+CTB2tDDqJ8E8tpz96CHh1R4S7vDRgTMQJC2E6Ef8yXhPhPxV
PvLDyXHbyiv5/c3t3dE4NHWxhKceO/cmq9Vhrr+ZZb+Fw5Bt0hStXR4Z8KdRwTR10Q3Dh5/e8mMrz2vM
Tk7E1rtCUxed4b+XtSyt/1Zw5oNspHDwdfEpIxn+jqczrVrR3TSOXywWJyW360/+QE3yOULyILPpl1XS
/xr6z2jI5mbzWmlnat4a50BIWkEC7JzNQmTGBiK6x1DD2pYTr0BRyRrJ8UhkmhRn4mPjnCHLxui1FbbQ
Cq/NhviRT3qH/iKs3tA4qQ/VCCWF9C+RSq3lTGlbFXLLYmBkCNkRYI1eAdfSST5EerYI/5bwfLLVv+PJ
ZB/skwfK9MsxdPocstek3Z73cCs+7Y8AAAD//76WoVOGCwAA
`,
	},

	"/style.css": {
		local:   "static/style.css",
		size:    1143,
		modtime: 1522847014,
		compressed: `
H4sIAAAAAAAC/4xT247TQAx971dYQkhQkW4aWHY7/RpPxklGO2sPE4d2Qfw72lxK2iYSebTnXGyf7DxX
8gV2inULvzcAAPKTUhXkZKDxzhEf+2pE5zzXWfJ1owb2eTxfN1SigX1xW7aiKq9T589m0yuOUifvtDGw
f8qnpu1UhdvrfvGYr9EWxTVy96Pz5csIt1i+1Ek6dlkpQZKBD87i49PzgEBmUVQvnJHzKmmElYEwGbCi
zSBaBUE1EKjShZmLi7lpnDz/OBSUzpph8DUbKImV0lB/2FpJjpKBfTxDK8E7sKGj4/Zh2dm70s3OLiJX
dpfRw7b+i2A2b3/q+0Fm5ZU93B3p23SkmTXPsdPRUkNDqu4ohqDNYjXsLUvofNca+D7V5wb/3enUeKWs
jViSAZZTwnhcyfiiOztFwvk2BnwzLEwrb7c3Tz0H//741sdkY4lkl6juAqb1/GLl9iWuoOfZH6/8Nb+s
dIGt6r9lNkOvUd+MpUoSTX/GiDscDmNShDVr/S8y8HyJk7ASqwFUTZ8cKmYxYEmNBEfpc6/2NwAA///I
LyiMdwQAAA==
`,
	},

	"/": {
		isDir: true,
		local: "static",
	},
}
