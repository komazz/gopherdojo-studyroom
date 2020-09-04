package downloader

import (
	"io"
	"net/http"
	"os"
)

// Run メインロジック
func Run() error {
	url := "http://flat-icon-design.com/f/f_object_174/s512_f_object_174_0bg.jpg"
	// リクエスト作成
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// リクエスト送信
	rep, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer rep.Body.Close()

	// ファイル作成
	dst := "file.jpg"
	dstfile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstfile.Close()

	if _, err := io.Copy(dstfile, rep.Body); err != nil {
		return err
	}

	return nil
}
