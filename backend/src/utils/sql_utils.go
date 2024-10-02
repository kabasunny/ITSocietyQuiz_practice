package utils

import (
	"os"
)

// SQLファイルの内容を読み込み、文字列として返す関数
func LoadSQLFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
