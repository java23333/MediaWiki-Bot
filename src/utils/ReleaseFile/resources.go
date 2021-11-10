// Code generated for package ReleaseFile by go-bindata DO NOT EDIT. (@generated)
// sources:
// resources/language/en-US.yml
// resources/language/ja-JP.yml
// resources/language/zh-CN.yml
package ReleaseFile

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _resourcesLanguageEnUsYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x94\xd1\x8e\xdb\x44\x14\x86\xef\x2d\xf9\x1d\x8e\x6c\x45\x50\xa9\x71\xa8\x54\x2e\xb0\xd0\x4a\xdd\x90\x52\x60\xb7\xac\x36\xa9\xaa\xaa\xbb\x52\x27\xe3\x63\xfb\x90\xc9\x19\x33\x33\x6e\xe2\x56\x95\xb8\xad\x04\xe2\x86\x3b\x78\x80\xde\x40\x5f\x80\xd7\x29\xed\x1d\xaf\x80\x66\x1c\x37\xe9\xa2\x02\x12\xb9\x89\x64\xcf\xf9\xcf\x7f\xbe\xff\x78\xd2\x13\xc1\x55\x2b\x2a\x84\x5a\x58\x58\x22\x32\xc8\x5a\x70\x85\x05\x38\x0d\xc8\xe3\x7b\xf3\x38\x1a\xce\x9c\xea\x82\x4a\xc2\x62\xde\x4a\x89\xd6\x96\xad\x52\x5d\x0e\xc9\x3f\x4a\x3c\x1d\x7d\xf4\x2c\x89\xa3\x34\x48\x41\xa1\xd1\x02\x6b\x87\x80\x5b\xb2\xee\xaa\xb4\x14\x8e\x34\xdf\x16\xa4\xb0\xc8\x21\xf1\xb5\x6f\x4b\xfa\x8a\x24\x8e\xe2\x28\x5d\xd4\x64\x41\x0d\x6d\x4b\x52\x08\x64\xa1\x31\xfa\x31\x15\x58\xc0\xb2\x83\x07\x0f\x16\xf0\x69\xed\x5c\x63\xf3\xc9\x64\xb3\xd9\x64\x5d\xe7\x32\xa9\xb5\x9a\x1c\xc5\xd1\xa9\x20\x9e\x19\xa3\xcd\x82\x1a\x9b\x43\x72\x66\xd0\x5a\x10\xdc\xc1\x0a\xbb\x30\xf7\x96\x5c\x96\x65\x49\x1c\xdd\x63\xb1\x54\x78\xab\x69\x54\x37\x47\x6b\x49\x73\x0e\x49\xff\xd0\x1f\x14\xfe\x05\x94\xda\x80\x00\xc6\x0d\xec\xce\xf8\x57\xa7\x64\x04\x5d\x87\x46\xa1\xb0\x08\xb2\x46\xb9\x02\x57\x23\xcc\xd1\x39\xe2\xca\x26\x71\x34\xd5\x5c\x52\x75\x9b\x14\xce\xb6\x12\x1b\xd7\xab\x07\x67\x40\x0c\x32\xbc\x6e\xcb\x00\xc5\x1f\x17\xcc\xda\x4d\x35\x33\x4a\x17\xe4\xdf\xf1\xa2\x88\x57\x6f\x1b\x83\xd4\xc6\xa0\x74\xaa\xbb\x62\xe1\xa0\xfd\x79\xcb\x5f\x7f\x95\x43\x72\xae\x97\xda\xc1\x7d\x5c\xde\xd1\x7a\x05\x06\x77\x56\x00\x83\x74\xf1\xe7\xef\xdf\x9b\x96\x99\xb8\x02\xcd\xd0\x68\xe3\x86\x50\xef\xb6\x45\x85\x39\x24\xe1\xbf\x4f\xe6\x00\x4d\xd8\x11\x18\x32\x78\x52\x67\x6b\x8d\x15\x19\x95\x69\x53\x65\x92\xf7\xbe\x34\x3b\x21\x5d\x80\x23\x8a\x35\x31\x59\x67\x84\xd3\x26\x8e\xee\xd3\x8a\x4e\x88\x57\x01\xc9\xbb\xe0\x83\xba\xf7\xf1\xef\x3a\x10\xac\xdd\x41\xd5\x2c\x70\xeb\x7c\xde\x7d\x41\xa5\xbd\xd2\x60\xb0\x22\x57\xb7\xcb\x4c\xea\xf5\x84\x3b\xc1\x52\xb8\x42\x4c\x4e\xb1\x20\xe1\x5d\x8c\x8f\xb5\x4b\x47\xb3\x8f\x47\x9f\xdc\x18\x1d\x7f\x36\x9a\xdd\x1c\x1d\x1f\x8f\x6e\xdd\x0c\xd1\x13\x5b\x67\x5a\xe9\x99\xd9\x9e\xc2\x54\xf0\x07\x0e\x4a\xe2\x02\x1e\xfa\x65\x6b\x44\x85\x97\x3e\xd2\x87\xef\xc5\x71\xf9\xb7\x55\x11\xfd\x28\xc4\x4d\xeb\xfc\x7a\xef\x12\x8d\xa3\xcf\xd1\x79\x4b\x5f\x70\xa9\x07\x30\x87\x0d\x9f\x8e\x6e\x3c\xeb\x9b\x79\x3c\xff\x59\x77\xf8\xb2\x10\xbc\x5b\x90\xba\x55\x45\xf8\xee\x96\x08\xa5\x6e\xb9\xb8\x0e\xcb\xb6\x2f\x2d\xb5\x52\x7a\xe3\x17\xc2\xa0\x6d\x95\xb3\xb0\x41\xb3\x3b\x15\x47\xe9\x7e\xe4\x38\x4a\x77\xac\x5b\x8b\x90\xac\xab\xfc\xcc\x6b\x3b\x72\x0a\x13\x0f\xff\xdb\x16\x4d\x07\xa2\x12\xc4\x7d\xda\x7e\xa8\x39\x0a\x23\xeb\x1c\x92\xff\x6d\xe6\x82\x3d\x82\x03\x0b\x8f\x3c\x9c\x03\x13\x8f\xae\x98\xe8\x21\xbc\x37\xa4\xc9\x30\x59\x1c\xa5\x1f\x1a\x2c\xc8\x83\x0b\xe3\x7a\xc9\xcb\xf1\xd1\x7e\xf4\x6b\x71\x94\x66\xe1\xb7\x1f\xec\x7c\x57\x91\x43\x72\xc1\xfb\xfa\x10\xd3\xf8\xa8\x0f\xee\xda\xc5\xce\xc4\xeb\x9f\x7f\x7b\xf5\xfc\x87\x57\x3f\xbe\x7c\xfd\xd3\x8b\x37\x2f\x7f\x7d\xf3\xe2\xbb\x3f\x7e\x79\x9e\x0f\x37\x69\x1c\xa5\xdf\x88\xf1\x97\x67\x71\x94\x3e\xa9\xc7\xd3\xbb\xfb\x9b\xf4\x84\xac\x6f\x30\x6d\x8d\x41\x76\xaa\x03\xf1\x58\x90\x0a\xdf\xcd\x70\x63\xe6\x3d\x97\xd0\x67\xe1\x31\xcc\xd6\x8d\xeb\x76\xc0\x03\x17\xbf\x17\xe8\x1f\x82\xe0\x02\x64\xb8\x78\x3c\x7b\x8f\x8a\xb0\x48\xe2\xe8\xaf\x00\x00\x00\xff\xff\x39\xf6\x80\x85\x39\x06\x00\x00")

