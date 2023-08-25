package utility

import (
	"fmt"
	"time"

	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

const SECRET_KEY = "secret_key"

func GenerateToken(userID string) (string, *models.JWTJTI, error) {
	jti := uuid.New().String()
	timeNow := time.Now()
	exp := timeNow.Add(time.Minute * 15).Unix()
	// claimsオブジェクトの作成
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     exp,
		"sub":     "auth",
		"jti":     jti,
	}

	// ヘッダーとペイロードの生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// トークンに署名を付与
	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", nil, err
	}

	return tokenString, &models.JWTJTI{
		JTI:        jti,
		ExpireDate: exp,
	}, nil
}

func GenerateRefreshToken(userID string) (string, *models.JWTJTI, error) {
	jti := uuid.New().String()
	timeNow := time.Now()
	exp := timeNow.Add(time.Hour * 1).Unix()
	// claimsオブジェクトの作成
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     exp,
		"sub":     "refresh",
		"jti":     jti,
	}

	// ヘッダーとペイロードの生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// トークンに署名を付与
	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", nil, err
	}

	return tokenString, &models.JWTJTI{
		JTI:        jti,
		ExpireDate: exp,
	}, nil

}

func ParseToken(tokenString string) (jwt.MapClaims, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return nil, false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, false
	}
	return claims, true
}

func ParseRefreshToken(tokenString string) (jwt.MapClaims, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return nil, false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, false
	}
	return claims, true
}
