package imgconv

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
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
	errUnkownDecode = fmt.Errorf("Unkown Decode Error")
	errUnkownEncode = fmt.Errorf("Unkown Encode Error")
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

// Convert 画像の拡張子変換をするメインロジック関数
func (c *Converter) Convert(src string) error {
	// 入力ファイルを取得する
	srcfile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcfile.Close()

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
	defer dstfile.Close()

	// 書き出す(encode)
	err = c.Encode(dstfile, img)
	if err != nil {
		return err
	}
	return nil
}

// Decode 画像をデコードする関数
func (c *Converter) Decode(srcfile *os.File) (image.Image, error) {
	var img image.Image
	switch c.SrcExt {
	case ".jpg":
		img, err := jpeg.Decode(srcfile)
		return img, err
	case ".png":
		img, err := png.Decode(srcfile)
		return img, err
	default:
		return img, errUnkownDecode
	}
}

// Encode 画像をエンコードする関数
func (c *Converter) Encode(dstfile *os.File, img image.Image) error {
	switch c.DstExt {
	case ".png":
		err := png.Encode(dstfile, img)
		return err
	case ".jpg":
		err := jpeg.Encode(dstfile, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
		return err
	default:
		return errUnkownEncode
	}
}

// createDstFileName 画像の出力先ファイル名を作成する関数
func (c *Converter) createDstFileName(src string) string {
	oldExt := filepath.Ext(src)
	newExt := c.DstExt
	dst := strings.Replace(src, oldExt, newExt, 1)
	return dst
}
