# Q syntax highlight json log

GO言語で作成しております。

指定JSONファイルの内容をシンタックスハイライトでコンソール出力します。

元々はプロパティの内容をわかりやすく出力するために作ったので、正規表現を使っていないゴリゴリなプログラムになっています。

## 機能

- 指定のJSONファイルをシンタックスハイライトでコンソールに出力します。
- 出力する階層を指定できます。指定階層以下は省略されます。

## コマンドについて

`Sample.json`のパスを指定する。

```
> QSyntaxHighlightJsonLog -input /usr/hoge/QSyntaxHighlightJsonLog/Sample.json
```

3階層まで出力する。(デフォルトは-1:指定なし)

```
> QSyntaxHighlightJsonLog -level 3
```

## ビルド環境について

このライブラリのビルドは<br>
https://github.com/constabulary/gb<br>
で行っております。
