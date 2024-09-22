package utils

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// ValidateTokenは、JWTトークンを検証し、トークンのクレームからユーザーIDを抽出
func ValidateToken(tokenString string) (string, bool, error) {
	// Bearer トークンの形式を確認
	if !strings.HasPrefix(tokenString, "Bearer ") {
		return "", false, fmt.Errorf("invalid token format")
	}

	// Bearer プレフィックスを取り除く
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// トークン検証用関数
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil // jwt.Parse 関数が返された秘密鍵を使用して、トークンを検証するため、取得した秘密鍵をバイトスライスに変換し返す
	}

	token, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		fmt.Println("Failed to parse token:", err) // デバッグ用のログ出力
		return "", false, fmt.Errorf("failed to parse token: %v", err)
	}

	// トークンのクレームを検証
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// クレームの内容を確認
		sub, subOk := claims["sub"].(string)
		exp, expOk := claims["exp"].(float64) // JSONの数値はすべて float64 として扱われるため、直接 int64 にキャストすることはできない

		if !subOk || !expOk {
			fmt.Println("Invalid token claims") // デバッグ用のログ出力
			return "", false, fmt.Errorf("invalid token claims")
		}

		// 有効期限の確認
		if time.Unix(int64(exp), 0).Before(time.Now()) {
			fmt.Println("Token has expired") // デバッグ用のログ出力
			return "", false, fmt.Errorf("token has expired")
		}
		return sub, true, nil
	}
	fmt.Println("Invalid token") // デバッグ用のログ出力
	return "", false, fmt.Errorf("invalid token")
}