func resourcesLanguageEnUsYmlBytes() ([]byte, error) {
	return bindataRead(
		_resourcesLanguageEnUsYml,
		"resources/language/en-US.yml",
	)
}

func resourcesLanguageEnUsYml() (*asset, error) {
	bytes, err := resourcesLanguageEnUsYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/language/en-US.yml", size: 1593, mode: os.FileMode(438), modTime: time.Unix(1636521310, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourcesLanguageJaJpYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\x5d\x6f\x1a\x47\x14\x7d\x5f\x69\xff\x03\xda\x15\x52\xf3\x00\x34\x52\xfa\x50\x1e\x22\xc5\x2e\x69\xda\xda\x49\x14\x3b\xca\x83\x9d\x87\xf5\x32\x86\xa9\xf1\x2c\x82\x45\x8a\x13\x45\x62\x76\x09\xfe\x60\x2d\x5c\x2b\xc4\xad\xed\x1a\xe3\x4f\x02\x02\x22\xb9\x91\xb0\x4d\xe5\x1f\x73\xd9\x65\xf3\xc4\x5f\xa8\x66\x30\xc6\x75\xdd\x56\xaa\x1b\x1e\xd0\x4a\x33\xf7\x9c\x73\xcf\x3d\x57\x23\xbb\xe5\xb4\x5b\xd9\x02\xda\x78\x19\xf5\x0d\x3f\x04\x5a\xb5\xf7\x16\x9d\x8d\xdf\x80\xae\x01\x2d\x8a\xc2\x88\x42\x22\x29\x25\x82\x46\xb5\x30\x9e\xc6\x28\x3c\x96\x52\x55\x94\x4c\x4e\xa7\x62\xb1\xb9\xa0\x47\xba\xa8\x7e\xe5\xfd\xf2\xf5\x95\x62\x49\x14\x64\x0e\x7a\x71\x09\x68\x05\x68\xe6\x2a\xa8\xaa\xe8\x58\x23\xf7\x15\x1c\x43\xe1\xa0\x47\x62\x48\x57\x2a\x24\x51\x10\x85\x51\x05\x93\x50\x22\xa1\x25\xc6\x71\x3c\x19\xf4\x48\x40\x33\x40\xd7\xc1\xb0\x80\xd6\xc1\xa8\x81\xd9\x02\x63\xd5\x59\x6a\x01\xfd\x05\x68\xd9\x59\x38\x05\x23\xe7\xf7\xfb\xfd\x92\x28\x3c\x25\xca\x54\x0c\xdd\x8b\xc7\x63\x73\x63\x28\x99\xc4\x1a\x09\x7a\xa4\x51\x9c\x50\x30\xd0\xaa\xf3\xee\x03\xd7\x9b\x39\x3f\xea\xbc\x3d\x72\xab\x39\xa0\x87\x40\x97\x7b\xf4\xdd\x96\xe5\x96\x6b\x76\x7d\x1d\x8c\x55\x30\x29\x18\x87\x60\x9a\x60\x34\x78\xd9\x01\xd0\x3c\xd0\x6d\xa0\x85\x9e\xd0\x61\x8d\x4c\xe3\xc8\x7d\x1c\x43\xa1\x17\x2a\x8a\xeb\x3d\xb2\xf3\x7a\xb3\x00\x46\x09\x8c\x3d\x30\xab\x9d\xc2\x07\xbb\xd9\x64\x05\x0a\x21\x9a\x3e\xac\x11\x82\x54\x9d\x8b\xba\x2c\xae\xb6\xcb\x59\xf2\x60\x56\xc0\x3c\xe2\xa4\xff\x51\xd8\x93\x14\x79\xf4\x03\xf3\xcd\xdc\x64\xb7\xcc\x85\x67\x68\xea\x81\xa6\xcd\x00\xad\xdb\xf9\xb5\xf6\x59\x09\xa8\xf5\x88\x05\x80\x63\xef\x72\x80\x0c\x18\xb9\x6e\xcb\xe2\xc3\x35\x7f\x65\x16\x9b\x0b\x40\x0f\xed\x7a\xd1\x2d\x59\xed\x66\x4d\x12\x85\x87\xa9\x70\x04\x71\xd8\x25\x30\xe7\x7b\xff\x7c\x5e\x72\x54\xd7\xe3\xc9\x60\x20\xf0\x32\xea\x9f\xd5\x50\x04\x27\x62\x7e\x2d\x11\xf1\xab\x84\x71\x18\x3b\x4c\xa7\x71\x0a\xc6\xf1\x95\x8e\x3a\xf5\x52\x67\x25\xeb\xa6\xdf\x00\xad\x7e\x4a\xef\x76\x3e\x96\xfe\xda\x8e\x28\x3c\xc3\x33\x78\x04\x93\x19\x9e\x88\xf3\xd4\xdc\x1c\x97\x0b\x7f\x80\x62\xf1\x71\xf4\x42\x0f\x7a\xa4\xf6\xef\x67\x40\x33\xce\xbb\x63\xa0\x0d\x4f\xbf\x9f\x08\xd6\xa3\xa9\x29\xbf\xaa\xcd\x06\xc8\x9c\x42\x54\x45\x0f\x2b\x81\x51\x14\xc6\x0a\x13\xe5\x1b\xd2\x74\xd9\x1b\xfa\xca\xfb\xf5\x6d\xef\xd0\x37\xde\xd0\x1d\xef\xd0\x90\xf7\xde\x1d\x0f\xd0\x43\xb7\x72\x06\xf4\xf8\x3a\x4a\x79\xe2\x6f\xcd\x7a\xce\x27\xd2\x98\x60\xf1\x8f\x2b\x11\xf4\x1c\xa8\xe5\x1e\xe4\x80\xee\x01\xfd\x09\x8c\x45\xbe\x00\x15\xa0\x39\x3e\xb3\x62\xb7\x65\xd9\x6f\xf6\xed\xa5\x0d\xa0\x8d\x7e\x78\x32\xfc\xf4\x3d\xd0\x2c\xfb\xf8\xc7\x9c\x88\xc2\xb7\x48\x67\x6d\x7c\x47\xa6\xb5\xbe\xb7\x13\xcc\xdc\xbe\x8e\x57\xde\xdb\xaf\x3f\xaf\x86\x9e\x23\x60\xae\xf3\x95\x6e\x82\xb1\x3a\xe0\xba\xc4\x02\xd4\xea\xb6\xac\xf6\xe9\x7e\xbb\x99\x03\x5a\xef\x7c\x5c\x71\xb6\x36\xff\x7c\xb9\xc8\x9c\xbd\x30\x4e\x14\x64\xcf\x6c\x24\x38\x00\x36\xb3\x3c\xd1\x5b\x6c\x36\x60\x18\x40\xb3\xed\x66\xda\x3e\x39\xb0\x0b\x79\xa0\x19\x7b\x65\x01\x8c\x3c\xd0\x8d\x6b\xb3\xc7\xfc\x19\x43\x4a\x42\x8d\xf2\xec\xff\x0f\x5a\x27\x09\xb3\xd9\xc3\xfc\xbd\xb1\xc6\x7f\x59\xc0\x40\xdf\x13\x51\x90\xbf\xb0\xb3\xcb\x8e\x35\x6f\xd7\xd7\xb9\x53\x8f\x99\x53\xbe\xbb\x03\xd7\x6e\x89\x82\xec\xe7\xbf\x41\xe3\x4f\x50\x18\x27\x90\xca\x36\x64\x92\x5c\x02\xe0\x39\xf1\xdd\xed\x65\xe4\xd6\x24\xe9\xc9\xe8\x6c\xd4\xed\xc5\x65\x3b\xdf\xe8\xbc\x2d\xbb\x8d\x9a\x5b\x4e\x3b\x9b\x8b\x41\x51\x90\x11\xf1\x3d\x1d\x13\x05\xf9\x47\xc5\xf7\xfd\xe3\xfe\x63\x31\x78\x1d\x46\x70\x92\x31\x5c\x5f\xde\x73\x8b\x13\x8c\x63\x3d\x86\x42\xb3\x71\x9d\x3d\x48\xce\xf6\xfc\xa7\x9d\x9f\xdb\xcd\x93\xce\xfb\x93\x6e\xcb\x72\xd6\xb6\x9d\xa3\x82\x53\xdc\x77\x1b\x3b\x92\x28\xfc\x11\x00\x00\xff\xff\x07\x54\x89\x88\xee\x06\x00\x00")

