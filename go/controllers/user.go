package controllers

import (
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/models"
	"github.com/PRTIMES-hackathon-2023-summer-team1/hackathon-backend/repository"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	userModelRepository repository.IUserRepository
}

func NewUserController(repo repository.IUserRepository) *UserController {
	return &UserController{userModelRepository: repo}
}

// passwordEncrypt パスワードを受け取り、ハッシュとエラーを返す
func passwordEncrypt(password string) (string, error) {
	if hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		return "", err
	} else {
		return string(hash), nil
	}
}

// ハッシュとパスワードが正しいかチェックする
func checkPassword(hashed, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func (t UserController) Signup(c *gin.Context) {
	user := &models.User{}

	// unmarshall
	err := c.ShouldBindJSON(user)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		c.JSON(400, "")
		return
	}

	// パスワードのハッシュ化(by bcrypt)
	user.Password, err = passwordEncrypt(user.Password)

	// データの挿入
	err = t.userModelRepository.Create(*user)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		c.JSON(400, "")
		return
	}

	c.JSON(200, "")
}

/* 認証方法未決定, 未実装 */
func (t UserController) Login(c *gin.Context) {
	user := &models.User{}

	// unmarshall
	err := c.ShouldBindJSON(user)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		c.JSON(400, "")
		return
	}

	// userIDを用いて登録されたユーザ情報を取得
	registered := &models.User{}
	registered, err = t.userModelRepository.Read(user.UserID)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		c.JSON(400, "")
		return
	}

	// パスワードのハッシュをチェック
	if !checkPassword(user.Password, registered.Password) {
		c.Error(err).SetType(gin.ErrorTypePublic)
		c.JSON(400, "")
		return
	}

	c.JSON(200, "")
}
