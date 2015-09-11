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

var _configsDfYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x8d\x4d\x6a\xc3\x30\x10\x85\xf7\x3e\xc5\x6c\x02\x09\x45\x74\xd3\xd5\x2c\x9a\x5d\x2f\x51\x17\x33\x91\x26\x74\x88\x15\x09\xcf\xb4\xc5\x35\xbe\x7b\xfd\x07\xa6\x8e\xb5\x9b\xf7\xbe\xf7\xc9\x39\x57\xc4\x14\x18\xe1\x42\xe6\x3f\x8b\x94\x4d\xd2\x5d\xb1\x00\x70\x10\xc7\x08\xe1\x75\x38\xc6\xf7\x26\x35\x6b\xab\xc6\x11\x96\xa4\xeb\xe0\xaa\xd5\xa5\x4e\xfe\xa6\x78\x08\xd0\xf7\x6e\x3e\xa6\x79\xc3\x99\xc9\x10\x5e\x26\x7a\x23\x1b\xa6\x41\xf4\x86\xc7\xe3\x19\x23\x65\x38\x9d\xdf\xcb\x9f\xf2\xb9\x74\x1f\x4f\xa7\xc1\xb3\x52\x96\x8c\xea\x4a\xe5\x97\xe7\x1f\xd6\xe6\x4b\x39\xec\x16\xf4\x4d\xb2\x3f\xf1\x94\xc9\x8b\xb5\x78\xc8\xff\x72\x19\x5d\x5b\x58\xae\x0d\x3f\x18\xb2\xbf\x5b\xb5\xe0\x93\xe3\x2f\x00\x00\xff\xff\xa3\xfb\x38\x1c\x42\x01\x00\x00")

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

	info := bindataFileInfo{name: "configs/df.yaml", size: 322, mode: os.FileMode(420), modTime: time.Unix(1441946541, 0)}
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

