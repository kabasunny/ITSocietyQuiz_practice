package services

import (
	"backend/dto"
	"backend/models"
	"backend/repositories"
	"backend/utils" // ValidateToken(tokenString string) (string, bool, error)
	"errors"

	"gorm.io/gorm"
)

type IAnswersService interface {
	SaveAnswers(input dto.AnswersInput, tokenString string) error
	// ValidateToken(tokenString string) (string, bool, error) // トークンの検証メソッド　utilsにて共通化処理とする
}

type AnswersService struct {
	repository repositories.IAnswersRepository
}

func NewAnswersService(repository repositories.IAnswersRepository) IAnswersService {
	return &AnswersService{repository: repository}
}

func (s *AnswersService) SaveAnswers(input dto.AnswersInput, tokenString string) error {
	// トークンの検証とEmpIDの抽出
	empID, valid, err := utils.ValidateToken(tokenString)
	if err != nil || !valid {
		return err
	}

	// 直近の回答を取得
	latestAnswer, err := s.repository.GetLatestAnswer(empID, input.QuestionID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) { // 特定のエラーErrRecordNotFoundは無視
		return err
	}

	// 連続正解数の更新
	streakCount := uint(0)
	if latestAnswer != nil { // 初回解答時は、nil
		if input.AnswerID == 0 {
			streakCount = latestAnswer.StreakCount + 1
		} else {
			streakCount = 0
		}
	}

	// 回答の保存
	answers := models.Answers{
		EmpID:       empID,            // トークンから抽出
		QuestionID:  input.QuestionID, // 引数から
		AnswerID:    input.AnswerID,   // 引数から
		StreakCount: streakCount,      // 更新後の値
	}

	err = s.repository.CreateAnswers(&answers)
	if err != nil {
		return err
	}

	return nil
}

// // トークンの検証メソッド　utilsにて共通化処理とする
// func (s *AnswersService) ValidateToken(tokenString string) (string, bool, error) {
// 	// Bearer トークンの形式を確認
// 	if !strings.HasPrefix(tokenString, "Bearer ") {
// 		return "", false, fmt.Errorf("invalid token format")
// 	}

// 	// Bearer プレフィックスを取り除く
// 	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // jwt.SigningMethodHS256 で生成されたトークンを、HMAC 系列で検証
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte(os.Getenv("SECRET_KEY")), nil
// 	})
// 	if err != nil {
// 		return "", false, fmt.Errorf("failed to parse token: %v", err)
// 	}

// 	// トークンのクレームを検証
// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		// クレームの内容を確認
// 		sub, subOk := claims["sub"].(string)
// 		exp, expOk := claims["exp"].(float64)

// 		if !subOk || !expOk {
// 			return "", false, fmt.Errorf("invalid token claims")
// 		}

// 		// 有効期限の確認
// 		if time.Unix(int64(exp), 0).Before(time.Now()) {
// 			return "", false, fmt.Errorf("token has expired")
// 		}
// 		return sub, true, nil
// 	}
// 	return "", false, fmt.Errorf("invalid token")
// }
