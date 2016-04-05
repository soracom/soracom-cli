// Code generated by go-bindata.
// sources:
// ../generators/assets/i18n/soracom-api.text.en.json
// ../generators/assets/i18n/soracom-api.text.ja.json
// DO NOT EDIT!

package cmd

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

var _GeneratorsAssetsI18nSoracomApiTextEnJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x59\xdf\x6f\xdb\xb6\x13\x7f\xcf\x5f\x71\x30\x50\xa4\x2d\xac\xa4\x69\xfb\xf4\x7d\xf8\x6e\x5e\x31\x0c\x5e\x7f\x24\x68\xda\x0d\x03\x02\x38\xb4\x44\xc7\x5c\x65\x52\x23\xa5\xb8\xc6\xd0\xff\x7d\x77\x14\x29\x53\xb2\x64\xc9\x4e\x0a\xf4\xcd\xa2\x78\x9f\x3b\xde\xdd\xe7\xee\x28\xff\x7b\x02\x30\x62\x45\xbe\x1c\xfd\x0f\xe8\x37\x3e\x65\xca\xe4\xd5\x13\x3e\x9b\x62\xb5\x62\x7a\x83\x4b\xa3\x2b\xae\x17\x4a\xaf\x0c\x90\x04\x97\xb9\x88\x59\x2e\x94\x84\x5c\x01\x8b\x63\x6e\x0c\xfd\xc2\x37\x70\x7d\xf9\x71\xf2\xe6\xf2\x3d\x4c\xae\xa6\x67\xa3\xb1\x47\x4a\xb8\x89\xb5\xc8\x48\xe4\x68\x34\xf8\xa4\x20\x2b\x05\x9b\x72\xf3\x0d\x30\xd0\x4a\xe5\x24\xae\x0a\x99\x8f\xc1\x64\x3c\x16\x8b\x0d\xdc\xf2\x15\x13\xe9\x2d\x30\x99\xc0\x6d\xc6\x8c\x59\x2b\x9d\xdc\xf6\x81\x49\x98\xe0\xe2\x5b\xbe\x09\x80\x58\xb9\x32\x4d\x1c\x98\x7b\xee\xc5\x82\xeb\xc9\x7b\x28\x0c\xd7\x01\x96\xc2\xfd\x2c\x57\x1a\xc1\xc6\x70\x4b\x2f\x3f\xb0\x15\xdf\xb5\x72\x22\xe9\xe8\x80\x6a\xec\x2b\x56\x3e\x7e\x52\x5f\xb8\x84\xb5\x48\x53\x98\x73\x10\x32\x4e\x8b\x84\x27\xf8\xc3\xfa\x4c\x73\x93\x29\x69\xf0\xc5\x02\x4c\x61\xdd\xb9\x28\xd2\x33\xb8\x76\xca\x69\x4f\x08\xea\x9f\x4b\x54\x74\xbc\xe6\xff\x14\xdc\xe4\x18\x9d\x45\xce\xf5\x9a\xe9\xc4\x6c\x23\x99\x31\x8d\x96\xe2\xba\x09\x32\x05\xd7\xe7\x2a\xd9\xd4\x56\x76\xa3\x3e\xa9\xfb\xc6\xa9\x19\x55\x12\xdf\xdc\xaf\x6f\x95\x32\x7f\x94\x86\x2e\x93\xb3\xbc\x30\x2f\x5f\xbc\x38\x4c\xe1\xd6\x19\x81\xce\x71\x13\xf6\xf5\x8b\x8b\x1e\xd8\xcf\x92\xcd\x53\x6e\x93\x75\xab\x80\xef\x3b\x47\x9c\x8a\xc6\x09\xfa\x99\x45\xa1\x11\xc6\x14\xdc\x47\xbd\x35\x09\xea\x31\x1e\x05\xa7\x79\x0c\xca\xdd\xc8\xe3\x38\x17\x45\x21\xeb\xa2\x68\x9b\xd1\x7d\x88\x6d\xc4\x8b\x22\xda\x18\x7d\xe1\x9b\x48\x24\x15\xa6\x5f\xeb\xc7\x6c\x23\x60\x14\x79\x0a\x12\xe6\x98\x16\x68\x47\x24\xb7\x34\x0c\xcd\x1e\x44\x44\xb3\x54\xeb\x9d\x88\x9c\x84\x39\xe1\x32\x62\x64\xe3\x3a\xf3\xf0\x33\x4c\x73\x9e\xcf\x72\x02\x0b\xcb\x6f\xa3\x1c\xd7\xd3\x66\x4a\x10\x18\x4d\xf0\x28\x60\x51\xc0\xa2\x00\xba\xc2\x06\xd3\x1f\x72\x5f\x66\xfc\xc6\x25\x6d\xea\x06\xa3\x03\x1b\x4e\xe9\x98\xfb\x24\xf1\xb8\xa7\x06\x28\xd0\xc0\x92\x04\x25\x0c\xfa\x89\x4a\x06\x4a\xc7\x5c\xdc\x0b\x79\x67\x37\xb7\x81\x8e\x21\x66\xe8\xb5\xf3\xfb\x8b\x73\x8a\xd8\x79\x9b\x2f\xce\xef\xb9\xa6\x60\x91\xa7\xd7\x22\x5f\x5a\xb0\xdc\x57\xa9\x22\x4b\xd0\xe8\xd0\x12\x8f\x51\x3b\x6c\x7b\x09\xd9\x53\x44\x76\x1d\x74\xf9\x76\x14\xbc\x0f\x2a\x46\x50\x33\xfa\x51\xa6\xf2\x9e\xa5\x22\x01\x5e\x73\x58\x0d\xb9\xb5\x2a\x35\x6b\xc7\x63\xa5\xc1\x0f\x93\x08\xba\x90\x70\x6a\x94\x66\xb1\x2a\xe9\x0b\x65\xdc\x2b\xfa\x45\x76\x7b\x64\xb7\x9f\x1e\x98\x09\xbb\x45\xb9\x46\xc4\x52\xd3\xc3\x98\xf8\x07\x61\x08\xf4\x5a\xd7\x11\xad\xdf\x4a\x23\x3b\x92\xb4\xd9\x60\xdc\xde\x86\x83\x2b\x6c\x2c\x31\x9d\xba\x84\x29\xdd\x27\xf8\x0f\xce\x03\x6b\xee\x59\x3f\xe2\xeb\x5e\xc4\xb2\x0a\xf3\xaf\x99\xd0\x3c\x79\x30\xa3\x1e\x1e\xce\xef\x18\xd0\xae\x74\x3e\x71\x27\x1c\xcd\xb1\x17\x6d\x43\xdc\x38\x6a\x78\xcc\x6b\x6c\x56\x80\xe5\x01\xdd\xa6\x74\x0e\x24\x47\x34\x15\x72\xa1\xba\xe7\x76\x2b\x14\x6e\xb5\x15\xa6\x6c\xaa\x68\x1f\x35\x61\xa1\x92\x71\x80\x4b\xc7\xaa\x09\xd0\xbc\x01\x0b\x91\xba\x79\x69\x6b\x39\xb2\x3f\x4b\xb9\xd3\xd4\x6b\xbe\xaf\x55\x50\xda\x67\xed\x98\x33\xb3\x84\x00\x66\x47\x81\x5c\x88\xbb\x42\xf3\xa1\xee\xc1\x1e\xa5\xb9\xad\x2c\xda\xd7\x18\x8f\x61\x87\x0b\xb3\xdf\x51\x87\x88\x67\x5a\x59\x9f\xd4\x6a\x0c\x2e\xae\x32\xaa\x3b\xa3\x28\x8a\xaa\xb9\xec\xcd\xbb\x29\x16\xe1\xbc\xc8\x00\x57\x71\xf4\x59\x62\x82\xd8\x09\xc4\x69\x63\x90\x20\x0f\x62\xcc\xb0\x0d\x3c\x31\x94\x5a\x58\xad\x13\x85\xd9\x27\x71\x56\xe3\x5f\x85\xc9\x61\x83\xd9\x45\x59\x9c\xa5\x2c\xe6\x70\xfa\xc4\x9c\xfd\x6d\x14\xd6\x56\xbc\x41\x88\x1c\xe7\xa9\x1b\x79\x95\x72\x86\x57\x08\xc3\x53\x84\x82\xf5\x52\xc4\xcb\xe6\x68\x85\xb7\x80\xa5\x4a\x6c\x01\x36\xdc\x0a\x5d\x9c\xc1\x54\x66\x45\xee\x67\xb8\x69\x62\xb5\xb8\x27\x78\x0e\x1f\x39\x86\x67\x85\x1d\x04\x73\xe5\x39\xdc\xc8\x97\x5e\xe0\xd2\x91\x82\x0e\x91\x90\x0e\x96\x9a\xa7\xd5\xe2\xaf\x65\x6b\x41\xa8\x2b\xc7\x90\x67\x37\xf2\x95\x97\xa5\xe9\xae\x4d\x6c\x8a\x89\xf8\x19\x67\x3a\xa0\x99\xae\x29\x7d\x23\xc3\xe2\x58\x1e\x93\x3c\xed\x0e\xfc\xf4\x22\x7a\xf5\x0c\xfe\x0f\xe1\x26\x85\x34\x5c\x6b\x91\x73\x3b\x4e\x97\x01\x23\x0f\xb3\x14\x1d\x9f\x6c\x4a\xd7\x62\xe7\xbb\xf4\xfb\xd0\x97\x3f\xc1\xd3\xbf\xce\xe5\x33\x68\x1f\x03\xef\x78\xed\xba\xbd\xef\x96\x60\x99\x17\x10\xcd\xa9\xaf\x27\xd4\x9e\x96\x32\x4c\xbe\xa3\xa8\xe0\x78\x4f\xbd\x63\x08\x6f\xde\x51\x7a\x29\x6a\xfc\x77\xf8\x0b\x9d\xef\x64\xbb\xa9\xd2\x2d\xd1\xb4\x62\x1b\xe3\xa1\x96\x78\x0a\x8e\xab\x19\x41\x43\x82\x11\x26\x2a\x6e\xd1\x88\x4e\x3d\x06\x1e\x04\xd4\xb0\x9b\xdf\xe3\xeb\xd9\x12\x13\x30\x0d\x6f\xcd\xc7\x9b\x6e\x01\xc1\x03\x1e\x6d\x78\x13\xa6\x61\xf6\x9d\xc6\x38\x3c\x86\xb9\x25\xd0\xd1\x66\x7a\xf1\x86\x79\xbe\x8d\x0e\x31\xf0\x3d\x93\xec\x6e\xdb\x79\xf7\xd8\xb2\xbb\xb3\x46\x58\xaa\x80\x33\xbc\x73\x9a\xa1\xb4\x0d\xcf\x16\x9c\x29\x98\x01\x5c\x71\x34\xfb\xe6\xc1\x43\x50\x3a\x28\x8c\x55\xef\x88\xe4\x23\x65\x31\x93\x31\x4f\xa1\x04\x18\x16\xc5\x16\xb1\x86\x39\x19\xdb\xac\x2c\x2b\x50\x4a\xe9\xcd\xe0\xca\xe2\xe4\xc0\xc9\xf5\x54\x95\x9d\xdd\x1d\x56\x94\x9d\x6c\x90\x77\xde\x34\x5b\xba\xd7\xe1\x30\xba\x2d\xea\x97\xac\xa5\xda\x9a\xcf\x71\xc3\xd0\x3c\xdb\x01\xff\x93\xcf\xaf\xd8\xa6\xa1\x63\x5f\x8e\x0d\x45\xe8\xc8\x2f\x6c\x28\x49\x11\xe7\xc3\x7b\x84\x17\xe8\x0b\x61\xb5\xad\xa1\x50\xab\x94\x3f\x46\x75\xb2\x38\x47\x17\x27\x27\xdd\xb0\xcd\x2c\x45\x96\xe1\xc4\x3b\x73\xb7\xe1\x47\x31\xd4\x83\x42\x05\x7a\xb4\xd5\x6d\x50\xcd\x23\xe0\xed\xeb\x98\x7b\x04\xc9\xa1\x72\x11\xf7\x0c\xc7\x1d\x22\xf5\x6a\x2b\xf4\x41\xe3\xd1\x16\x72\x22\x74\xab\x25\xbd\xb6\x34\x05\x5b\xc7\xb6\x39\x67\xab\x23\x0d\xfb\x05\x45\x8f\xb3\x6c\x47\xb2\x83\x88\xa6\x98\x13\xcc\x7c\x60\xb5\x77\x0d\x2f\x90\xea\x6d\x8e\xb5\xbd\x0d\xed\xf4\x0d\xf5\x10\xbd\xfe\xcb\xec\x1e\xad\xdf\xa7\x03\x7b\xc5\x0f\xeb\xc0\xad\x28\xad\x39\xe3\xaf\xf6\x87\x15\xf4\x16\xee\x06\x2a\x87\x7c\x4a\x3a\x0c\xa7\xdd\x74\xae\x57\xc2\x18\x11\xde\x00\x06\x64\x7d\xe5\xae\xed\xe7\x92\x9a\xd2\x0a\xb5\x8f\x04\x07\x00\x75\x70\x02\x2f\x64\x66\xe0\x77\x05\xcb\x3a\xb7\x7f\xc8\xc7\x90\xfa\xd6\xe6\x8d\x25\x50\x43\xdf\x37\x56\x4a\xce\xec\xff\x67\x35\x47\xee\xfc\x73\x86\xfd\x03\x19\x3f\x5b\xf2\x34\x23\x3d\xbf\x5f\x5f\x7e\x40\xe6\x6b\x2a\xd8\x18\xbf\x9f\xe9\xe2\x66\xaf\xb4\xf4\xcd\x83\xbe\x8d\xbb\x3f\xd2\x80\x80\xda\x7c\x70\xf2\xed\xe4\xbf\x00\x00\x00\xff\xff\xd7\x7b\xd8\xcf\xed\x1d\x00\x00")

