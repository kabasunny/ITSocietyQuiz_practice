package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// データをPythonのマイクロサービスに送信し、可視化データを取得する関数
// data: 可視化のために送信するデータ
// url: PythonマイクロサービスのエンドポイントURL
func GetGraphVisualizationData(data interface{}, url string) (map[string]interface{}, error) {
	// データをJSON 形式のバイト列に変換
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// JSONデータを含むHTTP POSTリクエストを作成
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	// HTTPクライアントを使用してリクエストを送信
	client := &http.Client{}
	resp, err := client.Do(req) // 結果を格納
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// レスポンスのステータスコードが成功を示しているか確認
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to send data, status code: %d", resp.StatusCode)
	}

	// JSONレスポンスをマップにデコード
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result) // Go構造体に変換
	if err != nil {
		return nil, err
	}

	// 可視化データを含む結果のマップを返す
	return result, nil
}
