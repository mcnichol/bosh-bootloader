// Code generated by go-bindata.
// sources:
// templates/bosh_director.tf
// templates/cf_dns.tf
// templates/cf_lb.tf
// templates/concourse_lb.tf
// templates/jumpbox.tf
// templates/vars.tf
// DO NOT EDIT!

package gcp

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

var _templatesBosh_directorTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x95\x51\x6f\xda\x30\x10\xc7\x9f\x9b\x4f\x71\xb2\xf6\xb0\x49\xc0\x68\x46\x80\x17\x3e\xc9\x84\x22\x27\x1c\x99\x57\x13\x47\x8e\x0d\x9d\x2a\xbe\xfb\xe4\x38\x4e\x4c\x81\xcc\xb0\xa2\xb6\x0f\x09\xd6\xff\xfe\xbe\xfb\xf9\xce\x11\x5a\x55\x5a\x01\x29\x51\x1d\x84\x7c\x49\x4b\xba\x43\x02\x6f\x11\x00\xc0\x9e\x72\x8d\xb0\x02\xf2\xe5\xad\x10\xa2\xe0\x98\xe6\x62\x57\x69\x85\x69\xab\x9e\x64\x19\x1f\xbb\x77\x13\x79\x24\xd1\x31\x8a\x9c\x67\xad\xb3\xdb\x6c\xfb\x80\xc6\xd9\xfe\xbc\x60\x9c\x89\xfa\x57\x2a\x2a\x2c\x53\x45\x8b\x40\xef\x2d\x93\x78\xa0\x9c\x4f\x4c\xf0\xd8\x04\x5f\x33\xde\x30\x89\xb9\x12\xf2\xc4\xfc\x29\xd4\xd9\x45\x5f\x70\xff\xad\x77\x55\x26\x5e\xaf\xfa\xee\xa9\x9c\x60\xb9\x4f\xd9\xe6\x38\x6e\xb5\x27\xf1\xac\x54\x28\x4b\xca\xef\xa9\xda\xc5\x7a\x69\x49\xac\x85\x96\x39\x02\xb9\x7c\xba\x04\x88\x77\xbe\x76\x2f\x13\xfd\xf4\x74\x9e\xae\x13\x45\x00\x54\x2b\x91\xe6\x12\xe9\xc9\x81\xd6\xb0\x82\x2d\xe5\x35\x5e\x2c\x28\x67\x1b\x69\x37\xf0\x4a\x31\x8b\xd6\xe0\xab\xd9\xca\xbe\x36\xd2\x11\x2c\x47\x30\xfd\x66\xcb\xd8\x53\xc9\x68\xc6\xd1\xf5\x9b\x67\xa6\xfe\x54\x68\xf8\xac\x80\xd4\x4a\xb2\xb2\x30\xf9\x6d\x70\x4b\x35\x57\x66\xf1\x79\x3a\x69\xfe\xbf\x3f\xcf\x87\x89\xf4\x75\xb4\x50\xec\x42\xcf\x04\xba\xbf\x33\x34\xad\x34\x02\x60\x55\x93\x5b\x2a\x69\x59\x78\x47\xee\xa5\x7d\x34\xb2\x76\x27\xdf\x2f\x60\xfc\x6a\xe4\xdb\x94\xb3\xf2\xe5\x1f\x87\xeb\x3a\x82\x00\xc1\x57\x8b\xff\xb4\x8c\xb3\x02\x3a\x99\x97\xdb\x8d\x97\x42\x04\x60\xd3\xb1\xb5\x9b\x66\xf8\x49\x1c\xfc\x29\x59\x1b\x01\xe5\x5c\x1c\xda\x86\xae\x84\x54\x56\x14\xc7\x64\x04\x64\xbe\x9c\x2f\xcd\x33\x4e\x92\x24\x21\x6b\xab\x91\x42\x89\x5c\x70\x93\x8b\xca\x2b\x93\xdd\xd1\xf8\x28\x2a\x0b\x54\x66\x46\xac\xc3\x69\x31\xdd\xf4\x93\x75\x28\xa6\x3e\x64\x98\x53\xaf\xfb\x08\x50\x01\xf9\x07\x42\x5b\xce\x66\x3f\x9a\xe7\x72\x36\xfb\x40\x88\xee\xa2\xbb\x11\x64\x17\x16\x00\xb3\xd3\x3e\x1a\xa8\x57\xcb\x7b\xa8\x77\x01\x72\x17\x5b\x38\x1b\x17\x31\x56\x22\x14\xd1\xc5\x90\x07\x92\xf2\x8a\xba\xda\x79\xb3\xd8\xf6\x5e\x9c\xc4\xc9\xd4\xbe\x2c\x16\x8b\xcf\x68\xb6\xf6\x03\x6a\xe0\x34\x0b\x83\x28\xdf\x89\x1f\x08\xd1\x7d\xd7\x87\xa7\xf7\x7f\x78\x75\xc7\x34\x1a\x9e\xa9\x9b\x5b\x33\xb0\x1d\x3f\xa9\x05\x3d\x56\x2c\xdf\xf5\xb0\x42\x86\xf9\x9a\x46\x6f\xee\x1a\xf8\xbf\x01\x00\x00\xff\xff\x71\x51\xee\xa5\x54\x0b\x00\x00")

