// Code generated by go-bindata.
// sources:
// configs/csv-example.yaml
// configs/df.yaml
// configs/rule-example.yaml
// configs/test.yaml
// configs/top-new.yaml
// configs/top-old.yaml
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

var _configsDfYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x8d\xcd\x4e\xc3\x30\x0c\x80\xef\x7d\x0a\x1f\x8a\x04\x87\x48\xfc\x0d\x24\x1f\x38\xf2\x1a\x95\x97\x78\x9a\xb5\x64\xa9\x62\x03\x1a\xd3\xde\x9d\xb4\x14\x8d\x48\xcd\x29\xfe\xbe\xf8\x8b\x73\xae\x4b\x39\x30\x42\x2e\x81\x0b\x87\x2e\x8f\x26\xf9\xa8\xd8\x01\x38\x48\x64\x7e\x8f\xf0\x56\x87\xe9\xbc\x4b\x64\x3d\xa9\x71\x82\x85\x9c\xcf\xb0\xd3\x61\x1b\xb3\x3f\x28\xde\xa4\x2a\x2e\x17\xf7\x3b\xce\x81\x3d\x53\xcd\xe2\xbf\xcd\x79\xb1\xe4\x2f\x45\xd8\xcc\x77\xa3\x6d\xe4\xeb\x1f\xb5\x18\x44\x0f\xd8\xdf\xd7\xd4\x95\xdd\xca\xd1\xee\xc0\xb2\x51\x1c\x54\xbe\x19\xfb\x87\x15\xff\xa1\x1c\x16\xfd\xb8\xa2\xe9\x93\xe4\x6f\xfd\xa9\xf1\x9e\x46\xf2\x62\x27\xec\x9f\x1b\x2e\x53\x11\xfb\x4d\x0b\x77\x85\x6b\xe1\xa5\x81\xa3\xb7\x61\x79\xfd\x3a\x89\x9f\x00\x00\x00\xff\xff\xe9\x8b\xc3\xb6\x5a\x01\x00\x00")

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

	info := bindataFileInfo{name: "configs/df.yaml", size: 346, mode: os.FileMode(420), modTime: time.Unix(1442033341, 0)}
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

var _configsTestYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\xca\x31\x0e\x82\x21\x0c\x47\xf1\x9d\x53\xfc\xf3\x85\xd1\x26\xc0\xd8\xc1\xcb\xd8\x0e\x24\x60\x0d\xe0\x44\xb8\xbb\xe0\xa6\xe3\xfb\xe5\x11\x91\xab\x26\xca\xb0\x26\xda\x54\x9c\xbd\x46\xb6\x67\x67\x07\x10\x44\x4b\xae\x79\x68\xeb\x8c\xeb\x76\x6d\x03\x1e\x56\xde\x75\x0f\xb8\x7f\x13\x98\xf3\x58\x60\x1f\xb0\xd6\x8f\x45\xf6\xf1\xdf\x12\xfb\x74\xec\x13\x00\x00\xff\xff\x7f\xdd\x9b\x22\x7a\x00\x00\x00")

func configsTestYamlBytes() ([]byte, error) {
	return bindataRead(
		_configsTestYaml,
		"configs/test.yaml",
	)
}

