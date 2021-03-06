// Code generated by go-bindata.
// sources:
// templates/report.html
// DO NOT EDIT!

package assets

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

var _templatesReportHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\xd9\x8e\xdb\x36\x17\xbe\xf7\x53\x30\x0c\xf0\xdf\xfc\xa1\x38\xf1\x4c\x83\x2c\x92\x81\x36\x99\x34\x6b\x1b\x64\xa6\x4b\x2e\x8f\xc8\x23\x89\x35\x45\xaa\x24\xe5\x59\x0c\xbf\x7b\x41\x6d\x23\xdb\x98\x00\x05\x3a\xba\xb0\xa8\xb3\x7e\x67\xa5\xd3\x47\xd2\x8a\x70\xd3\x20\xa9\x42\xad\x57\x8b\x34\xbe\x88\x06\x53\x66\x14\x0d\x8d\x04\x04\xb9\x5a\x10\x92\xd6\x18\x80\x88\x0a\x9c\xc7\x90\xd1\x36\x14\xec\x39\xbd\x63\x18\xa8\x31\xa3\x1b\x85\x57\x8d\x75\x81\x12\x61\x4d\x40\x13\x32\x7a\xa5\x64\xa8\x32\x89\x1b\x25\x90\x75\x1f\x4f\x88\x32\x2a\x28\xd0\xcc\x0b\xd0\x98\x3d\x7d\x42\x7c\xe5\x94\x59\xb3\x60\x59\xa1\x42\x66\xec\x91\x61\x89\x5e\x38\xd5\x04\x65\xcd\xcc\xf6\x91\x18\xb4\xa1\xb2\xee\x48\x42\x2b\xb3\x26\x0e\x75\x46\x95\x88\x06\x2a\x87\x45\x46\xb9\xb4\xc2\xf3\xb3\xe4\x84\x83\xf7\x18\x3c\x57\x75\xc9\x0b\xd8\x44\x19\x3f\x1e\x12\x25\x22\x9a\x68\x25\xa8\xa0\x71\xf5\x01\xcd\x5a\x19\x4f\xfe\x24\x97\xd0\x90\xaf\x18\xc3\x4d\x79\xcf\x5b\xec\x3b\xf3\xe1\x46\xa3\xaf\x10\xc3\xe8\xb2\x0a\xa1\xf1\x2f\x39\xf7\x01\xc4\xba\x81\x50\x25\xb9\xb5\xc1\x07\x07\x8d\x90\x26\x11\xb6\xe6\x13\x81\xff\x90\x9c\x24\x27\x0c\x74\x53\xc1\x92\x0b\xef\xef\x58\x49\xad\x4c\x22\xbc\xa7\x0b\x32\x3c\xca\x04\x2c\x9d\x0a\x37\x19\xf5\x15\x9c\x3e\x3f\x63\x6f\xaa\x6f\xcf\xac\x39\x7f\x56\x9c\xde\xde\x7e\xcc\xff\xfa\xed\xcb\x57\xb1\xac\x7e\xbd\xfd\xf9\x47\x79\x5e\x9c\xf1\x37\xb7\xff\xff\xe3\xc3\xd5\x4f\x9b\xf3\x6f\x9f\xb8\x5e\xaf\xdf\xfb\xf7\xa7\xaa\x6a\x8b\xbf\x5f\x54\xeb\x17\x1f\xcf\xf4\xef\xf6\x23\x25\xc2\x59\xef\xad\x53\xa5\x32\x19\x05\x63\xcd\x4d\x6d\x5b\x3f\x24\xa3\x8b\x6d\x35\xf8\x4f\x6a\x50\x86\x6c\x27\x34\x84\x34\x20\xa5\x32\x25\x0b\xb6\x79\x49\x96\x27\xcd\xf5\xab\x81\xb9\x8b\xca\x7c\xd0\x4e\x79\xdf\x5d\x8b\x34\xb7\xf2\x26\xbe\xe3\x37\xba\xae\x68\x8f\x18\x23\x6f\xd5\x35\x4a\x62\x60\x93\x83\x23\x8c\x75\x74\x03\x1b\x22\x34\x78\x9f\xd1\x81\xd1\xbf\x18\x5e\x37\x60\x24\xd3\xe5\x48\xd0\xaa\xac\x02\xc9\xcb\xfe\x40\x7b\xb4\xa9\x54\x93\x7e\x6c\x13\x50\x06\x1d\x2b\x74\xab\x24\x1d\xe3\x49\x61\xdf\x03\xcb\x1d\x18\x39\x16\xf1\x31\x5d\x7d\x51\x0d\x6a\x65\xd0\xa7\x1c\x26\xa5\xbc\x0d\xc1\x9a\x03\xcd\x60\xcb\x52\xa3\xa3\x24\x0e\x59\x46\x7b\x19\x4a\x24\x04\x18\x78\x11\x86\xd6\xd0\x78\x1c\xc9\xe0\xca\x38\x62\x8f\x7b\x13\x17\x6d\x13\x7b\x0c\xe5\xeb\xbe\xa9\xe9\x2c\xcd\xf1\x01\xa7\x80\xc5\x48\x9c\xd5\x93\xdf\x23\xa5\x5e\xac\x4f\x11\xca\x8c\x16\xa0\xa3\xc3\x8e\xaa\x21\x8f\xed\x7a\xd9\xc1\x89\xc9\x53\x25\x74\x93\xb6\x9a\x5c\xa5\xbe\x81\x7b\x42\x63\xdd\x4c\xad\x52\x1e\x45\xa6\x64\xf0\x3e\xd2\xe9\x7b\x2f\xeb\x7d\xb8\x63\x99\xee\xc2\x57\xf2\x5e\xfc\x33\x28\xad\x3e\x00\x12\x3b\xa2\x76\x0c\xda\x60\x49\x9d\xb3\x65\xfc\xd1\x25\x3b\x99\x29\x75\x73\x39\x53\x63\x2a\x60\xbd\xc7\x3f\xa8\x3a\x8b\x63\x3c\xab\xb8\xbd\x32\xe8\x66\xd5\xee\x83\xd4\xea\x3f\x74\xe1\xb0\xb1\x0f\xeb\xa1\x06\x1f\xd0\x91\xc7\x4f\x97\xa7\x0f\xe0\x88\x80\x08\x6a\x33\x36\x95\x68\x9d\xeb\x56\x70\x03\x25\xce\x30\xe8\xb8\xaa\xbe\x9f\xc9\x94\xb7\x7a\xf6\x55\x58\x57\x8f\xbe\x24\x2b\x34\x5e\xef\x17\x56\x99\xa6\x0d\xa3\x40\x14\x1e\xa7\x21\x36\xc5\x72\x9c\x3c\x8f\xe0\x44\x45\x49\xa3\x41\x60\x65\xb5\x44\x97\xd1\x8b\x81\x38\x1f\x83\x81\xb6\xe7\x62\x7f\xb4\xf3\x60\x48\x1e\x0c\xb3\x6d\x88\x4b\x80\xf9\x56\x08\xf4\x7e\xf2\xd4\xe6\xb5\x0a\x74\xd5\x1b\x3a\x9c\x84\x18\x5e\x04\x79\x37\x29\x52\x6d\x86\xd5\x34\x1e\x53\x6e\x60\x33\xae\xc7\xb8\x0e\x17\x69\xb7\x61\x9d\x8d\xeb\x22\x1e\xe9\x88\xa5\xa3\x4f\x7b\xac\x5f\xce\x84\x6c\xb7\x0e\x4c\x89\x24\xb9\x44\x1f\xfc\x6e\x37\xf9\x9e\x18\xe7\xce\x59\x17\x39\xf3\x30\xe7\x43\x0a\x4e\x1e\x56\xfc\x80\xcd\x7a\x74\x07\x52\x5d\x67\x8c\xd5\x26\xdd\xbd\x98\x51\x08\x64\xbb\x25\xc9\x27\x65\x90\xec\x76\xdd\xf9\xb5\xd5\x64\xb7\xa3\xab\x78\x7e\xab\x74\xa4\x1f\x34\x45\x7c\xfe\x67\x72\xdf\xbc\x7a\xd9\xbf\x0e\x78\xdb\x6d\xf7\xa7\xe5\x02\x0a\x24\xc9\x3b\x84\x78\xe1\xbc\xbb\xfc\xfc\x89\xcc\xe2\x3d\xc8\xf0\xbd\xa1\xc4\xfb\xe7\x38\x90\xc6\xe1\x2a\x15\x56\x62\x87\xf2\x33\x7a\x0f\x65\x0f\xb4\x23\xa6\x3c\x0a\x1c\x29\xed\x19\x0e\x78\x1d\x8e\x0d\xf3\x66\xf5\x7d\x88\x03\x61\x56\x37\xd4\x1e\x1f\xa4\x5c\xbf\x58\xf2\x15\x7d\xab\x83\xff\xd7\x90\x8c\x1c\x52\x3d\x9e\x17\x29\x8f\x0d\x19\x1b\x96\xf7\x57\x7a\xca\xbb\x3f\x96\xff\x04\x00\x00\xff\xff\xdf\x4d\xb8\x97\x68\x0a\x00\x00")

func templatesReportHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesReportHtml,
		"templates/report.html",
	)
}

func templatesReportHtml() (*asset, error) {
	bytes, err := templatesReportHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/report.html", size: 2664, mode: os.FileMode(420), modTime: time.Unix(1603914917, 0)}
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
	"templates/report.html": templatesReportHtml,
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
		"report.html": &bintree{templatesReportHtml, map[string]*bintree{}},
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
