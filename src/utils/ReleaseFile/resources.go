package ReleaseFile

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _resources_language_en_us_yml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x54\x5f\x8b\x1c\xc5\x17\x7d\x6f\xe8\xef\x70\xe9\x21\xfc\x7e\x81\x6c\x8f\xbe\x36\xb2\x20\xc3\xc6\x44\x77\xe3\xb2\x33\x61\x09\xd9\x85\xd4\x74\xdf\xee\xbe\x4e\xcd\xad\xb6\xaa\x3a\x33\xbd\x21\x0f\x3e\xf8\x07\xcc\xa3\x82\xa2\x06\x14\x91\xa0\x44\x41\x1f\x54\x0c\xec\x97\xc9\x0c\xeb\x53\xbe\x82\x54\xf5\xf4\xce\x64\x25\x51\x70\x5e\x06\xba\xee\x3d\xf7\xdc\x73\x4e\x55\x6f\x57\x70\x51\x8b\x02\xa1\x14\x06\xc6\x88\x0c\x69\x29\xb8\xc0\x0c\xac\x02\xe4\xad\x9b\xc3\x30\xe8\x6a\xf6\x54\x46\x39\x61\x36\xac\xd3\x14\x8d\xc9\x6b\x29\x9b\x04\xa2\x97\x42\xdc\xbb\xf4\xca\xfd\x28\x0c\x7a\x1e\x0a\x32\x85\x06\x58\x59\x04\x9c\x93\xb1\x17\xa1\x53\x61\x49\xf1\x55\x41\x12\xb3\x04\x22\xd7\x7b\xde\xd2\x76\x44\x61\x10\x06\xbd\x51\x49\x06\x64\x37\x36\x27\x89\x40\x06\x2a\xad\xee\x52\x86\x19\x8c\x1b\xb8\x75\x6b\x04\xaf\x95\xd6\x56\x26\xe9\xf7\x67\xb3\x59\xdc\x34\x36\x4e\x95\x92\xfd\xed\x30\xd8\x13\xc4\x3b\x5a\x2b\x3d\xa2\xca\x24\x10\xed\x6b\x34\x06\x04\x37\x30\xc1\xc6\xef\x3d\x27\x1b\xc7\x71\x14\x06\x37\x59\x8c\x25\xbe\x5e\x55\xb2\x19\xa2\x31\xa4\x38\x81\xa8\xfd\xe8\x0a\x85\x3b\x80\x5c\x69\x10\xc0\x38\x83\x55\x8d\x3b\xda\x23\x2d\xe8\x0a\x54\x12\x85\x41\x48\x4b\x4c\x27\x60\x4b\x84\x21\x5a\x4b\x5c\x98\x28\x0c\x06\x8a\x73\x2a\xae\x92\xc4\x9d\x79\x8a\x95\x6d\xd1\x3d\x33\x20\x86\xd4\x1f\xd7\xb9\x17\xc5\x95\x0b\x66\x65\x07\x8a\x19\x53\xeb\xe1\x9f\xe3\x22\x89\x27\xe7\x83\x21\x55\x5a\x63\x6a\x65\x73\x81\xc2\xc6\xf8\x83\x9a\xdf\x7e\x2b\x81\xe8\x40\x8d\x95\x85\x43\x1c\x5f\x53\x6a\x02\x1a\x57\x54\x00\x3d\x74\xf6\xec\xc9\x03\x5d\x33\x13\x17\xa0\x18\x2a\xa5\x6d\x67\xea\x8d\x3a\x2b\x30\x81\xc8\xff\xb7\xce\x6c\x48\xe3\x33\x02\x9d\x07\x27\x65\x3c\x55\x58\x90\x96\xb1\xd2\x45\x9c\xf2\x9a\x97\x62\x2b\x52\xeb\xc5\x11\xd9\x94\x98\x8c\xd5\xc2\x2a\x1d\x06\x87\x34\xa1\x5d\xe2\x89\x97\xe4\x79\xe1\x3d\xba\xe3\xf1\xcf\x38\xe0\xa9\x5d\x43\x59\x8d\x70\x6e\x9d\xdf\x6d\x43\xa1\x1c\x52\x47\xd0\x94\xa4\x78\xac\x6c\x3c\x6f\x4e\xfa\x5d\xb6\x4c\xdf\xe7\xb6\x3f\x50\xd3\xa9\xe0\xac\xef\xad\x26\x36\x56\xd7\xa9\xd3\xc8\xb4\x5b\x0f\x04\xff\xcf\x42\x4e\x9c\xc1\x6d\x17\xae\x4a\x14\x78\xec\x2c\xbc\xfd\xc2\xf5\x8f\xff\x16\x0d\xd1\x52\x27\xae\x6a\xeb\xe2\xbc\x72\x30\x0c\xde\x40\xeb\x84\xb8\xce\xb9\xea\x84\xd8\x1c\x78\xef\xd2\xab\xf7\xdb\x61\x4e\x8e\x7f\x8d\xdb\xdd\x24\x04\xc7\x16\x52\x55\xcb\xcc\xdf\xb3\x31\x42\xae\x6a\xce\xae\xc0\xb8\x6e\x5b\x73\x25\xa5\x9a\xb9\x00\x68\x34\xb5\xb4\x06\x66\xa8\x57\x55\x61\xd0\x5b\xaf\x1c\x06\xbd\x95\xb6\xb5\x41\x88\xa6\x45\xb2\xef\xb0\x2d\x59\x89\x91\x13\xfb\xdd\x1a\x75\x03\xa2\x10\xc4\xad\xbb\x6e\xa9\x21\x0a\x9d\x96\x09\x44\xff\x99\xcc\x11\x3b\x09\x36\x28\xdc\x71\xe2\x6c\x90\xb8\x73\x81\x44\x2b\xc2\x0b\x4d\xea\x77\x9b\x85\x41\xef\xff\x1a\x33\x72\xc2\xf9\x75\x1d\xe4\xf1\xd6\xf6\x7a\xf5\xcb\x61\xd0\x8b\xfd\x6f\xbd\xd8\xc1\xaa\x23\x81\xe8\x88\xd7\xfd\xde\xa6\xad\xed\xd6\xb8\xcb\x47\x2b\x12\x83\x5a\x6b\x64\x2b\x1b\x10\x77\x05\x49\x1f\xf4\xf3\x18\x26\xdd\x0b\x1a\x06\xbd\x77\xc4\xd6\x9b\xfb\x61\xd0\x3b\x29\xb7\x06\x37\xd6\x2f\xe8\x2e\x19\x37\xe8\x65\x30\x49\xab\x8f\x9f\x37\x72\x72\xec\x4c\x2b\xdb\xac\x84\xf7\xfa\xb8\x7c\xa0\xfb\x08\x82\x33\x48\xfd\x83\xe3\x3c\x70\x92\x11\x66\xbe\x73\xe0\xee\x19\xb1\xb9\x2e\x25\x16\x42\x1e\x2a\x9d\xb9\x47\xf4\xec\xf4\x93\xc5\x17\x0f\x9f\x9e\x7e\xbd\x7c\xef\xa7\xa7\xbf\x3d\x5e\x3c\xfe\x6c\xf1\xe5\xa3\x3f\xbf\x7a\xb8\xfc\xe5\xd3\xc5\x07\xef\x2f\x7e\xfc\xfd\xd9\x93\x07\x8b\x5f\x7f\x3e\xfb\xe6\x87\xb3\xd3\x0f\x97\x7f\x7c\xbb\xf8\xfc\xd1\xf2\xe3\xef\x96\x1f\x7d\x1f\x85\xc1\x5f\x01\x00\x00\xff\xff\xe5\x13\x02\x0f\x86\x06\x00\x00")

