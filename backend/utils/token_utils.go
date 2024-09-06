package utils

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// ValidateTokenは、JWTトークンを検証し、トークンのクレームからユーザーIDを抽出します
func ValidateToken(tokenString string) (string, bool, error) {
	// Bearer トークンの形式を確認
	if !strings.HasPrefix(tokenString, "Bearer ") {
		return "", false, fmt.Errorf("invalid token format")
	}

	// Bearer プレフィックスを取り除く
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // jwt.SigningMethodHS256 で生成されたトークンを、HMAC 系列で検証
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return "", false, fmt.Errorf("failed to parse token: %v", err)
	}

	// トークンのクレームを検証
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// クレームの内容を確認
		sub, subOk := claims["sub"].(string)
		exp, expOk := claims["exp"].(float64)

		if !subOk || !expOk {
			return "", false, fmt.Errorf("invalid token claims")
		}

		// 有効期限の確認
		if time.Unix(int64(exp), 0).Before(time.Now()) {
			return "", false, fmt.Errorf("token has expired")
		}
		return sub, true, nil
	}
	return "", false, fmt.Errorf("invalid token")
}
