package middlewares

import (
	"backend/src/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc { // 管理者用は、API利用が多いので、ミドルウェア内でトークンの検証を行う
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
			ctx.Abort() // 現在のリクエストの処理を中止し、残りのハンドラをスキップする
			return
		}
		sub, valid, err := utils.ValidateToken(tokenString)
		if err != nil || !valid {
			fmt.Println("Error in ValidateToken:", err)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort() // トークンが無効な場合もリクエストを中止
			return
		}
		fmt.Println("Token is valid for user:", sub)
		ctx.Next() // 次のミドルウェアまたはハンドラに進む
	}
}