var _configsTestYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x54\x5f\x6f\x9b\x3e\x14\x7d\xcf\xa7\xb0\xf4\xfb\x45\x6a\x34\xb2\x69\x9b\xf6\xc2\xc3\x26\x42\x59\x1a\x2d\x04\x06\xa4\x51\xb5\x4e\xc8\x09\x6e\x62\x25\xfc\x11\x36\x89\xba\xaa\xdf\x7d\xb6\x81\x04\x6e\x5c\x69\x3c\xa0\x7b\x8f\xcf\xbd\xf7\x1c\xb0\x3d\x1e\x8f\x07\x69\x9e\x10\x13\xad\x31\xdf\xec\x06\x79\xc1\x69\x9e\x31\x73\x80\xd0\x18\xa5\x12\x32\xd1\x57\x91\xc8\xc7\x2f\xf3\x0d\x61\x8c\x30\x13\x35\xc8\xcb\x0b\x2a\x04\x18\xf3\x9c\xe3\x83\x39\x4c\xd0\xeb\x2b\x52\xb1\x01\x19\x65\x95\x65\x34\xdb\x36\x9c\x26\xbb\x62\x31\x5e\x6d\xf6\x0d\x47\xc5\xd7\x8c\x03\x21\xc5\xa5\x51\x9b\x76\x79\x59\x95\xc6\x7c\x57\x12\x9c\xb0\x56\x52\x9d\xe9\x3c\xcd\x73\x9c\x20\xeb\xb8\xed\x5a\x3a\x08\x2c\xfe\x68\x0e\x9f\x44\xad\x01\xf1\x4f\x6f\xe0\x9f\x6b\x5c\x37\xc3\xf6\x97\xa8\x62\x78\x4b\xba\x43\x36\x45\x15\x2b\x50\xbc\x49\x69\x0e\x0b\x29\x54\x86\x86\x96\xc4\x9e\x59\xc3\x11\x91\xa1\x63\xd0\xe4\x40\x1a\x8a\x0c\x75\x42\xc2\x1d\x2e\x49\x32\xa7\x6b\xf5\x7f\x5b\xf5\x74\x1d\x97\x84\xd1\x84\x64\xdc\x1c\xa6\x24\x55\x3f\xa8\x01\x8c\x3e\x2f\xc1\x1c\x9f\x39\x32\x01\xeb\x07\x9a\xed\x49\x42\x2f\x7d\x5a\xe0\xbd\x4e\x8e\x4b\xd2\x80\x6c\xdb\xed\xd6\xb4\x11\x85\x42\x8e\x42\xdb\x6d\xd5\xf4\xaa\x37\x96\x9e\xf9\x2f\x06\xba\xfc\xa2\xa4\x47\xcc\xc9\x99\xde\xe4\x6f\xb0\x99\xfa\x6e\x67\x72\x9d\x6a\x2d\xf9\xbb\x67\x26\x6c\x75\xfc\x14\x02\x91\xbf\xf8\x52\x2e\x13\x74\xd3\x2e\x9d\x68\xb7\xb5\xca\x46\x06\x2c\xcf\xfa\x0d\x54\xaa\x9d\x7f\xef\x76\x46\x1f\xd3\xf8\xc8\xe8\x9f\x8b\x4b\x95\x19\x3d\xc2\x53\x89\x53\x02\x68\x0a\x3b\xe5\xe5\x5e\x57\xc0\x4e\x58\x1c\x38\x66\xde\x3c\x26\xef\x1e\xd5\x6b\x34\x52\x1f\xa5\xc6\xaf\xc9\x79\xc5\xf5\x6c\xb9\xa0\x35\xb1\x20\x5c\x4e\x3f\xef\x8a\x02\x6f\xf6\x44\x34\x51\x87\x9b\xf0\xb8\xc9\x63\x9a\xd5\x7d\x3f\x88\xd7\xaf\xc9\x0f\x77\xfa\xfb\x9b\x6a\x4e\xb3\x8e\x8a\x6e\x81\x98\xa8\xaf\x10\x0b\xb5\x92\x9d\xb8\x27\xc4\x69\x44\xfe\xec\x56\xb5\x20\x59\x62\x5e\xee\x3d\x05\x95\xf9\x49\x48\xf9\xa2\x62\x8e\xd7\xe2\xd0\x9d\x85\x8b\x71\xff\x8b\xca\xfa\x1a\x68\x01\xdb\x73\x5d\x6b\x01\xc0\xa1\xbc\x13\x7a\x48\x34\x73\x9d\x3e\xf2\x5f\x74\x07\x80\xd5\x4f\x00\xf8\x5e\x10\x85\x7d\xcc\x75\xdc\x3e\xe0\x2f\x83\x29\x50\xe4\xfa\x01\xa8\xf2\xa7\x81\x0f\x90\x2b\x23\x61\x64\x45\x40\xe2\xc4\xf3\x42\x28\x40\x5a\x8b\xa1\x17\x05\x7a\xd1\x1d\x9c\xbb\x84\x43\xbe\x5b\xcb\x39\xec\x68\x7b\x2b\xe0\x31\x9c\x86\xce\x22\xba\x02\x03\xc7\xbe\x07\x9a\x1f\xc2\x49\x08\x7d\x3c\x84\xae\x65\x83\x6f\x6b\x87\x60\x86\x6f\x4d\x9d\xd9\x02\x28\x99\xdd\xce\x1d\xc8\xf3\x56\x4e\x00\x4c\x85\x35\xf2\x37\x00\x00\xff\xff\x92\x80\xcd\x0e\x5c\x07\x00\x00")

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

	info := bindataFileInfo{name: "configs/test.yaml", size: 1884, mode: os.FileMode(420), modTime: time.Unix(1441946246, 0)}
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

	info := bindataFileInfo{name: "configs/top-new.yaml", size: 2037, mode: os.FileMode(420), modTime: time.Unix(1441913405, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _configsTopOldYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x55\x4d\x8f\xda\x30\x10\xbd\xef\xaf\xc8\x05\x89\x94\xc0\x6a\xb7\x6a\x0f\x39\x14\xf5\x43\xea\xa1\xbb\x15\x6a\xd5\x5e\x00\x45\x26\x19\x82\x45\x1c\xbb\xb6\x43\xb4\xa5\xfc\xf7\xda\x8e\x03\x4e\xe2\x54\xdd\x43\xd6\xf3\xde\xf3\x7c\xd8\xc3\x78\x3e\x9f\xdf\x11\x9a\x41\x1c\xec\x90\x4c\x0f\x77\x94\x49\x4c\x4b\x11\xdf\x05\xc1\x3c\xd8\x23\x5c\xc4\xc1\x3b\xb5\xd6\x7f\x2b\x4e\x53\x10\x02\x2c\x49\xb4\xde\xc7\x06\x16\x39\x9f\x03\xa6\xc0\x44\x52\x89\x8a\x78\x92\x05\x97\x4b\x60\xd6\x51\x5f\xc1\xab\xb2\xc4\x65\x6e\x35\xd6\x1a\xa8\x84\xac\xd2\xa3\xd5\x98\xf5\x50\x51\x00\xb0\x9b\xa3\xd6\x74\x75\x65\x45\x12\x79\xe0\x80\x32\xd1\xa6\xd4\x58\xbe\x9a\x9e\x28\xca\x82\xf7\xa7\xdc\x2d\xa9\x50\x58\xf2\x10\x4f\xf6\x6a\x6f\xd4\xc7\x1f\x47\xf0\xd7\x0d\xee\x8b\xf1\x71\xf5\x23\xa8\x04\xca\xc1\x0d\x92\xb2\x2a\x31\xa0\xfa\x02\x8f\x27\x4c\x27\xaa\x97\x91\x57\x24\x5e\x84\xd5\xa8\x55\xe4\x53\xe0\xac\x00\x2b\xd1\x4b\x5f\x22\xdf\x0f\x88\x43\xf6\x84\x77\xe6\x7e\xdb\xec\xf1\x2e\xe1\x20\x70\x06\xa5\x8c\x27\x04\x88\xb9\x20\x0b\x44\x5d\x5d\x86\x24\xba\x6a\xb4\xd1\xe3\x0b\x5c\x1e\x21\xc3\x37\x3f\x2d\xb0\xf0\xa5\xf3\x0c\xe4\x1b\xe4\x6d\x2f\x5a\x37\x6a\xa3\x4a\xc7\xa0\x6d\x5b\x59\x5f\x4d\x63\xf9\x95\xff\x53\x80\xab\x67\x1c\x9f\x90\x84\xab\xdc\xda\x23\x6a\x61\xce\xed\x2a\x6e\x4c\x6f\x49\xab\xc3\x8b\x50\x65\x39\xf5\x30\x85\xe8\x2b\xbe\x6d\xd7\x46\x30\x6d\xa9\x1a\xbb\xae\x8d\x15\x46\xfd\xed\x65\xd7\x81\x31\xbd\xf1\x7f\x3e\x3b\xa1\x4f\x24\x39\x09\xfc\xfb\x56\xa5\xb1\xa2\x8e\x60\xcf\x11\x81\x9e\xcc\x60\x35\xe5\x47\xdf\x06\x51\x23\xf5\x83\x13\xf1\x74\x93\xcd\x36\xe6\x13\x86\xe6\x50\x1a\x7c\x28\xa6\x95\xf4\xab\x35\xe1\x2d\xe2\x2b\x48\x1d\xfd\xda\x15\x0c\xa5\x47\x50\x4e\xcc\x8f\x1b\x64\x62\xed\x04\x97\x8d\xdf\x7b\xf5\x59\x7f\xf8\xf2\xfc\x79\xbb\x34\xce\x71\xe9\x64\xe1\x6e\x50\x11\xfd\x3b\x14\xe1\xcd\xe4\x13\x16\x47\xb7\x39\x33\x65\x27\x7a\x96\xf8\xdd\x68\x26\xea\xa9\x6b\x8e\xa5\x84\x91\x4c\x2d\xd9\xc4\xe6\xc0\x00\xc9\x38\x78\x63\x3c\xf4\x32\xd1\xad\x80\x9b\xb0\xe1\x7a\x36\xdf\x2e\x9b\x59\xd3\x4e\x01\x4a\x08\x2a\x15\x3d\x5d\xc6\x9b\x3a\x5c\x9c\x1f\xa2\x87\xb7\x97\xb0\xa3\x61\xa9\x4c\xd4\xb4\x30\x2e\x16\xda\x4d\x87\x95\x98\x80\xa6\xce\x8f\x17\xed\x63\xf1\x27\x0e\x87\x46\x77\x4b\x33\x67\xe3\xe9\x7a\x93\x6d\xee\xb7\xb3\x21\x59\xff\xfa\x07\xc9\x28\xb7\x7d\xe1\xa9\x47\x35\x62\xdb\x8d\x4e\x01\x15\xcf\x87\x68\x4a\x18\x17\x1e\x71\xce\x59\x33\xfd\x1d\x4c\x9f\x60\x0f\x13\xd2\x8c\x81\xba\x03\xee\x28\x15\x26\xb9\x57\x4b\x7d\x6b\xb3\xe5\x66\x6d\xff\x6f\xbb\x85\xe8\xe9\xab\xce\xad\x9d\xfc\x0e\x4a\xe5\x01\x74\x5e\x5d\xa6\x1a\x26\xb0\x47\x55\xd1\x1e\x84\x8a\xd0\xbb\xd6\xda\x4f\x10\x91\x0b\x3d\xed\xc6\x48\x0e\xe9\xc9\x4f\xaa\xd7\x63\x27\xb2\x51\x8e\x20\xd5\x74\xfe\x5c\xc4\x48\x2e\x4c\x3d\x3e\xed\x40\x18\x90\xfa\x25\x1a\xdb\x47\x6b\xf5\xee\xf9\xbb\xb1\x79\x12\xcd\xa5\xfc\x0d\x00\x00\xff\xff\x32\x6b\x4f\xe7\xbe\x08\x00\x00")

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

	info := bindataFileInfo{name: "configs/top-old.yaml", size: 2238, mode: os.FileMode(420), modTime: time.Unix(1441912961, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _configsTopYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x54\x5f\x6f\x9b\x3e\x14\x7d\xef\xa7\xb0\xf4\xfb\x45\x6a\x35\xb2\x69\x9b\xf6\xc2\xc3\xa6\x34\x65\x69\xb4\x10\x18\x90\x56\xd5\x3a\x21\x27\x76\x13\x2b\xe1\x8f\xb0\x49\xd4\x55\xfd\xee\xb3\x8d\x49\xe0\xc6\x95\xc6\x03\xf2\x3d\xf7\xdc\x7b\xcf\x31\xc6\xc3\xe1\xf0\x22\x2b\x08\x75\xd1\x12\x8b\xd5\xe6\xa2\x28\x05\x2b\x72\xee\x5e\x20\x34\x44\x99\x82\x5c\xf4\x55\x06\xea\x09\xab\x62\x45\x39\xa7\xdc\x45\x06\x79\x79\x41\xa5\x04\x53\x51\x08\xbc\x73\x07\x04\xbd\xbe\x22\xbd\x76\x20\xa3\xaa\xf3\x9c\xe5\x6b\xc3\x31\xd1\x19\x8b\x8b\x7a\xb5\x35\x1c\xbd\x3e\x67\xec\x28\x2d\x4f\x8d\xda\xb0\xcb\xcb\xeb\x2c\x15\x9b\x8a\x62\xc2\x5b\x49\x4d\x64\xf3\x34\x2b\x30\x41\xa3\xfd\xba\x6b\x69\x27\xb1\xf4\xa3\x3b\x78\x92\xb5\x0e\xc4\x3f\xbd\x81\x7f\x6e\x70\xdb\x8c\x71\xb8\x40\x35\xc7\x6b\xda\x1d\xb2\x2a\xeb\x54\x83\xf2\x4d\x2b\x77\x50\x2a\xa1\x6a\xe9\x58\x49\xfc\x99\x1b\x8e\x5c\x39\x36\x06\x23\x3b\x6a\x28\x6a\x69\x13\x12\x6f\x70\x45\xc9\x8c\x2d\xf5\xf7\x6d\xd5\xb3\x65\x5a\x51\xce\x08\xcd\x85\x3b\xc8\x68\xa6\x3f\x90\x01\x9c\x3e\x8f\x60\x81\x8f\x1c\x15\x80\xfc\x8e\xe5\x5b\x4a\xd8\xa9\x4f\x0b\xbc\xb7\xc9\xf1\x69\x16\xd1\x75\x7b\xdc\x4c\x1b\x59\x28\xe5\x68\xb4\x3d\x56\xa6\x57\x73\xb0\xec\xcc\x7f\x31\xd0\xe5\x97\x15\xdb\x63\x41\x8f\x74\x13\xbf\xc1\xe6\x7a\xdf\x8e\xe4\x26\xb4\x5a\x0a\x37\xcf\x5c\xda\xea\xf8\x29\x25\xa2\x3e\xf1\xa9\x5c\x05\xe8\xb2\x4d\x1d\x58\xb7\xb5\x8e\xae\x1c\x58\x9e\xf7\x1b\xe8\xd0\x3a\xff\xce\xef\x8c\xde\x67\xe9\x9e\xb3\x3f\x27\x97\x3a\x72\x7a\x84\xa7\x0a\x67\x14\xd0\x34\x76\x28\xaa\xad\xad\x80\x1f\xb0\xfc\xe1\xb8\x7b\xf9\x48\xde\x3d\xea\xd7\xd5\x95\xde\x94\x06\x3f\x27\x17\xb5\xb0\xb3\x55\xc2\x6a\x62\x4e\x85\x9a\x7e\x3c\x15\x25\x5e\x6d\xa9\x6c\xa2\x7f\x6e\x2a\x52\x13\xa7\x2c\x6f\xfa\x7e\x90\xaf\x5f\xd7\x3f\xfc\xc9\xef\x6f\xba\x39\xcb\x3b\x2a\xba\x05\x72\xa2\xbd\x42\x26\xac\x4a\x6e\x18\xdf\x76\x0f\x27\x91\x71\xaa\xee\x12\x7b\x1b\x95\x71\x00\xfb\x50\x31\x21\xe8\x1b\x4a\x4d\xb2\x99\xbd\x91\xd5\xf2\x26\x40\xe1\xf4\x46\xf7\xa0\x39\x71\x4f\x77\xae\x86\xaa\xe2\x20\xb7\xe1\x8b\x5e\x0b\xbc\x94\x3f\xfc\x51\xaa\x9c\xf7\xbf\xac\x6c\xae\xa0\x16\x18\x07\xbe\x3f\x9a\x03\x70\xa0\xee\xa3\x1e\x92\x4c\x7d\xaf\x8f\xfc\x97\xdc\x02\xe0\xfe\x27\x00\xc2\x20\x4a\xe2\x3e\xe6\x7b\x7e\x1f\x08\x17\xd1\x04\x28\xf2\xc3\x08\x54\x85\x93\x28\x04\xc8\x99\x91\x38\x19\x25\x40\xe2\x75\x10\xc4\x50\x80\xb2\x96\x42\x2f\x1a\x0c\x92\x5b\x38\x77\x01\x87\x7c\x1f\x2d\x66\xb0\xe3\x38\xb8\x07\x1e\xe3\x49\xec\xcd\x93\x33\x30\xf2\xc6\x77\x40\xf3\x43\x7c\x1d\x43\x1f\x0f\xb1\x3f\x1a\x83\xbd\x1d\xc7\x60\x46\x38\x9a\x78\xd3\x39\x50\x32\xbd\x99\x79\x90\x17\xdc\x7b\x11\x30\x15\x37\xc8\xdf\x00\x00\x00\xff\xff\x93\x42\x68\x2a\xd8\x07\x00\x00")

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

	info := bindataFileInfo{name: "configs/top.yaml", size: 2008, mode: os.FileMode(420), modTime: time.Unix(1441946520, 0)}
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

