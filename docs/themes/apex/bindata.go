// Code generated by go-bindata.
// sources:
// files/css/index.css
// files/js/index.js
// files/js/smoothscroll.js
// files/views/index.html
// DO NOT EDIT!

package apex

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

var _cssIndexCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x58\xdd\x6e\xdb\xb8\x12\xbe\xcf\x53\x10\x0e\x0a\xc4\xad\xa5\x4a\x4e\xec\xba\x0a\x50\x9c\x73\xda\x04\x3d\x17\xdd\x2d\x36\x5b\xec\xde\x52\x12\x6d\x73\x43\x89\x02\x45\xe5\xa7\x8b\xbe\xfb\x0e\x39\x94\x4c\x5a\x72\xbb\x45\x50\x58\x1c\x72\xfe\xbf\x99\x21\x9b\x29\x29\x35\xf9\xfb\x8c\x90\x28\x62\xb4\x65\x19\x29\xba\x9c\x17\x51\xce\xbe\x72\xa6\x2e\xe2\xcd\x72\x41\x92\x05\x89\x53\xf8\x4d\xe7\xd7\xf6\xdc\x23\x2f\xf5\x3e\x23\x9b\x24\x69\x9e\x90\xb2\x67\xb4\x64\x0a\x7e\xf8\x6e\xaf\x33\x72\x85\x3b\x76\x4b\x2b\x5a\xdc\xf3\x7a\x97\x91\x24\x4e\x56\x8a\x55\xd7\x01\x39\xaa\x58\xc9\xbb\xca\xec\x4e\x6c\x0a\xaa\x76\xcc\xec\x6d\xec\x1e\x6c\xbe\x7e\x19\x45\x25\x55\xf7\x19\x39\x5f\xa6\xcb\xd5\xf2\xed\xf5\xcb\xd7\x96\xc7\x11\x93\x24\x41\x19\xb9\xe8\x80\xf3\xfc\xf2\xf2\xbf\x6f\x6f\x6f\x91\x24\x8c\x75\xd1\x4e\xd1\x67\xd8\xd8\x52\xf3\xe7\x8c\xcc\x77\x86\xb2\xdd\xe2\xb9\xad\x59\x6d\xd6\x9b\x9b\xb7\xeb\x9e\x70\xa4\x13\xb9\x5a\x26\x58\xa1\xb9\xac\x2d\xff\x03\x55\x17\xa8\xd6\x85\xe9\xb0\x6d\x04\x3e\xee\xb9\x66\xd7\x67\xdf\xce\xce\x5e\xda\x70\xe7\xf2\x29\x6a\xf9\x57\x1b\x99\x5c\x2a\x13\x3e\x20\xd9\x03\x7b\x5d\x89\x05\x10\xcb\x67\x7b\x72\x2b\x6b\x1d\x6d\x69\xc5\x05\xd8\x1d\xd1\xa6\x11\x2c\x6a\x9f\x5b\xcd\xaa\x05\xf9\x9f\xe0\xf5\xfd\x27\x5a\xdc\xd9\xf5\x2d\x9c\x5c\x00\x07\x21\xb3\x3b\xb6\x93\x8c\x7c\xf9\xff\x6c\x41\x66\xbf\xc9\x5c\x6a\x69\xbe\x7e\x7d\x7a\xde\xb1\x7a\xe6\xce\x7c\xc9\xbb\x5a\x77\x86\xfe\x9e\xd6\x9a\x2a\x26\x84\x59\xdc\x72\x45\xc9\x1d\xad\xdb\xfe\xdc\x07\x25\x79\xe9\x28\x64\xf6\x91\x89\x07\xa6\x79\x41\xc9\x2f\xac\x63\x40\x69\x61\x03\x9c\x55\x1c\xc3\xf7\xc8\xf2\x7b\x0e\x06\x1b\xab\xdb\x0a\xd0\xb5\xb7\x3e\x82\x06\x4e\x05\x07\x88\x95\xd7\xbd\x53\xe0\x3f\xe4\x28\x5d\x23\x8c\x2c\xe9\xd1\x41\xe8\x12\xf3\x98\x03\x10\x76\x4a\x76\x75\x39\x04\x78\x67\xc3\x5b\x48\x21\x55\x4f\x73\xf9\xb1\x1b\x10\x10\x36\x00\x31\x35\x94\x0a\x40\xc4\x6b\x40\x91\x59\x34\xb4\x2c\x11\x8d\x36\xd4\x59\x36\x64\x09\xb3\x32\xd2\xe7\x27\x79\xac\xd9\xcf\xf1\x1c\x93\x97\x2e\xc8\x1e\x8a\x65\x7f\x09\xff\xae\xac\x50\x34\x20\xd2\xb2\xc9\xc8\x9b\x15\x3a\xeb\x68\x90\x18\x2d\x2b\x67\x9b\x1f\x93\x78\xe9\xaa\x21\x88\xca\x1a\xa3\x12\xfa\x38\x14\xce\x89\x98\x58\xa3\x46\x86\x8c\x35\xa2\x18\x3c\xfd\x8a\x34\x07\xec\x85\x07\x46\xea\xd7\x96\x89\xda\xf3\xce\x04\x5e\xef\x01\x0f\xda\x1c\xd6\xec\x49\x47\x25\x2b\xa4\xa2\x26\x4e\x19\xa9\x65\x8d\x75\xd0\x78\x36\x65\x64\x09\x81\x21\x53\xde\xad\xa7\x5c\xb3\x6b\x2c\x68\xf4\x50\x70\xd2\x6a\x25\xeb\xdd\x02\xe4\xe2\x97\x6f\xd0\x18\x27\x41\x5c\x57\x49\xd2\x4b\xa1\x46\x00\xfd\x09\xde\xab\xc4\x07\xd6\x90\xd1\x4b\xcc\xf3\x50\xda\x48\x4d\xc1\xc9\x12\xbe\x59\x49\xce\xcb\xb2\x1c\x94\x66\x7b\xf9\xc0\x94\x55\x8d\x9f\x63\x03\x86\xce\x72\x24\xd2\x8b\xa7\xcf\x1b\x9c\x8a\x02\x49\x76\x81\x61\xeb\x44\x90\x84\x55\x62\x92\x80\x3f\x97\xae\xbb\x1f\x55\x0c\x70\x1c\x31\xa5\xc8\x94\x1e\x98\xf0\x18\xf8\x15\xc8\xee\xf3\xfb\xfd\x5c\x8e\xd1\xed\x50\x89\x12\xbf\x93\xda\xde\xad\x13\xb9\x6d\x14\x3b\x51\xe2\x87\xb9\x30\x2e\xf0\x21\xe7\x43\x18\xfa\xb8\xb8\x00\x2b\x0a\xf3\xab\x05\xfc\x22\xd5\xc4\x7f\x2b\xe4\x63\xf4\x94\x91\xb6\x50\x52\x88\xde\xa0\x0c\x9a\xb2\xec\x54\xc1\xc8\x7b\x59\x32\xf2\x59\x99\x8e\xfc\x89\xd5\x42\x2e\x48\x25\x6b\xd9\x36\xb4\x60\x47\x45\x19\x6f\x26\x2b\xae\x0f\x09\xc4\xe3\x1d\x98\x5b\x32\x03\x1c\xfc\xf2\x72\x8f\x68\x6b\xa5\x80\xe6\x7d\xfe\xe1\xe6\x66\x79\xb3\x3e\x12\x9f\xc4\x6f\xfa\x9a\x3e\xb8\x07\x3c\xe9\xb4\x8b\x0e\xd2\x76\x92\x45\xd6\x5c\x83\xbe\x47\x45\x9b\x53\x8d\x2a\x18\x5d\x43\x53\x00\xcb\x4b\xa6\x29\x17\x2d\x18\xdd\x76\x15\x20\x04\x27\x5d\xd1\xa9\xd6\x84\xbe\x91\xbc\xd6\x4c\xd9\x78\x76\xda\x38\xdf\xc3\x9c\x90\x0e\x26\x8d\x6b\xbb\x1e\xf6\x0f\xf2\x1a\x1f\xfd\x82\x6d\x35\xba\x84\x61\x98\xce\xa9\x3b\x96\xae\x1c\x76\xe3\x3f\xc0\xa5\xc6\xd5\x51\xc9\xdb\x46\x98\x1b\xc3\x56\x30\xeb\xfe\x5f\x5d\xab\xf9\xf6\x19\xe0\x06\x36\x9a\xb4\x16\x0c\x8d\x35\x9c\xef\x81\x48\xc1\x5e\xe4\x75\x17\x25\xd4\x6a\x17\x56\xed\x91\x48\xc3\x77\xc7\x4b\x96\x53\xe4\x32\x54\xb0\x06\xfe\x68\xa7\xe5\x41\x2e\xa8\xf1\xa5\xbe\x59\xbd\xc0\x3d\xc0\x50\x67\x37\x1a\xd9\x72\xec\xb1\xfd\x10\x06\x4b\x8b\xfb\xe7\xeb\x60\xef\x40\xb3\x73\x60\x95\x04\x03\xc9\xd2\x0a\x2a\x8a\x0b\xb4\x3a\xb8\xd9\xcd\x61\x2e\x18\x6c\x7c\x6f\x02\x0f\x16\xbd\x23\x31\xe0\xa4\x42\xcb\x8e\x9a\xe3\x21\xd4\xc1\xd1\x77\xae\xf1\x1e\x8c\x85\x3b\x09\x8c\x8d\x87\x53\x99\x9f\x6e\xc5\x70\x7d\xac\x7b\x01\xd6\x4a\xb2\x4c\x92\xaa\x9d\x30\x7a\xca\x5e\x30\x22\xa6\x85\x51\x7a\x98\x80\x41\x2f\xf9\xf7\xae\x83\xa8\x2c\x67\x5b\xe9\x3a\xcf\x80\x98\xd9\x2c\x4c\x09\xcd\x01\x9e\x9d\xb6\x0e\xb9\xec\xa6\x49\xf2\xc2\x2c\x87\xa2\xef\x2b\x12\x03\x18\xb9\x6b\x04\x22\xf7\xe8\xaa\x14\x9d\x1c\x5c\x0f\xbc\xe5\x39\x17\x5c\x03\xf8\xf6\xbc\x2c\x59\x3d\xc4\x0b\xac\xac\x4c\xcb\xa2\x82\xfd\x79\x91\xcc\x03\x7a\x24\x15\xb7\x3d\xdc\xa8\x1b\xf0\x1e\x06\x9a\x0a\x01\x33\x1c\xc2\xec\xd4\x9a\xc7\xc4\xa9\x98\x9c\x98\x70\x3f\x88\xa4\xe5\xf2\xe3\xe9\x7b\x63\xbf\x05\x9b\x76\x27\x75\x22\x3f\x5a\x30\xff\x64\x59\x13\x70\x8d\xef\xea\xc8\xd8\xd1\xfa\xe4\x3e\x35\x53\x85\x82\xfa\x3e\xd3\x1d\x1b\x5d\xbc\xd2\x24\x1c\xac\x91\x57\x86\x9e\x91\xaf\xc8\x34\x7b\x32\x62\xc5\xf1\x16\xff\xce\xb5\x60\xd3\x13\x37\x18\x1f\xcb\x78\xe9\xc6\x07\xb2\xc4\xe8\x91\xaf\x08\x51\x85\xdd\x67\x20\x2a\x64\xef\xa9\xf6\x5e\x67\x23\xe3\xc7\xa4\xa2\x4f\xfd\xf3\x70\x95\x0c\x2e\xa1\x1e\x94\xe3\xeb\xe9\xd1\xbc\x09\x4f\x9a\xa9\xd0\xd0\x3a\xcc\x53\x2e\x64\x71\xef\x1f\x8a\xdb\x2e\x37\x46\x9c\xbe\x0b\x78\x77\x8a\x60\xe8\x6d\xdc\xcc\xb3\x2e\x78\x68\xe9\x4c\xd3\x2f\x00\xb7\x41\x93\x1e\x66\x4c\xaf\x77\x50\x2a\x18\x5c\xe4\x94\x9d\x84\x76\x74\xa2\xf6\xfe\xf1\x3a\xff\xa1\x8a\xa0\xb5\xe4\x52\x4c\x3f\x8c\x0e\x01\x6c\x2b\x53\x64\x07\xfd\xfe\xd1\x2b\xc4\xd4\xb1\xbe\xbe\x4d\x1e\x9b\x5a\xc3\x2e\x15\x63\x6c\xf4\x4f\x80\xf8\x16\x1e\x6e\x0e\x14\x43\x0b\x1a\x12\xfa\x1f\xf3\x68\xa7\xe6\x7e\xc3\x58\x0d\x2f\xbb\x92\x5c\x78\x89\xdf\x18\x2c\xcf\x91\x35\x7c\xc9\x4e\xfc\x57\xc1\xb2\xaf\x06\xb3\x37\xba\x63\x4d\xfa\xf8\xcd\xbc\xbe\x8f\x26\x2d\x4e\xcd\xa8\xe4\x0a\x9f\x63\xb6\xf1\x77\x55\x8d\x42\x8e\x7a\x2a\x0a\x18\xc6\xa6\x7f\xc5\x75\x39\x74\x84\x79\xa8\xae\x1f\xc0\x23\x81\x3f\x94\xf0\xed\xec\x9f\x00\x00\x00\xff\xff\x28\x13\x27\x20\x68\x11\x00\x00")

