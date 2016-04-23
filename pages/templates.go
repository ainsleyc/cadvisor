// Copyright 2016 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// generated by build/assets.sh; DO NOT EDIT

// Code generated by go-bindata.
// sources:
// pages/assets/html/containers.html
// DO NOT EDIT!

package pages

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

var _pagesAssetsHtmlContainersHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x5a\xdb\x72\xe3\x36\xd2\xbe\x96\x9f\xa2\x7f\xd6\x7f\x91\xad\x32\xa9\x39\xdd\x6c\x56\x56\x95\x22\xcf\x6c\xb4\xf1\xc8\x53\x96\x9d\x54\x2e\x41\x12\xa2\x30\x86\x08\x06\x00\x2d\x6b\x5d\x7e\xf7\x6d\x00\xa4\xc4\xa3\xe4\x53\x4d\x94\x72\x24\x12\xe8\xaf\xbf\x3e\x01\x0d\x72\x46\xff\xe7\xfb\x27\x00\x53\x91\x6d\x25\x4b\x56\x1a\x3e\xbc\x7b\xff\x09\xfe\x2d\x44\xc2\x29\xcc\xd2\x28\x80\x09\xe7\x70\x65\x86\x14\x5c\x51\x45\xe5\x1d\x8d\x83\x13\x14\xb9\x60\x11\x4d\x15\x8d\x21\x4f\x63\x2a\x41\xaf\x28\x4c\x32\x12\xe1\x57\x31\x72\x0a\xbf\x53\xa9\x98\x48\xe1\x43\xf0\x0e\x7e\x32\x13\xbc\x62\xc8\xfb\xc7\xbf\x10\x61\x2b\x72\x58\x93\x2d\xa4\x42\x43\xae\x28\x42\x30\x05\x4b\x86\x8a\xe9\x7d\x44\x33\x0d\x2c\x85\x48\xac\x33\xce\x48\x1a\x51\xd8\x30\xbd\xb2\x6a\x0a\x90\x00\x21\xfe\x2c\x20\x44\xa8\x09\xce\x26\x38\x3f\xc3\xab\x65\x75\x1e\x10\x6d\xf8\x9a\xcf\x4a\xeb\xec\xe7\xe1\x70\xb3\xd9\x04\xc4\x72\x0d\x84\x4c\x86\xdc\xcd\x53\xc3\x8b\xd9\xf4\xf3\x7c\xf1\xd9\x47\xbe\x46\xe2\x26\xe5\x54\x29\x90\xf4\xaf\x9c\x49\x34\x34\xdc\x02\xc9\x90\x4d\x44\x42\xe4\xc8\xc9\x06\x84\x04\x92\x48\x8a\x63\x5a\x18\xb6\x1b\xc9\x34\x4b\x93\x53\x50\x62\xa9\x37\x44\x52\x44\x89\x99\xd2\x92\x85\xb9\xae\xb9\xaa\xe4\x86\x16\x57\x27\xa0\xb3\x48\x0a\xde\x64\x01\xb3\x85\x07\xbf\x4c\x16\xb3\xc5\x29\x62\xfc\x31\xbb\xfe\xf5\xf2\xe6\x1a\xfe\x98\x5c\x5d\x4d\xe6\xd7\xb3\xcf\x0b\xb8\xbc\x82\xe9\xe5\xfc\x7c\x76\x3d\xbb\x9c\xe3\xd5\x17\x98\xcc\xff\x84\xdf\x66\xf3\xf3\x53\xa0\xe8\x28\x54\x43\xef\x33\x69\xf8\x23\x49\x66\x9c\x68\xe2\x06\xb0\xa0\xb4\x46\x60\x29\x1c\x21\x95\xd1\x88\x2d\x59\x84\x76\xa5\x49\x4e\x12\x0a\x89\xb8\xa3\x32\x45\x73\x20\xa3\x72\xcd\x94\x09\xa5\x42\x7a\x31\xa2\x70\xb6\x66\x9a\x68\x7b\xa7\x65\x54\x70\xe2\xfb\xe3\x93\x93\xd1\x4a\xaf\xf9\x18\x27\x8f\x56\x94\xc4\x63\x1b\x82\x91\x66\x9a\xd3\x71\x34\x89\xef\x98\x42\xcd\x3e\x3c\x3c\x04\xe7\x4c\x65\x9c\x6c\xe7\x64\x4d\x1f\x1f\x47\x43\x37\xc5\x4d\xc7\xec\x84\x0b\xa2\xa9\xd2\x36\x13\x30\x37\x62\xc3\x00\xd6\x2c\x45\xb2\x78\x31\x5d\x2c\xc0\x68\xb3\xb3\x39\x4b\x6f\x31\x5c\xfc\xcc\x53\x7a\x8b\xb1\x5b\x51\xaa\x3d\x58\x49\xba\x3c\xf3\x50\xcf\x95\x10\xfa\xf1\x51\x19\xde\xd1\x30\xc4\x0b\xf4\x3b\xc9\xfc\x8f\xc1\x7b\xfc\x0f\x11\x83\x48\x29\x6f\x7c\xb2\xd7\x7c\x99\x19\x0b\x09\x37\xc6\xad\xe9\x6b\xf5\x58\x90\x1e\x6d\xcf\x41\x8c\x44\x6a\x92\x1d\x6b\xab\x45\xf8\xa0\xab\xfe\x43\xee\xc8\x22\x92\x0c\x0b\x6b\x67\x89\x72\xd7\x4a\x46\x6d\x3d\xdf\xff\xca\xa9\xdc\xfa\x48\xf7\x5d\xf0\xc1\x32\xfe\x8e\xda\x46\x43\x27\xf3\x04\x80\x2e\x17\xf7\x43\xe8\x6d\x46\xcf\x3c\x4d\xef\xf5\xf0\x3b\x32\x75\x77\xbd\x6e\xe4\xc4\xae\x4f\xfe\x77\x45\x32\xd6\x80\x7c\x31\x66\xc5\xad\x6f\x44\x32\x5a\x11\xa9\xdb\x68\xa3\x61\x59\x0f\xa3\x50\xc4\xdb\x42\x41\xcc\xee\x20\xe2\x44\xa9\x33\x6f\xc7\xc4\xe5\x9d\xaf\x56\x62\x13\x11\x5c\x35\x61\x5c\xac\x63\x23\xd2\xcc\x0d\x6f\x2f\xcc\x7d\xb5\xf6\xdf\x7f\xf0\x80\xc5\x67\x1e\x17\x89\xf0\x76\x62\x43\xb2\xfb\x59\xd3\x57\x8a\x8c\x4f\x06\xd5\x81\x0c\x97\x01\xdf\x90\xa5\xd2\x0c\x99\x4a\x7e\x3f\x6e\x17\x2c\xde\x44\xb9\x21\x0a\x9a\x6f\xc1\x4b\xf1\x50\xa2\x68\x24\xf3\x75\xe8\xa4\x1f\x1e\x24\xae\x2d\x14\xfe\x3f\xc3\x95\x31\xd5\xd3\x9d\x99\x3f\x9f\x41\xf0\xad\x7e\x4f\x3d\x3e\x5a\x85\x9c\x8d\x2b\xc6\x36\x25\x83\x0b\xac\x1b\x34\x7e\xdc\x31\x74\x8d\x41\x32\xec\x08\x3a\x1f\x51\x1c\x01\x9a\xc6\x06\x78\x34\x14\x7c\xef\x14\x4b\xdc\x5d\x3c\x3c\xb0\x25\x04\x33\xe5\x9c\x7a\xc4\x57\x50\x7c\x46\xab\x4f\x7b\x92\x41\x30\x8c\x45\x74\x6b\x3c\x76\x6e\xbf\x61\x6f\x93\x23\x83\xb3\xbb\x55\x3b\x72\x55\x22\x8b\x3c\x8c\xaa\x1e\x79\x5d\xec\x3e\x8e\x6b\x78\xc8\xe4\x63\x35\x70\x15\x61\x8e\x5b\x92\x9f\x48\x91\x67\x8d\xc8\xa9\x0a\x80\x0d\x5b\x93\xe1\xa0\x96\x9c\xb5\xf9\x65\xb0\xda\x4a\x7c\xa6\xe9\xda\x06\xb1\x36\x7f\x1f\xc1\x46\xf0\x2a\x5e\xeb\x77\xa1\xf3\xa0\x8b\xc1\x02\x2b\x32\x7f\x0b\x07\x9e\x4b\x86\x9b\x22\x38\xbc\xa6\x03\x73\x7e\xd4\x7f\x2e\x35\x94\x15\xb7\xfe\x6b\xf0\x73\x29\xef\x60\xa0\xc3\x45\x23\x95\x61\x7f\x50\x68\x31\x30\x3e\x27\x21\xe5\xd6\x77\x55\xec\xe0\x37\xba\x35\xae\x33\xd3\xc7\xd0\x1c\xfc\x9d\xf0\xdc\x56\x6e\xb3\x2e\xea\x5e\x73\xc6\xee\xb9\x0d\x5e\x46\x6d\xa1\x85\x44\x5f\x8e\x42\x39\x2e\x08\x19\xa8\x3e\x67\x0d\xf6\xbe\xb2\xea\x5b\xbe\xea\x67\xf5\x5c\x7f\x55\xf0\xdb\xfe\xaa\x0e\xd6\xfd\x35\xd8\xb9\x0b\x43\x9f\x73\x6b\x4d\xe9\xc9\xe2\x46\x5f\xb6\x76\xd5\xb8\xb3\x6a\xb6\x46\x17\x1d\xcf\x50\xd8\x7d\xfa\x53\x15\x2a\x1f\x93\xb3\x0e\xda\x25\x6b\x65\xa4\xca\xcb\xa1\x99\xfd\xc2\xe5\x89\xcf\xac\x8c\xd9\xb7\x6a\xb3\x4c\x08\xf1\xef\xa4\x0b\xa3\xcb\x36\x3c\x30\x88\x5c\x46\x54\x4d\xee\x08\xe3\xa6\x6d\x7e\x83\x1a\x9c\x29\xc1\x6d\xeb\xd9\xa8\x3f\xa7\x72\x9a\xe5\x55\x65\xbd\x89\x56\xf1\x44\x6f\xfe\x00\x89\x34\xa6\x01\x36\xe9\x85\x46\xdf\xf6\xa6\x80\x49\x42\xb9\xfb\xed\x8d\xa7\xdf\x6e\x5c\xf8\xf7\x88\xc5\xe2\x8d\x1d\xb5\xa1\x83\xeb\x1e\x36\xcb\x3b\xc3\x0f\xab\x3c\x54\x47\xd8\x4e\x98\x38\x96\x39\x9a\x49\x96\x6a\x77\xb3\xad\x0c\x6a\x30\x79\xca\x76\x30\xaa\x0a\xd3\x66\x5e\x0d\x62\x87\x2d\x5f\xc9\xfd\x1b\x99\x83\x48\x60\xa1\x1a\x16\x4d\x45\xdd\xa0\xbd\xc6\x7e\x9b\x22\xf1\x2a\x93\xd4\xed\xeb\xcd\xc1\x63\xb2\xd8\x98\x03\x89\x68\x07\xc9\x68\x68\x28\x04\xfc\x7f\xb4\xc2\x6d\x6e\x96\x2e\x45\x30\xcf\xd7\x56\xae\x5c\x63\xda\xec\xcb\xa5\x66\x77\xed\x8c\xf8\x4a\xd7\x42\x6e\x7f\x6c\xc2\x3b\x9d\x07\x72\xde\x4d\x08\xdc\xd3\x02\x0b\xf3\x7a\xf7\x56\xc0\x9a\x15\xc0\xfe\x4b\x0f\x28\xee\x4f\x9a\x42\xfe\x06\x6f\x1d\x90\x7f\x49\x56\x15\x38\x6f\x54\x28\x5d\x45\xd2\x36\xfa\x68\x8d\xf4\x9a\x5b\x48\xbe\xc2\xd0\xc5\x86\x64\x6f\xb5\xc8\x21\x54\xe7\xb2\xd0\xb6\xb8\xa2\xf5\x05\x56\x57\xa4\x8f\x58\xde\x2c\xbd\xc2\xba\x5a\x17\xfa\xe2\xcd\xec\x46\x99\xd6\xa8\xbf\x13\xb7\x95\x57\xd4\x1f\x5a\xb2\x26\x72\x7b\xa0\x0d\x30\xb3\x8c\x06\x96\x26\xed\x46\xa0\x3e\xad\x28\xe6\x4b\xec\x72\xee\x18\xdd\x1c\x6e\x0f\xa0\xd2\x21\xe4\x86\xb1\x9f\x90\x3c\xa1\x5e\x1d\xd2\x9c\x66\x77\x2d\xc3\xdf\x62\xcd\x37\x29\xb0\xd9\x50\xc7\xba\x9d\xaa\x39\x59\x29\xe2\x6b\x91\x3d\xc9\xa0\x9e\x3e\xe3\x07\x9a\x69\x5b\x8e\xa7\x18\xd8\x61\x4d\x43\xc1\xa7\xf1\xb5\xd0\x84\x43\x99\x87\x9f\x6c\x66\x56\xfc\x13\x65\x39\x7a\x06\xa7\xf8\x2e\xf0\xf6\xa1\xc6\xde\x29\x76\xae\x79\xf4\x64\xa0\x90\x17\x5c\x08\x12\xc3\x04\xb3\xea\x00\x1e\xc7\x39\x75\xa0\xdd\x13\xa9\x2a\x33\xcb\xc9\x3c\x7c\xb4\x9b\x6a\x1f\x18\x8e\xfb\x66\xff\xef\xe4\xd7\x0d\xf9\x8b\xa4\xe4\x36\x16\x9b\xb4\x0f\xd3\x41\x85\xe5\xb4\x5e\xd0\x76\x6a\x1c\xdd\x9d\x7f\x60\x9a\x94\x1b\xf5\x0f\xca\x94\xb5\x55\x77\x3c\x0c\xa1\x1c\x36\xee\x54\x08\x48\xb1\x81\xee\x03\xcf\xc1\x10\x36\xa6\xb5\x97\xe3\x7f\xda\xb3\x65\xcd\x54\x29\x12\xf3\x7c\xbc\xa5\xa4\xe5\x93\x62\xa2\x1f\x12\x09\xd5\x0b\x3f\x36\x07\x55\xe9\x95\xeb\x88\x1b\x58\x09\xed\x3b\x57\x74\x22\x43\x7d\xaf\x52\xd2\x17\x29\xc7\xa9\xbf\x0a\x0d\x65\xc0\xdc\x21\xb9\x43\xb2\xed\xcd\xe7\xd0\x65\xd8\x6a\x36\xc8\xa2\x77\xe2\x97\xb0\x9d\xa2\xdc\x53\xe9\x0e\x06\x9d\xbc\xbb\x6f\xb6\x23\xf7\xd1\xab\x66\x97\x79\xf4\xda\x58\x7d\x9e\x59\x94\x73\xaa\x37\x42\xde\x3e\xb3\x2a\x07\xaf\x2f\xc7\x42\x71\xb1\xd9\x3f\xa7\x10\x07\xcd\xd1\x58\x8a\xcc\x24\x7f\xbb\x40\xc2\x5c\x6b\xb1\x8b\x57\xa8\x53\xc0\x3f\x3f\xa6\x4b\x92\x73\x0d\xa5\x1c\xae\xe8\x49\x82\x9c\x8a\xe7\xd9\x4e\xc8\xf9\x39\x75\x2c\x7d\x45\x39\x8d\xec\x11\x60\xa7\x0c\x62\xa2\x49\x21\x5a\xe1\x00\x44\x32\xe2\xaf\x88\xca\x44\x96\x67\x67\x9e\x96\x39\x2d\x6e\xd2\x7b\xb4\x23\xa6\x08\xbb\x24\x5c\xd1\x8e\x14\x73\xe9\xd5\xad\xb8\x8c\x75\x77\x7e\xd5\x12\x33\xc2\x33\x6d\x65\xee\xa0\xcc\x04\x67\x59\xcb\x4b\x78\x44\xea\x54\xe9\x35\x1d\x8c\xb5\x91\xe6\x1e\x48\x61\x2c\x76\xbf\xad\x61\xb6\xbb\xe4\x34\x0e\xb7\x07\x3d\xd6\xce\xf9\xe2\xf1\xd0\x81\xb4\x7d\xce\x82\xbc\xc2\x96\x3a\x59\x65\xb9\x6e\xaf\x82\xbb\x65\xb9\xa4\x17\x6e\x35\x36\x39\xad\xed\xfb\x05\x6a\x3f\x4b\x29\xec\xe3\xe3\xd6\x16\x50\xea\xa2\x76\x46\xbf\xb2\x86\xf1\x8d\x0a\xfd\xa2\xfe\xb6\x2d\xf3\x0b\xe3\x54\x6d\x15\x9e\x51\x9e\xde\x41\x2e\x77\x32\x6e\xef\xeb\x6c\x22\xfb\x91\x7a\x96\xa9\x69\xae\xb4\x58\x7f\xa5\x5a\xb2\xe8\xb9\xfe\x38\xb2\x58\x0d\x0e\x79\x60\xe2\xde\x70\x9b\x3c\x86\x42\x7b\x73\xc5\x3a\x94\x2b\x8d\x5e\xca\x1a\x81\x45\x64\x71\x8e\xe6\xc3\xa1\x23\x56\x75\xa8\x32\x70\xf8\xc5\x5c\x29\x8c\x27\x4c\xa9\xbf\x61\x6c\x7e\x7a\x78\x08\x76\xef\x63\xdc\xfb\xab\x53\xf3\x0a\xba\xde\xcd\xdb\x5b\xad\xe6\xcd\xde\x75\x2f\x86\xec\xcf\xf2\x2d\x91\xfd\xb7\x0c\xe6\x13\x4b\xb2\x71\x0f\x5b\x8d\x9a\xfa\x73\xdd\x62\x52\xfd\x3d\xa0\x7b\xfd\x87\xee\xb5\x2f\xca\xff\x17\x00\x00\xff\xff\xb4\xe6\x8c\x7e\x8b\x21\x00\x00")

func pagesAssetsHtmlContainersHtmlBytes() ([]byte, error) {
	return bindataRead(
		_pagesAssetsHtmlContainersHtml,
		"pages/assets/html/containers.html",
	)
}

func pagesAssetsHtmlContainersHtml() (*asset, error) {
	bytes, err := pagesAssetsHtmlContainersHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pages/assets/html/containers.html", size: 8587, mode: os.FileMode(416), modTime: time.Unix(1461437445, 0)}
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
	"pages/assets/html/containers.html": pagesAssetsHtmlContainersHtml,
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
	"pages": {nil, map[string]*bintree{
		"assets": {nil, map[string]*bintree{
			"html": {nil, map[string]*bintree{
				"containers.html": {pagesAssetsHtmlContainersHtml, map[string]*bintree{}},
			}},
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
