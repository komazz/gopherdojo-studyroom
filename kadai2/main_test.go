package main

import (
	"flag"
	"os"
	"path/filepath"
	"testing"
)

type args struct {
	src string
	dst string
}

var (
	testDir         = "testdata/img"
	testDirInvalid  = "testdataa"
	testfileJPG     = "azarashi.jpg"
	testfilePNG     = "osaru.png"
	testfileInvalid = "hito.jpg"
)

func TestExec(t *testing.T) {
	tests := []struct {
		name  string
		args  args
		src   string
		isErr bool
	}{
		// 異常系
		{
			name: "Error: TooFewArguments",
			args: args{
				src: "",
				dst: "",
			},
			src:   "",
			isErr: true,
		},
		{
			name: "Error: File: UnsupportedExtension",
			args: args{
				src: "",
				dst: "",
			},
			src:   filepath.Join(testDir, testfilePNG),
			isErr: true,
		},
		{
			name: "Error: Dir: UnkownDirectory",
			args: args{
				src: "",
				dst: "",
			},
			src:   testDirInvalid,
			isErr: true,
		},
		{
			name: "Error: Dir: UnsupportedExtension",
			args: args{
				src: "png",
				dst: "pdf",
			},
			src:   filepath.Join(testDir, testfileInvalid),
			isErr: true,
		},
		{
			name: "Error: File: Fail to Convert",
			args: args{
				src: "png",
				dst: "pdf",
			},
			src:   filepath.Join(testDir, testfileJPG),
			isErr: true,
		},
		// 正常系
		{
			name: "Success: File: jpg → png",
			args: args{
				src: "jpg",
				dst: "png",
			},
			src:   filepath.Join(testDir, testfileJPG),
			isErr: false,
		},
		{
			name: "Success: File: png → jpg",
			args: args{
				src: "png",
				dst: "jpg",
			},
			src:   filepath.Join(testDir, testfilePNG),
			isErr: false,
		},
		{
			name: "Success: Dir: png → jpg",
			args: args{
				src: "png",
				dst: "jpg",
			},
			src:   testDir,
			isErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := setFlag(t, tt.src, tt.args)
			if err != nil {
				t.Error(err)
			}

			err = exec()
			if !tt.isErr && err != nil {
				t.Error(err)
			}
		})
	}
}

// setFlag コマンドライン引数をセットする
func setFlag(t *testing.T, src string, args args) error {
	t.Helper()

	// set non-flag
	if src != "" {
		os.Args[1] = src
	}

	// set flag(s)
	if args.src != "" {
		err := flag.CommandLine.Set("s", args.src)
		if err != nil {
			return err
		}
	}

	// set flag(d)
	if args.dst != "" {
		err := flag.CommandLine.Set("d", args.dst)
		if err != nil {
			return err
		}
	}
	return nil
}
