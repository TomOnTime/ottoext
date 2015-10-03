// Code generated by go-bindata.
// sources:
// js/.gitignore
// js/bundle.js
// DO NOT EDIT!

package fetch

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _jsGitignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xcf\xcb\x4f\x49\x8d\xcf\xcd\x4f\x29\xcd\x49\x2d\xe6\x02\x04\x00\x00\xff\xff\x99\x1d\xdc\xd1\x0e\x00\x00\x00")

func jsGitignoreBytes() ([]byte, error) {
	return bindataRead(
		_jsGitignore,
		"js/.gitignore",
	)
}

func jsGitignore() (*asset, error) {
	bytes, err := jsGitignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/.gitignore", size: 14, mode: os.FileMode(420), modTime: time.Unix(1443770584, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _jsBundleJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe4\x57\xdd\x8e\xdb\x36\x13\x7d\x15\x59\x17\x06\x89\xf0\x13\xbc\xc9\xd7\x1f\x58\x61\x8c\x20\x08\x1a\xb4\x49\xb6\x48\x93\x2b\xc3\x30\x18\x69\xbc\x66\x22\x93\x2e\x35\xda\x8d\x61\xeb\xdd\x3b\x34\x25\x4b\x5a\xbb\x41\xd2\xe6\xa2\x68\xb1\x0b\x5b\xe4\x90\x87\x67\xce\xfc\x50\x1e\xad\x2a\x93\xa1\xb6\x86\x01\xdf\xb7\xcf\x11\x32\xc7\xf7\x7a\xc5\xcc\xdc\x2d\xb8\x03\xac\x9c\x89\xfc\x73\x02\x9f\xb6\xd6\x61\x99\xde\x2a\x17\x69\xe9\xa7\xe4\xbe\x99\x9b\xee\x6b\xa1\xf3\xa9\x13\x85\x55\x39\xe4\xd3\xd1\x55\x9d\x36\x5b\xc1\x6f\xcd\x54\x51\x30\xdd\x22\x08\x2d\xba\x67\xe4\x34\x08\xdb\xe4\x68\xd2\x19\x6a\x7f\x8c\x91\xfb\x13\x10\x26\x1b\x09\x02\x93\x4c\x1a\xfa\xdc\xca\x38\x16\xc8\x26\xbc\x66\xf3\xce\x0d\x81\xc2\xf0\x7d\x5c\x95\x10\x95\xe8\x74\x86\x71\x6a\xd8\x8f\x5c\x18\xf6\x9d\xff\xf8\xde\x7f\xfc\xc0\x6b\xd1\xdf\x71\x6f\xfd\x49\x87\xc6\x48\x4a\x8c\x18\x44\xda\x94\xa8\x4c\x06\x76\x15\x21\xe7\xb8\x76\xf6\x2e\x32\x70\x17\xbd\xdd\x6d\xe1\xb9\x73\xd6\xb1\xf8\x99\x32\xc6\x62\xe4\x7d\x8d\x54\x94\x15\xaa\x2c\x23\x45\xff\x51\x0b\x19\xf3\xfa\xfa\xfd\x07\xc8\x30\xc9\x61\xa5\x0d\xfc\xea\xec\x16\x1c\xee\x18\x8a\x78\xb9\x84\xf2\x95\xcd\xab\x02\x62\xb1\xbf\x55\x45\x05\xd3\xd1\xa4\xe6\x47\xb1\x9d\x3c\xf1\xed\xc5\x09\x02\xbf\x15\x1d\x1d\xa4\x9a\xa4\xe6\x31\x26\x05\x98\x1b\x5c\xa7\xe6\xc1\x03\xbe\x0f\x9b\x71\x6e\x16\xa9\x4b\xc0\x54\x1b\x70\xea\x7d\x01\xb2\x3f\x38\x1c\x46\x57\xc2\x25\x99\x35\x2b\x7d\x53\x05\x3b\xc5\x21\x3e\x72\x88\xb5\x89\xdc\x78\xcc\x5c\x72\xe7\x34\x36\x36\x2e\x2e\x7b\x01\x04\xf3\x11\x76\xc2\xf1\xba\x6e\x62\x76\xe2\x4d\x71\xa1\xf9\x7d\x9b\x4e\xe3\x31\x30\x8a\xa1\xb3\x68\x91\x04\xa4\xa0\x09\x77\x9c\xa3\x45\x02\xeb\x9a\x51\x52\xfc\x89\xd3\x78\x72\x6b\xad\xcb\xd4\xe7\x29\xf3\x4f\x02\x68\x23\x7d\x27\xcb\x35\x50\x2e\xb9\x92\x32\x47\x60\x3f\x6c\x40\x7e\xa0\xc4\xd3\x02\x2e\x62\x7b\xf4\x23\x96\xd2\x93\xf0\x91\x1d\x8f\x4d\x55\x14\x23\x9a\xe0\xad\xae\x9a\x30\xa2\xe6\x50\x4b\x5a\xea\x45\xfa\xd4\x39\xb5\x4b\x74\x79\xfc\x66\x96\x1f\x0e\xcc\xca\xb9\x5d\x70\x61\x13\xda\xf6\x5c\x65\x6b\xd6\x2f\xad\xc6\x6b\x97\xa8\xed\x16\x4c\xce\x34\x91\xad\x3b\x91\x1c\x29\x37\xdf\x93\x72\xd3\x38\x2c\x88\x45\x48\x80\xbe\x7a\xad\xd7\x90\x18\xeb\x36\xaa\xd0\x25\xbc\x56\x1b\x2f\x47\xda\x44\x63\xad\xca\xeb\x3b\xd3\xe5\x54\x5f\x0c\x92\x95\x48\x0e\xa6\x7c\xf9\xce\x17\xf7\x44\xf3\xb5\xba\xad\xca\x35\x33\x44\x4f\x04\x4e\x39\x14\x80\x70\xce\x89\xef\x83\x25\x1a\x22\x9c\x13\x5c\x9c\xa0\x6e\x00\x2f\xe1\x84\xf4\xbd\xe0\x59\x5b\xf9\x83\x03\xcc\x62\x76\x7f\x62\x3e\x59\x4c\x6f\xad\xce\xa3\x49\xff\xa8\xa7\x45\x71\xe9\xb4\x8b\xa0\x17\x58\x1f\x0e\xf3\x8e\x39\x89\xfb\x57\x98\x0f\x33\xe5\x3e\xed\x4e\xe3\xf2\x92\x30\x9f\x0d\xfa\x79\x2c\x0d\xb1\x5d\xb4\x89\x34\x58\x7f\x06\xdd\xe5\x24\x24\x68\x5f\xda\x3b\x70\xcf\x54\x09\x8c\x08\x51\x46\x00\x95\x5f\x8a\x73\x0a\xfc\x4a\x55\x05\xc6\x0b\x49\x09\xdb\xb6\x65\xd9\x37\x0c\x1b\xe9\x59\xeb\x3d\x55\xad\xfb\xc7\xb4\x52\xba\xb7\xd8\x15\xd5\xe9\xa9\xb9\x44\xca\xab\xe1\x4d\x28\x95\xbb\xa1\xae\x68\xb0\x6c\x9a\xe8\x63\x79\x75\x38\x84\xbc\x92\xb2\xb3\xce\xaf\x16\xb3\x7d\x3d\xed\x8f\x85\xa1\xc6\xb2\x01\x5c\xdb\x9c\xb0\x4f\x5b\xcc\x2c\xfe\xe9\xf9\xdb\x78\x6a\x04\xe9\x96\xb4\xb5\x58\x75\x0b\x4a\x0f\x54\x8a\x8c\xac\xef\x6d\xbe\x13\xab\xce\x94\xcd\x7c\x2b\x9a\x66\xa2\x20\xa3\x83\x5c\x3b\xf2\x56\xac\xbb\x05\xc5\x2c\xde\x28\x53\xa9\x22\x9e\x16\xbe\x0f\xba\xd0\x07\x15\x05\xb0\x2f\xb1\x0a\xde\xe5\x94\x44\x16\xd7\xe0\xde\xbd\x79\x29\x6e\xdb\xc1\xab\x40\x79\xdb\x8e\x5f\x34\x14\x77\xed\xc4\x9b\xe6\xe0\x90\x6f\x95\x2b\x64\x1e\x3a\x46\x70\x56\xde\x86\x51\xdb\x73\x7d\x18\x35\xdb\x36\x5d\xa5\x65\x2d\x77\x35\x14\x65\xd3\x28\x3c\x06\xa4\x7d\x0c\x7b\x09\xa3\x6a\x30\xbc\x2c\x72\x75\x0f\x6f\x5d\x0f\xf3\xd3\xfe\x8b\xf2\xf3\x3f\x78\xd5\xdb\x50\x94\xea\xf3\x37\xbe\xf9\x3b\xf5\x49\x75\x9f\x50\x54\xb1\xa2\x02\xe9\x4a\x48\xcf\x1e\x4e\x26\x53\x4d\xd5\xd9\x5a\xdf\xc2\x27\x1c\x16\x68\x7c\xfd\x4b\x7c\x2c\x51\x73\x2a\xe0\x41\x95\xd2\x49\x59\xea\x86\xef\x20\x3e\x69\xdf\x95\xe1\x85\x36\x34\xea\x63\x1a\xfb\x82\x3e\xcf\x75\xcb\x56\xcd\x36\xfb\x51\xaa\x27\x92\x28\x8d\xc7\x8f\x26\x93\x27\x2a\xcc\x06\x62\x72\x30\xf2\x34\x65\x15\x66\xbc\xaa\xb2\x8f\x9a\xd0\xdd\xc7\x62\x8a\x32\x92\xf7\xff\xf3\x66\xca\xc6\x26\x1c\xba\x7b\xd7\x40\xc2\x38\xbb\x19\x82\xd4\x01\xaf\xbd\xc4\x3c\x49\x8a\xfc\x86\x6e\x13\x76\xe9\x22\x65\x10\xfc\x3b\xbe\xd3\x34\x17\xda\x87\x92\x6a\xe0\x1c\xbc\x7f\xf7\xfa\xf3\x19\x4f\xa8\xcb\x98\x4b\x6f\x4b\x3f\xff\x76\xfd\x3a\xd9\x2a\x47\x87\x86\xb7\xa5\x8b\x37\x93\xfa\x36\x95\x1f\xf2\x2b\x74\x1e\x3f\x21\x5c\x08\xcd\x67\x25\x38\xee\x6c\x16\x2c\x97\x5b\xa7\x6f\x15\xc2\x72\xb9\x02\xcc\xd6\x4b\xf8\x04\x59\x85\xc0\xa8\x16\x3a\x2a\xa6\x2b\x8b\x19\xd2\x68\x0a\xf4\x5b\xab\xf6\x7f\x5f\xdb\x2c\xc4\x40\x05\x77\xba\xdb\x1e\x86\x6a\x7a\xc4\xd3\x2f\xd5\x65\x10\xd2\xde\xa6\xe4\x45\x9b\xa3\x54\x9b\x35\x0f\x3f\xe1\x50\x5c\x0e\x26\xc5\x85\x7f\x05\xf4\x1b\xf8\xbd\x82\x12\x8f\x7c\xbf\x39\x74\xb9\xb5\xa6\x84\xa3\x08\xdf\x18\xfb\x18\x59\x02\xfe\xff\x97\x01\x2f\x78\xfa\x47\x00\x00\x00\xff\xff\xf9\x49\xbc\x36\x62\x0f\x00\x00")

func jsBundleJsBytes() ([]byte, error) {
	return bindataRead(
		_jsBundleJs,
		"js/bundle.js",
	)
}

func jsBundleJs() (*asset, error) {
	bytes, err := jsBundleJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/bundle.js", size: 3938, mode: os.FileMode(420), modTime: time.Unix(1443839543, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if (err != nil) {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"js/.gitignore": jsGitignore,
	"js/bundle.js": jsBundleJs,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"js": &bintree{nil, map[string]*bintree{
		".gitignore": &bintree{jsGitignore, map[string]*bintree{
		}},
		"bundle.js": &bintree{jsBundleJs, map[string]*bintree{
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
        if err != nil {
                return err
        }
        err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
        if err != nil {
                return err
        }
        err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
        if err != nil {
                return err
        }
        return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        // File
        if err != nil {
                return RestoreAsset(dir, name)
        }
        // Dir
        for _, child := range children {
                err = RestoreAssets(dir, filepath.Join(name, child))
                if err != nil {
                        return err
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

