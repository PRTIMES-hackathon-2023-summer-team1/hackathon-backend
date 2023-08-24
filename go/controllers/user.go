package controllers

import (
	"net/http"

	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/models"
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/repository"
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/utility"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userModelRepository repository.IUserRepository
}

func NewUserController(repo repository.IUserRepository) *UserController {
	return &UserController{userModelRepository: repo}
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

	// データの挿入
	err = t.userModelRepository.Create(*user)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}

	c.JSON(http.StatusOK, "")
}

/* 認証方法未決定, 未実装 */
func (t UserController) Login(c *gin.Context) {
	user := &models.User{}

	// unmarshall
	err := c.ShouldBindJSON(user)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}

	// userIDを用いて登録されたユーザ情報を取得
	registered := &models.User{}
	registered, err = t.userModelRepository.Read(user.UserID)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}

	// パスワードをチェック
	if !utility.IsValidPassword(registered.Password, user.Password) {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}

	c.JSON(http.StatusOK, "")
}

func (t UserController) IsAdmin(c *gin.Context) {
	userId := c.Query("user-id")
	authority, err := t.userModelRepository.IsAdmin(userId)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(APIError{http.StatusBadRequest, err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "authority": authority})
}
