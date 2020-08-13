package imgconv

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Converter 画像の拡張子変換を行う型
type Converter struct {
	SrcExt string
	DstExt string
}

var (
	supportExt = map[string]string{
		"png":  ".png",
		"jpg":  ".jpg",
		"jpeg": ".jpg",
	}
)

// ValidExt サポートされている拡張子か確認する関数
func ValidExt(ext string) bool {
	_, ok := supportExt[ext]
	return ok
}

// NewConverter Converter型を生成する関数
func NewConverter(srcExt, dstExt string) *Converter {
	return &Converter{
		SrcExt: supportExt[srcExt],
		DstExt: supportExt[dstExt],
	}
}

// SrcFileList 変換する候補のファイルを再帰取得する関数
func (c *Converter) SrcFileList(srcPath string) ([]string, error) {
	var srcFileList []string
	err := filepath.Walk(srcPath, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == c.SrcExt {
			srcFileList = append(srcFileList, path)
		}
		return nil
	})
	return srcFileList, err
}

// Convert 画像の拡張子変換をするメインロジック関数
func (c *Converter) Convert(src string) error {
	// 入力ファイルを取得する
	srcfile, err := os.Open(filepath.Clean(src))
	if err != nil {
		return err
	}
	defer func() error {
		if err := srcfile.Close(); err != nil {
			return err
		}
		return nil
	}()

	// 読み出す(decode)
	img, err := c.Decode(srcfile)
	if err != nil {
		return err
	}

	// 出力ファイルを作成する
	dst := c.createDstFileName(src)
	dstfile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() error {
		if err := dstfile.Close(); err != nil {
			return err
		}
		return nil
	}()

	// 書き出す(encode)
	err = c.Encode(dstfile, img)
	if err != nil {
		return err
	}
	return nil
}

// Decode 画像をデコードする関数
func (c *Converter) Decode(srcfile io.Reader) (image.Image, error) {
	var img image.Image
	var err error
	switch c.SrcExt {
	case ".jpg":
		img, err = jpeg.Decode(srcfile)
	case ".png":
		img, err = png.Decode(srcfile)
	}
	return img, err
}

// Encode 画像をエンコードする関数
func (c *Converter) Encode(dstfile io.Writer, img image.Image) error {
	var err error
	switch c.DstExt {
	case ".png":
		err = png.Encode(dstfile, img)
	case ".jpg":
		err = jpeg.Encode(dstfile, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
	}
	return err
}

// createDstFileName 画像の出力先ファイル名を作成する関数
func (c *Converter) createDstFileName(src string) string {
	oldExt := filepath.Ext(src)
	newExt := c.DstExt
	dst := strings.Replace(src, oldExt, newExt, 1)
	return dst
}
