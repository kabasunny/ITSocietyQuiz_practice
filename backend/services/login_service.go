package services

import (
	"backend/models"
	"backend/repositories"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type ILoginService interface {
	Login(empID string, password string) (*string, bool, int64, error)
	GetUsersFromToken(tokenString string) (*models.Users, error)
}

type LoginService struct {
	repository repositories.ILoginRepository
}

func NewLoginService(repository repositories.ILoginRepository) ILoginService {
	return &LoginService{repository: repository}
}

func (s *LoginService) Login(empID string, password string) (*string, bool, int64, error) {
	foundUsers, err := s.repository.FindUsers(empID)
	if err != nil {
		return nil, false, 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUsers.Password), []byte(password))
	if err != nil {
		return nil, false, 0, err
	}

	token, err := CreateToken(foundUsers.EmpID, foundUsers.Username)
	if err != nil {
		return nil, false, 0, err
	}

	roles, err := s.repository.FindUsersRole(empID)
	if err != nil {
		return nil, false, 0, err
	}

	isAdmin := false
	for _, role := range roles {
		if role.RoleID == 1 {
			isAdmin = true
			break
		}
	}

	todaysAnswersCount, err := s.repository.FindTodaysAnswersCount(empID)
	if err != nil {
		return nil, false, 0, err
	}

	return token, isAdmin, todaysAnswersCount, nil
}

func CreateToken(empID string, Username string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      empID,
		"username": Username,
		"exp":      time.Now().Add(time.Hour).Unix(),
		// "exp": time.Now().Add(time.Minute).Unix(), // exp有効性確認
	}) // Unixタイムスタンプは、1970年1月1日のUTC午前0時0分0秒からの経過秒数を表すため、int64型では約292億年分の秒数をカバー

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY"))) // 秘密鍵で署名する
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

func (s *LoginService) GetUsersFromToken(tokenString string) (*models.Users, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// jwt.Parseの戻り値の入る変数tokenと、func(token *jwt.Token)のtokenは同じもの
		// jwt.Parse関数は、JWTトークンを解析し、その結果をtoken変数に格納される...
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil // 秘密鍵を関数の外部から取得し、検証関数に渡すための一般的なパターン
	})
	if err != nil {
		return nil, err
	}
	var Users *models.Users
	// claimsにJWTトークンのクレームはinterface{}型で格納される
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// 現在の時間がトークンの有効期限（expクレーム）を超えているかどうか
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return nil, jwt.ErrTokenExpired
		}
		// データベースからユーザー情報を取得
		Users, err = s.repository.FindUsers(claims["sub"].(string))
		if err != nil {
			return nil, err
		}
	}
	return Users, nil
}
