package utils

import (
	"backend/src/models"
	"encoding/csv"
	"errors"
	"os"
	"strconv"

	"github.com/lib/pq"
)

func ParseCSV(filePath string) ([]*models.Questions, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var questions []*models.Questions
	for _, record := range records[1:] { // 1行目はヘッダーなのでスキップ
		if len(record) < 7 {
			return nil, errors.New("CSVファイルのフォーマットが正しくありません")
		}

		difficulty, err := strconv.ParseUint(record[6], 10, 32)
		if err != nil {
			return nil, err
		}

		questions = append(questions, &models.Questions{
			Question:   record[0],
			Options:    pq.StringArray{record[1], record[2], record[3], record[4]}, // 選択肢の数に応じて調整
			Supplement: record[5],
			Difficulty: uint(difficulty),
		})
	}
	return questions, nil
}
