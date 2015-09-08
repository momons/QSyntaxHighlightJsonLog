package JsonLog

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// キー色
var colorKey string = "\x1b[36m"

// 文字列色
var colorString string = "\x1b[32m"

// 数字色
var colorNumber string = "\x1b[33m"

// Boolean色
var colorBoolean string = "\x1b[35m"

// null色
var colorNull string = "\x1b[31m"

// 色初期化
var colorInit string = "\x1b[m"

// タブスペース
var tabSpace string = "  "

// ログ出力
func Output(
	data interface{}, // 変換データ
	maxLevel int, // 最大階層 (-1:制限なし 他:階層に達すると省略)
) {
	// コンバート
	logStr := ConvertSyntaxHighlightString(data, maxLevel)
	// ログ出力
	fmt.Println(logStr)
}

// JSON文字列をログ出力
func OutputJsonString(
	jsonString string, // 変換文字列
	maxLevel int, // 最大階層 (-1:制限なし 他:階層に達すると省略)
) {

	// JSONに変換
	var jsonData interface{}
	reader := strings.NewReader(jsonString)
	dec := json.NewDecoder(reader)
	dec.Decode(&jsonData)

	// ログ出力
	Output(jsonData, maxLevel)
}

// 変数の内容をJSONを文字列に変換
func ConvertSyntaxHighlightString(
	data interface{}, // 変換データ
	maxLevel int, // 最大階層 (-1:制限なし 他:階層に達すると省略)
) string {
	return convertSyntaxHighlightStringSub(data, "", 0, maxLevel)
}

// 変数の内容をJSONを文字列に変換 サブルーチン
func convertSyntaxHighlightStringSub(
	data interface{}, // 変換データ
	beginStr string, // 初期文字列
	level int, // 現在の階層
	maxLevel int, // 最大階層 (-1:制限なし 他:階層に達すると省略)
) string {

	jsonStr := beginStr

	switch data.(type) {
	case string:
		jsonStr += (colorString + "\"" + strings.Replace(data.(string), "\"", "\\\"", -1) + "\"" + colorInit)
	case float64:
		// jsonStr += (colorNumber + strconv.FormatFloat(data.(float64), 'f', 4, 64) + colorInit)
		jsonStr += (colorNumber + fmt.Sprintf("%v", data.(float64)) + colorInit)
	case bool:
		if data.(bool) {
			jsonStr += (colorBoolean + "true" + colorInit)
		} else {
			jsonStr += (colorBoolean + "false" + colorInit)
		}
	case nil:
		jsonStr += (colorNull + "null" + colorInit)
	case []interface{}:
		jsonStr += "[\n"
		// 最大階層に達した
		if maxLevel >= 0 && level >= maxLevel {
			// 省略
			jsonStr += (addTabSpace(level+1) + "count : " + strconv.Itoa(len(data.([]interface{}))) + "\n")
		} else {
			for count, value := range data.([]interface{}) {
				jsonStr = (convertSyntaxHighlightStringSub(value, jsonStr+addTabSpace(level+1), level+1, maxLevel))
				if count < len(data.([]interface{})) - 1 {
					jsonStr += ",\n"
				} else {
					jsonStr += "\n"
				}
			}
		}
		jsonStr += (addTabSpace(level) + "]")
	case map[string]interface{}:
		jsonStr += "{\n"
		// 最大階層に達した
		if maxLevel >= 0 && level >= maxLevel {
			// 省略
			jsonStr += (addTabSpace(level+1) + "count : " + strconv.Itoa(len(data.(map[string]interface{}))) + "\n")
		} else {
			count := 0
			for key, value := range data.(map[string]interface{}) {
				jsonStr += (addTabSpace(level+1) + colorKey + "\"" + key + "\"" + colorInit + " : ")
				jsonStr = convertSyntaxHighlightStringSub(value, jsonStr, level+1, maxLevel)
				if count < len(data.(map[string]interface{})) - 1 {
					jsonStr += ",\n"
				} else {
					jsonStr += "\n"
				}
				count++
			}
		}
		jsonStr += (addTabSpace(level) + "}")
	default:
	}

	return jsonStr
}

// タブスペースを付与する
func addTabSpace(level int) string {
	retStr := ""
	for i := 0; i < level; i++ {
		retStr += tabSpace
	}
	return retStr
}
