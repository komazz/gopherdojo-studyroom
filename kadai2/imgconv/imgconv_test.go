package imgconv

import (
	"path/filepath"
	"testing"
)

var (
	testDir    = "../testdata/img"
	testDirErr = "../testdata/err"
)

func TestValidExt(t *testing.T) {
	tests := []struct {
		name    string
		ext     string
		isValid bool
	}{
		// 異常系
		{
			name:    "gif",
			ext:     "gif",
			isValid: false,
		},

		// 正常系
		{
			name:    "png",
			ext:     "png",
			isValid: true,
		},
		{
			name:    "jpg",
			ext:     "jpg",
			isValid: true,
		},
		{
			name:    "jpeg",
			ext:     "jpeg",
			isValid: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			valid := ValidExt(tt.ext)
			if valid != tt.isValid {
				t.Errorf("Expected %t, Got %t", tt.isValid, valid)
			}
		})
	}
}

func TestSrcFileList(t *testing.T) {
	tests := []struct {
		name      string
		converter Converter
		src       string
		isErr     bool
	}{
		// 異常系
		{
			name:      "Error: get jpg file list",
			converter: *NewConverter("jpg", "png"),
			src:       filepath.Join(testDirErr, "read_permission"),
			isErr:     true,
		},

		// 正常系
		{
			name:      "Success: get jpg file list",
			converter: *NewConverter("jpg", "png"),
			src:       testDir,
			isErr:     false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := tt.converter
			_, err := c.SrcFileList(tt.src)
			if !tt.isErr && err != nil {
				t.Error(err)
			}
		})
	}
}

func TestConvert(t *testing.T) {
	tests := []struct {
		name      string
		converter Converter
		src       string
		isErr     bool
	}{
		// 異常系
		{
			name:      "Error: Fail to Read",
			converter: *NewConverter("jpg", "png"),
			src:       filepath.Join(testDirErr, "read_permission.jpg"),
			isErr:     true,
		},
		{
			name:      "Error: Fail to Write",
			converter: *NewConverter("jpg", "png"),
			src:       filepath.Join(testDirErr, "write_permission/write_permission.jpg"),
			isErr:     true,
		},

		// 正常系
		{
			name:      "Success: File: jpg → png",
			converter: *NewConverter("jpg", "png"),
			src:       filepath.Join(testDir, "azarashi.jpg"),
			isErr:     false,
		},
		{
			name:      "Success: File: png → png",
			converter: *NewConverter("png", "jpg"),
			src:       filepath.Join(testDir, "osaru.png"),
			isErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.converter
			err := c.Convert(tt.src)
			if !tt.isErr && err != nil {
				t.Error(err)
			}
		})
	}
}
