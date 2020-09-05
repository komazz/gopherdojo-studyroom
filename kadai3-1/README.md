## 仕様
- 標準出力に英単語を出す（出すものは自由）
- 標準入力から1行受け取る
- 制限時間内に何問解けたか表示する
  - 制限時間は15秒

## 開発条件
- 制限時間にはtime.After関数を用いる
  - context.WithTimeoutを使用しました
- select構文を用いる
  - `typegame.Start()`にて使用


## 使い方
```shell
$ cd {Path_To_Repository}/kadai3-1
$ make build # ビルド
$ ./run # ゲーム開始
$ make test # テスト
```

## 実行例
```shell
$ ./run
orange
>> orange
...

------------
answer: orange input: orange
answer: banana input: yo
------------
1/2 正解
------------
```
