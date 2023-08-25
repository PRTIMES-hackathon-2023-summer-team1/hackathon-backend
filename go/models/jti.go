package models

import "time"

type JWTJTI struct {
	JTI        string `gorm:"primaryKey" json:"jti"`
	ExpireDate int64  `json:"exp"`
}

func (jit *JWTJTI) IsValid() bool {
	return jit.ExpireDate > time.Now().Unix()
}
