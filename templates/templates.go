// Code generated by go-bindata.
// sources:
// templates/swift/service.tpl
// templates/swift/struct.tpl
// DO NOT EDIT!

package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
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
	name    string
	size    int64
	mode    os.FileMode
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

var _templatesSwiftServiceTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func templatesSwiftServiceTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSwiftServiceTpl,
		"templates/swift/service.tpl",
	)
}

func templatesSwiftServiceTpl() (*asset, error) {
	bytes, err := templatesSwiftServiceTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/swift/service.tpl", size: 0, mode: os.FileMode(420), modTime: time.Unix(1499221380, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSwiftStructTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x53\x4b\x8b\x13\x41\x10\xbe\xcf\xaf\x28\x96\x1c\x36\xb0\x4c\xee\x81\x10\x64\x65\x2e\xe2\x03\x14\x3c\x88\x87\xde\x4c\x4d\x68\xe9\xd4\x0c\xdd\x3d\xbb\x3b\x34\x7d\xf0\xe4\x1b\x3c\x04\xaf\x5e\x44\xbc\xa8\x78\xf1\x81\x82\x7f\x26\xce\xee\xcd\xbf\x20\xd5\x99\xcc\x8c\x11\x1f\xb9\x38\x87\x64\xba\xba\xbe\xaa\xef\xfb\xaa\x66\x34\x02\xe7\x60\x60\x0c\x8c\x27\x10\x83\xf7\x51\x1b\x89\xaf\x88\x05\x82\xf7\xb1\x39\x91\x99\x8d\x46\x23\xbe\xaa\x5f\xbf\xa8\x9f\xdd\x5b\x7d\x7e\x7f\xb6\x7c\x07\x87\x79\x8a\x89\xaa\xe0\x6c\xf9\xbc\xbe\xff\xf4\xfb\x97\xc7\xe7\x6f\x3f\xac\x3e\x3e\x39\x7f\x79\xb7\x7e\xf0\xe8\xdb\xc3\x57\xab\xaf\x6f\xea\xe5\x27\x46\x46\x72\x51\xe4\xda\x42\x92\x97\x94\x0a\x2b\x73\xda\x44\x2e\x8b\xa2\x10\x47\x0a\xa3\xa8\x28\x8f\x94\x9c\x81\xb1\xba\x9c\xd9\x2d\x0e\xe3\x36\x0f\x5c\x04\x00\x7c\xad\x05\xcd\x11\x06\xf2\x00\x06\x59\x60\x9f\x48\x54\xa9\x01\xef\x03\x98\x38\xc6\x25\x66\x39\x59\x3c\xb5\x27\x5a\x14\x05\xea\x38\xc9\xf5\x42\x58\x4c\x13\xa9\x30\x0d\xe5\x07\xd9\xa6\x4d\x00\x6a\x34\xa5\xb2\xbf\x41\x5f\x13\xda\xa0\xbe\x51\x15\x01\x16\xfe\xbd\x0f\x8c\x1a\xfa\xc7\x42\x37\xed\x99\x75\x57\x6f\x93\xeb\x1c\xc8\x8c\xb1\x57\x0b\xb6\x41\x28\xf0\x7e\xea\x1c\xa0\x32\x81\xc1\xa4\x8f\xb9\x88\x99\x60\x2e\x01\x86\x94\xf6\x5f\xa2\x7e\x57\x49\xd2\xee\x0f\xc1\xfd\x1a\x9d\xee\x0b\xaa\xc6\x70\x81\xaa\xe9\xb0\xf1\x8e\x9f\xf6\x65\x5e\x0a\x9d\x82\x42\x0b\x8d\x44\x98\xb0\xd7\x37\xd7\x07\x06\x0f\xd7\xdc\x1c\x68\xb4\xa5\x26\x20\xa9\xc0\xb7\xf8\x7f\x18\x04\x7b\xbb\xfb\x2c\xfa\x1d\xd6\x35\x82\x3b\x0d\xf4\xd6\x61\x9e\x4a\x9a\x5f\xc2\xca\xc4\xbd\x84\xdb\x9d\x3d\x0c\x5d\xff\x6e\x8f\xe7\x8e\xc9\x29\x38\xd2\xf3\xa3\xd1\xd6\x29\x8f\xf1\x94\xf7\x33\x64\xb5\xd6\x48\xfa\x0f\x7a\xf9\xf9\x8b\xca\x66\x4b\x9a\xe3\xcf\x9a\x3b\xdd\x9b\x65\xd0\xf2\x58\x58\x04\xa4\x72\x01\x5d\xc1\x31\x5c\xb7\x5a\xd2\xfc\xa0\x8b\x81\xfb\xb3\xbc\xb6\xfe\x4c\xf0\x42\xb8\x9d\x3f\xaf\x09\xec\x31\xaa\x0d\xec\x6d\x8f\xcb\x47\x3f\x02\x00\x00\xff\xff\x16\x94\xdc\xf4\x91\x04\x00\x00")

func templatesSwiftStructTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSwiftStructTpl,
		"templates/swift/struct.tpl",
	)
}

func templatesSwiftStructTpl() (*asset, error) {
	bytes, err := templatesSwiftStructTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/swift/struct.tpl", size: 1169, mode: os.FileMode(420), modTime: time.Unix(1499221380, 0)}
	a := &asset{bytes: bytes, info: info}
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
	if err != nil {
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
	"templates/swift/service.tpl": templatesSwiftServiceTpl,
	"templates/swift/struct.tpl": templatesSwiftStructTpl,
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
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"swift": &bintree{nil, map[string]*bintree{
			"service.tpl": &bintree{templatesSwiftServiceTpl, map[string]*bintree{}},
			"struct.tpl": &bintree{templatesSwiftStructTpl, map[string]*bintree{}},
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

