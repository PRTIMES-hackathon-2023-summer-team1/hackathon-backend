package utility

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const SECRET_KEY = "secret_key"

func GenerateToken(userID string) (string, error) {
	// claimsオブジェクトの作成
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	// ヘッダーとペイロードの生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// トークンに署名を付与
	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken トークンの認証とuser_idの返却
func ParseToken(tokenString string) (string, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		log.Println(err)
		return "", false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)

	//トークン認証
	if !ok || !token.Valid || !claims.VerifyExpiresAt(time.Now().Unix(), false) {
		return "", false // fail
	} else {
		return userID, true // success
	}
}
