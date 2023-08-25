package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/models"
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/repository"
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/utility"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	userModelRepository repository.IUserRepository
	Cache               repository.IJTIRepository
}

func NewUserController(repo repository.IUserRepository, cache repository.IJTIRepository) *UserController {
	return &UserController{userModelRepository: repo, Cache: cache}
}

func (t UserController) Signup(c *gin.Context) {
	user := &models.User{}

	// unmarshall
	err := c.ShouldBindJSON(user)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}

	// パスワードのハッシュ化(by bcrypt)
	user.Password, err = utility.EncryptPassword(user.Password)

	user.UserID = uuid.New().String()

	// データの挿入
	err = t.userModelRepository.Create(*user)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}

	c.JSON(http.StatusOK, "")
}

func (t UserController) Login(c *gin.Context) {
	user := &models.User{}

	// unmarshall
	err := c.ShouldBindJSON(user)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}

	// ユーザーの取得
	registered := &models.User{}
	registered, err = t.userModelRepository.ReadByEmail(user.Email)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}

	// emailをチェック
	if registered.Email != user.Email {
		err := errors.New("email incorrect")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}

	// パスワードをチェック
	if !utility.IsValidPassword(registered.Password, user.Password) {
		err := errors.New("password incorrect")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}

	// トークンを生成
	tokenString, tokenJti, err := utility.GenerateToken(registered.UserID)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}

	// トークンのJTIと有効期限を保存
	err = t.Cache.Create(tokenJti)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}

	// リフレッシュトークンを生成
	refreshTokenString, refreshTokenJti, err := utility.GenerateRefreshToken(registered.UserID)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}

	// リフレッシュトークンのJTIと有効期限を保存
	err = t.Cache.Create(refreshTokenJti)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":         tokenString,
		"refresh_token": refreshTokenString,
	})
}

func (t UserController) Refresh(c *gin.Context) {
	// トークンを生成
	refreshToken := c.GetHeader("Authorization")
	if refreshToken == "" {
		c.Error(errors.New("refresh token is empty")).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, "refresh token is empty"})
		return
	}

	// remove bearer
	if len(refreshToken) < 7 || refreshToken[:7] != "Bearer " {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Unauthorized",
		})
		return
	}
	refreshToken = refreshToken[7:]

	// parse token
	climes, ok := utility.ParseRefreshToken(refreshToken)
	if !ok {
		c.Error(errors.New("failed to parse refresh token")).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, "failed to parse refresh token"})
		return
	}

	// check expired
	ok = climes.VerifyExpiresAt(time.Now().Unix(), true)
	if !ok {
		c.Error(errors.New("expired token")).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, "expired token"})
		return
	}

	// check sub
	sub, ok := climes["sub"].(string)
	if !ok {
		c.Error(errors.New("failed get sub")).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, "failed get sub"})
		return
	}
	if sub != "refresh" {
		c.Error(errors.New("invalid token")).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, "invalid token" + climes["sub"].(string)})
		return
	}

	// check user_id
	userID, ok := climes["user_id"].(string)
	if !ok {
		c.Error(errors.New("failed get user_id")).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, "failed get user_id"})
		return
	}

	// check jti
	jti, ok := climes["jti"].(string)
	if !ok {
		c.Error(errors.New("failed get jti")).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, "failed get jti"})
		return
	}

	// check jti
	isValid, err := t.Cache.IsValid(jti)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}
	if !isValid {
		c.Error(errors.New("expired token")).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, "expired token"})
		return
	}

	// delete jti
	err = t.Cache.Delete(jti)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}

	// トークンを生成
	tokenString, tokenJti, err := utility.GenerateToken(userID)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}
	err = t.Cache.Create(tokenJti)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}

	// リフレッシュトークンを生成
	refreshTokenString, refreshTokenJti, err := utility.GenerateRefreshToken(userID)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}
	err = t.Cache.Create(refreshTokenJti)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":         tokenString,
		"refresh_token": refreshTokenString,
	})
}

func (t UserController) IsAdmin(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		err := errors.New("userId is empty")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}

	authority, err := t.userModelRepository.IsAdmin(userID)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "authority": authority})
}