func templatesBosh_directorTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesBosh_directorTf,
		"templates/bosh_director.tf",
	)
}

func templatesBosh_directorTf() (*asset, error) {
	bytes, err := templatesBosh_directorTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/bosh_director.tf", size: 2900, mode: os.FileMode(420), modTime: time.Unix(1510187537, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesCf_dnsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x96\x41\x8f\x9b\x30\x10\x85\xef\xfc\x8a\x91\xd5\x53\xa5\xa0\x95\x7a\xce\xa1\x52\xcf\xbd\xf4\x58\x45\xc8\xb1\x27\x04\xc9\xd8\xd6\xcc\x40\x9a\xae\xf8\xef\x95\x09\x10\xd2\x2e\x6a\x38\x64\x95\xc3\x72\x01\xcc\xf3\xbc\x37\x9f\x11\xa6\xd5\x54\xe9\xbd\x43\x50\x7c\x66\xc1\xba\xb0\xa1\xd6\x95\x57\xf0\x9a\x01\xc8\x39\x22\x6c\x41\xb1\x50\xe5\x4b\x95\x75\x59\x46\xc8\xa1\x21\x83\xa0\xca\x10\x4a\x87\x85\xf5\x5c\xd4\xda\xeb\x12\x6d\xf1\x3b\x78\x54\xa0\xd0\xb7\xfd\xf0\xe5\x36\x15\xf2\xba\x46\x18\x8e\x2d\xa8\x4f\xaf\xad\xa6\x3c\xc9\x2a\xdb\x6d\x7a\x59\x06\x90\xa6\x8c\xc2\x49\x74\x93\xaa\xcb\x7b\x1d\xb2\xa1\x2a\x4a\x15\x7c\xd2\x7d\xfb\xfe\x03\x52\x09\x38\x04\x02\x39\x22\xdc\x54\x07\xf4\x6d\x45\xc1\xd7\xe8\xa5\x6f\x20\x34\x12\x1b\xf9\xab\xdd\x3e\x2e\x23\xb5\x48\x7c\x49\xdc\x6a\xd7\xe0\x25\xc6\x42\xa3\xf9\xbc\xcd\x3c\x05\x1f\x2b\x74\xcb\xa4\x08\x4d\x20\x5b\x30\x8a\x02\x75\xaa\x9c\x35\x9a\xec\xc6\x7a\xfe\x87\xd3\x16\xd4\xe7\xfc\x4e\xf3\x91\x5c\x77\xc1\x13\xd1\x5b\x2e\x7a\x3a\x3f\x47\x73\x13\xea\xd8\x08\x16\xa5\x0b\x7b\xed\x0a\x6d\x2d\x21\x73\x6e\x0e\x9b\xe1\x52\xed\xc6\x05\x9f\xfc\xbf\xa6\x72\x22\xee\xba\x72\x5f\x5e\x5e\xb2\x0c\x60\x9e\x64\x25\xa3\x4e\xa5\x02\x44\x56\x8b\xe6\x3e\xe0\x34\xf9\xbf\x11\xf3\xe1\xdc\xa9\xdd\x7d\x80\xcd\x61\xc3\x7c\xdc\x44\x0a\xbf\xce\x6f\x01\x66\x3e\x3e\x00\xf1\x2c\xf8\xd5\xfd\x69\xe8\xbe\x95\x6e\x35\x58\x31\x71\xe9\xa5\x15\x13\x1f\xcb\x34\x79\x53\x68\x04\xe9\x29\xa1\x5e\xe3\xad\xa6\x6a\x43\x8c\x0e\x69\x89\xec\xf0\xf8\xb1\x74\x4f\x4f\xf4\x21\xb8\x89\xb5\x9a\xa6\x0b\x65\x49\x58\x6a\x09\x8b\x44\x67\x92\x0f\xaa\x2b\xf7\xac\x13\x2f\x6f\x5b\x27\xfe\xc0\x79\xe7\x0e\x45\x68\x8f\xcd\x7e\x86\x71\x3b\x0d\x3e\x92\xe1\x60\xbb\x9b\xfd\xe4\x4d\xfc\xde\x1b\xdc\xd0\xed\x0d\xba\x3f\x01\x00\x00\xff\xff\xea\xd9\xbe\x05\x97\x0a\x00\x00")

