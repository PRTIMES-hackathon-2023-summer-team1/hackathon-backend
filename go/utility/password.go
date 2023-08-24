package utility

import "golang.org/x/crypto/bcrypt"

// パスワードを受け取り、ハッシュとエラーを返す
func EncryptPassword(password string) (string, error) {
	if hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		return "", err
	} else {
		return string(hash), nil
	}
}

// ハッシュとパスワードが正しいかチェックする
func IsValidPassword(hashed string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}
