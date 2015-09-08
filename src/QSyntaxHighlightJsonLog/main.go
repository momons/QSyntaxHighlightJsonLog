// メイン
package main

import (
	"QSyntaxHighlightJsonLog/JsonLog"
	"QSyntaxHighlightJsonLog/ConsoleLog"
	"fmt"
	"os"
	"io/ioutil"
	"flag"
)

// コマンド情報
// 入力ファイルパス
var commandInputFilePath string
// 表示階層
var commandOutputLevel int

// 終了コード
var exitCode = 0

// メイン
func main() {

	// コマンドライン取得
	setupCommand()

	// ファイルパスを出力
	ConsoleLog.Info(fmt.Sprintf("JSONファイルパス: %s", commandInputFilePath))

	// ファイル読み込み
	data, err := ioutil.ReadFile(commandInputFilePath)
	if err != nil {
		ConsoleLog.Error(fmt.Sprintf("JSONファイルの読み込みに失敗しました。%v", err))
		exitCode = 1
	} else {
		// ログ出力
		JsonLog.OutputJsonString(string(data), commandOutputLevel)
	}

	os.Exit(exitCode)
}

// コマンドライン設定
func setupCommand() {
	// 入力ファイルパス
	flag.StringVar(&commandInputFilePath, "input", "JSONファイルパス", "JSONファイルパスを指定して下さい。")
	// 表示階層
	flag.IntVar(&commandOutputLevel, "level", -1, "出力レベルを指定して下さい(-1:制限なし)。")

	flag.Parse()

	// JSONファイルパス
	if commandInputFilePath == "JSONファイルパス" {
		commandInputFilePath = ""
	}
}