func templatesCf_dnsTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesCf_dnsTf,
		"templates/cf_dns.tf",
	)
}

func templatesCf_dnsTf() (*asset, error) {
	bytes, err := templatesCf_dnsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/cf_dns.tf", size: 2711, mode: os.FileMode(420), modTime: time.Unix(1509726301, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesCf_lbTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x59\x4b\x8f\xdb\x36\x17\x5d\x7f\xfe\x15\x84\xf0\x2d\x5a\x20\xf2\xd8\x1e\x37\x75\x17\x59\x15\xdd\xa6\x5d\x74\x17\x0c\x04\x8a\xba\xb2\x09\x73\x44\x96\xa4\xec\x18\x83\xf9\xef\x05\x49\xc9\xd6\x93\x7a\xd8\x41\x31\xc9\xc2\x1a\xf1\xde\x73\xa9\x73\x1f\x3a\x92\x4e\x58\x52\x1c\x33\x40\x81\x52\x2c\x22\x20\x35\x4d\x29\xc1\x1a\x02\xf4\xb6\x40\x48\x5f\x04\xa0\x2f\x28\x50\x5a\xd2\x6c\x1f\x2c\xde\x17\x8b\x5e\x8f\x48\x48\x7a\x32\xbf\x47\xb8\xf4\x7a\xf3\x5c\x8b\x5c\xa3\x40\xf2\x5c\x83\x8c\x62\x4c\x8e\x90\x25\x91\x02\x79\xa2\xa4\x08\x7a\xc2\x2c\xb7\x7e\xff\x7f\xdb\x73\xbe\x67\x10\x11\xfe\x2a\x72\x0d\x4d\xf3\xa5\x43\x09\x59\x1c\x16\x2b\x61\xb9\x92\xe1\x57\x78\xef\x8a\xc8\xe2\x88\x0a\x17\xc7\x17\x69\xcf\x78\x8c\x59\x84\x93\x44\x82\x52\x4b\x92\x86\xe5\x61\xf1\x5b\x07\x57\xea\x10\x09\xc9\xbf\x5f\xc6\xe2\x57\x80\x95\x3a\x84\xd6\xb7\x1b\x5a\x13\x11\x4d\xdb\x7b\x05\x5b\x13\x11\x3a\xe7\x6e\xf0\xb3\x9a\x01\x7a\xee\x21\x81\x48\x48\x0e\x79\x3c\x19\xd1\xb9\xd5\x31\x25\x28\x9e\x4b\x02\x28\x68\x78\xa5\x54\xc2\x19\x33\x16\xa0\xa0\x3c\x0c\x49\xea\xa2\x99\xa4\x23\xf7\xcf\x06\x3c\x61\xb9\x84\xec\x14\xd1\xe4\x3d\x24\x69\xc8\x05\x64\xc1\x02\xa1\x04\x04\x64\x89\x8a\x78\x86\xbe\xa0\x6f\xcd\x00\x19\xe8\x33\x97\xc7\x65\x1c\xb3\xb0\x38\x0e\x5e\x0c\xb8\x3b\xbe\x82\x0f\xbb\x95\x45\xb8\x40\x08\x33\xc6\xcf\x05\x23\x42\x72\xcd\x09\x67\x06\x46\x13\x11\xb8\x93\x5c\x6a\xe5\xb0\xbf\x05\xbb\x55\xf0\x09\x05\xdb\xed\xb3\x0d\xfc\x6e\x00\x1c\x1b\x91\xc4\xd9\x1e\x94\x35\x5a\x2d\xed\xff\xa7\x55\xf0\x62\x0c\x34\x96\x7b\xd0\x91\xc6\x7b\xb7\x7c\x77\xef\xbc\x78\xd3\x50\xef\x8f\x00\x05\xb7\x0e\xa9\xe4\xa2\x23\x0b\xfe\xec\x16\xb0\x29\x97\x67\x2c\x13\x9a\xed\x23\x99\x33\x70\xf0\x07\xad\x45\x78\x5b\x09\xdd\xca\x88\xbc\x1b\x47\xc3\x32\x15\xe5\x7e\x67\xb7\x7c\xc9\x33\xea\x2b\x83\x22\x0d\x26\xa4\x1b\x08\xcb\x72\xe7\x2c\x2e\xba\x5c\x01\x4b\x23\x46\xb3\xa3\xc5\x33\x89\x77\x69\x35\x78\xbb\xd5\x7d\xfc\xa8\xd9\x04\xa9\xff\x80\x21\x55\xa7\x48\x8d\xe3\xc8\xf4\x85\x97\xa4\x56\x0e\x2a\xf5\x53\x46\x68\xf1\xd2\x26\xc6\xda\x3b\x63\x3b\x34\x14\x91\x54\x68\x6a\xa7\x46\x20\x01\x33\x76\x41\x18\x31\x8e\x13\x14\x63\x86\x33\x02\x12\xc5\xb9\x46\x8c\x2a\x0d\x09\xc2\x0a\xe1\x0c\x19\x10\x74\x05\xc9\x25\x8b\x5e\xb1\xe8\xe5\xa6\x58\xaf\x11\x92\x4b\x16\x9a\x73\x55\x4a\x46\x5e\xbd\x6a\x5e\xbe\xf2\x5c\x7f\x3f\x09\xaa\x9b\x85\xd2\x61\x0a\x15\xaa\x9b\x8b\xbb\x09\x41\xa8\x21\x46\x7a\x86\x60\xc3\xca\xe0\x9a\x3f\xab\x58\xfe\xb9\xd7\x52\x49\x41\x01\x71\x23\x34\x12\x12\x52\xfa\xbd\xc5\x65\x47\x15\xe5\x0a\xa4\x61\xe4\x44\x13\x48\xcc\x25\xa0\x42\x43\xa1\x23\x5c\xd0\x93\x3d\x53\x89\x86\x04\xa6\xd2\x36\xc4\x4d\x69\xb9\x30\x29\x65\xf0\x93\x89\xe5\xd1\x64\x3f\xdb\x1d\x54\xe1\xbc\xae\xce\x9c\xd1\x14\xc8\x85\x30\x40\x6f\x8b\xff\x11\x09\x06\x2b\x86\x94\x4b\x88\x12\x50\x5a\x72\xb3\x01\x2d\x73\xb0\x37\x2a\x1f\x73\x45\x2a\x1b\xc5\x58\x24\xd3\x7f\xcf\x28\x26\xb8\xe5\x2f\xc5\x39\xd3\xe5\x4d\xec\x4e\x91\x38\xb6\xa5\x0e\x80\x99\x3e\x44\xe4\x00\xe4\xe8\xf6\x2f\xf2\x98\x51\x12\xba\x85\xb0\x58\xe8\xec\xa8\x9e\x91\xeb\x00\xec\x35\xd9\x39\x55\x0d\x61\xa8\x76\x43\xaf\x8d\xb4\x5b\xed\x56\x66\x55\xc2\x3f\x39\x28\x1d\x09\xac\x0f\x95\x38\x4f\x0e\x27\x18\xcc\x46\x2b\xe8\x63\xae\xab\x9c\xd6\x3d\x1b\x1f\xde\xf7\x48\xe9\x67\x6a\xc2\xb7\xc7\xce\x22\xaa\x3a\x7c\x0c\x19\xe8\x84\xe0\x6e\xe5\xd3\x81\xeb\xe7\xd5\x72\xb3\x5e\x5b\x2d\xb8\xd9\x18\xfb\xe7\x5f\x96\xeb\xdf\xdc\x89\xf5\x67\xeb\x5a\x15\x87\xe8\x81\xf2\xb0\xfd\xf8\x53\x44\x12\x9c\xb3\xa1\x87\xb9\x8a\x69\xfd\x31\xe8\xf6\xec\xd6\x5b\x0a\x35\xdd\x79\xf5\x1c\x98\x22\x37\xbb\x09\x65\xd6\x05\xde\x5f\x63\x57\xeb\x8f\xf3\xb0\xb1\xd9\x6c\x36\xb7\xfa\x1a\x7c\x8c\x18\xc8\x9a\xff\xee\x59\xab\x8e\x99\xa9\x33\x4d\x00\x4a\x51\x9e\x45\x38\x4d\x69\x46\xb5\xbd\x07\x7e\xfd\xf3\xeb\x1f\x03\x79\xed\x12\xcd\xfd\xe9\x1d\xda\x47\x4d\xe8\x4e\x2b\xf0\x5e\x75\x6b\x60\x6c\x3e\x9c\x16\xaf\x26\xef\xef\xdf\xff\x6a\x28\xf4\xc7\xbd\x57\x98\xdf\xb4\x95\xf7\x0b\x23\xba\xb6\xde\x59\x37\xdf\x51\xad\x55\x31\xff\x08\x6d\xb5\x5e\x6d\xb6\xe1\xf3\xe6\xd7\xcf\xbb\xf9\xcd\xd5\x62\xd7\xdf\x5d\xb5\xa1\xd8\xc9\xee\x10\xaf\x33\x14\x83\x27\x8b\x5e\x01\x54\x4f\x67\x9f\x66\xb8\x53\x31\xb4\xe6\xcd\x2c\x56\xbc\x13\xc7\x08\xb8\x0a\x29\x36\xb1\xb6\x1a\xda\xd9\x6d\x31\xd8\x99\xe3\x4f\x0b\x84\xfc\x79\xee\x1c\x64\xbe\x3c\x0c\xf3\x3f\x71\x94\x55\x36\xed\x9d\x65\x95\x26\x78\xc4\x44\x1b\xf1\x36\x73\xfe\x28\x3b\xab\xc9\xc2\xe3\x3c\xf0\xae\xcb\x18\x4c\xab\xcf\x51\x88\x93\xeb\x71\x64\x29\x76\x88\xff\x51\x73\xa7\xb3\x1e\xcf\xaa\x78\xad\x34\xaa\x1a\xaf\xd6\xd3\x6b\xf1\xac\xfc\x35\x68\x5f\x17\x3d\xa0\xf8\x9a\x6f\xbd\x67\xd1\x31\x89\x8d\x1f\x40\xc6\x6e\xf5\x43\xb8\x68\x7e\x01\x98\xdb\x85\xc5\x97\x80\xf6\xc7\x9b\x06\xb0\xb9\x83\x0e\x01\x97\x62\xe3\x8a\x5a\xf1\x9d\xa0\x53\x9c\x73\x3d\x6f\xed\x9c\x39\xab\xab\xe8\x2f\x75\xc6\x44\x91\x31\xf9\x01\x71\xb7\xdd\x16\xea\x62\xaa\xb8\xa8\x31\x3d\x5a\x56\xb4\xc8\xe8\x63\x62\xd2\xd8\x9b\x80\xfa\x40\xed\xdf\x95\x58\x5f\x72\x27\x77\x63\x41\xb1\xbf\x1f\x4d\x06\xef\xef\xc8\x8e\x2f\x68\xff\x06\x00\x00\xff\xff\x43\xa8\x44\x67\xd9\x1d\x00\x00")

