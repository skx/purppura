// Code generated by go-bindata.
// sources:
// data/index.html
// data/login.html
// data/purppura.js
// DO NOT EDIT!

package main

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

var _dataIndexHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\x7b\x6f\xdb\x38\x12\xff\xbf\x9f\x62\xaa\x14\x67\x1b\xad\x24\xe7\xb1\x6d\x92\x4a\x06\x72\x49\x37\x49\x9b\xb6\x79\xb5\x7b\xdd\xc5\xe2\x40\x89\xb4\x44\x87\x22\x15\x92\x92\xed\x3d\xf4\xbb\x1f\x28\x4a\xb2\xe4\xa4\xb9\x6c\xaf\x58\x2c\x0c\x44\x7c\x0c\x67\x7e\xf3\x9b\xd1\x0c\x95\xe0\xe9\xd1\xc7\xc3\xeb\x2f\xe7\x6f\x20\xd5\x19\x9b\x3c\x09\xcc\x03\x18\xe2\x49\xe8\x10\xee\x4c\x9e\x00\x04\x29\x41\xd8\x0c\x00\x02\x4d\x35\x23\x93\x03\x46\xa4\x56\xf0\xdb\xd8\x1f\xfb\xe3\xdf\x03\xdf\xae\x5a\x89\x8c\x68\x04\x71\x8a\xa4\x22\x3a\x74\x0a\x3d\x75\x77\x9d\xee\x16\x47\x19\x09\x9d\x92\x92\x79\x2e\xa4\x76\x20\x16\x5c\x13\xae\x43\x67\x4e\xb1\x4e\x43\x4c\x4a\x1a\x13\xb7\x9a\xbc\x00\xca\xa9\xa6\x88\xb9\x2a\x46\x8c\x84\x9b\x3d\x45\xa9\xd6\xb9\x4b\x6e\x0b\x5a\x86\xce\xa1\x55\xe2\x5e\x2f\x73\xd2\x51\xa9\xc9\x42\xfb\xc6\xa1\xd7\x2d\xa2\x4f\xd7\x3f\xaf\x00\x3d\x75\x5d\x38\x43\x9a\x28\x0d\xb1\xc8\x72\xca\x08\x06\xc4\x31\x64\x94\xd3\x29\x25\x18\x0e\xaf\xae\xc0\x75\x6b\x69\x46\xf9\x0d\x48\xc2\x42\x47\xe9\x25\x23\x2a\x25\x44\x3b\x90\x4a\x32\x0d\x1d\x03\x46\xed\xfb\x7e\x86\x16\x31\xe6\x5e\x24\x84\x56\x5a\xa2\xdc\x4c\x62\x91\xf9\xed\x82\xbf\xed\x6d\x7b\xaf\xfc\x58\xa9\xd5\x9a\x97\x51\xee\xc5\x4a\x39\x40\xb9\x26\x89\xa4\x7a\x19\x3a\x2a\x45\xdb\xbb\x3b\xee\x3f\x3f\x7f\xa1\xf4\xea\xf4\x67\xf2\x6e\x13\x1f\x67\x6f\x2f\x0f\x6e\x96\x71\x71\x72\x70\x72\x99\x6c\x6f\x7d\xcc\x3e\xc5\xf3\xf9\x2b\xc1\xb7\x2f\xbf\xe0\x64\xe7\x33\x7a\x7e\x9e\x5d\x5d\xab\x3f\xfc\x77\x2f\x77\xcb\x08\xbf\x99\xa5\x3b\x85\x03\xb1\x14\x4a\x09\x49\x13\xca\x43\x07\x71\xc1\x97\x99\x28\x54\x97\x82\xd9\x45\x41\xe4\x12\xa8\x02\x69\xf8\x94\x04\x43\xb4\x84\x16\xde\x8a\x01\x15\x4b\x9a\x6b\x50\x32\x5e\x79\x1c\x0b\x4c\xbc\xd9\xad\xd1\x50\x79\x6a\x87\xee\xa6\xb7\xb9\xe5\xed\x54\x9e\xcd\xee\x75\x8c\x97\x07\x68\xfc\xfc\xe5\x45\xb2\x17\xb3\xf9\x97\xc3\xe3\xe3\xf3\xfc\xe8\xe2\xec\x73\x7e\xf6\x81\x8f\xa7\x97\xe8\xf2\xe3\xec\xe4\x56\x6d\x6e\xeb\x9d\xe3\x64\xb6\xfd\xe6\x8f\x9f\xc6\xff\xe2\xc7\x17\xb7\xb1\x2f\x77\xdf\xa7\xfc\xf2\xe8\xd7\x6f\x3b\x16\xf8\x16\xe7\x63\x83\xfc\x16\x95\xe8\xca\x7a\xf6\xa0\xa7\x8f\x8d\xed\x6c\x3d\xb4\xf7\x13\x70\x1d\xff\x74\x7a\x41\xa3\xf1\xd6\xab\xdb\x72\x39\xbb\x7a\x3f\x3d\x99\x7d\x7c\x8f\xce\x6e\xa6\xc5\x2f\x9f\x17\xbf\x2e\x3e\x9d\xf3\xc3\xb7\x07\xaf\xd8\x56\x76\xf8\xcb\x87\xd3\xfc\x78\x2f\x3b\x3e\x3c\xda\x9d\x1f\x7f\x38\x8d\xcf\x8f\x5e\x5d\x2f\xd0\xa3\x09\xa8\x7d\xd1\xcb\x9c\xd4\xaf\xc4\x0c\x95\xc8\xae\x3a\xd6\x45\x3f\x2f\x64\x9e\x17\x12\x19\xac\x77\xce\x9b\x6c\xef\x1e\x37\xa9\x6a\xf7\xa0\xaa\x1a\xf0\x9f\x7a\x12\xa1\xf8\x26\x91\xa2\xe0\xd8\x8d\x05\x13\x72\x7f\x63\xba\x67\x7e\xaf\xeb\xfd\xaf\xf5\xd3\x53\x9a\x94\xc4\x9d\x0a\xa1\x89\x6c\x4f\xe7\x08\x63\xca\x93\x7d\xd8\x19\xe7\x0b\x18\x37\x87\xac\x26\xd8\xd8\xdb\x6b\xf5\xdc\xb1\x03\x6b\x86\x22\x21\x31\x91\xae\x16\xf9\x3e\x6c\xe5\x0b\x50\x82\x51\x0c\x1b\x64\xcf\xfc\x1e\x46\xe3\x71\x54\xba\xa6\xe2\x75\x90\xd5\x46\x22\x86\xe2\x9b\xe6\xf4\x54\x70\xed\xce\x09\x4d\x52\xbd\x0f\x91\x60\xb8\xaf\x36\xf0\x2b\xda\x1e\x13\x82\x86\xca\x39\xe5\x58\xcc\x3d\xc1\x99\x40\x18\x42\x98\x16\x3c\xd6\x54\xf0\xe1\xa8\xc1\x41\xa7\x30\x64\x22\x46\x66\xd5\x4b\x91\x4a\xe1\x69\x18\xc2\x60\x30\x82\x67\xc3\x01\xfa\xcd\x96\xa2\x01\x3c\x87\xbe\xd0\x73\x18\x38\xbf\x0f\x46\x9e\x46\xd1\x70\xa0\x52\x31\x1f\x8c\x1a\xb0\xd5\x39\x8c\x34\x72\xb5\x48\x12\x66\xd0\xa1\xa8\x12\x8e\x28\xc7\xc3\x41\xcc\x68\x7c\x33\x78\xb1\xc2\x42\x46\x2d\x2b\x7d\x23\x21\x3c\x1b\x12\x4f\x23\x99\x10\x3d\xf2\x90\xd6\x72\x38\x30\x80\x06\x23\x4f\x15\x91\xd2\x72\xb8\xd9\x5a\xfd\x3a\x7a\xfd\xa4\x1e\xfa\x3e\x54\x36\x28\x4f\x40\x70\xd0\x29\x01\x55\x44\x33\x12\x6b\x98\x53\xc6\xc0\xc2\xaa\xd6\x31\xd1\x88\x32\xe5\xb5\xd0\xc1\x80\x65\xc4\x81\x91\x87\x09\x23\x09\xd2\xc4\xac\x49\xf0\x2a\x95\xce\x0b\x00\xa7\x19\xb5\x0e\x00\xac\x3c\x78\x36\xd4\x29\x55\x23\x2f\x66\x42\x11\xa5\x87\x03\x2d\x07\x23\x8f\x93\x45\x33\xb4\xd6\x87\xdf\x00\xfe\x29\xc7\x48\x13\x10\x85\x04\x64\x3b\x22\x29\x4d\x21\xdd\x1e\x83\x22\xb1\xe0\xb8\xc5\x5a\x54\x92\xff\xb6\x52\x2b\x75\x8a\xe8\x53\xae\x89\x2c\x11\x1b\x76\xa2\x7d\x47\x1c\xbe\xbe\x80\xed\xf1\x78\x3c\x1e\xdd\x49\xb2\xf6\x45\x0d\xfc\xa6\x4d\x07\x91\xc0\xcb\x3a\xf1\x38\x2a\x21\x66\x48\xa9\xd0\xe1\xa8\x8c\x90\x04\xfb\x70\x31\x99\xa2\x82\xad\x92\x2f\xc0\xb4\x95\x34\xfd\x13\x51\x4e\xa4\x3b\x65\x05\xc5\xad\x4c\x5f\xaa\x56\x64\x5f\x95\x8e\x8c\x01\x50\x68\x6d\x82\x59\xa5\xbb\x9d\x38\x6b\xc7\xea\xb8\xc6\x82\x31\x94\x2b\x82\x1d\xe8\x65\x61\xb3\xde\x2c\x57\x69\x15\x3a\x1b\xf6\xb4\x03\x48\x52\xe4\x92\x45\x8e\x38\x26\x38\x74\xa6\x88\x19\xd9\x6a\xd5\xa0\x97\x82\xb5\xa6\x7a\xd0\xcc\xbb\x98\x23\xde\x80\x51\xd2\x15\x9c\x2d\x9d\xc9\xb5\x85\xc3\x51\x49\x93\x2a\xa9\x03\xdf\xc8\x3d\x70\x94\xc6\x82\xbb\x95\xfa\xbf\x4a\x34\xf0\x2d\x95\xbd\x35\xb4\xc6\x6b\x24\x11\xc7\xcd\xad\xc4\x77\x26\xe7\x75\x61\x0f\x7c\xd4\x09\xa3\x8f\x69\xb9\x16\x55\x8a\x5b\xc2\xd6\x54\x36\xb1\x68\x83\xd5\x0f\x76\xc1\x3a\xf2\x4d\x7a\x71\x54\xae\xf3\xce\xe8\x24\x40\x2b\x60\x27\x22\x23\x06\x54\xe0\x33\xda\xf7\xb2\x60\x8f\xd2\xdf\x0c\xa5\x29\xc1\x77\x8d\x41\xc7\x1a\x13\x89\x28\xb4\x33\x39\xab\x9e\xff\xdb\x6a\x8f\x9f\xce\x24\xf0\x39\x2a\x1f\x7c\x63\xba\xef\x4a\xba\x55\x5f\x94\x03\x3f\xdd\xea\xac\xdf\x71\xc8\xd5\x28\x32\x57\x04\x1c\x3a\xd9\xf2\x1a\x45\x7d\x7e\x19\x6d\xc4\x51\xac\x69\x49\x1c\xe3\xd8\x9d\x92\x5d\xbb\xba\x21\x11\x35\x6f\xd3\xe4\xb2\x7a\xde\xeb\xaa\x0d\xc4\x37\x35\xa0\xf8\x86\x8b\x39\x23\x38\x31\x7a\x0e\x3a\xb3\xef\xd1\x96\x13\x6e\x9a\xba\x33\x39\xb7\x83\xbb\x3a\x2c\xf5\xf7\x96\x18\x8d\x22\xb7\xbe\xcc\xf7\x39\x69\x12\xb6\xf6\xb6\x2b\x9f\x23\x4e\x60\x8a\x30\x01\xca\xa1\x61\xac\x9f\x1c\xf9\xe4\x1f\x3c\x52\xf9\xeb\xc0\xcf\xd7\x76\xaa\x7e\xd2\xd1\x5c\x97\xe0\xae\x01\xd3\x8c\xcc\x5f\x57\x69\x49\x73\x82\xeb\x59\x2a\x4a\x22\xeb\xb1\xbd\x7d\xb4\x5b\xa6\x13\x10\x5e\xe1\xec\x9b\x33\x06\xe5\x24\xd0\xe9\x64\x23\xf0\x75\x5a\x8d\xae\x44\x21\x63\xb2\x9a\xda\x5e\xd8\xce\xcf\x90\xd2\xf0\x41\x68\x7b\x73\x3d\x58\x6d\x1c\x54\xcd\x43\xd9\xb9\xaf\xe5\x9a\x63\x7e\x05\xa5\x9f\xf1\xbd\x1a\xd0\x21\xb5\x97\x00\xf7\x52\xfb\x3d\x84\x76\xb5\xfe\xed\x68\xfd\x81\xec\x35\x09\xff\xc3\x88\xab\x15\xfe\xed\x38\xab\x2a\x0c\x5c\xd3\x8c\xfc\x30\x1a\xef\xab\xbb\x4f\xee\xa1\xea\x7b\x26\xb6\x80\xd7\x97\x7c\xc3\xab\x1d\xb6\x7c\xd6\x3b\xdd\x6f\x01\x07\xa4\x30\xf5\x6c\x75\x2b\x78\xf8\xba\xf4\x8d\x8b\x52\x2c\x98\x9b\x61\xf7\xe5\x03\x6d\xf3\x9e\xf6\xb5\xda\x6c\xaf\x58\x36\x20\x70\x28\x30\x59\x2f\xc3\x6b\xfd\xb5\xf9\x64\x4d\xa8\x4e\x8b\xa8\xfa\x4a\x55\x37\x8b\xf6\x2b\xcf\x99\x1c\x94\x88\xb2\x2a\x83\x04\x07\x2b\xf5\x27\xdb\xe2\x9f\xf6\x11\x1e\xe7\xe4\xa1\xc8\x97\x55\x4b\x7f\xa4\x8b\x55\xc0\xbc\x1b\x92\xe5\xde\x94\xfa\xce\xe4\xca\xcc\xe1\x1d\xc9\xf2\xff\xa7\xcf\xdb\x04\xb0\xf7\x6a\x7b\x9d\x0e\x7c\xfb\x0f\xb2\xff\x06\x00\x00\xff\xff\x86\x12\x93\x4a\x31\x13\x00\x00"

func dataIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_dataIndexHtml,
		"data/index.html",
	)
}

func dataIndexHtml() (*asset, error) {
	bytes, err := dataIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/index.html", size: 4913, mode: os.FileMode(420), modTime: time.Unix(1522422914, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataLoginHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\x7b\x53\x1b\x39\x12\xff\x7f\x3f\x45\xaf\xb2\x75\x0b\xb5\x37\x1e\xc0\x24\x80\x99\x71\x15\x67\xb2\x90\x84\x84\x87\x09\x7b\xec\x7f\x9a\x91\x66\x2c\xa3\x91\x84\xa4\x19\xec\x6c\xe5\xbb\x5f\x69\x5e\xb6\xc7\x86\x23\x77\x55\x57\x57\x54\xd9\x6a\xa9\xd5\xea\xfe\xf5\xd3\x04\x3f\x9f\x5e\x8e\x6e\xef\xaf\xde\xc3\xc4\x66\x7c\xf8\x53\xe0\xbe\x80\x63\x91\x86\x88\x0a\x34\xfc\x09\x20\x98\x50\x4c\xdc\x02\x20\xb0\xcc\x72\x3a\xbc\x90\x29\x13\x70\x43\x1f\x73\xa6\x29\x09\xfc\x6a\xb7\xe2\xc8\xa8\xc5\x20\x70\x46\x43\x54\x30\xfa\xa4\xa4\xb6\x08\x62\x29\x2c\x15\x36\x44\x4f\x8c\xd8\x49\x48\x68\xc1\x62\xea\x95\xc4\xdf\x81\x09\x66\x19\xe6\x9e\x89\x31\xa7\xe1\x2e\x5a\x16\x34\xb1\x56\x79\xee\x9d\x22\x44\xa3\x4a\x88\x77\x3b\x57\x74\x49\xa4\xa5\x33\xeb\x3b\xad\x8f\x21\x9e\x60\x6d\xa8\x0d\xbf\xde\xfe\xee\x1d\x36\x72\x7e\xf6\x3c\xb8\xc0\x96\x1a\x0b\xb1\xcc\x14\xe3\x94\x00\x16\x04\x32\x26\x58\xc2\x28\x81\xd1\x78\x0c\x9e\x57\x73\x73\x26\x1e\x40\x53\x1e\x22\x63\xe7\x9c\x9a\x09\xa5\x16\xc1\x44\xd3\x24\x44\x4e\x19\x33\xf0\xfd\x0c\xcf\x62\x22\x7a\x91\x94\xd6\x58\x8d\x95\x23\x62\x99\xf9\xed\x86\xdf\xef\xf5\x7b\x07\x7e\x6c\xcc\x62\xaf\x97\x31\xd1\x8b\x8d\x41\xc0\x84\xa5\xa9\x66\x76\x1e\x22\x33\xc1\xfd\xc3\x7d\xef\x1f\x77\xf7\x8c\x8d\x3f\xfc\x4e\x3f\xed\x92\xb3\xec\xe3\xcd\xc9\xc3\x3c\xce\xcf\x4f\xce\x6f\xd2\xfe\xde\x65\xf6\x35\x7e\x7a\x3a\x90\xa2\x7f\x73\x4f\xd2\xfd\x3b\xfc\xdb\x55\x36\xbe\x35\xdf\xfc\x4f\xef\x0e\x8b\x88\xbc\x9f\x4e\xf6\x73\x04\xb1\x96\xc6\x48\xcd\x52\x26\x42\x84\x85\x14\xf3\x4c\xe6\x66\x19\x82\xe9\x75\x4e\xf5\x1c\x98\x01\x5d\xfb\x0d\xa2\x39\xb4\xea\x2d\x10\x30\xb1\x66\xca\x82\xd1\xf1\xc2\xe2\x58\x12\xda\x9b\x3e\x3a\x09\xa5\xa5\xd5\xd2\xdb\xed\xed\xee\xf5\xf6\x4b\xcb\xa6\x1b\x0d\x13\xc5\x09\xde\xf9\xed\xdd\x75\x7a\x14\xf3\xa7\xfb\xd1\xd9\xd9\x95\x3a\xbd\xbe\xb8\x53\x17\x5f\xc4\x4e\x72\x83\x6f\x2e\xa7\xe7\x8f\x66\xb7\x6f\xf7\xcf\xd2\x69\xff\xfd\xb7\xb7\x3b\xff\x14\x67\xd7\x8f\xb1\xaf\x0f\x3f\x4f\xc4\xcd\xe9\x9f\xcf\x1b\x16\xf8\x95\x9e\xaf\x75\xf2\x47\x5c\xe0\x71\x65\xd9\x8b\x96\xbe\xd6\xb7\xd3\xae\x6b\x37\x03\x70\x1b\xbf\xfd\x70\xcd\xa2\x9d\xbd\x83\xc7\x62\x3e\x1d\x7f\x4e\xce\xa7\x97\x9f\xf1\xc5\x43\x92\xff\x71\x37\xfb\x73\xf6\xf5\x4a\x8c\x3e\x9e\x1c\xf0\xbd\x6c\xf4\xc7\x97\x0f\xea\xec\x28\x3b\x1b\x9d\x1e\x3e\x9d\x7d\xf9\x10\x5f\x9d\x1e\xdc\xce\xf0\x6b\x00\xa8\x8d\x71\xe1\x0a\x76\xae\x68\x9d\x12\x2e\xd6\x2a\x43\xa1\xcc\x6d\xf8\xab\x26\x22\x1c\x3f\xa4\x5a\xe6\x82\x78\xb1\xe4\x52\x0f\xde\x24\x47\xee\xef\xb8\x3e\xff\x5e\x7f\xf7\x8c\xa5\x05\xf5\x12\x29\x2d\xd5\xed\x6d\x85\x09\x61\x22\x1d\xc0\xfe\x8e\x9a\xc1\x4e\x73\xa9\x92\x04\x6f\x8e\x8e\x5a\x39\x6b\xef\x40\xe7\xa1\x48\x6a\x42\xb5\x67\xa5\x1a\xc0\x9e\x9a\x81\x91\x9c\x11\x78\x43\x8f\xdc\xdf\xcb\xda\xf4\x04\x2e\x3c\x57\x97\x96\x34\xab\x1f\x89\x38\x8e\x1f\x9a\xdb\x89\x14\xd6\x7b\xa2\x2c\x9d\xd8\x01\x44\x92\x93\x55\xb1\x81\x5f\xc2\xb6\x1a\x0f\x4b\x18\x4e\x71\x81\xab\xdd\x1a\xca\x5f\xb6\x80\xc8\x38\xcf\xa8\xb0\xb0\xdd\xd3\x14\x93\xf9\x56\x92\x8b\xd8\x32\x29\xb6\xb6\xe1\x2f\xf8\x65\xeb\x57\x26\x54\x6e\x07\xee\xfe\xa0\x60\x86\x45\x9c\x0e\x12\xa6\x8d\xfd\x75\xbb\x97\xc8\x38\x37\x5b\xdb\xc7\xf0\x7d\xbb\x79\xbf\x8d\xe2\xc0\x6f\xea\x6c\x10\x49\x32\x6f\x5c\x17\x08\x5c\x40\xcc\xb1\x31\x21\x12\xb8\x88\xb0\x86\xea\xcb\x23\x34\xc1\x39\xb7\xad\x93\x01\x02\xc2\x5a\x5e\x57\x1f\x31\x13\x54\x7b\x09\xcf\x19\x59\xe2\x5a\xe5\xab\x85\x55\x58\xae\x70\x39\x45\x72\x6b\xa5\xa8\x11\xa9\x08\xd4\xb9\x68\x65\x9a\x72\xea\xc0\xe7\x58\x19\x4a\x10\x10\x6c\x71\xbd\xed\xd4\xa8\xf6\x9b\x6d\xac\x53\x6a\x43\xf4\xa6\xba\x8d\x00\x6b\x86\x3d\x3a\x53\x58\x10\x4a\x42\x94\x60\xee\x78\xcb\x5d\x67\x81\x96\xbc\x7d\xaa\xa3\x9c\x73\x98\xc2\xa2\x51\xc7\x68\x4f\x0a\x3e\x47\xc3\xdb\x4a\x21\x81\x0b\x96\x62\xe7\x97\xc0\x77\x7c\x2f\x5e\x66\xb1\x14\x5e\xf9\xc4\xff\x96\x39\xf0\x2b\x50\x3b\xbb\xb8\x83\x71\xa4\xb1\x20\x4d\x17\xf2\xd1\xf0\x2a\xd7\x4a\xe5\x1a\x07\x3e\x5e\x71\xab\x4f\x58\xb1\xe6\x67\x46\x5a\x00\x3b\x62\x1b\xdf\xb4\xce\xeb\xba\x3f\xe7\x4b\x37\x9a\xb0\x13\xb8\x58\xf7\x04\x67\xc3\x00\x2f\x14\x3c\x97\x19\x75\xca\x05\x3e\x67\x5d\x8b\x73\xfe\x82\xce\x2b\x64\xe0\x0b\x5c\xd4\xf5\xed\x99\xe0\xee\x06\xbf\x33\x96\xbb\xe9\x24\x92\x33\x04\x65\x72\x87\x28\xc3\x3a\x65\xa2\xac\x33\x6f\x77\xd4\xec\xb8\xc5\x21\xc3\x25\xa3\xb3\xdf\xcb\x88\xf7\xae\x59\xc8\x24\x31\xd4\x7a\xfd\x92\x36\x99\x77\xd8\x2c\xea\x83\xbd\x67\xb3\x49\x61\x41\x39\x94\x9f\x1e\x13\x89\x44\xd0\x31\xbf\xcb\x5b\x26\x1e\x13\xe9\x3a\xa4\x6b\x9c\xe5\x94\x85\xaa\xe1\x6b\xcd\xd5\x1b\xbc\x5f\xcb\xa8\x41\xa8\xeb\x76\x89\x42\x7f\x47\xcd\xd0\xaa\x70\x57\x72\xba\xca\xae\x0a\x20\xcc\x28\x8e\xe7\x03\x21\x05\x45\x0b\xa0\x3d\xcc\x69\x39\xe2\x55\xd2\x4a\x0a\xca\x4f\x8f\x60\x91\x52\xdd\x60\xb7\xbb\xe7\x72\x60\x4d\x47\x80\x20\x91\x3a\x5b\x08\x74\x54\x2b\xce\x11\xde\x44\x6a\xf6\xcd\x79\x9c\x23\xd0\xd2\xe9\x52\xf1\xe0\xb2\xec\x86\xc8\x2f\xef\x21\xc8\xa8\x9d\x48\x12\xa2\xab\xcb\xf1\xed\x1a\x9e\xab\xc6\xd4\x21\x11\x49\x6b\x65\x36\x80\xbd\xb7\x4b\x80\x94\xf5\xdb\x73\x9d\x4b\x6d\x90\xd2\xcd\xef\x05\xb3\x87\x09\x91\x02\x0d\x03\xd6\x1c\xa6\x7c\xae\x26\xae\x02\x40\xbb\xf2\x72\x43\xcb\x62\xc0\x9e\xa9\x1e\xe5\x0b\xa5\xd4\x3a\x73\x33\x8a\x96\x5a\xd2\x2a\x32\x75\x89\x44\xf5\xc0\x5d\x31\x17\x98\xe7\x34\x44\x08\x14\xc7\x31\x9d\x48\x4e\xa8\x0e\xd1\xbd\xcc\x35\xb8\xc7\x1d\x13\x6c\x82\x67\x93\x6f\xfe\x5f\x50\xe3\x32\x7e\xf8\x01\xd4\x14\x36\xe6\x49\x6a\xd2\x20\xb7\xa0\x5f\x40\x6f\xc1\xb4\x0e\x5c\x7b\xf6\xdf\xe0\xe6\x52\x6f\x77\x39\xf5\x4a\x2d\x4a\x14\x40\xe5\x9c\x7b\xda\xcd\x29\x9b\xd1\x5b\x29\x7e\x75\x3e\x41\xd3\x21\x37\x5e\x69\x01\xa9\x20\x30\x79\x94\x31\x5b\x25\x6e\xb3\xae\xec\x6e\xa8\x3a\x6e\x2e\xaa\x64\xaa\x1f\x8b\xac\x80\xc8\x8a\xaa\x98\x6d\xd4\xec\x19\xf3\x37\x66\xba\xef\x2c\xfe\x37\x65\xeb\xe5\x8e\x50\x12\x0d\xa5\x86\x7f\x13\x91\x51\xc7\x81\xaf\x86\xff\x09\x51\xcd\x60\xf5\x48\xe9\x80\xa9\x96\x4b\xfe\x29\x4f\x96\x27\xcf\xa6\x02\x2d\xc6\x8b\x16\x95\x57\x35\xa8\x25\x1f\xba\x8e\xb3\xda\x48\x56\xfa\x6d\xb7\x19\x73\xb6\x74\xd8\xce\x6b\x63\x99\xeb\x98\xc2\x48\x12\xba\xa1\xd9\x2e\xb7\xe5\xe6\x17\x4e\xca\xec\x24\x8f\xca\x1f\x35\xe6\x61\xe6\xab\x7a\x96\x40\xc3\x93\x02\x33\x8e\x23\x4e\xc1\xa5\x5e\xc9\xb5\xa9\x87\xaf\x76\xf0\xae\xb7\x7e\xcc\xc6\x6e\x1d\x7a\xc6\xc8\x91\x54\xf3\x32\x37\x5e\x69\x62\xe9\xb0\xde\x03\xcd\x54\x2f\x61\x3e\x1a\x8e\x1d\x0d\x9f\x68\xa6\x7e\xd0\xa0\x25\xc2\xc5\xae\x0b\x80\x6a\x54\xaf\x26\xf4\xc0\xaf\xfe\x69\xf2\xaf\x00\x00\x00\xff\xff\xe8\xbf\x39\xf6\x45\x11\x00\x00"

func dataLoginHtmlBytes() ([]byte, error) {
	return bindataRead(
		_dataLoginHtml,
		"data/login.html",
	)
}

func dataLoginHtml() (*asset, error) {
	bytes, err := dataLoginHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/login.html", size: 4421, mode: os.FileMode(420), modTime: time.Unix(1522422867, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataPurppuraJs = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x57\x5d\x6b\xe3\x46\x14\x7d\xcf\xaf\x38\x4c\x17\x2c\xd5\x89\xe4\x85\x3e\xad\x2d\x43\xda\x50\x68\xa1\x5b\x68\xfa\x16\x9b\x32\xd1\x5c\x5b\x53\x2b\x1a\x31\x73\x15\xd7\x6c\xfc\xdf\xcb\x8c\x6c\x47\x8e\xe2\xb0\xd9\xd2\x76\x1f\xa2\x97\x38\xf7\xfb\xe3\x9c\x19\x69\xd1\x54\x39\x6b\x53\xa1\xa9\x95\x64\xfa\x43\x96\x64\xd9\x45\xf1\xd9\xa7\x33\x00\x78\x97\x2c\x89\x7f\xbe\xfe\xf5\x63\x04\x91\xd2\x3d\x55\xec\x04\xce\xb1\xf7\x8a\xa0\x24\x4b\xc4\x68\xad\xfd\x73\x2f\x2d\x0a\x64\xf8\xb4\x1d\x1f\x64\xc5\xcd\xc0\x4a\xed\x48\x0d\xe6\xc8\x30\xea\xca\x6b\xaa\x94\xae\x96\x7d\x85\xcc\x57\x95\x59\x97\xa4\x96\x07\xb7\x83\x3a\x4d\xf1\x43\x49\xd2\xc2\x34\x16\x2c\x6f\x4b\x82\x59\xa0\x96\x8e\xd1\xd6\x88\x0b\xdc\x6e\x60\xe9\xce\xdc\xeb\x6a\x09\x59\x96\xb0\x66\xed\x60\x16\xdd\x18\x5c\x50\xeb\x7d\x0e\xfa\x2b\xa7\x9a\xb1\x30\x36\x48\x17\xda\x3a\x46\xb4\x2e\x74\x5e\x40\xbb\x20\x2b\x48\x2a\xb2\x17\xd6\xac\xe3\xe4\x10\xc5\x3b\x44\xbe\xe5\x15\x6d\xa0\x2b\x14\xdd\x51\x84\x01\x46\xe2\x1b\x81\x61\xd0\x0f\x21\x76\xf3\x15\x71\xb2\xd0\x95\x8a\x04\xdb\x0f\x4b\x8e\x46\xb1\x88\x93\x50\x2d\x45\xf1\xe3\xd8\xb6\x8f\x1d\xbf\x4b\x48\xe6\x45\x3b\xee\xee\xf8\x57\xb4\x39\xc7\xbd\x2c\xc3\x0a\x8e\x12\xa7\x29\xbe\x6f\xee\xea\x50\x7a\x6e\x9a\x8a\xfd\x88\xb8\xf0\xcd\x6c\xea\x30\xaf\xb6\x94\xf3\x43\xd3\xac\xb9\xa4\xe4\x28\x88\xef\xcc\xb1\xe4\xc6\x21\xf3\x69\x6e\x06\xd7\xe1\xbf\xc1\x7c\x7c\x64\x57\xdc\xb4\x56\x7e\x4f\x87\xdf\xc3\xf7\xbd\x8a\x7e\xf7\xf9\xb5\x43\x61\xac\xd5\x0a\x17\x58\x13\xd6\xb2\x62\xb0\x81\x2b\xcc\x1a\xa4\xb9\x20\x0b\xb1\xd6\x7e\x65\x1e\x33\x90\x2c\xce\x9f\xc6\x11\xa5\xdf\x75\x65\x58\x2f\x34\xa9\x60\x02\x63\xbd\xa0\xf0\xfb\x56\xb4\x83\x15\x4c\xd5\xb6\xb6\xa9\x9f\x74\xa6\x17\x88\x0e\xad\x65\x10\x2d\x3e\x45\x7f\x8e\xfb\x31\x28\x64\xf8\x45\x72\x91\x58\xd3\x54\x2a\x6a\xa7\xf1\x71\x57\xc1\x25\x0f\xe6\xf8\x16\xef\x47\xa3\x51\xdc\xf3\xf6\x9e\xb5\xb4\x8e\x7e\x2c\x8d\xe4\x48\xc5\xe3\x67\x4d\x2a\x5a\xe3\x4a\x32\x45\x50\x40\x3c\xee\x57\x91\xa6\xd0\x0a\x0f\x70\xa6\xb1\x39\xf9\x1f\xcd\xed\x9f\x94\x33\x1e\x70\x3c\x8e\x07\xc8\x00\x8f\x67\xfb\x60\x64\x10\x13\xb6\xd3\x09\xab\xa9\x47\x66\x68\xe4\xa7\xab\xc1\xdc\xe3\x73\x92\x7a\x69\xbf\x40\xc6\x30\xb8\x75\x5c\xae\x43\x19\x9f\xe9\x86\xbc\x94\xce\x65\x33\x91\x97\x3a\x5f\xcd\xc4\x54\xe0\x10\xa7\xed\xe2\x35\xf9\x15\x3e\xcf\x76\x22\x51\x58\x5a\x64\x33\x91\x76\x8e\x93\xd4\x87\x88\x0e\x6d\xc7\x3e\xd6\x4c\x4c\x65\xbe\x9a\xa4\x72\x2a\x4e\xc4\x43\x27\x58\xee\x0f\x9f\x13\x61\x82\xee\xa5\x40\x6d\xd9\xcf\xae\xf7\x52\xa9\x00\xd7\xc0\xcc\x9e\xc1\xfe\x28\x39\xe2\xe1\x33\x87\xca\x80\x6f\x8d\xda\x0c\xe2\x44\xd6\x9e\x06\x11\xc7\x2f\x26\x53\xc4\x52\x97\x2e\xf9\xc7\xf9\x7a\x01\xfc\xb3\x2f\xc2\x43\x0e\x8e\x37\x25\x65\x33\xa1\xb4\xab\x4b\xb9\xf9\x50\x99\x8a\xc6\x33\x11\xa0\x18\xc6\x12\x90\x62\x4a\x57\xcb\x2a\x9b\x89\xef\xbc\xaa\x7e\x44\xdc\x55\xa8\x74\x0f\x94\x7a\xe7\x92\xb2\x9d\x8a\xe3\xe4\xdb\x17\xa9\xde\xbd\x59\x4e\x10\xfe\x05\xaa\xb5\xd4\x72\x5f\x37\xb7\xbe\x8c\x5a\x1d\x84\x87\xe3\xf0\x04\xc2\x83\xce\x23\xfc\xd5\x94\x78\x83\xfe\xbf\x02\xfd\x57\x60\x7f\x77\x2f\xbe\xf6\x9e\xfb\xcd\xaf\xfc\xed\x92\xfb\xaf\x88\xf8\x85\x77\xdc\x1b\x07\x7b\xcf\xff\x74\xfd\xf4\x5e\x7c\xaf\x89\x4f\xbd\x63\x2b\x93\x37\x77\x54\x71\x12\x74\x1e\xb8\x97\xa1\x3f\xdc\xf8\xcc\xdd\x4f\xa7\x21\x44\xba\x93\x3d\xf9\x3a\xea\x68\x1e\x3f\xa8\x86\x10\x73\xd1\xa1\xd8\x76\xc7\x49\xff\x77\x7b\xf6\x77\x00\x00\x00\xff\xff\x5b\x55\x3c\x14\xf9\x0d\x00\x00"

func dataPurppuraJsBytes() ([]byte, error) {
	return bindataRead(
		_dataPurppuraJs,
		"data/purppura.js",
	)
}

func dataPurppuraJs() (*asset, error) {
	bytes, err := dataPurppuraJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/purppura.js", size: 3577, mode: os.FileMode(420), modTime: time.Unix(1522242333, 0)}
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
	"data/index.html":  dataIndexHtml,
	"data/login.html":  dataLoginHtml,
	"data/purppura.js": dataPurppuraJs,
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
	"data": &bintree{nil, map[string]*bintree{
		"index.html":  &bintree{dataIndexHtml, map[string]*bintree{}},
		"login.html":  &bintree{dataLoginHtml, map[string]*bintree{}},
		"purppura.js": &bintree{dataPurppuraJs, map[string]*bintree{}},
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