func resourcesLanguageJaJpYmlBytes() ([]byte, error) {
	return bindataRead(
		_resourcesLanguageJaJpYml,
		"resources/language/ja-JP.yml",
	)
}

func resourcesLanguageJaJpYml() (*asset, error) {
	bytes, err := resourcesLanguageJaJpYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/language/ja-JP.yml", size: 1774, mode: os.FileMode(438), modTime: time.Unix(1636521310, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourcesLanguageZhCnYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\x51\x4f\xdb\x56\x18\x7d\xb7\xe4\xff\x10\xd9\xb2\xb4\x3e\xc4\x59\xa5\xee\x61\x79\xa8\x54\x58\xba\x6e\x83\xae\x2a\x54\x7d\x28\x7d\x30\xce\xc5\xb9\xc3\x5c\x5b\x8e\x23\x95\x56\x95\xb2\xa9\x25\x44\x4d\x20\x88\x8c\x66\x6d\xb6\x36\xa0\x54\x99\x26\x6c\x44\xba\xc6\x73\x42\xf8\x33\xbe\xd7\xf6\x13\x7f\x61\xba\xbe\xa4\xac\x0c\x34\xcd\x8f\xbe\xf7\x7c\xe7\x3b\xe7\x7c\xf7\x13\x23\x77\x3f\xea\x95\xf1\xe0\x90\xbc\x7e\x4f\x9a\x7f\x05\x9e\xff\xb8\x90\x9e\xbe\xcd\x73\x33\x0a\xd2\x4a\x8a\x06\x66\x8d\x3c\x5c\x82\x20\x3f\x57\x52\x55\x50\x2c\x2e\x95\x74\x7d\x35\x9b\x12\xfe\x8d\x7b\x22\x7d\xfe\x54\xe0\x39\x31\xc1\xb3\xe3\xc0\xab\xe3\xfd\x16\x6e\xf7\xce\x97\x53\x15\x1b\x1a\xe8\xa6\x02\x75\x90\xcf\xa6\x04\x0a\x3d\x87\x10\x78\x8e\xe7\x66\x15\x88\x72\x96\x65\x58\xf3\xd0\x2c\x66\x53\x02\xa9\x55\x83\xe1\x90\x3c\xdb\x8c\x9b\x4e\x5c\x2e\xe3\x8a\x2f\xcb\xb2\x2c\xf0\xdc\x3d\xa4\x2c\xea\xe0\x86\x69\xea\xab\x73\xa0\x58\x84\x06\xa2\xb7\x5f\xbe\x25\xfd\x9f\x71\x63\x6b\x16\x5a\x0a\x0c\x9b\xfd\xc8\x1d\x90\x9d\x83\xf0\xd5\xb3\xd3\x3b\x27\xa3\x1a\xfd\xb5\x57\x26\x6f\xba\x91\x33\x0e\x8f\x1c\x81\xe7\xa6\x0d\xb4\x04\xb5\x9b\x50\x07\xb9\x47\x2a\x30\x6d\x56\x2b\x7e\x5e\x0f\x8f\x1c\xb2\x53\x09\x86\x1f\xf0\xe8\x27\xec\x79\xf4\xaa\x82\x90\x61\x4f\x1b\x08\x01\xd5\x4e\x48\x3e\xb2\x92\xfd\xbd\xb0\xe3\xc4\xdb\x63\xb2\xd1\x8d\x2a\xfd\xe4\xf0\x42\xba\xbb\x25\xf4\xfd\x77\x14\xd6\xf6\xf1\x2f\xbd\xc0\xf7\xef\x83\xc5\x5b\x86\xb1\x4c\x36\xba\xa4\xf9\x01\x0f\x0e\xf1\xa8\x8c\x1b\x2e\xc5\x1e\x37\xa2\x4e\x2d\xf0\x37\xa8\x5b\xe1\x1f\x2e\xde\xdc\x13\x78\xee\x76\x29\xaf\x01\x8a\x5f\xef\x07\x5e\x99\xac\xf7\x13\xe3\x44\xd6\x45\xe4\x1c\xc7\x2f\x9d\x82\x6d\x9b\xc5\x6c\x26\xf3\xb8\x20\xaf\x18\x40\x83\x96\x2e\x1b\x96\x26\xab\xa7\xfa\xa3\x1f\x9b\x61\x7f\x18\x3a\x9d\xb0\xb1\x86\xb7\x5a\x3c\x77\x1f\x2e\xc3\x19\x88\x96\x13\xe7\x3f\x2a\x62\xb5\x28\xf7\x85\xb0\x84\xf6\x16\xd0\xcd\x79\xf0\xc8\xce\xa6\x84\xe0\xe8\x38\x6c\xf6\x22\xf7\x3d\x69\x6d\x44\xee\x00\x57\xeb\x78\x5c\x4e\x4d\x5a\xd1\xa0\x5d\x28\x2d\xca\xaa\xb1\x92\x41\xab\x0a\x52\x15\x3b\xaf\x64\x66\x41\x1e\x2a\x94\x3c\x3d\x65\xd8\xa2\x94\xfb\x42\xfa\xf2\xaa\x34\xf5\x95\x94\xbb\x26\x4d\x4d\x49\x37\xae\xa5\xc8\x9b\x6e\xd8\x7e\xc1\x04\xe2\x76\xef\xc1\xa5\xc2\x1e\x06\xde\x3e\xa9\x8e\xe9\x30\xad\x1f\x3c\xa0\x53\x64\x2a\x1a\x78\x88\xb7\xdf\x7d\x92\xc1\x78\x1b\x3f\xef\x92\x96\x8b\x1b\xef\x58\x5e\x3c\xf7\x35\xb0\x69\x07\xdf\xa0\x25\x63\x22\x9f\x32\x51\xd9\x9f\x16\x7d\x22\x5d\x7d\xfa\xdf\x05\x4f\xc3\x38\xec\x90\x76\x95\x54\xc7\x78\xfd\x20\xee\xfc\x19\xff\xba\x7b\x32\xaa\x05\x47\x6b\xa4\xe5\xb2\x9f\x81\xbf\x16\x0c\xbb\x81\xf7\x22\x1c\x6e\x93\xdf\xda\x3c\x27\x9e\x35\xcd\x73\x62\xe4\x0e\x98\x9d\xa9\x15\x2d\xcb\x0a\x90\xb7\x95\x78\xb7\x95\x8a\x2b\x75\xb2\x73\x40\xa9\xdd\x5d\x16\x1c\xed\x7c\x0e\x28\x96\x5a\xa0\xc9\xfd\x4f\xe2\x05\xc4\x5e\xe2\x84\x8e\x8a\xbc\x9c\x90\x89\xbb\x34\x84\xcc\x44\x02\xcf\x89\x9f\xc5\x95\x3a\x76\x5e\xe1\xc6\x56\x22\xec\x0e\x15\x96\xbe\x7e\x26\xf2\x0a\xcf\x89\x72\xf2\x9d\xa9\xb8\x0b\xf2\xd0\x02\x2a\x1d\xa6\x05\xf4\x8f\x02\x49\x16\xe9\xeb\x2c\x81\x2b\x0b\x88\xb5\x11\xbe\x76\xe8\x8c\x6d\xba\xc9\xd4\xd1\x55\x42\xda\xd5\x2c\xcf\x89\x00\xa5\xef\xcd\xf1\x9c\xf8\x83\x92\xfe\xf6\xce\x64\x3d\x9d\xed\xa3\x19\x58\xa4\x0c\x17\xc3\x99\x1f\x09\xc1\x3c\xb4\x75\x90\x5b\x31\x6d\xba\xfc\x98\x19\x81\xe7\x87\xbf\xfb\x27\xa3\xda\xe9\x93\x9f\x98\xf2\x77\x00\x00\x00\xff\xff\x6a\x4e\x0b\x9f\x54\x05\x00\x00")

func resourcesLanguageZhCnYmlBytes() ([]byte, error) {
	return bindataRead(
		_resourcesLanguageZhCnYml,
		"resources/language/zh-CN.yml",
	)
}

func resourcesLanguageZhCnYml() (*asset, error) {
	bytes, err := resourcesLanguageZhCnYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/language/zh-CN.yml", size: 1364, mode: os.FileMode(438), modTime: time.Unix(1636521310, 0)}
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
	"resources/language/en-US.yml": resourcesLanguageEnUsYml,
	"resources/language/ja-JP.yml": resourcesLanguageJaJpYml,
	"resources/language/zh-CN.yml": resourcesLanguageZhCnYml,
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
	"resources": &bintree{nil, map[string]*bintree{
		"language": &bintree{nil, map[string]*bintree{
			"en-US.yml": &bintree{resourcesLanguageEnUsYml, map[string]*bintree{}},
			"ja-JP.yml": &bintree{resourcesLanguageJaJpYml, map[string]*bintree{}},
			"zh-CN.yml": &bintree{resourcesLanguageZhCnYml, map[string]*bintree{}},
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