func templatesCf_lbTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesCf_lbTf,
		"templates/cf_lb.tf",
	)
}

func templatesCf_lbTf() (*asset, error) {
	bytes, err := templatesCf_lbTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/cf_lb.tf", size: 7641, mode: os.FileMode(420), modTime: time.Unix(1509726301, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesConcourse_lbTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x93\xcf\x6a\xdc\x30\x10\xc6\xcf\xd5\x53\x0c\x43\x8f\xd5\x12\xdc\x1c\x72\xe9\xa9\xf4\x9a\xf6\xd0\x5b\x09\x42\x6b\x8f\xbd\x22\x8a\x46\x48\xf2\x9a\x12\xfc\xee\x45\xb6\xd6\xeb\x34\xff\x16\x96\x85\x9c\x3c\x8c\x47\xdf\x8c\x7e\xdf\x88\xfb\xe4\xfb\x04\x58\xb3\xab\xb9\x0f\x91\x54\xd2\xa1\xa3\xa4\x3c\xb3\x45\x78\x14\x9f\xf6\xda\xf6\x04\xdf\x00\x3f\x3f\x76\xcc\x9d\x25\x55\xf3\x83\xef\xd3\x93\xca\xcd\x1c\xcb\x29\x76\xfa\x81\x46\x14\xa3\x10\xcf\xd5\xed\x56\x19\x9f\x75\x01\x00\x5e\x97\xd6\x4d\x13\x28\xc6\xcd\x72\x50\x1e\x32\xe5\x3b\xeb\x07\x8a\xdc\x87\x9a\x00\xff\x3b\xdf\x9a\x40\x83\xb6\x16\x01\x0f\xa1\x5c\xb4\xe6\xf6\x79\xca\x3c\xc4\xd4\x7e\xaf\xc3\x86\xdc\x5e\x99\x66\x3c\xd6\x49\xf6\xe4\x30\x97\x52\x1a\x38\xdc\xbf\x38\x69\xf9\xb7\xd9\x6e\xad\x3c\xc4\x05\x80\x00\xd0\xd6\xf2\x50\x6e\xeb\x03\x27\xae\xd9\x66\x99\x54\x7b\x9c\x93\x1c\x52\x9c\xc7\xf8\x83\x37\x57\xf8\x05\xf0\xfa\xfa\x6b\xfe\x54\x55\x55\xe1\x9d\x00\x18\xb3\x50\x61\x9d\x74\x17\xa7\xd2\xe3\x65\xee\xde\x04\x51\x70\xe1\xca\x03\xb9\xe4\x16\x0c\xaf\x33\x78\x1b\xf3\x93\x5d\xc1\xd5\x0e\x9c\xa8\x2d\x00\x22\xc5\x68\xd8\x29\xdd\xb6\xc6\x99\xf4\x37\xd7\xdf\xfe\xbc\xfd\xf1\x8e\xbf\x1c\x06\x1d\x1a\xe3\x3a\x15\x7a\x4b\x08\x18\xe3\x4e\x1e\xb3\x72\xce\xae\x7d\x7e\xc7\xeb\x18\x77\xb8\x70\x5e\x55\x9f\xb8\xf3\x91\x6c\xab\xac\x71\xf7\x63\x56\xc9\xae\xaa\xa0\x5d\x47\x93\xca\x64\xa5\x00\x30\x5e\xad\x97\xe0\xf7\xf7\x5f\x25\x5b\x1c\x79\xb9\xe5\xd9\x6f\xe1\x19\xab\x5d\x4a\x3e\x9e\x45\x6b\x52\xb8\x18\xaf\xfc\x02\x3e\x18\xae\xb3\x69\x5d\x0c\xd6\xcd\xd5\xa5\x59\xfd\x0b\x00\x00\xff\xff\x18\xdb\x2d\x3e\x24\x06\x00\x00")