func configsTestYaml() (*asset, error) {
	bytes, err := configsTestYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "configs/test.yaml", size: 122, mode: os.FileMode(420), modTime: time.Unix(1442033311, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _configsTopNewYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x54\x5f\x6f\x9b\x3e\x14\x7d\xef\xa7\xb0\xf4\xfb\x45\x6a\x35\xb2\x69\x9b\xf6\xc2\xc3\xa6\x34\x65\x69\xb4\x10\x18\x90\x56\xd5\x3a\x21\x27\x76\x13\x2b\xe1\x8f\xb0\x49\xd4\x55\xfd\xee\xb3\x8d\x49\xe0\xc6\x95\xc6\x03\xf2\x3d\xf7\xdc\x7b\xcf\x31\xc6\xc3\xe1\xf0\x22\x2b\x08\x75\xd1\x12\x8b\xd5\xe6\xa2\x28\x05\x2b\x72\xee\x5e\x20\x34\x44\x4f\x98\xed\x5c\xf4\x55\xae\xd5\x13\x56\xc5\x8a\x72\x4e\x4d\x32\x53\x7c\x5b\x16\x19\xe4\xe5\x05\x95\x12\x4c\x45\x21\xf0\xce\x1d\x10\xf4\xfa\x8a\xf4\xda\x81\x8c\xaa\xce\x73\x96\xaf\x0d\xc7\x44\x67\x2c\x2e\xea\xd5\xd6\x70\xf4\xfa\x9c\xb1\xa3\xb4\x3c\x35\x6a\xc3\x2e\x2f\xaf\xb3\x54\x6c\x2a\x8a\x09\x6f\x25\x35\x91\xcd\xd3\xac\xc0\x04\x8d\xf6\xeb\xae\xa5\x9d\xc4\xd2\x8f\xee\xe0\x49\xd6\x3a\x10\xff\xf4\x06\xfe\xb9\xc1\x6d\x33\xc6\xe1\x02\xd5\x1c\xaf\x69\x77\xc8\xaa\xac\x53\x0d\xca\x37\xad\xdc\x41\xa9\x84\xaa\xa5\x63\x25\xf1\x67\x6e\x38\x72\xe5\xd8\x18\x8c\xec\xa8\xa1\xa8\xa5\x4d\x48\xbc\xc1\x15\x25\x33\xb6\xd4\xdf\xb7\x55\xcf\x96\x69\x45\x39\x23\x34\x17\xee\x20\xa3\x99\xfe\x40\x06\x70\xfa\x3c\x82\x05\x3e\x72\x54\x00\xf2\x3b\x96\x6f\x29\x61\xa7\x3e\x2d\xf0\xde\x26\xc7\xa7\x59\x44\xd7\xed\x59\x34\x6d\x64\xa1\x94\xa3\xd1\xf6\x58\x99\x5e\xcd\xc1\xb2\x33\xff\xc5\x40\x97\x5f\x56\x6c\x8f\x05\x3d\xd2\x4d\xfc\x06\x9b\xeb\x7d\x3b\x92\x9b\xd0\x6a\x29\xdc\x3c\x73\x69\xab\xe3\xa7\x94\x88\xfa\xc4\xa7\x72\x15\xa0\xcb\x36\x75\x60\xdd\xd6\x3a\xba\x72\x60\x79\xde\x6f\xa0\x43\xeb\xfc\x3b\xbf\x33\x7a\x9f\xa5\x7b\xce\xfe\x9c\x5c\xea\xc8\xe9\x11\x9e\x2a\x9c\x51\x40\xd3\xd8\xa1\xa8\xb6\xb6\x02\x7e\xc0\xf2\x87\xe3\xee\xe5\x23\x79\xf7\xa8\x5f\x57\x57\x7a\x53\x1a\xfc\x9c\x5c\xd4\xc2\xce\x56\x09\xab\x89\x39\x15\x6a\xfa\xf1\x54\x94\x78\xb5\xa5\xb2\x89\xfe\xb9\xa9\x48\x4d\x9c\xb2\xbc\xe9\xfb\x41\xbe\x7e\x5d\xff\xf0\x27\xbf\xbf\xe9\xe6\x2c\xef\xa8\xe8\x16\xc8\x89\xf6\x0a\x99\xb0\x2a\xb9\x61\x7c\xdb\x3d\x9c\x44\xc6\xa9\xba\x4b\xec\x6d\x54\xc6\x01\xec\x43\xc5\x84\xa0\x6f\x28\x35\xc9\x66\xf6\x46\x56\xcb\x9b\x00\x85\xd3\x1b\xdd\x83\xe6\xc4\x3d\xdd\xb9\x1a\xaa\x8a\x83\xdc\x86\x2f\x7a\x2d\xf0\x52\xfe\xf0\x47\xa9\x72\xde\xff\xb2\xb2\xb9\x82\x5a\x60\x1c\xf8\xfe\x68\x0e\xc0\x81\xba\x8f\x7a\x48\x32\xf5\xbd\x3e\xf2\x5f\x72\x0b\x80\xfb\x9f\x00\x08\x83\x28\x89\xfb\x98\xef\xf9\x7d\x20\x5c\x44\x13\xa0\xc8\x0f\x23\x50\x15\x4e\xa2\x10\x20\x67\x46\xe2\x64\x94\x00\x89\xd7\x41\x10\x43\x01\xca\x5a\x0a\xbd\x68\x30\x48\x6e\xe1\xdc\x05\x1c\xf2\x7d\xb4\x98\xc1\x8e\xe3\xe0\x1e\x78\x8c\x27\xb1\x37\x4f\xce\xc0\xc8\x1b\xdf\x01\xcd\x0f\xf1\x75\x0c\x7d\x3c\xc4\xfe\x68\x0c\xf6\x76\x1c\x83\x19\xe1\x68\xe2\x4d\xe7\x40\xc9\xf4\x66\xe6\x41\x5e\x70\xef\x45\xc0\x54\xdc\x20\x7f\x03\x00\x00\xff\xff\x8d\xf9\x36\xb0\xf5\x07\x00\x00")

func configsTopNewYamlBytes() ([]byte, error) {
	return bindataRead(
		_configsTopNewYaml,
		"configs/top-new.yaml",
	)
}

func configsTopNewYaml() (*asset, error) {
	bytes, err := configsTopNewYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "configs/top-new.yaml", size: 2037, mode: os.FileMode(420), modTime: time.Unix(1442020009, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _configsTopOldYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x55\x4d\x8f\xdb\x20\x10\xbd\xef\xaf\xf0\x25\x52\xdc\x38\x59\xed\x56\xed\xc1\x87\x46\xfd\x90\x7a\xe8\x6e\xb5\x6a\xd5\x5e\x92\xc8\x22\x66\xd6\x41\x09\x86\x02\x8e\xb5\x4d\xf3\xdf\x0b\x18\x27\xd8\xc6\x55\xf7\xe0\x65\xde\x7b\xcc\x07\x4c\x86\xf9\x7c\x7e\x43\x19\x86\x34\xda\x22\x95\xef\x6e\x18\x57\x84\x95\x32\xbd\x89\xa2\x79\x44\x0d\x94\x46\xef\xb4\x61\xfe\x9e\x04\xcb\x41\x4a\x90\x69\xe4\x90\xd3\x29\xe2\x1a\xcc\x14\x53\xe8\x90\x4e\x70\x74\x3e\x47\x76\x9d\xf4\x15\xa2\x2a\x4b\x52\x16\x4e\xe3\xac\x81\x4a\xaa\x2a\xdf\x3b\x8d\x5d\x0f\x15\x07\x00\x7e\x75\xd4\x9a\xbe\xae\xac\x68\xa6\x76\x02\x10\x96\x6d\x4a\x8d\x15\xaa\xe9\x81\x21\x1c\xbd\x3f\x16\x7e\x49\x07\x8d\x65\x77\xe9\xe4\x59\xef\x4d\xfa\xf8\xfd\x08\xfe\xba\xc1\x43\x31\x3e\x3e\xfd\x88\x2a\x89\x0a\xf0\x83\xe4\xbc\xca\x2c\xa8\xbf\x20\xd2\x09\x37\x89\x9a\x65\x12\x14\xc9\x17\xe9\x34\x7a\x95\x84\x14\x04\x1f\xc0\x49\xcc\x32\x94\xc8\xf7\x1d\x12\x80\x1f\xc8\xd6\xde\x6f\x9b\x3d\xd9\x66\x02\x24\xc1\x50\xaa\x74\x42\x81\xda\x0b\x72\x40\xd2\xd5\x61\xa4\xd0\x45\x63\x8c\x1e\x7f\x20\xe5\x1e\x30\xb9\xfa\x69\x81\x45\x28\x9d\x47\xa0\xdf\xa0\x68\xdb\xcd\xb9\xd1\x1b\x75\x3a\x16\x6d\xdb\xca\xf9\x6a\x1a\x2b\xac\xfc\x9f\x02\x7c\x3d\x17\xe4\x88\x14\x5c\xe4\xce\x1e\x51\x4b\x7b\x6e\x17\x71\x63\x06\x4b\x7a\xda\xbd\x48\x5d\x96\x57\x0f\xd7\x88\xb9\xe2\xeb\x76\x63\x44\xd3\x96\xaa\x89\xef\xda\x5a\x71\xd2\xdf\x5e\x76\x1d\x58\x33\x18\xff\xe7\xa3\x17\xfa\x48\xb3\xa3\x24\xbf\xaf\x55\x5a\x2b\xe9\x08\x9e\x05\xa2\xd0\x93\x59\xac\x66\x62\x1f\xda\x20\x6b\xa4\x7f\x70\x32\x9d\xae\xf1\x6c\x6d\x3f\x71\x6c\x0f\xa5\xc1\x87\x62\x56\xa9\xb0\xda\x10\xc1\x22\xbe\x82\x32\xd1\x2f\x5d\xc1\x51\xbe\x07\xed\xc4\xfe\xb8\x41\x65\xce\xce\x48\xd9\xf8\xbd\xd5\x9f\xd5\x87\x2f\x8f\x9f\x37\x4b\xeb\x9c\x94\x5e\x16\xfe\x06\x1d\x31\xbc\x43\x13\xc1\x4c\x3e\x11\xb9\xf7\x9b\x13\x6b\x3b\x33\xb3\x24\xec\xc6\x30\x49\x4f\x5d\x0b\xa2\x14\x8c\x64\xea\xc8\x26\xb6\x00\x0e\x48\xa5\xd1\x1b\xeb\xa1\x97\x89\x69\x05\xd2\x84\x8d\x57\xb3\xf9\x66\xd9\xcc\x9a\x76\x0a\x30\x4a\x51\xa9\xe9\xe9\x32\x5d\xd7\xf1\xe2\x74\x97\xdc\xbd\x3d\xc7\x1d\x0d\xcf\x55\xa6\xa7\x85\x75\xb1\x30\x6e\x3a\xac\x22\x14\x0c\x75\xba\x3f\x1b\x1f\x8b\x3f\x69\x3c\x34\xba\x5b\x9a\x39\x9b\x4e\x57\x6b\xbc\xbe\xdd\xcc\x86\x64\xfd\xeb\x1f\x24\x67\xc2\xf5\x45\xa0\x1e\xdd\x88\x6d\x37\x7a\x05\x54\xa2\x18\xa2\x39\xe5\x42\x06\xc4\x85\xe0\xcd\xf4\xf7\x30\x73\x82\x3d\x4c\x2a\x3b\x06\xea\x0e\xb8\x65\x4c\xda\xe4\x5e\x2d\xcd\xad\xcd\x96\xeb\x95\xfb\xbf\xe9\x16\x62\xa6\xaf\x3e\xb7\x76\xf2\x7b\x28\x53\x3b\x30\x79\x75\x99\x6a\x98\xc0\x33\xaa\x0e\xed\x41\xe8\x08\xbd\x6b\xad\xc3\x04\x95\x85\x34\xd3\x6e\x8c\x14\x90\x1f\xc3\xa4\x7e\x3d\xb6\x12\x8f\x72\x14\xe9\xa6\x0b\xe7\x22\x47\x72\xe1\xfa\xf1\x69\x07\xc2\x80\x34\x2f\xd1\xd8\x3e\x56\xeb\x77\x2f\xdc\x8d\xcd\x93\x68\x2f\xe5\x6f\x00\x00\x00\xff\xff\xe2\x1e\xce\x1a\xa1\x08\x00\x00")

func configsTopOldYamlBytes() ([]byte, error) {
	return bindataRead(
		_configsTopOldYaml,
		"configs/top-old.yaml",
	)
}

func configsTopOldYaml() (*asset, error) {
	bytes, err := configsTopOldYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "configs/top-old.yaml", size: 2209, mode: os.FileMode(420), modTime: time.Unix(1442000714, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _configsTopYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x54\xdf\x6f\xda\x30\x10\x7e\xef\x5f\x61\x69\x45\x2a\x5a\xd8\xb4\x4d\x7b\xc9\xc3\x26\x4a\x33\x8a\x46\x48\x96\x84\xa2\x6a\x9d\x22\x17\xbb\x60\x41\x7e\x28\x36\xa0\xae\xea\xff\xbe\x8b\x93\x40\x72\xb8\xd2\x78\x40\xbe\xef\xbe\xbb\xfb\x3e\xdb\xf1\x60\x30\xb8\x48\x32\xc6\x6d\x92\x15\x8c\x17\x9c\x5d\x64\xb9\x12\x59\x2a\xed\x0b\x42\x06\x24\xa1\x6a\xb9\xb6\xc9\x37\x08\xca\x9f\x5f\x64\x4b\x2e\x25\x97\x36\xa9\x91\x97\x17\x92\x03\x18\xab\x4c\xd1\xad\xdd\x63\xe4\xf5\x95\xe8\xb5\x85\x19\xc5\x2e\x4d\x45\xba\xaa\x39\x75\x74\xc6\x92\x6a\xb7\xdc\xd4\x1c\xbd\x3e\x67\x6c\x39\xcf\x4f\x8d\x9a\xb0\xcd\x4b\x77\x49\xac\xd6\x05\xa7\x4c\x36\x92\xaa\xc8\xe4\x69\x9a\x51\x46\x86\xfb\x55\xdb\xd2\x16\xb0\xf8\x93\xdd\x7b\x82\x5a\x0b\xe3\x9f\xdf\xc0\xbf\x54\xb8\x69\xc6\xc8\x9f\x93\x9d\xa4\x2b\xde\x1e\xb2\xcc\x77\xb1\x06\xe1\x9f\x17\x76\x2f\x2f\x85\x96\x4b\xcb\x48\x92\xcf\xb2\xe6\xc0\xca\x32\x31\x04\xdb\xf2\x9a\x52\x2e\x4d\x42\xc2\x35\x85\x33\x9e\x8a\x47\x7d\xbe\x8d\x7a\xf1\x18\x17\x5c\x0a\xc6\x53\x65\xf7\x12\x9e\xe8\x03\xaa\x01\xab\xcb\x63\x54\xd1\x23\xa7\x0c\x50\x7e\x2b\xd2\x0d\x67\xe2\xd4\xa7\x01\x3e\x98\xe4\xb8\x3c\x09\xf8\xaa\xb9\x6e\x75\x1b\x28\x04\x39\x1a\x6d\xae\x55\xdd\xab\xba\x58\x66\xe6\xff\x18\x68\xf3\xf3\x42\xec\xa9\xe2\x47\x7a\x1d\xbf\xc1\x96\x7a\xdf\x8e\xe4\x2a\x34\x5a\xf2\xd7\xcf\x12\x6c\xb5\xfc\xe4\x80\x94\x47\x7c\x2a\x2f\x03\x72\xd5\xa4\x0e\xa2\xdd\x5a\x47\x7d\x0b\x97\xa7\xdd\x06\x3a\x34\xce\xbf\x73\x5b\xa3\xf7\x49\xbc\x97\xe2\xef\xc9\xa5\x8e\xac\x0e\xe1\xa9\xa0\x09\x47\x34\x8d\x1d\xb2\x62\x63\x2a\x90\x07\x0a\x1f\x9c\xb4\xaf\x1e\xd8\xfb\x07\xfd\xd7\xef\xeb\x4d\xa9\xf0\x73\x72\xb6\x53\x66\x76\x99\x30\x9a\x98\x71\x55\x4e\x3f\xde\x8a\x9c\x2e\x37\x1c\x9a\xe8\x8f\x9b\xab\xb8\x8e\x63\x91\x56\x7d\x3f\xc2\xdf\xef\xeb\x9f\xee\xf8\xcf\x77\xdd\x5c\xa4\x2d\x15\xed\x02\x98\x68\xae\x80\x84\x51\xc9\x8d\x90\x9b\xf6\xe5\x64\x10\xc7\xe5\x5b\x62\x6e\x53\x66\x2c\xc4\x3e\x14\x42\x29\xfe\x86\xd2\x3a\x59\xcd\x5e\x43\x35\xbc\x04\xc4\x9f\xdc\xe8\x1e\x3c\x65\xf6\xe9\xcd\xd5\x50\x91\x1d\x60\x1b\xbe\xea\xb5\xa2\x8f\xf0\xc1\x1f\xa5\xc2\xbc\x4b\xa8\xac\x9e\xa0\x06\x18\x79\xae\x3b\x9c\x75\xc1\xab\x27\x78\xaf\x54\x9f\x5c\xf6\xca\x77\xa9\x43\x8f\x26\xae\xd3\x45\xde\x45\xb7\x08\x58\xfc\x42\x80\xef\x05\x51\xd8\xc5\x5c\xc7\xed\x02\xfe\x3c\x18\x23\x65\xae\x1f\xa0\x2a\x7f\x1c\xf8\x08\x39\x33\x14\x46\xc3\x08\x49\xbc\xf6\xbc\x10\x09\xe8\x58\x8c\x91\xa7\x6e\xd2\x8b\x6e\xb1\x8e\x39\x1e\xfa\x63\x38\x9f\x62\x8b\x23\x6f\x81\x3c\x87\xe3\xd0\x99\x45\x67\x60\xe0\x8c\xee\x90\x87\xfb\xf0\x3a\xc4\xbe\xee\x43\x77\x38\x42\x7b\x3d\x0a\xd1\x0c\x7f\x38\x76\x26\x33\xa4\x64\x72\x33\x75\x16\x66\x87\xbe\xb7\x70\x02\x64\x2e\xac\x90\x7f\x01\x00\x00\xff\xff\x04\x18\xb6\x44\xfa\x07\x00\x00")

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

	info := bindataFileInfo{name: "configs/top.yaml", size: 2042, mode: os.FileMode(420), modTime: time.Unix(1442027677, 0)}
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
	"configs/test.yaml": configsTestYaml,
	"configs/top-new.yaml": configsTopNewYaml,
	"configs/top-old.yaml": configsTopOldYaml,
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
		"test.yaml": &bintree{configsTestYaml, map[string]*bintree{
		}},
		"top-new.yaml": &bintree{configsTopNewYaml, map[string]*bintree{
		}},
		"top-old.yaml": &bintree{configsTopOldYaml, map[string]*bintree{
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