func GeneratorsAssetsI18nSoracomApiTextEnJsonBytes() ([]byte, error) {
	return bindataRead(
		_GeneratorsAssetsI18nSoracomApiTextEnJson,
		"../generators/assets/i18n/soracom-api.text.en.json",
	)
}

func GeneratorsAssetsI18nSoracomApiTextEnJson() (*asset, error) {
	bytes, err := GeneratorsAssetsI18nSoracomApiTextEnJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../generators/assets/i18n/soracom-api.text.en.json", size: 7661, mode: os.FileMode(420), modTime: time.Unix(1459775849, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _GeneratorsAssetsI18nSoracomApiTextJaJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe4\x5a\x5f\x6f\xdb\xd6\x15\x7f\xf7\xa7\x20\x04\x14\x4d\x00\x2b\x4e\xda\x3e\xed\x61\x98\xdb\xbd\x78\x5d\x9a\x62\x79\x18\x06\x18\x30\x28\xeb\x26\xe6\x22\x91\x02\x29\x25\x10\x86\x00\x26\x99\xcc\x8a\x27\x37\xfe\x17\xcb\xaa\xdd\xca\x7f\x65\xd5\xaa\x24\xbb\x4e\x57\xc5\xf6\xe2\x0f\x73\x4d\x5a\x7e\xea\x57\xd8\xb9\x97\xa4\xf8\x4f\xb2\x48\xc9\x0e\x36\xf4\xc5\x10\xe5\x7b\xcf\x39\xbf\x73\xcf\xf9\x9d\x73\x0f\xf5\x8f\x21\x86\x89\xb0\x99\xf4\x54\xe4\x77\x0c\xf9\x0c\x4f\x29\x41\x4a\xb7\x9f\xe0\x59\xca\x24\x93\xac\x98\x85\xaf\x22\xa3\x5f\x8f\x31\x58\xd9\xc2\x4a\x03\x2b\x27\x58\x79\x87\xe5\x7a\x6b\x7f\xae\x55\x39\xc5\xca\x62\x6b\x33\x8f\xe5\x17\x78\x5a\x36\x16\xd5\xb0\x7a\x8a\xe5\x0a\x43\x9f\xd4\x1c\x79\x82\x5d\xea\x11\xac\xbc\x28\x1e\xd3\xc5\x05\x2c\xbf\xc7\x72\x11\x4f\x2b\x91\x61\x4b\x59\x1c\x49\x93\x22\x97\x4a\x73\x02\x1f\x50\xa1\x25\x03\xab\x55\xa2\x04\x54\x91\xf5\x55\xac\x94\x89\x36\x78\x94\xf7\xcc\x2d\xb0\x4c\xf9\x97\xb6\xf1\x56\x9b\x87\x2f\x1b\x0c\x4a\xb2\x5c\x82\x21\x36\xa6\x58\x49\x7a\x26\x88\x71\x30\xde\x92\x6e\x98\xbf\x67\xeb\x93\xeb\xf6\x4e\xe2\xaf\x2f\x51\x76\x2c\x4e\x77\x9b\x4f\xb0\xf9\xe1\xe8\x7d\xc0\x5a\xa6\x58\xff\x0d\x7f\xfd\x5b\x85\x14\x12\xd9\xb4\x20\x8e\xc5\x87\x99\x8c\x84\xc4\xaf\xd8\x24\x1a\x6e\xeb\x07\xa8\x8b\x58\xdd\xa7\x68\x2b\x04\x2d\xb1\xbe\xaa\xcd\x03\x18\x19\xcb\x65\x2c\xbf\xc6\xf2\x06\x96\xdf\x50\x3f\x2b\x96\xf4\xbc\x9e\x9b\xd7\x66\x4b\xd4\x9f\x25\x53\x53\xcf\x53\x90\xf3\xad\xb3\x65\x22\x4a\xc9\x83\x53\xc0\x44\x02\x76\x5a\x3e\x3f\xd9\xbd\x2c\xce\x39\x50\xe7\xb5\xb3\x97\xad\x32\x68\xdf\xf7\x18\xa6\x17\x15\xb0\x0d\x20\x61\x79\x09\xd6\x33\x3d\x4f\x1d\xab\xab\x58\x55\xb1\x3a\x4d\x97\x54\xcf\x4f\x56\xcf\x9b\xdf\x50\xa3\xcb\x5e\xcc\xe6\xd9\x6e\xfb\x31\xdb\x81\x92\x62\x45\xf0\x5d\x1a\x89\x92\x23\x56\xe1\xfb\x98\x10\xcf\xba\xbe\xf1\x07\x95\x89\xcd\xad\x34\xd2\xde\xf0\xdc\xfc\xf4\xbc\xad\x4b\x44\x52\x4a\xe0\x25\xe4\x51\x25\xa5\xd9\x74\x46\xfa\xe4\xee\xdd\x40\xfa\x8c\x53\x72\xa8\x19\xf6\x4a\xfa\xec\xee\xbd\x40\x92\xb4\x9d\x43\xfd\x4d\xe1\x2a\x83\x27\x13\x9c\xc7\xd4\x0f\x98\xc3\x03\x65\xf1\x38\x1f\x3a\x8f\xa3\x51\x3b\x93\xa3\x51\x67\x2e\xe9\xf9\x19\xad\xfe\xad\x19\x63\xee\x40\x1a\xe7\x5d\x99\x0e\x80\x96\x2b\xd4\x8e\x72\x37\x25\x24\xcd\xa3\x4f\x50\x36\xca\xc5\x4d\x55\xd6\x37\x3d\x55\x79\x79\xe1\x0a\x28\x16\x41\x80\x96\x61\x78\x24\x1c\x11\xe5\x29\x49\xb8\xa1\x0d\x40\x13\x45\x9a\xf0\x15\xf7\x29\x2b\x2f\x3a\x70\x44\x5d\x57\x5f\x6a\x1b\x87\x84\x2c\x36\x2b\x17\x3b\xc7\x26\x5f\xb4\x4f\x7c\xc8\x19\x7f\x66\xf4\x45\x38\x49\xca\xa0\x09\xcb\xd6\x09\xc8\x1d\x94\x9e\x48\x0b\x4f\x10\xef\x2c\x2a\x9e\x22\xe3\x0e\x51\xac\x2e\x50\x58\x0d\x1a\x08\xaf\x28\xd6\x13\x4a\x1e\x39\x7a\x4c\xf5\x41\x82\xd1\x2b\xdc\xf4\xa5\x29\xdf\x72\x4f\x89\xfa\xb2\x4e\x97\x81\x8a\x33\xac\xec\x60\x75\xc3\xab\x77\xb9\x04\x4e\x35\x0f\x7d\x5a\xc6\x0a\x90\xe4\xb7\x58\xfd\x91\xae\x39\xa3\x07\x0d\x12\x36\xa9\xa2\x2a\x09\x67\x82\xe5\x47\x9a\x00\xd5\x4b\x42\xb6\xab\x4e\x9b\xed\x95\x72\xb5\x55\x59\x6d\x9d\xbe\xb2\xfc\x5d\x26\x27\x09\x56\xb9\xb5\x6b\xaf\x57\xb4\xf7\x05\x22\x61\x5a\x1e\x79\x7a\x6f\x84\x84\xe3\x48\x27\xbf\x8f\x3c\x45\x22\xf7\x28\x6b\x52\xf4\xa2\xb6\x00\x76\x1d\x69\x33\xc7\x86\xe1\xfa\xca\x01\xfd\xf0\xc2\xef\x98\x56\xa5\x66\x05\x75\x07\xbf\x76\x66\xc5\x2b\x78\xd1\x7f\x16\x0f\xbe\x8c\x38\xfe\xef\x60\x44\x07\x27\xf6\x96\xd2\xc5\xc5\x79\xbd\xb6\x4d\x6d\x87\x74\x50\xb0\x32\x4b\x41\xac\x61\x65\xc9\x11\xba\x4e\xfa\x74\x53\xb2\x97\x42\x6f\x30\x42\x7f\x23\x31\x2a\x09\x22\x3b\x29\x24\x69\xbf\xc4\x18\x21\xd9\xa6\xb4\x28\x0d\xd7\x28\x0d\x57\x88\x51\xe0\xfb\xef\x29\xeb\x83\xf8\xbd\x60\x01\xea\xeb\x14\xfc\xe5\xd1\x45\x53\x86\x01\x1f\x96\xa7\xf4\x9d\x75\x4a\xc5\xc4\x1b\x7e\x30\xfa\xda\x5b\x0b\x6a\x78\x16\x0b\xa3\xda\xfa\x5c\xed\xd0\x38\x76\x0e\x8f\x70\xa6\xfe\xcf\x10\xc3\x79\x73\x8e\x72\xc0\xbe\xa7\xf5\xb5\x8a\x75\xbb\x9a\x11\x07\x78\x59\xa1\xa3\xce\xcf\x02\x90\x91\xbb\x82\xae\xbf\xd2\x66\xdf\xe9\xeb\xa5\xcb\xe2\x3c\xe9\xa6\x73\x33\x76\xb6\x78\xeb\xa8\x33\x58\x6f\x92\x8c\x06\x08\xc3\xff\x83\x40\xec\x96\xf7\x43\xa6\x4b\x23\x31\x2e\x91\xb0\x43\xd3\xe3\x5b\xa7\x5f\x5b\xfb\x3f\x5d\xce\x2c\x10\x1e\x5c\xd9\x32\x38\xd7\xec\x87\x80\x55\xb7\x8e\xa1\xaf\x32\x03\x07\x4a\x8b\x0a\xa1\xb4\x49\xf9\x17\x98\xb1\x49\x4c\x25\x9d\x59\x83\x5a\xfb\x9d\xd5\xce\x7a\x7b\xdd\xae\xb7\xde\x1b\x57\xec\xf1\x08\xb0\x72\x2a\x81\x4c\xed\x3d\xdd\x12\x63\xa5\x29\xc6\x38\xdb\xd6\xf6\xba\x56\xcf\x33\xb7\x6c\x01\xb7\x19\xa2\x9a\x1c\x36\x9c\x5f\xc1\xd0\xee\x28\x41\x5d\x0d\xe0\x1f\x71\x8f\x33\x22\x0a\xa2\xdf\x59\x1e\xb4\x7a\x09\xb0\x5d\x2c\x1e\x68\x5b\xea\xaf\xa7\x39\xaa\xb2\xe6\x74\xca\xaf\xa4\x42\xd5\xcd\x3a\xa1\xbc\x30\xfb\xd8\xe0\x47\x71\xb3\xca\x52\xa2\xf0\x88\x4b\x20\x77\x81\x81\x2f\x93\x29\x52\x74\x22\xd1\x68\x94\x79\xf8\xe0\x2f\xa3\x5f\x3c\xb8\xcf\x7c\xf1\x67\xd2\xbc\x59\x1d\x00\xa9\xd3\xf0\xa1\x00\xb7\x82\xe8\x38\xff\x91\x04\x8d\x3b\x10\xcb\x36\xad\xdc\x0d\x9a\x71\xfb\xc0\x36\x84\xf9\xe4\x05\x4a\x38\x07\xe7\xff\x59\x37\x0f\x01\x92\x5e\xfe\x8e\xde\xda\xab\x4e\xe3\x99\x8f\x3f\x92\xee\xfc\x5d\x12\xf8\x8f\x49\x93\x68\x2f\xb7\x6f\x66\xd6\x9d\x49\x5f\x79\xa7\x1f\xc1\xc6\xc5\x4b\xb9\xa9\xcf\x7e\xdf\xe5\xd2\x33\xce\xdf\xbb\xc3\x8c\xba\x26\x25\xe6\x13\x6d\x42\x5f\xee\x6a\xb3\x6b\x46\x70\x83\x2f\xf5\x6f\x2a\xda\x6e\x05\xfc\x37\xce\x7f\x72\x87\x09\xd3\xa4\x54\xfc\xa4\xe0\x94\x3d\xce\x7f\x7a\x87\xf1\xdf\xbf\xcc\xfb\x67\x3b\xad\x3c\xe6\xf8\x0d\x18\xfb\x23\x25\x4b\x5b\x86\x36\x3f\xe7\xa7\x4f\x0a\x60\x9c\x77\x16\x43\x09\x25\xd0\x24\x3d\xcd\x6e\xde\x62\x6e\xdd\x8b\x7e\x7a\x9b\xf9\x3d\xe3\xdc\x26\x40\x7f\xf2\x4c\xe4\xd2\xc8\xe0\x58\x6f\xa4\x31\xe4\xc8\xc9\xec\xa5\x48\x87\x54\x55\xad\xb6\xaa\xad\x57\x2c\xd9\x76\xa8\x9d\x37\x67\xf5\x35\xa0\x85\x39\xfb\x24\x65\x00\x58\x62\x6e\xfd\x6d\x04\x72\xb5\xf3\x0d\xee\x31\x72\xcd\xff\xae\x1a\x26\xf8\x2d\x6b\xbb\xd4\xba\x32\x06\x6c\x67\xc2\xc8\xe9\x42\xec\x93\x42\x86\xf4\x1d\x41\x38\xe4\xa2\x78\x72\x99\xff\x49\x6f\x42\x2f\x7d\x46\x33\xe6\x94\xd2\xe5\x51\x07\x86\xf5\x4e\x21\x0c\x92\xb5\xb7\x18\xa2\x42\x50\xca\x87\xd0\xed\xf5\x8c\x88\xe2\x88\x4f\x73\x6c\xd0\xc2\xe7\x48\x0e\x67\x15\x5a\x9a\x03\x5e\x08\x55\xc8\x42\x0a\xf2\xd8\x8d\x9e\x82\xd5\x13\x53\x2c\x1f\x4f\x38\x07\x7d\x57\x16\x07\x08\x9e\xa2\xe9\x2f\xf5\xb5\x79\x89\x50\x7f\x30\x86\x8e\xfd\x23\x19\x54\xae\x07\xd8\x63\x11\x62\x35\x20\xa0\x03\x6b\x22\x56\x18\x10\x41\x48\x41\x1e\x93\xad\xc9\x54\x30\xa3\xfd\x04\x3e\x90\xe9\x7d\x89\x73\x91\x1a\xb9\x75\x4e\x3c\x41\x59\x29\x30\xb5\x75\x2a\x42\xee\xd7\x03\x21\x20\x05\x03\x15\x5e\x7e\x17\x2a\x84\x1b\x6d\xc0\x84\xd1\xde\xfc\x53\x5b\x22\x63\x04\xfd\xa8\xa2\xaf\xcc\x0c\x74\x4e\xfd\xc8\xf2\x18\x9e\x62\xb3\x49\x9a\xf3\x9c\x04\xb1\x96\x0d\x14\x6c\xf2\xac\xbe\xdc\xd0\x5f\x91\x5a\xaa\x1d\xee\xea\xb5\xb7\xae\xc1\x43\xef\xe8\x0a\xb8\xbf\x8b\xa5\x49\x94\x9e\x12\xe2\xc1\x72\xd9\xa1\xc9\x6c\xa2\x7a\xb9\x28\x90\xd5\x01\x65\xf9\x72\xe2\x19\x8a\x01\x8a\xa0\x09\xf1\x57\x14\xfb\x9a\x85\xe6\x0d\xda\x46\x25\x47\x07\x50\x3e\x0b\xec\x89\xcc\x7b\x52\xc1\xe4\xc6\x15\xc7\xe0\x87\x74\x4d\x0a\xba\xa4\x04\xb4\xd5\xf1\xcc\x64\x3a\x44\x52\x9c\x37\xa7\x5b\xe5\xbd\x70\xc1\xd4\x7b\xa7\xc7\x2c\x51\x48\xa0\x60\xc1\xa3\xd6\xda\xa3\xb7\x01\xa8\x34\x94\x14\x8f\xad\xd2\x14\x97\x4a\x71\xfc\xe3\x09\x36\x1e\x17\x91\x24\x05\x33\x5c\x9b\x39\x6e\xcd\xfd\xa2\xbd\xcc\x5d\x43\x43\xd1\x8f\x2c\x2f\x88\x34\x1b\x2c\x08\x2e\x7e\x3e\x6c\x55\xae\xc3\xe8\xd0\x82\xdc\x85\x8b\x13\x83\x66\xe8\x28\x27\x42\xf6\xd4\xfb\x36\xbc\xc3\x9b\xbb\x3e\x25\x76\xbc\x57\xc4\x10\x9b\x0c\x0a\xe5\x73\x58\x7b\xcd\x58\xfa\x16\xd9\x85\x50\xa4\x4c\x8c\xc8\x8f\x05\xee\x4c\x7f\xc6\xea\x8a\x35\x1d\xf9\x81\x36\x94\xf3\x90\x8b\xcc\xad\x87\x63\xf7\x6f\x33\x03\x36\x48\x83\xcb\xf6\xa0\x23\x6f\x1e\x03\xe1\xf2\x5f\xae\x07\x00\xd2\x97\xb0\xc1\xfa\xbc\xae\xc3\x81\x6b\xe9\xf2\xae\x43\x7a\xc7\x64\xb2\x5e\x5c\x0c\x00\xd3\x3b\x39\xb9\x76\xa4\xfd\x28\xe8\x0c\x16\x89\x49\x4e\x92\x38\xe7\xed\xbe\x3f\xbc\xc4\x90\x12\x99\x9d\x29\xbf\x60\x75\x8f\xce\xe8\x6f\x02\x75\x7f\x6a\xba\x10\xcd\x53\xc8\xc4\x80\xb3\x59\x23\xed\xb1\xd2\x34\xb4\x06\x19\xc6\xf8\x3b\x84\xc0\x22\xbc\x53\x06\x87\x59\x64\xa6\x9c\x14\xf8\x09\xfa\x33\x1d\xd7\xa1\xf9\x7e\xa0\x03\x6d\x85\x20\x42\xc7\x8f\x12\x29\x43\xbf\xe7\x97\x0d\x75\xe6\x73\xd8\x42\x48\xcc\x7a\x6f\x43\x9c\xc8\xfc\xe9\xe1\x83\xaf\x18\x3a\xfd\x22\xb6\xfc\x81\xcc\x51\xc9\x0f\x25\x18\xac\x28\xd6\x5b\xdf\x86\xb1\x08\xae\x22\x5a\xad\xa0\xe5\x0a\xe4\x65\xa2\xfd\x3b\x8d\xee\x5e\x1f\x7a\x3e\xf4\xdf\x00\x00\x00\xff\xff\xae\x0a\x50\x91\x18\x27\x00\x00")