func templatesConcourse_lbTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesConcourse_lbTf,
		"templates/concourse_lb.tf",
	)
}

func templatesConcourse_lbTf() (*asset, error) {
	bytes, err := templatesConcourse_lbTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/concourse_lb.tf", size: 1572, mode: os.FileMode(420), modTime: time.Unix(1510093133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesJumpboxTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\xcf\xc1\x8a\x83\x30\x10\x06\xe0\xf3\xe6\x29\x86\xb0\x57\x15\x84\x5c\x84\x7d\x96\x90\x35\x83\xb5\x44\x13\x26\x33\x22\x88\xef\x5e\x4a\xad\x56\xe8\xa5\xbd\x86\xfc\xdf\xff\x0f\x61\x8e\x42\x2d\x82\xee\x62\xec\x02\xda\x36\x0e\x49\x18\xad\xf3\x9e\x30\x67\x0d\xfa\x2a\x43\xfa\x8f\x73\xd1\x27\x0d\x8b\x02\x18\xdd\x80\xf0\x07\xfa\x77\x99\x1c\x95\x38\x4e\xb6\xf7\x6b\xf1\xf2\x4b\xad\x4a\x45\xe1\x24\xbc\x87\xad\x50\x78\xa4\x01\x26\x17\x64\x03\xde\x77\x96\x87\x55\x6e\x4f\x6b\x53\xd7\x27\x17\x67\x46\x1a\x5d\xb0\xcf\x55\x5f\xba\x27\xd4\xf7\x84\x2d\x47\x3a\x8e\x5f\xd4\xcf\xce\x5e\x98\x53\x6e\xaa\xea\xb3\xd9\xc6\x18\x73\x2f\xb9\x05\x00\x00\xff\xff\xb2\xf6\x55\xa8\x69\x01\x00\x00")

func templatesJumpboxTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesJumpboxTf,
		"templates/jumpbox.tf",
	)
}

func templatesJumpboxTf() (*asset, error) {
	bytes, err := templatesJumpboxTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/jumpbox.tf", size: 361, mode: os.FileMode(420), modTime: time.Unix(1509726301, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesVarsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xd0\x41\x0a\xc3\x20\x10\x05\xd0\x75\x3c\x85\x0c\x59\xb4\x9b\xde\xa0\x67\x29\x36\x99\xca\x14\x99\x91\x89\x08\xad\x78\xf7\x22\x16\xcc\xa6\xa4\x4b\xf9\x4f\x3e\xf3\xb3\x53\x72\xf7\x80\x16\xa2\xca\x13\x97\x74\xa3\x15\x6c\x31\x53\x7a\x45\xb4\x57\x0b\x5b\x52\x62\x0f\xa6\x1a\x33\xac\xa2\x27\xe1\x63\xf7\x16\xc6\x63\x85\x9c\xff\x6a\x5d\x14\x57\xe4\x44\x2e\x6c\xbf\x70\x54\xc9\xb4\xa2\x5a\xf0\x22\x3e\xf4\xf2\xdd\xbf\xc6\xe7\xf2\xa0\x80\x27\x98\x4b\x76\x7a\xd9\x85\x15\xce\x15\xcc\xf4\x1d\xa2\xd3\x46\xc6\x32\x2d\xee\xb7\x8f\xb4\xbf\x6b\xab\xff\x04\x00\x00\xff\xff\x39\xda\x2a\x22\x4d\x01\x00\x00")

func templatesVarsTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesVarsTf,
		"templates/vars.tf",
	)
}

func templatesVarsTf() (*asset, error) {
	bytes, err := templatesVarsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/vars.tf", size: 333, mode: os.FileMode(420), modTime: time.Unix(1509726301, 0)}
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
	"templates/bosh_director.tf": templatesBosh_directorTf,
	"templates/cf_dns.tf": templatesCf_dnsTf,
	"templates/cf_lb.tf": templatesCf_lbTf,
	"templates/concourse_lb.tf": templatesConcourse_lbTf,
	"templates/jumpbox.tf": templatesJumpboxTf,
	"templates/vars.tf": templatesVarsTf,
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
		"bosh_director.tf": &bintree{templatesBosh_directorTf, map[string]*bintree{}},
		"cf_dns.tf": &bintree{templatesCf_dnsTf, map[string]*bintree{}},
		"cf_lb.tf": &bintree{templatesCf_lbTf, map[string]*bintree{}},
		"concourse_lb.tf": &bintree{templatesConcourse_lbTf, map[string]*bintree{}},
		"jumpbox.tf": &bintree{templatesJumpboxTf, map[string]*bintree{}},
		"vars.tf": &bintree{templatesVarsTf, map[string]*bintree{}},
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

