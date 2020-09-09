## 仕様
- 分割ダウンロードを行う
  - Rangeアクセスを用いる
  - いくつかのゴルーチンでダウンロードしてマージする
  - エラー処理を工夫する
    - golang.org/x/sync/errgourpパッケージなどを使用
  - キャンセルが発生した場合の実装を行う
    - `ctrl + c`(Interrupt)を検知して終了するように実装
    - `context.WithTimeout`にてタイムアウト処理実装


## 使い方
```shell
$ cd {Path_To_Repository}/kadai3-2
$ make build  # ビルド
$ ./run http://flat-icon-design.com/f/f_object_174/s512_f_object_174_0bg.png # 実行
$ make test   # テスト
```
