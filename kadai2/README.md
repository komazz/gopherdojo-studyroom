# 仕様
- ディレクトリを指定する
- 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
- ディレクトリ以下は再帰的に処理する
- 変換前と変換後の画像形式を指定できる（オプション）


# 開発条件
- テストのしやすさを考えてリファクタリングしてみる
  - エラーハンドリングをリファクタ・makefileを追加
- テストのカバレッジを取ってみる
- テーブル駆動テストを行う
- テストヘルパーを作ってみる
  - `main_test.go`にてコマンドライン引数をセットするヘルパーを作成


# オプション
|オプション|説明|デフォルト|対応|
|:---:|:---:|:---:|:---:|
|-s|変換前の拡張子を指定|jpg|jpg・png|
|-d|変換後の拡張子を指定|png|jpg・png|


# 使い方
```shell
$ cd {Path_To_Repository}/kadai1
$ ./testdata.zsh # テストデータ生成&初期化
$ make build # ビルド
$ make test # テスト
```

```shell
$ # 単体ファイル(jpg->png)
$ ./exec ./testdata/img/azarashi.jpg
$ # 単体ファイル(png->jpg)
$ ./exec -s png -d jpg ./testdata/img/osaru.png
$ # ディレクトリ(jpg->png)
$ ./exec ./testdata/img
$ # ディレクトリ(png->jpg)
$ ./exec -s png -d jpg ./testdata/img
```


# テストデータのディレクトリ構造
```
testdata
├── err
│   ├── read_permission
│   ├── read_permission.jpg
│   └── write_permission
│       └── write_permission.jpg
└── img
    ├── azarashi.jpg
    ├── tanuki.jpg
    ├── osaru.png
    └── img
        ├── azarashi.jpg
        └── tanuki.jpg
```


# io.Readerとio.Writerについて
## 概要
- データの入出力(読み書き)を抽象化するインターフェース
- 標準パッケージでは`bytes.Buffer`、`os.File`、`image.Image`などで実装されている


## 利点
- 関数が外部の仕様を知りすぎない
  - 例えばio.Readerを引数に持つ関数は、その引数がio.Readerを満たしていることだけ知っていればよく、その内部実装にまで依存しなくていい
- 仕様変更に強い
  - 読み込み元・書き込み先が変更になっても、その関数の実装の変更なしに、io.Reader/io.Writerを満たしている型と容易に付け替えられる
- テスト時にモック・スタブを付け替えやすい


## 参考
- [io.Reader](https://golang.org/pkg/io/#Reader)
- [io.Writer](https://golang.org/pkg/io/#Writer)
- [bytes.Buffer](https://golang.org/pkg/bytes/#Buffer)
- [os.File](https://golang.org/pkg/os/#File)
- [image.Image](https://golang.org/pkg/image/#Image)
