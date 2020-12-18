package middleware

import (
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"os"
)

func FirebaseAuth(c *gin.Context)  {
	// Firebase SDKのセットアップ
	opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Printf("error i")
		firebaseUninitialized(c)
		return
	}
	auth, err := app.Auth(context.Background())
	if err != nil {
		fmt.Printf("error initializing app: %v", err)
		firebaseUninitialized(c)
		return
	}

	authorizationHeader := c.Request.Header.Get("Authorization")
	if authorizationHeader == "" {
		unauthorized(c)
		return
	}

	if authorizationHeader[:7] != "Bearer " {
		unauthorized(c)
		return
	}
	tokenString := authorizationHeader[7:]

	// JWTの検証
	token, err := auth.VerifyIDToken(context.Background(), tokenString)
	if err != nil {
		// JWTが無効ならHandlerに進まず別処理
		fmt.Printf("error verifying ID token: %v\n", err)
		unauthorized(c)
		return
	}

	log.Printf("Verified ID token: %v\n", tokenString)
	c.Set("uid", token.UID)
}

func unauthorized(c *gin.Context)  {
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": http.StatusText(http.StatusUnauthorized),
	})
	c.Abort()
}

func firebaseUninitialized(c *gin.Context)  {
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": http.StatusText(http.StatusInternalServerError),
	})
	c.Abort()
}