func resources_language_en_us_yml() ([]byte, error) {
	return bindata_read(
		_resources_language_en_us_yml,
		"resources/language/en-US.yml",
	)
}

var _resources_language_ja_jp_yml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x95\x51\x4f\xdb\x56\x14\xc7\xdf\x2d\xf9\x3b\x44\x8e\x2a\xad\x0f\x24\xdb\x6b\x1e\x2a\x4d\xa8\x5d\xbb\x41\x5b\x15\x2a\x1e\xa0\x0f\x26\xb9\x84\x3b\xcc\x75\x14\x1b\x09\xa8\x2a\x71\xec\x10\x52\xe2\x0a\x86\x4a\xbb\x02\x0d\x0b\xa4\x24\x4b\x94\xa4\x1a\xad\x16\x96\x4c\xf9\x30\x27\x76\xdc\x27\xbe\xc2\x74\x6f\x12\xc2\x68\xb7\x49\xeb\x96\x87\x28\x92\x7d\xfe\xe7\x7f\x7f\xe7\x7f\x72\x83\x7e\x69\xcd\x2f\xe7\x10\xea\xab\xf3\x23\xa3\x77\x11\x2a\x6e\xe1\xa9\xb7\xff\x0e\xe1\x25\xc2\xa1\x2c\x8d\xa9\x2c\xbe\xa4\xc6\xc9\xb8\x1e\xa3\x73\x94\xc4\x26\x96\xa2\x51\x62\x18\x73\x4b\x9a\xb6\x12\x09\x28\x17\xd5\x8f\xaf\x7d\xf9\xe4\x4a\xb1\x22\x4b\x41\x21\x7a\xf1\x12\x42\x19\x21\x75\x55\x34\xaa\x9a\x54\x67\xb7\x54\xaa\x91\x58\x24\xa0\x70\xa5\x2b\x15\x8a\x2c\xc9\xd2\xb8\x4a\xd9\xcd\x64\x52\x4f\x4e\xd2\x84\x11\x09\x28\x08\x29\x84\x3d\xb4\x1c\x84\x1a\x5a\x55\xb4\x5b\x68\xed\x78\x9b\x2d\x84\x57\x08\x25\x2f\xd3\x44\x2b\x1b\x0a\x85\x42\x8a\x2c\x3d\x64\xea\xac\x46\xbe\x4e\x24\xb4\x95\x09\x62\x18\x54\x67\x91\x80\x32\x4e\x93\x2a\x45\xa8\x78\x2f\xde\x0a\xbf\xa9\xfe\xa3\xee\xf3\x53\xbf\x92\x45\x28\x22\x3c\xeb\xb5\x3f\x6f\x39\x7e\xa9\xea\xd6\xf6\xd0\xda\x41\x1b\xd0\x2a\xa2\x6d\xa3\x55\x17\x65\x27\x08\x5b\x08\x3f\x21\xec\xf6\x8c\x8e\xea\x6c\x8e\xc6\x6f\x51\x8d\xdc\x5c\x8e\x92\x84\xd9\x6b\xd6\xaf\xb7\x77\xd1\xca\xa3\x55\x40\xbb\xd2\xdd\x7d\xeb\x36\x1a\xbc\x40\x65\x4c\x37\x47\x75\xc6\x48\xd4\x14\xa6\x2e\x9b\xab\x1e\x8b\x2e\x5b\x68\x97\xd1\x3e\x15\x4d\xff\xa5\xb1\x07\x4b\xec\xde\x77\x9c\x9b\x7d\xc0\xdf\xb2\x33\x53\x64\xf6\xb6\xae\x2f\x20\xd4\xdc\xad\x97\x9d\x76\x1e\xc1\xb9\xc7\x03\x20\xb4\x8f\x85\x40\x0a\xad\xec\x79\xcb\x11\xc3\xb5\x5f\x73\xc4\x76\x06\xa1\xe8\xd6\x0e\xfd\xbc\xd3\x69\x54\x15\x59\xba\xbb\x14\x8b\x13\x21\xbb\x89\xf6\x46\xef\x5b\xcc\x2b\x38\x6f\x9a\x09\x23\x12\x0e\xaf\xce\x87\x16\x75\x12\xa7\x49\x2d\xa4\x27\xe3\xa1\x28\xe3\x3d\xac\x23\xee\xd3\x6a\xa2\x75\x76\xe5\x44\xdd\x5a\xbe\xbb\x9d\xf6\xd7\xd6\x11\x2a\x1f\xd6\x8e\xbb\xef\xf3\x1f\x1f\x47\x96\xa6\xe8\x02\x1d\xa3\x6c\x41\x24\xa2\x9f\x9a\xcf\xd7\x15\xc6\x6f\x13\x2d\x31\x49\x96\xcd\x48\x40\xe9\xfc\xde\x46\x48\x79\x2f\xce\x10\xea\x81\xc1\x79\x8c\x79\xaa\xb3\x59\xdd\x0c\x2d\xaf\xac\x86\xb5\x7e\x94\x8d\xf0\xf7\xea\xc8\xb7\xf7\xc3\xa3\xfa\xe2\xa2\xca\x62\xe1\x00\x42\xd1\x2f\xb7\x11\xce\x3e\xd5\x22\x38\xfd\x97\x70\x1e\x89\x09\xd4\xa7\x79\xdc\x13\x6a\x9c\x3c\x42\x70\xfc\x93\x2c\x42\x01\xe1\x07\xb4\x9e\x8a\xc0\x97\x11\xb2\x62\x46\x87\xe7\x2d\xc7\x5d\x7f\xe3\x6e\xee\x23\xd4\x07\x61\x49\x89\xa7\x3f\x23\xa4\xf9\x8f\xbf\xcd\x85\x2c\x7d\x43\x4c\xce\xf2\x0e\x9b\xd3\x07\x2c\xa7\x39\xcc\x81\x8f\xc7\xd7\xbe\x7a\xf2\xff\x7a\xe8\x11\x41\x7b\x4f\xac\x70\x03\xad\x9d\x61\xaf\x4b\x5d\x10\x9c\xf3\x96\xd3\x69\xbe\xe9\x34\xb2\x08\xb5\xee\xfb\x6d\x2f\x77\xf0\xe7\x97\x0f\x39\xd9\x0b\x70\xb2\x14\x0c\x2c\xc6\x23\x43\x61\x3b\x2d\x12\x9c\xe3\xb3\x41\xcb\x42\x48\x77\x1a\x6b\xee\x6f\x27\xee\xee\x16\x42\xca\xdd\xce\xa0\xb5\x85\xb0\xff\xc9\xac\x71\x3e\x13\x44\x4d\x46\xe7\x45\xd6\xff\x03\xaf\x33\x8c\x63\x0e\x70\xbe\x9f\xed\xf1\x1f\x16\x2e\x3c\x60\x22\x4b\xc1\x2f\xdc\xf4\x33\xcf\xd9\x70\x6b\x7b\x82\xd4\x7d\x4e\x6a\xe4\xc6\x90\xda\x75\x59\x0a\x86\xc4\x67\x78\xf0\x07\x24\x46\x93\x24\xca\x37\x62\x86\x5d\x12\x10\x39\x19\xb9\xd1\xcb\xc8\xf5\x19\xd6\xb3\xd1\x69\x6e\x8a\xbd\xc9\xa0\x95\xed\xfd\x91\x47\x64\x29\x48\xd8\xc8\xc3\x09\x59\x0a\x8a\x35\x19\xdc\x0c\xc3\xab\x60\x8c\x1a\x62\xe1\x3e\xae\xed\x71\x12\xd2\x93\xd4\xd4\xc8\xcd\xc5\x84\xb9\x22\x86\xd0\x47\x75\x71\x4d\xf0\x8b\x00\x8a\xb8\x06\x5e\xe1\xa0\xfb\xee\xe8\xf2\xfe\x8b\xf2\x51\x9d\x99\x2a\x65\xc6\x1d\x4d\x23\x71\x55\x9b\xd2\x93\x31\x7e\x91\xf8\xed\xe7\xee\x7e\xae\xd3\xce\x7b\x50\xef\x34\xaa\x6e\xf5\x47\xf7\xa0\xf4\xe1\x75\xce\x3b\xdd\x75\xd3\xeb\x6e\xed\x8c\xe7\xfb\xd7\x5f\xfc\xa3\x8a\xdf\xde\xf0\x9a\x05\xf7\x55\xc9\xcb\x9e\x78\x99\xb2\x22\x4b\x7f\x04\x00\x00\xff\xff\x17\xa1\xf9\xde\x3f\x07\x00\x00")

