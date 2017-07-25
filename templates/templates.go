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

var _templatesSwiftServiceTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\x4d\x6f\xd3\x40\x10\xbd\xfb\x57\x8c\x62\x1f\x6c\x29\x75\x7a\xb6\xf2\x41\x85\x28\xe2\x50\xa8\x5a\xc4\xa5\xea\x61\xbb\x1e\xa7\x0b\xf6\xae\xd9\x5d\x53\xac\xc8\x07\x4e\x7c\x4b\x1c\x2a\xae\x5c\x10\xe2\x02\x88\x0b\x1f\x02\x89\x3f\x53\xd2\xde\xf8\x0b\x68\xfd\xd1\x24\x4e\x03\xcd\x25\xf1\xf8\xcd\x9b\xf7\xe6\x8d\xd2\xeb\xc1\x64\x02\x8e\x52\x10\x0c\xc0\x87\xa2\xb0\xce\x2b\xfe\x4d\x92\x20\x14\x85\xaf\x8e\x58\xa4\xad\x5e\xcf\xbc\x9a\x7e\x78\x3b\x7d\xfd\xf8\xe4\xc7\x97\xd3\xe3\xcf\x70\x55\x84\xb8\x19\xe7\x70\x7a\xfc\x66\xfa\xe4\xd5\x9f\x9f\x2f\xce\x3e\x7d\x3d\xf9\xf6\xf2\xec\xdd\xa3\xe9\xd3\xe7\xbf\x9f\xbd\x3f\xf9\xf5\x71\x7a\xfc\xdd\x74\x5a\x2c\x49\x85\xd4\xb0\x29\x32\x1e\x12\xcd\x04\x6f\x2a\x5b\x24\x4d\xc9\x41\x8c\x96\x95\x66\x07\x31\xa3\xa0\xb4\xcc\xa8\x6e\x69\x80\x89\x05\x00\xa6\xd8\xe9\x18\x91\xd5\xc3\x1a\x48\xc2\xc7\x08\x0e\xeb\x82\x93\x18\x07\xa6\x65\x0b\xf5\xa1\x08\x55\x03\x6b\xa0\x0e\x15\x19\xd7\x06\x14\x23\x07\x27\xf1\x37\xe4\x38\x4b\x90\xeb\x65\xa4\x44\x95\x0a\xae\xb0\x61\xa4\x82\x6b\x7c\xa8\x8f\x24\x49\x53\x94\xfe\x36\x91\x0a\xe5\xed\x3c\x45\x43\xb3\x83\x3a\x93\xbc\x7c\xaa\x79\x94\x26\x9a\x51\x88\x32\x4e\x4b\x1b\x49\xe3\xc2\x3d\x1f\xd3\x8c\x62\x11\x70\x84\xf5\x46\xdc\x9c\x90\x79\xd4\x9c\xcb\xa8\xd4\xb4\x42\x7c\xbb\xcf\x91\x78\x3f\x43\xa5\x2f\xe1\x23\xf2\x6b\x07\x46\x70\xd4\x08\x0e\x4a\xfd\x35\x4b\x83\xe8\x9a\x14\xc0\xc4\x00\x17\xaa\x45\x1e\xb6\x45\xcd\xca\x54\x24\x69\x8c\xe6\x00\x02\xb8\x82\x8a\x92\x94\xf1\x31\xb8\x3b\xa8\xb2\x58\xf7\xab\x69\xd5\xf2\x9b\x71\x43\x0f\xd6\x86\x70\x47\xb0\xd0\xab\xaf\x60\xe9\x12\xcc\x27\x46\x0d\x29\xd1\x87\x30\x80\x4e\x7d\x3b\x2d\xbb\xd7\x51\x6f\x1b\x80\x93\x40\x51\x74\x56\x33\xcd\x72\xa9\x53\x59\x6f\xbf\xbd\x6c\x1e\x95\x26\x49\x12\x18\xc0\xc2\x5e\x7d\xc2\x73\x20\x6a\x04\x7b\xbb\x5a\x32\x3e\x0e\x60\x83\xe7\xfb\x30\x1a\xc1\x5e\xb0\x6f\xfd\x63\x99\x65\x29\x56\xb8\x6a\xca\x7f\xdb\x5b\x5e\xed\xfe\x2e\x2a\xc5\x04\xbf\xc1\x95\x26\x9c\xe2\xd0\xf6\xed\x3e\xe3\x0f\xc4\x3d\x1c\xda\xae\x59\x68\x50\xae\xb5\x5b\x8d\x08\xaa\xaf\x2e\xcc\xe7\x38\x81\x2a\x33\xbc\x75\x70\x17\xa9\x06\xc6\xad\x85\xf4\x59\x54\x6a\x94\x30\x00\x2d\xf3\x51\x2b\x62\x97\xf0\x3c\x58\x64\x30\x41\x2f\x9d\xd6\x6c\x62\x7d\x2d\xbe\xca\x28\x45\x0c\x5d\xe9\x79\x8b\xf0\xa2\xda\xd1\xa5\x48\x22\xc2\xe2\x4c\xa2\x6b\xf7\xaf\x49\x29\xe4\xd0\xf6\xbc\x45\xae\x19\x49\xd1\x85\x1a\x6d\x3c\xa3\x81\x2f\x79\x5d\x3d\xa0\xc4\xcf\x91\x17\xd5\xcf\xe2\xe2\x3f\xb6\x3a\xb8\xe2\x6f\x00\x00\x00\xff\xff\x96\xda\x46\x41\xa0\x05\x00\x00")

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

	info := bindataFileInfo{name: "templates/swift/service.tpl", size: 1440, mode: os.FileMode(420), modTime: time.Unix(1500961941, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSwiftStructTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x56\x3b\x8f\x13\x3d\x14\xed\xe7\x57\x5c\x45\x29\x36\xd2\xf7\x4d\xa8\x23\x45\x2b\x58\x18\x89\x82\x87\x04\xdd\x2a\x85\x37\xe3\x09\x96\x26\x1e\x73\xed\xd9\x30\x1a\x4d\x41\xc5\x5b\xa2\x88\x68\x69\x10\xa2\x01\x44\xc3\x43\x20\xf1\x67\x42\x76\x3b\xfe\x02\xb2\x33\x33\x3b\x0f\x27\x4a\xc4\x0a\x51\xd0\x24\xb2\x7d\x7d\x7c\xcf\xf1\xf1\x49\xfa\x7d\x48\x53\xe8\x4a\x09\x83\x21\xb8\x90\x65\x4e\x39\xe3\x5e\x27\x53\x0a\x59\xe6\xca\x19\x0b\x94\xd3\xef\xeb\xa5\xe5\xdb\x57\xcb\x17\x0f\x16\x5f\x3f\x9e\xcc\x3f\xc0\x41\xe4\x53\x2f\x4c\xe0\x64\xfe\x72\xf9\xf0\xf9\xcf\x6f\x4f\x4f\xdf\x7f\x5a\x7c\x7e\x76\xfa\xfa\xfe\xf2\xd1\x93\x1f\x8f\xdf\x2c\xbe\xbf\x5b\xce\xbf\xe8\x9d\x0e\x9b\x8a\x08\x15\x78\x51\xcc\x7d\xa2\x58\xc4\x8b\x99\x6b\x44\x08\x72\x14\x52\xc7\x11\xf1\x51\xc8\xc6\x20\x15\xc6\x63\xd5\xe8\x61\x50\xd6\x41\xea\x00\x80\x5e\xee\x74\x74\xbb\xab\xc1\xff\x80\x84\x4f\x28\x74\xd9\x7f\xd0\x0d\x0c\x17\x8f\xd1\xd0\x97\x45\x49\x51\xd6\xe5\x1a\x70\x30\x34\xe0\xe3\x88\x2b\x7a\x4f\xcd\x90\x08\x41\xd1\xf5\x22\x9c\x12\x45\x7d\x8f\x85\xd4\x37\x07\x77\x83\xa2\x81\x3a\x08\x52\x19\x87\x6a\x0d\xcc\x4d\x82\x92\xe2\xed\x44\x98\xfd\xe6\x3b\xdf\x9f\x33\x3c\x26\x68\xe8\xf1\x82\x9b\x1e\xac\x20\x8b\xf2\x34\x05\x16\xe8\xed\x37\x84\x16\x8b\x84\x90\x65\xfb\x69\x0a\x94\xfb\x55\xd2\xf9\xb0\x0a\xce\x38\x53\x7b\xbd\x5c\xa5\x96\x52\x7f\x5a\xad\xdf\x55\xac\x8a\xc1\x02\xa0\x77\x6b\x9a\x04\x24\x94\x8d\xcb\x29\x55\x85\x61\x55\xd6\xcb\x34\x20\xfa\x7c\x0b\x6c\x45\x52\xcb\x94\x45\x5c\xc2\x93\x01\x5c\xe4\xc9\x7e\x0f\x52\x73\xc6\x38\x8a\xb9\xa1\x16\x52\x6e\x57\xd2\x72\x03\x2c\x00\x4e\x8b\xbd\x17\xaa\xab\x93\x98\xa0\x0f\x21\x55\x30\x25\x02\x86\x40\x78\x02\x44\xee\xc3\xe1\x2d\x85\x8c\x4f\xcc\xd9\x23\xa0\x9a\x7a\x0a\xea\x0e\x46\xb3\xf2\x71\x5c\x41\x8c\xd0\xf5\xd9\xd8\x08\x84\xc9\x41\xc4\x8f\x29\x2a\x8f\xe8\x3b\xd2\x8d\xf7\x60\x43\x5b\x3b\x98\xe3\xdc\x4d\x72\x5e\x46\x29\x71\x98\xbc\x44\xe4\xba\xb6\xae\x9a\xc5\x15\x4e\xfd\xe1\x59\xc1\x5a\x4f\x71\xbd\xe1\x74\x69\x7e\x74\x96\x4d\x89\x38\xec\xe8\x92\x92\x73\x67\x64\xae\xd2\xfa\xe0\xe9\xca\xcc\x0a\x13\x5b\xc1\xca\x76\x36\xc4\x5e\x3d\x17\x5a\xfe\x6e\xbc\x11\x1b\xb9\xb2\xe3\xb2\x8a\x05\xc6\x81\xda\x7f\x3b\xb0\xa8\xc4\x4e\x5b\x1d\x51\xae\x65\xb9\x7b\x6b\xc5\x36\x27\x0b\x8c\x04\x45\xd5\xf0\xb1\x55\x84\x33\xf0\xb5\x4c\x9b\x4a\x54\x38\xee\xa4\xf7\x5f\xce\x72\x83\x15\xda\x4b\xdb\x24\xa0\xfe\x6c\xfe\x7e\x15\x21\xb8\x75\x04\x96\xf9\x6d\x49\x3c\xa4\x2a\x46\x0e\x87\x83\x91\xb3\xe9\xbe\x5a\x81\xa5\x1b\xd1\x61\x07\xc3\x7a\x38\xee\xf5\xfe\x85\xdc\x56\x21\xa7\xc5\x6b\xbf\xed\x76\x90\x55\x2c\x7e\x16\x54\xf5\xc9\x0d\x7f\x57\x5c\xc2\x13\x7b\x46\x59\xdd\xd7\xb8\xb0\xdc\x1c\xba\xd5\xf5\x16\xcd\x9c\x5f\x01\x00\x00\xff\xff\xd1\x71\x66\x79\xc9\x0a\x00\x00")

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

	info := bindataFileInfo{name: "templates/swift/struct.tpl", size: 2761, mode: os.FileMode(420), modTime: time.Unix(1500947588, 0)}
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