func GeneratorsAssetsI18nSoracomApiTextJaJsonBytes() ([]byte, error) {
	return bindataRead(
		_GeneratorsAssetsI18nSoracomApiTextJaJson,
		"../generators/assets/i18n/soracom-api.text.ja.json",
	)
}

func GeneratorsAssetsI18nSoracomApiTextJaJson() (*asset, error) {
	bytes, err := GeneratorsAssetsI18nSoracomApiTextJaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../generators/assets/i18n/soracom-api.text.ja.json", size: 10008, mode: os.FileMode(420), modTime: time.Unix(1459775849, 0)}
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
	"../generators/assets/i18n/soracom-api.text.en.json": GeneratorsAssetsI18nSoracomApiTextEnJson,
	"../generators/assets/i18n/soracom-api.text.ja.json": GeneratorsAssetsI18nSoracomApiTextJaJson,
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
	"..": &bintree{nil, map[string]*bintree{
		"generators": &bintree{nil, map[string]*bintree{
			"assets": &bintree{nil, map[string]*bintree{
				"i18n": &bintree{nil, map[string]*bintree{
					"soracom-api.text.en.json": &bintree{GeneratorsAssetsI18nSoracomApiTextEnJson, map[string]*bintree{}},
					"soracom-api.text.ja.json": &bintree{GeneratorsAssetsI18nSoracomApiTextJaJson, map[string]*bintree{}},
				}},
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