func resources_language_ja_jp_yml() ([]byte, error) {
	return bindata_read(
		_resources_language_ja_jp_yml,
		"resources/language/ja-JP.yml",
	)
}

var _resources_language_zh_cn_yml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\x51\x6f\xda\x56\x18\x7d\xb7\xe4\xff\x80\x6c\x55\x5a\x1f\x30\xdb\x2b\x0f\x95\x26\x94\xae\xdd\x92\xae\x6a\x52\xe5\xa1\xe9\x83\x03\x37\xe6\x2e\xe6\x1a\xd9\x46\x4a\x52\x55\xa2\x53\x13\x40\x81\x84\x28\x2c\xa5\x0d\x4d\x4b\x32\x22\xb6\x0a\x3b\x0a\x5d\x71\x0d\x81\x3f\xe3\x7b\x6d\x3f\xe5\x2f\x4c\xd7\x86\xb2\x66\x89\xa6\xf9\xd1\xf7\x9e\xef\x7c\xe7\x7c\xdf\x3d\xbc\x67\x76\xbc\x76\x1e\xf7\xce\xc9\xe1\x47\x52\xfb\xec\x58\xf6\x46\x3a\x9a\x78\xc0\x32\xb3\x22\x92\x72\xa2\x04\xe6\x94\x14\x5c\x81\x20\x35\x9f\x4b\x26\x81\xa6\xad\xe4\x64\x79\x3d\x1e\xe1\xfe\x8d\x7b\x76\xeb\xdb\xe7\x1c\xcb\xf0\x01\x3e\x3c\x76\xac\x0a\xee\xd4\x71\xa3\x7d\xb5\x5c\x52\xd4\xa1\x82\xee\x8a\x50\x06\xa9\x78\x84\xa3\xd0\x2b\x08\x8e\x65\x58\x66\x4e\x84\x68\x46\x55\x15\x75\x01\x66\xb5\x78\x84\x23\xe5\x92\xd3\xef\x93\x97\xbb\x7e\xcd\xf0\xf3\x79\x5c\xb0\x05\x41\x10\x38\x96\x79\x8c\xc4\x65\x19\x7c\x9f\xcd\xca\xeb\xf3\x40\xd3\xa0\x82\xe8\xed\x57\xef\x49\xf7\x37\x5c\xdd\x9b\x83\xaa\x08\xdd\x5a\xd7\x33\x7b\xe4\xe0\xcc\x7d\xf3\x72\x7c\xe7\x72\x50\xa6\xbf\x4e\xf2\xe4\x5d\xcb\x33\x86\xee\x85\xc1\xb1\x4c\x42\x41\x2b\x50\xba\x0b\x65\x30\xb3\x96\x04\x59\x3d\xac\xe5\x6f\x56\xdc\x0b\x83\x1c\x14\x9c\xfe\x27\x3c\xf8\x15\x5b\x16\xbd\x2a\x22\xa4\xe8\x09\x05\x21\x90\xd4\x03\x92\x2f\xac\xa4\x73\xe2\x36\x0d\x7f\x7f\x48\x76\x5a\x5e\xa1\x1b\x1c\x5e\x4b\xf7\x28\x87\x7e\xfe\x89\xc2\x1a\x36\x7e\xdd\x76\x6c\x7b\x11\x2c\xdf\x53\x94\x55\xb2\xd3\x22\xb5\x4f\xb8\x77\x8e\x07\x79\x5c\x35\x29\x76\x54\xf5\x9a\x65\xc7\xde\xa1\x6e\xb9\x1f\x4c\xbc\x7b\xc2\xb1\xcc\x83\x5c\x4a\x02\x14\x5f\xec\x3a\x56\x9e\x14\xbb\x81\x71\x7c\xd8\x85\x67\x8c\xfc\x57\x46\x5a\xd7\xb3\x5a\x3c\x16\xdb\x48\x0b\x19\x05\x48\x50\x95\x05\x45\x95\x84\xe4\x58\xbf\xf7\xa2\xe6\x76\xfb\xae\xd1\x74\xab\x5b\x78\xaf\xce\x32\x8b\x70\x15\xce\x42\xb4\x1a\x38\xff\x45\x51\x58\x8b\x72\x5f\x0b\x0b\x68\xef\x01\x39\xbb\x00\xd6\xf4\x78\x84\x73\x2e\x46\x6e\xad\xed\x99\x1f\x49\x7d\xc7\x33\x7b\xb8\x54\xc1\xc3\x7c\x64\xd2\x8a\x96\x86\x0a\x5a\x56\x74\x61\x6d\x7d\x23\x96\x50\x32\x19\x11\xa5\x62\x11\xf2\xae\xe5\x36\xb6\x43\x01\xb8\xd1\x7e\x72\x63\xe3\x4f\x1d\xab\x43\x4a\x43\xba\x2c\xc5\xb3\x27\x74\x4b\xb2\xa2\x04\x9e\xe2\xfd\xd3\xaf\x3c\x1e\xee\xe3\xcd\x16\xa9\x9b\xb8\x7a\x1a\xce\x83\x65\x7e\x00\x3a\x95\x77\x1f\xad\x28\x13\x79\x94\x89\xca\xfa\xba\xe8\xb3\x5b\xdf\x3d\xff\xef\x82\x63\xb3\xcf\x9b\xa4\x51\x22\xa5\x21\x2e\x9e\xf9\xcd\xbf\xfc\xb7\xc7\x97\x83\xb2\x73\xb1\x45\xea\x66\xf8\xd3\xb1\xb7\x9c\x7e\xcb\xb1\xb6\xdd\xfe\x3e\x39\x6a\xb0\x0c\x3f\x6d\x9a\x65\x78\xcf\xec\x85\x76\x45\x32\x52\x3c\x2c\x40\xde\x17\xfc\xe3\x7a\xc4\x2f\x54\xc8\xc1\x19\xa5\x36\x8f\xc3\xc1\xd0\xce\xe7\x81\xa8\x26\xd3\x74\x32\xff\x93\x78\x09\x85\x2f\x6d\x42\x47\x45\xde\x4c\x18\x8a\xbb\x71\x08\xb1\x89\x04\x96\xe1\xbf\xf1\x0b\x15\x6c\xbc\xc1\xd5\xbd\x40\xd8\x43\x2a\x2c\x7a\x67\x2a\xf2\x36\xcb\xf0\x42\xf0\x4d\x55\x3c\x02\x29\xa8\x82\x24\x5d\x96\x25\xf4\x8f\x02\xc1\x2c\xa2\x77\xc2\x09\xdc\x5e\x42\x61\x1b\xee\xa1\x41\x77\x68\xd7\x0c\xb6\x8a\x46\x05\x69\x94\xe2\x2c\xc3\x03\x14\x7d\x3c\xcf\x32\xfc\x2f\x62\xf4\xc7\x87\x93\xf8\x99\xe6\xcd\x2c\xd4\x28\xc3\xf5\xf0\xd0\x8f\x80\x60\x01\xea\x32\x98\xc9\x64\x75\x1a\x6e\xa1\x19\x8e\x65\xbb\x7f\xd8\x97\x83\xf2\xf8\x49\x4f\x4d\x49\x28\x48\x17\x21\xd2\xee\xcb\x32\x90\x44\x79\x51\x51\x53\x34\x9e\xbc\x51\x0d\x1f\x1e\x39\xa3\x26\x79\x61\x3a\x56\x27\x0c\x32\xff\xed\x11\x4d\xa1\xad\x4d\x6c\x7c\xbe\x1c\x94\x71\xef\xdc\x3b\xfe\xe0\x8d\x0a\xa4\xff\x3b\x7e\xdd\x26\xdb\xa7\xa4\xf8\x27\xc7\x32\x7f\x07\x00\x00\xff\xff\xd4\xfc\xc6\xa9\x89\x05\x00\x00")

func resources_language_zh_cn_yml() ([]byte, error) {
	return bindata_read(
		_resources_language_zh_cn_yml,
		"resources/language/zh-CN.yml",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"resources/language/en-US.yml": resources_language_en_us_yml,
	"resources/language/ja-JP.yml": resources_language_ja_jp_yml,
	"resources/language/zh-CN.yml": resources_language_zh_cn_yml,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"resources/language/en-US.yml": &_bintree_t{resources_language_en_us_yml, map[string]*_bintree_t{
	}},
	"resources/language/ja-JP.yml": &_bintree_t{resources_language_ja_jp_yml, map[string]*_bintree_t{
	}},
	"resources/language/zh-CN.yml": &_bintree_t{resources_language_zh_cn_yml, map[string]*_bintree_t{
	}},
}}
