// Code generated by go-bindata.
// sources:
// configs/csv-example.yaml
// configs/df.yaml
// configs/rule-example.yaml
// configs/top.yaml
// DO NOT EDIT!

package main

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

var _configsCsvExampleYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd5\xd5\xe5\xca\x2f\x28\xc9\xcc\xcf\x2b\xb6\xe2\x52\x50\xd0\x55\xc8\x48\x4d\x4c\x49\x2d\x2a\xb6\x52\x28\x29\x2a\x4d\x05\x8a\x28\x28\xa4\xa4\xe6\x64\xe6\x66\x96\x80\x05\x95\x74\x94\xc0\x62\xc9\xf9\x39\xa5\xb9\x40\x2d\x0a\x76\x60\xae\x82\x42\x75\x35\x48\xcc\xd0\x4a\xc5\x40\xa1\xb6\x16\x21\xa6\x02\x31\xce\x08\x55\x10\xcc\x05\x04\x00\x00\xff\xff\x7f\xcb\x5e\x7d\x7a\x00\x00\x00")

func configsCsvExampleYamlBytes() ([]byte, error) {
	return bindataRead(
		_configsCsvExampleYaml,
		"configs/csv-example.yaml",
	)
}

func configsCsvExampleYaml() (*asset, error) {
	bytes, err := configsCsvExampleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "configs/csv-example.yaml", size: 122, mode: os.FileMode(420), modTime: time.Unix(1441735447, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _configsDfYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x8f\xc1\x4a\x03\x31\x10\x86\xef\x79\x8a\xb9\x14\x5a\x24\x78\xf1\x34\x07\x8b\x07\x05\xf1\xa0\x08\x1e\x8a\x2b\x21\xcd\xce\xea\xb0\xbb\x4d\xd8\x99\x2a\x6b\xe9\xbb\xbb\x9b\x16\x8a\xed\xe6\x34\x99\xff\xfb\xbf\x10\x6b\xad\xa9\xa9\x47\x58\x3d\xbf\xbd\xba\xbb\x97\x47\xf7\x74\xbf\x32\x9f\x5d\xdc\x26\x04\x25\xd1\x3c\x9a\x36\x96\x84\xb0\xf6\x1a\xbe\x4c\x4c\xca\x71\x23\x68\x00\x2c\x54\x9e\x1b\x84\xdb\x61\x1e\xcf\x03\x37\x24\xbd\x28\xb5\x39\x6c\x47\x7e\x2a\x85\xe3\x66\xb7\x83\x4a\xdc\xba\x89\xa1\x16\x9c\x95\xb0\xdf\xdb\xc3\x25\xd7\x3b\x4a\xe4\x15\xe1\x26\xd3\x67\xb2\xa1\x5a\xb2\xd4\x38\x9f\x2f\xb1\xf5\x09\x16\xcb\xf7\xe2\xa7\xb8\x2e\xec\xc7\xd5\x62\xf0\x9c\x28\x8d\xea\x1b\x27\xfc\x4b\x87\x17\x4e\xc9\x56\xa8\x9c\x0c\xfc\xf7\xf0\xab\xc9\x24\xf8\xe4\x03\x6b\x8f\xb3\xf4\x6f\xcf\xa3\xeb\x1c\xe6\xaa\xa3\x0b\x43\x0a\x1b\x75\x47\x3c\x3b\xfe\x02\x00\x00\xff\xff\xd4\xdc\x99\xb9\x82\x01\x00\x00")

func configsDfYamlBytes() ([]byte, error) {
	return bindataRead(
		_configsDfYaml,
		"configs/df.yaml",
	)
}

func configsDfYaml() (*asset, error) {
	bytes, err := configsDfYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "configs/df.yaml", size: 386, mode: os.FileMode(420), modTime: time.Unix(1441679313, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _configsRuleExampleYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x55\xc1\x6e\xdb\x38\x10\xbd\xfb\x2b\x78\x31\x10\x6f\xe4\x64\x37\x7b\xd3\x61\x83\xec\x6e\x51\x14\x49\x5a\x23\x45\x0a\x04\x96\x2b\xd0\xe2\x44\x26\x6c\x89\x2c\x49\x59\x48\x83\xfc\x7b\x39\x14\x65\x4b\x32\x15\xd4\x07\x99\xf3\xe6\x71\xe6\x0d\x39\x24\xe7\xf3\xf9\x64\x0b\x2f\x31\x79\xfa\xf2\xf8\x90\xde\x2c\x3e\xa5\xb7\x1f\x9e\x26\xb9\x12\x95\x8c\x89\x01\x6d\xdc\x70\x52\x08\x06\x31\x59\x53\x93\x6d\x26\x42\x1a\x2e\x4a\x1d\x4f\x08\x99\x93\x02\xa1\x98\xfc\x63\x0d\xfc\x2d\x94\xc8\x40\x6b\xd0\x31\xf1\xc8\xeb\x2b\x91\x16\x4c\x8d\x30\x74\x17\x4f\x19\x79\x7b\x23\x6e\x1c\x0d\x19\xaa\x2a\x4b\x5e\xe6\x9e\xe3\xad\x13\x96\x36\x55\xb6\xf5\x1c\x37\x3e\x65\xec\x00\xe4\x31\x50\x6b\x76\x79\x65\x55\xa4\x66\xa3\x80\x32\xdd\x4a\x6a\xac\x50\x4d\x77\x82\x32\x72\xb3\xcf\xbb\x25\xed\x2c\x96\xfe\x15\x4f\x9f\xed\xdc\x68\x88\x5f\x8d\xe0\x7f\x37\x78\x28\xc7\x7f\x8b\x47\x52\x69\x9a\x43\x37\x49\x26\xab\xd4\x81\xf6\x0b\x2a\x9e\x4a\x14\x8a\xc3\x28\x48\xd2\x2f\xda\x73\xec\x28\x0a\x31\x38\xdb\x81\xa7\xe0\x30\x24\xe4\xeb\x86\x2a\x60\x77\x7c\xed\xf6\xb7\x55\xcf\xd7\xa9\x02\xcd\x19\x94\x26\x9e\x16\x50\xb8\x0d\xf2\x40\xd4\xe7\x31\x6a\xe8\x81\x83\xc6\xc0\xbf\xe3\xe5\x16\x18\x3f\xc6\x69\x81\x8b\x90\x9c\x7b\x28\x1e\x20\x6f\xdb\xcd\x87\xb1\x13\xad\x1c\x87\xb6\x6d\xe5\x63\x35\x8d\x15\x66\xfe\x4e\x01\x5d\xbe\x54\x7c\x4f\x0d\x1c\xe8\xde\x1e\x61\x6b\xb7\x6e\x07\x72\x63\x06\x4b\x5a\x6c\x5e\xb4\x2d\xab\x53\x8f\xb4\x08\x6e\xf1\x71\x3a\x1a\xe4\x6c\xc0\xa8\x79\x37\x83\xb3\x66\xd1\x30\x4a\xd9\x8f\xe3\xcc\xa0\x8c\x6f\xf7\x1d\x05\xfb\x22\xdd\x6b\xfe\xf3\x58\xac\xb3\xa2\x1e\xe1\x59\xd1\x02\x06\x34\x87\xd5\x42\x6d\x43\x13\x74\x4d\xed\xb9\xf3\x07\xec\xec\xcf\x19\xf1\xc0\x29\x4b\x54\x66\x40\x43\x24\x28\xfb\x33\x18\xcc\x77\x68\x07\x49\xb3\x2d\xd8\xd9\xee\x54\x83\x49\xbd\x9d\xf2\x32\x3e\x4b\xd8\x79\x72\x69\x3f\xcb\x7f\x6f\xef\x3f\xae\xae\x67\xae\xef\xcb\x4e\xfa\xee\x04\x9b\x31\x3c\xc3\x3a\x82\x4a\xfe\xe7\x7a\xdb\xed\x4a\x66\xed\x14\x2f\x91\x70\x18\xf4\x44\x03\x76\xad\xb8\x31\x30\xa2\xd4\x3b\x83\xb9\x71\xbb\x39\x8b\xbf\xe3\xc4\xd9\xf2\x7c\xbe\xba\x6e\xee\x95\xf6\xc4\x8b\xa2\xa0\xa5\x15\xb2\x4c\xea\xe4\x22\x99\x27\x7a\x75\x3e\xeb\x31\x64\x66\x52\x7b\x2f\xb8\xcc\x17\x18\xa4\xe7\x35\xbc\x00\x74\xbd\x5e\xbd\xc5\xee\x6b\x83\xe0\x5f\x9f\xd5\x5c\xa2\x98\x84\x25\x97\xc3\x04\xe8\xac\x7f\xbc\xe3\x94\x42\xd9\x6d\x1b\x29\xc0\xb6\x57\xdb\x63\x1d\xcd\x95\xca\x4f\xd1\xac\x90\x4a\x07\xc8\xb9\x92\x4d\x4b\x75\x30\x5c\xb3\x01\xa6\x8d\x3b\xe3\x75\x0f\x5c\x0b\xa1\x9d\xb8\x3f\xae\x71\x53\x50\x5e\xb2\x3c\x8c\x56\xfd\x62\xf0\x7a\xb5\xcb\xd5\x5e\xed\x1d\x54\x98\x0d\xa0\xb6\xbe\xa7\x3a\x15\xf1\x4c\xab\xdd\xf8\x62\x64\xa2\x1e\x5d\x27\x9d\x6b\xbc\xd2\xc6\xdd\x0a\xb2\xfd\x98\xdb\x3e\x13\x6b\xcd\xde\xf1\x16\xd4\x36\xdd\x98\x2a\x3d\xaa\x4a\xda\xb7\x06\x0f\xfe\x88\x1b\x9f\x9e\xf1\xb9\xa2\xb6\x4f\x5d\xb8\x2d\x9b\x57\x10\xb7\xea\x57\x00\x00\x00\xff\xff\xd2\x90\xbd\x51\xb6\x08\x00\x00")

func configsRuleExampleYamlBytes() ([]byte, error) {
	return bindataRead(
		_configsRuleExampleYaml,
		"configs/rule-example.yaml",
	)
}

func configsRuleExampleYaml() (*asset, error) {
	bytes, err := configsRuleExampleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "configs/rule-example.yaml", size: 2230, mode: os.FileMode(420), modTime: time.Unix(1441338602, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _configsTopYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x55\x4d\x8f\x9b\x30\x10\xbd\xef\xaf\xe0\x12\x69\xd3\x90\x5d\x75\xab\x5e\x38\x34\xea\x87\xd4\x43\x77\xab\x55\xab\xf6\x12\x22\xe4\xe0\x09\xb1\x82\xb1\x6b\x9b\xa0\xed\x2a\xff\xbd\x63\x63\x12\x20\x46\x6a\x0e\xc4\xf3\xde\xf3\x7c\xd8\xc3\xb0\x5c\x2e\x6f\xb8\xa0\x90\x44\x5b\x62\xf2\xfd\x8d\x90\x86\x89\x4a\x27\x37\x51\xb4\x8c\x76\x84\x95\x49\xf4\x01\xd7\xf6\xf7\xac\x44\x0e\x5a\x83\x27\xb9\xd5\x87\xd8\xc8\x23\xaf\xaf\x91\x44\x30\x33\xc2\x90\x32\x99\xd1\xe8\x74\x8a\xdc\x3a\x1e\x2b\x54\x5d\x55\xac\x2a\xbc\xc6\x5b\x57\x2a\x6d\xea\xfc\xe0\x35\x6e\x7d\xad\x28\x01\xe4\xc5\x51\x67\xf6\x75\x55\xcd\x33\xb3\x57\x40\xa8\xee\x52\x6a\xad\x50\x4d\x8f\x82\xd0\xe8\xe3\xb1\xe8\x97\x54\x22\x96\xbd\x4d\x66\x3b\xdc\x1b\x8f\xf1\x87\x09\xfc\x5d\x8b\x87\x62\x7c\x7e\xfe\x15\xd5\x9a\x14\xd0\x0f\x92\xcb\x3a\x73\x20\x3e\x41\x25\x33\x69\x13\xb5\xcb\x38\x28\xd2\x2f\xda\x6b\x70\x15\x87\x14\x8c\x96\xe0\x25\x76\x19\x4a\xe4\xe7\x9e\x28\xa0\x8f\x6c\xeb\xee\xb7\xcb\x9e\x6d\x33\x05\x9a\x51\xa8\x4c\x32\xe3\xc0\xdd\x05\x79\x20\x1e\xea\x28\x31\xe4\xac\xb1\xc6\x88\x2f\x59\x75\x00\xca\x2e\x7e\x3a\xe0\x2e\x94\xce\x13\xf0\x1f\x50\x74\xbd\xe8\xdd\xe0\x46\x4c\xc7\xa1\x5d\x5b\x79\x5f\x6d\x63\x85\x95\xff\x53\x40\x5f\x2f\x15\x3b\x12\x03\x67\xb9\xb7\x27\xd4\xda\x9d\xdb\x59\xdc\x9a\xc1\x92\x9e\xf7\x2f\x1a\xcb\xea\xd5\x23\x11\xb1\x57\x7c\xd9\x6e\x8d\xe8\xb6\xa3\x1a\xd6\x77\xed\xac\x79\x3c\xde\x5e\x0d\x1d\x38\x33\x18\xff\xf7\x53\x2f\xf4\x91\x67\x47\xcd\xfe\x5e\xaa\x74\x56\x3c\x10\xec\x14\xe1\x30\x92\x39\xac\x11\xea\x10\xda\xa0\x1b\x82\x2f\x9c\x4e\x6e\x53\xba\x48\xdd\x63\x3e\x77\x87\xd2\xe2\xd7\x62\x51\x9b\xb0\xda\x12\xc1\x22\xbe\x83\xb1\xd1\xcf\x5d\x21\x49\x7e\x00\x74\xe2\x5e\x6e\x30\x99\xb7\x33\x56\xb5\x7e\xef\xf1\xb1\xfe\xf4\xed\xe9\xeb\x66\xe5\x9c\xb3\xaa\x97\x45\x7f\x03\x46\x0c\xef\x40\x22\x98\xc9\x17\xa6\x0f\xfd\xe6\xa4\x68\x67\x76\x96\x84\xdd\x58\x26\x1e\xa9\x1b\xc5\x8c\x81\x89\x4c\x3d\xd9\xc6\x56\x20\x81\x98\x24\x7a\xef\x3c\x8c\x32\xb1\xad\xc0\xda\xb0\xf3\xf5\x62\xb9\x59\xb5\xb3\xa6\x9b\x02\x82\x73\x52\x21\xbd\x4e\x9b\xf4\x2e\xd5\x9b\xc5\x7c\xc0\xcb\xdc\x64\x38\x29\xdc\xf6\x3b\xeb\x62\xc0\x1a\xc6\xc1\x52\xaf\x0f\xa7\xc4\x3d\xd1\x85\xfd\x1b\xaa\xda\xb1\x6a\x43\xd0\xf4\x7e\x1c\xc0\x92\xcd\x9f\x76\xdc\x0e\x51\x29\x94\xbf\xff\x79\xba\x18\x66\x8d\xed\xd6\xf5\x5c\x2f\xd5\x5a\x15\xd7\x68\xce\xa5\xd2\x01\x71\xa1\xe4\x38\xa8\xb4\xe7\x34\xc2\xb4\x71\x2f\x7b\x33\x00\xb7\x42\x68\x97\xda\x9b\x95\xbd\x9b\xc5\x2a\x5d\xfb\xff\xcd\xb0\x38\x3b\x63\xf1\x84\xba\xf9\xde\x43\x85\xd9\x83\xcd\x6b\xc8\xd4\xd7\x09\xec\x48\x5d\x4e\x1d\x43\x2e\x9a\x89\xf3\xd1\x85\xb6\x33\x6d\x8a\x54\x90\x1f\xc3\x24\x7e\x23\xb6\x9a\x4e\x72\x9c\x60\x6b\x85\x73\xd1\x13\xb9\x48\xfc\xc4\x74\xaf\xfd\x15\x69\xbf\x37\x53\xfb\x44\x83\x5f\xb7\x70\xdf\xb5\x1f\x3e\x77\x29\xff\x02\x00\x00\xff\xff\x20\xa7\x40\xd0\xa4\x08\x00\x00")

func configsTopYamlBytes() ([]byte, error) {
	return bindataRead(
		_configsTopYaml,
		"configs/top.yaml",
	)
}

func configsTopYaml() (*asset, error) {
	bytes, err := configsTopYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "configs/top.yaml", size: 2212, mode: os.FileMode(420), modTime: time.Unix(1441770561, 0)}
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
	"configs/csv-example.yaml": configsCsvExampleYaml,
	"configs/df.yaml": configsDfYaml,
	"configs/rule-example.yaml": configsRuleExampleYaml,
	"configs/top.yaml": configsTopYaml,
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
	"configs": &bintree{nil, map[string]*bintree{
		"csv-example.yaml": &bintree{configsCsvExampleYaml, map[string]*bintree{
		}},
		"df.yaml": &bintree{configsDfYaml, map[string]*bintree{
		}},
		"rule-example.yaml": &bintree{configsRuleExampleYaml, map[string]*bintree{
		}},
		"top.yaml": &bintree{configsTopYaml, map[string]*bintree{
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