func cssIndexCssBytes() ([]byte, error) {
	return bindataRead(
		_cssIndexCss,
		"css/index.css",
	)
}

func cssIndexCss() (*asset, error) {
	bytes, err := cssIndexCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "css/index.css", size: 4456, mode: os.FileMode(420), modTime: time.Unix(1504200578, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsIndexJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x53\x4d\x6f\xda\x40\x10\xbd\xfb\x57\x4c\xd5\x83\x6d\x52\x2d\x9c\xdb\x12\x89\x44\x39\x44\x6a\x7b\x28\x52\x2f\x08\x89\x65\x3d\xe0\x15\xeb\x1d\xba\x5e\x83\x50\x95\xff\xde\xd9\x0f\x42\x83\xa8\x82\x64\xb4\xb6\xdf\xbc\x79\xef\xcd\x78\x3c\x1a\x15\x30\x82\x39\xfa\x61\x2f\xf8\x34\x2e\x0a\x45\xb6\xf7\xa0\x3d\x76\x3d\x4c\xa1\x21\x35\x74\x68\xbd\xf8\x3d\xa0\x3b\xcd\xd1\xa0\xf2\xe4\x66\xc6\x54\xe5\x42\x37\xcb\xb2\xce\x78\xa3\xed\xee\x1d\xbc\xf8\x8e\x76\x00\xc9\x25\x45\x6e\xfb\xd8\xa2\xda\x81\xde\xc0\x0a\xcd\x0a\x74\x0f\x34\xf8\x74\x6d\xe0\xa0\xf1\x98\x15\x6d\x06\xab\xbc\x26\xcb\x88\x07\x34\x74\x9c\x2b\x47\x4c\x88\xa6\x86\x3f\x05\x80\x63\xf1\xce\x02\x1a\xb1\x45\xff\x40\x83\x6d\xb4\xdd\x3e\x1a\xcd\x2a\x7e\x72\xf7\xaa\x16\x6b\xf2\x9e\x3a\xb8\x87\x49\xf1\xf2\xda\x7c\xc6\x9c\x07\xe9\x31\x5a\x85\x95\x5e\x5d\x77\x93\x19\xf0\xcc\xef\x2b\x9d\x7a\x45\x9b\x62\x43\xee\x49\xaa\xb6\x42\x98\xde\x03\x0a\x65\x64\xdf\x7f\xd3\xbd\x17\x0e\x3b\x3a\x60\x55\xc6\x52\x2c\xeb\xfa\x5c\xb2\xd0\xcb\x7f\x60\xb2\x69\x2e\x98\x5b\x92\x7c\x8b\xa0\xc8\x39\x96\x0f\x5d\x48\x2d\x6a\xe4\xb6\xe1\x4d\x40\x72\xe8\x9e\xfd\xf5\xa0\x6d\x04\x87\xb4\xf6\xe4\xfc\xff\x3c\x54\x59\x3e\xf2\x64\x79\x4a\x93\x82\x6f\x02\x5d\xf5\x85\xef\xbf\xa6\x69\x0b\x83\x76\xeb\x5b\x7e\x72\x77\x97\xe0\x10\x66\x53\xbd\x4d\x3d\x42\xd9\x4e\x7d\x86\x00\xac\x1d\xca\x5d\x3c\xbf\x14\xe1\xe2\xbf\xab\xec\x6e\x79\xec\x23\x1f\xf4\xfb\x13\x3b\xe0\x89\x9d\xb2\xf6\xa3\xb6\x0d\x1d\x43\x46\x4f\x07\xb6\x18\x02\x43\x8b\xae\x2a\x53\x41\xf9\x09\x62\xec\x17\x67\x97\x7d\x9a\x35\x0d\xf4\x1d\x91\x6f\x33\x3b\xd3\xbe\xc7\xaa\x8c\x56\xbb\x33\x69\x70\x94\xf6\x19\x0d\xc7\x84\xc2\x4b\xc7\x4b\x15\x1c\x8d\xc7\x79\xc7\xc9\x9a\x53\x91\x92\xe1\x95\xb3\xd4\xe0\x0f\xd9\x21\x7c\x98\x42\x39\x2b\xeb\xbc\x8d\xb9\x62\x70\xe6\x95\x91\xcf\x81\x32\x6e\xe9\xcc\x7b\xa7\xd7\x03\xab\x2f\x5b\x87\x9b\xf0\x49\x44\xbc\xb4\xaa\x25\xf7\xa6\x07\x97\x2d\x26\xcb\x48\xff\xf1\x9a\x3e\x99\xcc\x3f\x7e\x74\xfb\xe3\x0b\x14\xb5\x48\xd8\x67\xeb\xe9\x17\xef\x4a\x95\x66\xb7\xc6\x56\x1e\x34\xb9\xcf\x50\xa6\xdc\xca\x30\x40\x1e\x57\x5d\xfc\x0d\x00\x00\xff\xff\x79\x82\x0e\x63\x1a\x04\x00\x00")

func jsIndexJsBytes() ([]byte, error) {
	return bindataRead(
		_jsIndexJs,
		"js/index.js",
	)
}

func jsIndexJs() (*asset, error) {
	bytes, err := jsIndexJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/index.js", size: 1050, mode: os.FileMode(420), modTime: time.Unix(1504199128, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsSmoothscrollJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5a\x5f\x73\xdb\xb8\x11\x7f\xf7\xa7\xd8\xbc\x9c\xa4\x9c\x42\xd9\x6d\xef\xa6\x23\x55\x6d\x4f\x89\x5a\xbb\xb5\xa3\x1b\x4b\x93\xda\xd3\xb9\xe9\xd0\x22\x24\xe1\x42\x11\x3a\x12\xd4\x9f\x89\x7d\x9f\xbd\x8b\xbf\x04\x40\xd2\x39\x4f\x33\x93\x97\xe6\xc5\x04\x76\xb1\xd8\x5d\xfc\x76\xb1\x58\x65\xf0\xfa\x0c\x5e\x43\xb1\x65\x8c\x6f\x8a\x65\xce\xd2\x14\x76\x2c\x3d\xad\x28\x7e\xbc\x81\xfd\x79\xf4\xfb\xe8\x3b\xc1\xb1\xe1\x7c\x57\x0c\x07\x03\x1a\x6f\x93\xb2\xe0\x71\x16\xad\x29\xdf\x94\x0f\x11\x65\x03\x77\xb5\xe0\xfd\xdd\xf9\xc5\xf7\xd0\x5d\xf6\xe0\x9d\xe4\x84\x7f\xc6\x05\x27\x59\x1f\xfe\x41\x72\xb2\xa5\x71\x01\x37\x24\xa3\xcb\x0d\x49\x53\x8a\x7b\xdc\x5c\x2d\xe0\x9a\x2e\x49\x56\x10\x5c\x3c\x38\x3b\xeb\xae\xca\x6c\xc9\x29\xcb\xba\x87\x3e\x24\x7d\x28\xb3\x84\xac\x68\x46\x92\x1e\x7c\x3a\x03\xe8\x94\x05\x81\x82\xe7\x74\xc9\x3b\xa3\x33\x9c\x10\x26\x00\x6e\x1b\xa7\x28\x9b\x14\x6a\x70\x18\xc2\x81\x66\x09\x3b\xc0\x3a\x65\x0f\x71\x0a\xec\xe1\x67\xb2\xe4\x8a\x98\x0c\x21\x61\xcb\x72\x4b\x32\x3d\x61\xb7\x18\x56\x9f\x92\x32\x90\x1b\x0c\xac\x4f\x70\x64\xb4\xb3\x73\x5d\xa5\x97\xe4\xcb\x09\x2f\xf3\x0c\x0e\x1b\x92\x81\x72\xc8\x84\x6c\xe2\x3d\x65\x39\xd0\x8c\x93\x7c\x15\x2f\x09\xd0\x02\x8a\x72\xb7\x63\x39\x57\xbb\x00\x5d\x41\xb7\xe3\xb3\x77\x90\x1f\x92\xc8\xa8\x39\x4d\x89\xf8\x13\x15\xfc\x94\x12\xb3\x1f\xe8\xed\x46\x72\xf4\x74\xa6\x94\x78\xad\x68\xaf\xb5\xe1\x85\x1e\x0e\xe4\xdf\x7d\x9c\x83\x96\x05\x63\x38\x44\x97\x8b\x9b\x6b\x33\x7e\x7c\xc4\x09\x3d\x18\x59\xee\xf9\xdb\xdb\xd9\xf5\xf5\x7f\x16\x57\x37\x53\x5c\xf1\x87\xef\xff\x38\x0a\xf7\x51\x9e\x85\x75\xcc\x37\x24\xa7\xd9\x1a\x58\x4e\xd7\x34\x43\x9f\x6b\x40\x6d\x09\xdf\xb0\xa4\xae\x88\xe5\x1b\x5b\x83\xd4\x0a\x3c\xbb\x48\xaf\x95\x4a\xa9\xef\x05\xeb\x7b\x5c\x93\x53\xc5\x37\x39\x19\x1a\x49\xe7\x5a\x86\xf1\xd9\x2e\x67\x9c\xf1\xd3\x8e\x38\x32\xd5\x97\xe6\xf0\xc5\x5e\x65\x9c\x7d\xa0\xe4\xd0\x2e\xc0\x70\x28\xb7\xd7\xfc\xa1\xf0\x03\x9c\x6e\x85\x33\x94\xf1\x35\xdb\x33\x44\xa6\x38\x80\x1d\x62\x82\xe5\xdb\x38\x43\x5c\x7c\xf3\x8d\x3f\x11\x21\x93\x56\xed\x2f\x75\x4a\xf4\x80\xf8\xee\x7a\xd3\x3d\x18\xc2\xbb\x98\x4b\xb2\x55\xcb\xea\xb5\xdc\xc4\xd9\x9a\x14\x60\xc3\xbc\xa0\x12\xc7\x34\x2b\x68\x42\x00\x03\x95\x28\x83\xcd\x82\xbf\x2a\xdd\x7d\x67\x59\xe2\x2e\xce\xe3\x2d\x7c\x7a\x5f\x6e\x1f\x48\xfe\x04\xc7\x36\xc2\xc9\xb3\xdd\x46\x8f\x27\xb4\x7b\xec\xc3\xa9\x82\x35\xdf\xd0\x42\xfb\xfa\x9a\xac\x04\x52\x8f\xa3\x3a\x69\xc1\x76\x48\x39\x05\xe8\xb7\xe6\xaa\xd8\x28\xf0\x6f\x51\xa6\x1c\xd8\x0a\xe2\xdd\x2e\x3d\x89\x33\x21\x98\x27\x60\x8b\x78\xad\xd4\xe1\x0c\x62\xc8\xa4\xca\xa1\xf9\x82\xbb\xcd\xb8\x8f\x96\x60\xb6\x33\xa4\x66\xab\x85\xac\xee\xc7\x30\x80\xe1\x3c\xfa\x0e\x65\x74\x2f\x44\x3a\x44\xb5\xa2\x25\x2b\xba\xf2\xe3\xc7\x2b\x9c\xff\xd8\xeb\xb5\x19\x89\x10\xa0\x4b\x3c\xf1\x42\xe4\x90\x58\xe7\x71\x78\x30\x39\xa7\xd8\xb0\x32\x4d\x70\x2c\x6d\xa7\x24\xa9\x1d\xad\x64\x98\xc4\x34\x9d\x95\x2d\x47\xfb\x38\x93\xf1\xed\x9e\xb0\xb5\x75\xc2\x58\x4a\xe2\xac\xc5\x58\x4f\x78\xf7\x58\x59\x2d\xf2\x9d\x88\x27\x3c\x93\x23\xbc\x1a\x8f\xa1\xa3\x72\x48\x47\xd3\xd5\x3f\x0c\xd3\x23\x8c\x91\x9a\x95\x32\xf1\xfa\x94\xc8\x1a\x29\x58\xbc\xac\xfd\x0c\x5f\x27\x2e\x39\xab\x6f\x13\x30\x61\x44\xe0\xbd\xc5\x3b\x95\xc6\x32\xb9\xaf\x68\x5e\x70\x88\xf3\x35\x46\x2f\x17\xf1\xa2\xb4\x1e\x78\xfa\x21\x1f\x8a\xb1\xf2\x30\xd7\x8b\x2d\xfb\xa0\x65\x0a\x62\x5d\x59\x0d\x03\x9e\x97\xc4\x00\x5d\x9f\x75\xe0\xac\x71\x9b\xb3\x30\x77\x84\x56\x28\x30\xb4\x1a\x21\xef\x15\xa9\x9f\x31\x04\xbf\x12\x4f\x73\x25\x21\xd4\x72\x85\x97\x4a\x5d\x4d\x14\xcd\x37\x39\x66\x35\x92\xe7\xb8\x5a\xde\x80\xae\x2c\xe1\x32\xff\xd6\x03\xbd\x20\x23\x07\x58\xa0\x81\x53\xb1\xb0\xdb\xb1\x8b\xc4\x8a\x3d\x5e\xea\x49\xa7\x15\xfe\xe8\xc5\xc4\x24\xb4\xf8\x21\x25\x80\xc8\x15\x56\x89\x60\x6f\x4d\x67\x62\xd1\xdc\x2e\xf9\x51\xae\xa8\x41\x9f\x25\xe4\x09\x05\x34\x84\xb7\x4f\x09\x20\xdf\x24\xbb\x4b\xd2\xea\x0c\x44\xee\xa7\xc5\x84\x25\xa7\x91\x33\xb3\x89\x8b\x6a\xd5\x7c\x87\x45\x42\x40\xfd\x40\x0b\x8a\xa4\xd9\x1e\x93\x7d\x6a\x93\x3b\x60\x19\xe3\x9c\x2e\x11\x37\x29\x49\x23\xe5\x04\xa1\xa8\xe5\x93\xe7\x53\x10\x0e\x4b\x86\x19\x43\xaa\x8a\x92\xa9\xd8\xae\xb0\x2c\x4a\x2f\x29\x43\x22\x28\x89\x1e\x1c\x3d\xa1\x41\x4b\x18\x3b\x20\xc4\x9d\x97\x98\x63\x32\x7e\x49\xe8\x7a\xc3\xe1\x4f\x62\x46\x9d\x8d\x9e\x79\x7c\x6c\x62\xff\x17\x4d\x30\x67\x39\xdc\x72\xc2\xdb\x36\x30\xdf\xdb\xf6\x10\xad\x09\x7f\xcb\xb6\xbb\x12\x81\x35\x17\xf5\x11\x3a\xbc\x2f\x93\x46\x2f\x62\x76\x85\x88\x88\xbd\x12\xd3\xb1\xe0\x45\x98\x52\x44\x4d\xf7\x95\x36\x1d\xa3\xe8\x55\xb7\xc1\x4c\x31\x5f\x57\x43\xa4\xe5\xb3\xc0\x75\x4d\x3e\x6a\xb4\x40\x2a\x68\xd7\xeb\xd0\x22\x69\x1b\xd2\x0b\x92\xae\x30\x87\xec\xd9\x47\x92\x38\xb7\xd6\x26\xe6\x7d\x58\xd3\x3d\x06\x5b\x2c\xce\x96\x93\x23\x4e\x60\xb9\xbd\x2b\x64\x74\x95\xeb\x8d\x0e\x0f\xbc\xf7\x6a\xc9\x1f\xd9\x42\xe0\x9b\x64\xaf\x65\xb5\x24\x76\x5c\xd8\xd5\x1c\x3e\xb4\xb1\xe8\x11\x06\x63\xfd\xd1\xed\xb9\x10\xc6\x38\x2e\x3d\x4c\x2f\xcb\x5c\x80\xf4\xae\x61\xee\xde\x9d\x23\x69\xbc\x2b\xd0\xe4\x31\x26\x41\x21\xfb\x8d\xd1\x0c\x4b\xe1\x38\xe7\x0b\x9c\xeb\xc1\xc0\xad\x52\x47\x4e\x46\x8a\xf7\x8c\x26\x56\x86\x10\x50\xc0\x06\xa1\x48\x72\xe1\x3a\xcc\x7b\x19\xb1\x65\xa3\xd9\xc7\x7c\xfd\x19\x2e\xb0\xec\xba\xc0\x92\x4a\xcf\x78\x82\x45\x29\x21\x6e\x73\x51\x4e\x60\xed\xe0\x6e\x61\xb5\x47\x93\x85\x3c\x71\xe5\x6b\x7a\x05\x18\x63\x3f\x32\x78\x06\xdd\xc1\xb7\x60\x5c\x1b\x1d\x43\x73\xef\x7a\x78\x54\x9e\x2f\x8d\xcf\x42\x39\xf7\xae\x9c\x53\x28\xe7\xde\x91\x63\x04\x69\xba\xc2\x46\xb4\x8c\xf1\x61\x63\xd7\x58\x44\xf7\xad\xe2\xf6\xeb\xbe\xe7\x3a\xc6\x94\xfc\x2c\x27\xe2\xe6\x3a\x10\x04\xff\x9e\xc8\x5c\x9e\x93\x18\x1f\x7d\x09\xb0\x32\xc7\x1a\xb9\xe0\x58\xfa\x0b\x34\x39\xd7\x9c\x75\x8a\xa8\x09\x2a\x27\xe0\x0d\x6d\xcd\x74\x29\x27\xf7\x6a\x3b\x44\x39\xf9\xa5\x44\xb1\x3f\x64\x74\x2b\x05\xff\x0d\x01\x4d\xba\x02\xab\xba\x54\xee\x9b\x95\xbd\x5e\x75\x7d\xb5\x84\x9b\x34\xa3\x30\x6f\xc8\x03\xbe\x75\xeb\xe5\x55\x2d\x9e\x24\x79\xae\xdf\xc1\x4d\x71\xf5\x58\xbb\x57\xfe\xd7\x2a\xda\xd9\x52\xa6\x3d\xbf\x90\x16\x11\x54\x1d\x9e\x1b\x57\x0a\x4e\xb5\x19\x2f\xf6\x94\x59\x35\x9e\x85\x17\xe3\xd5\xd1\xeb\x77\x8f\x46\x80\x97\x40\xe4\xe9\x7a\xb7\x8a\x7b\x74\xce\xfd\x8d\x4f\xa2\x2a\xf1\xeb\x80\x18\xdb\x57\xde\x9d\x7a\x0e\xee\xe2\x35\xb9\x9b\xad\x56\x78\x9f\x05\xcc\xf7\x0e\xf3\x7d\xc5\x7c\x1f\x32\xeb\x03\x1b\xdb\x27\xa8\x5e\x53\x5d\x0c\x04\x6b\x9c\x36\x15\x49\xda\xa0\xa3\xbd\xbd\xc4\x8b\xa5\x41\x2d\x4b\xc7\x67\x4b\x83\x22\xde\x73\xa8\xa9\xb8\xd2\x6e\x4d\x19\xdb\xc9\x37\x36\x5e\x24\x88\xc8\x95\xc0\xb8\x79\xbe\x8a\xac\xdc\xa4\xf2\xd0\xf9\xee\x07\x5b\x0f\xf5\xdf\xbe\xaf\xb1\x38\xe3\x61\xf5\x19\x50\xef\x34\xe9\x2e\x98\xbf\xd7\xf3\xf7\xd5\xfc\x71\x88\x90\xb4\x23\x7c\xb1\x9f\x8c\x6d\x61\x51\x67\x90\x3f\xbb\xbd\xfa\xfb\xd5\xfb\x1f\xae\xe1\x66\xba\xb8\x9c\xbd\x9b\xc3\xec\xc3\xf4\xf6\xf6\xea\xdd\x74\x6e\x43\xe0\x4c\xfb\xc4\x76\x09\x44\xe1\x5a\xb5\x09\x24\xd9\xd2\xc6\x0e\x05\x07\xb6\xb9\x54\x21\xd0\x5e\x14\xe1\xf3\x09\x51\xab\x12\xd7\x2f\x25\xcd\x6d\xd5\x2a\xb0\xec\x3f\x6d\x4c\x35\x5d\xfc\xfb\xfc\xa7\x9e\x8b\xec\x00\x5e\x2a\xab\xba\xb5\x4b\xdf\x19\xb8\x52\xa2\x54\xbc\x7b\x11\xc0\xee\x64\x2b\x33\xc7\x97\xb0\xc7\x7b\xf1\x93\x65\xed\x8d\x82\xfa\xbd\x09\x5c\xd7\xd3\x05\x2c\x2e\xa7\x30\xbf\x99\xcd\x16\x97\xef\xa7\xf3\x39\x4c\xa6\x78\x0a\xaf\x0c\xb4\x9c\x34\x13\x18\xe1\x98\xa0\x22\xbb\x1a\xff\xfa\x6b\xcd\xa4\x56\x22\x9a\x70\xe6\x29\x6c\xfb\x2a\x03\xa7\xc9\xe3\x9d\xec\xe4\xf4\x95\x8f\x73\x72\xfa\xff\x81\x8a\x32\xa3\x3d\x3d\xf7\x9e\x3b\x6f\x77\x69\x3d\x59\xf7\x5a\xe1\xd0\xda\xd7\x13\x59\xa0\x8d\xa8\xb3\x42\xeb\xda\xf1\x33\x2b\xbf\x1a\xcc\x4c\x2f\xb3\x06\x33\xd5\x07\xeb\x7b\x33\x2f\x42\xdb\x17\xc5\x9b\xa8\x0c\x52\xd5\xa7\xab\x29\xe1\x56\x0f\x5c\x36\xec\xc2\x8d\x47\x5f\x06\xb6\x75\x97\xd4\xc6\xaa\x89\xa2\x34\x15\xcf\x41\xd5\xee\xeb\x60\x95\x2f\xe7\x86\x61\xdf\xb1\x61\xb1\x34\xc1\x5f\x2b\xa6\x86\x41\x5f\xf2\xc5\xd8\xd5\xa9\xad\x9d\xdc\x8c\x41\xe1\x55\xf4\xe7\x79\xe0\xd6\x51\x43\xef\x48\xb1\x39\xed\x23\x17\x71\x8e\xf6\x4e\x05\x01\xd2\x2f\x43\xb9\xd4\x04\xfb\x73\x2e\x62\x3b\xcd\xab\x82\xdb\x77\x89\xcb\x69\x42\x45\xb3\xfb\xf5\x34\x54\xb5\x41\x43\x21\xe6\x6a\x1a\x9c\xb5\xdb\x31\xfe\x56\x4a\x0e\xd1\xe0\x34\x8e\xbf\xfd\x0c\xd8\x9f\x7e\xeb\xc9\x79\x3f\x06\x7c\x8e\xe9\x2b\x5f\x58\x46\x8d\xcf\x86\x8e\x2b\xd0\xef\xaf\x0a\xc0\xe7\xf8\xae\x1d\x7a\x3c\x5f\xf6\x72\xf2\x5f\x2d\xaa\x89\x26\x3c\xd7\xd4\x5b\x13\x9a\x7b\x9d\x06\xd5\xfc\xba\x45\x84\x17\xb6\xac\xae\xf8\x45\xaf\x68\xc2\xd0\x1a\xac\xa1\xdf\xca\xee\x93\xe0\xf4\x7b\x15\x4b\x3b\x2f\x24\x48\xdc\xb4\xae\x72\x8f\x24\xd4\xf8\x55\xe3\x63\x47\xfe\x48\xb8\x27\x71\x6a\xda\x93\xe6\xf7\x97\x9d\xd3\x88\x84\x17\xa6\xba\x70\xef\xe7\x68\x7e\x98\x38\xc6\xaa\x10\x7f\xe3\x7a\x30\xa8\xd9\x5a\x85\xa9\x88\x72\x65\x89\x14\xe0\x8b\xaa\x2a\x3c\x0f\x26\x95\x3f\x74\xef\x96\x66\xb0\x47\x8c\x8a\x46\x71\x55\x93\xd8\x3c\xd8\x90\x9e\x9e\xd3\x57\xe6\xa4\x40\x8b\xe6\x4c\x64\xda\xe5\xbf\x25\x09\x35\x9d\xe1\x0b\x95\x0e\xfd\x5e\x53\x3a\x70\xe6\xcb\x95\x36\x09\x4c\xc7\x9e\x73\x15\x90\xa3\xd0\xb3\x68\xbc\x0d\xd0\xb4\x25\xdb\x6e\x59\xf6\xb3\xea\x0b\x6f\x59\x52\xa6\x24\xb2\x4b\xe0\x93\xfd\xdd\x7b\x58\xfd\x4f\x01\xb5\x8d\xeb\x26\x94\xa3\x7e\x7f\x96\xa3\xea\xa7\x72\xa5\xcf\x53\xaf\xab\x3a\x2b\x7d\xfb\x43\x3c\x52\xfe\x1b\x00\x00\xff\xff\x13\x45\xe8\xd6\x8c\x20\x00\x00")

func jsSmoothscrollJsBytes() ([]byte, error) {
	return bindataRead(
		_jsSmoothscrollJs,
		"js/smoothscroll.js",
	)
}

func jsSmoothscrollJs() (*asset, error) {
	bytes, err := jsSmoothscrollJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/smoothscroll.js", size: 8332, mode: os.FileMode(420), modTime: time.Unix(1504199128, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x54\x4d\x6f\xe3\x20\x10\xbd\xef\xaf\x60\xd9\xf3\x1a\xed\x6d\x0f\x38\x97\xb4\x55\x2f\x55\x23\x25\x52\xd5\x23\xb1\x27\x81\x14\x63\x0b\xc6\x55\xa2\xd4\xff\xbd\xf8\x23\x15\xe0\x34\xea\x29\x33\x79\xef\xc1\x9b\x07\x98\xff\xbe\x7b\x5e\x6e\x5e\x57\xf7\x44\x62\xa5\x17\xbf\xf8\xf8\x43\x08\x97\x20\xca\xbe\xf0\x65\x05\x28\x48\x21\x85\x75\x80\x39\x6d\x71\xf7\xf7\x3f\x9d\x20\x54\xa8\x61\x71\x3e\x67\x9b\xbe\xe8\x3a\xf2\x41\x7c\xb3\x6e\xb7\x38\xf6\x9c\x8d\x8c\x91\xad\x95\x79\x23\x16\x74\x4e\x1d\x9e\x34\x38\x09\x80\x94\x48\x0b\xbb\x9c\xa2\x84\x0a\x98\x68\xe0\xc8\x0a\xe7\x98\x32\x25\x1c\x33\x5f\x0d\x3b\x71\x76\xb1\xc3\xb7\x75\x79\x9a\x96\x2b\xd5\x3b\x29\xb4\x70\x2e\xa7\x2f\x56\x34\x0d\xd8\xc9\x56\x8c\x2d\x6b\x83\x42\x99\x00\x8d\xf1\xb5\x2a\x61\x2b\x42\x34\xc6\x9f\xc0\xb4\x11\x48\xfc\x8c\x56\x98\x3d\x90\x6c\x25\xf6\xe0\xba\x2e\x02\x63\xb5\x42\xa8\x12\xf5\x40\x11\xd3\xdc\x7f\xfa\xbc\x74\xbb\xef\x3a\x1a\xe4\xc8\x99\x48\x35\x9c\xf9\x55\x53\x1b\x60\xca\x68\xf7\x84\x34\xb5\x57\xc7\xee\x63\x01\x83\xdf\x8e\xfd\xe8\x13\x07\x9b\x58\x0f\x09\x83\x55\x52\xf8\x35\x66\x34\x4f\x74\x8d\x30\x17\x26\xc2\x11\xe3\xe9\x7a\xf4\xa6\xc4\xf9\x2b\x74\x51\x85\xd7\x69\x2e\x9c\xe5\x32\xfb\xe3\xe6\x69\x85\x13\xf5\x38\x25\xaa\xcc\x69\x70\x28\xa9\x4b\xf9\x2f\x9a\xc4\xb7\x09\xc3\xa3\x53\xb6\xe9\x56\x73\x63\xb3\xf3\x0b\xdc\x3c\xd4\xf5\x90\xec\x7c\x40\x57\x58\xd5\x20\x71\xb6\x88\xde\xcd\xc1\x31\x57\x79\x91\xf4\x78\xad\x75\x76\x70\xbd\x78\x24\xff\x50\x3f\x3e\xbb\x6b\xc2\xc0\xc5\x57\xc9\xd9\xf8\x1a\x7d\x0a\xc3\x67\xe3\x33\x00\x00\xff\xff\x2b\xb5\x57\x66\x4e\x04\x00\x00")

func viewsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_viewsIndexHtml,
		"views/index.html",
	)
}

func viewsIndexHtml() (*asset, error) {
	bytes, err := viewsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/index.html", size: 1102, mode: os.FileMode(420), modTime: time.Unix(1504199128, 0)}
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
	"css/index.css": cssIndexCss,
	"js/index.js": jsIndexJs,
	"js/smoothscroll.js": jsSmoothscrollJs,
	"views/index.html": viewsIndexHtml,
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
	"css": &bintree{nil, map[string]*bintree{
		"index.css": &bintree{cssIndexCss, map[string]*bintree{}},
	}},
	"js": &bintree{nil, map[string]*bintree{
		"index.js": &bintree{jsIndexJs, map[string]*bintree{}},
		"smoothscroll.js": &bintree{jsSmoothscrollJs, map[string]*bintree{}},
	}},
	"views": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{viewsIndexHtml, map[string]*bintree{}},
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

