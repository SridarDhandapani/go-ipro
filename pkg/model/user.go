package model

import (
	"strconv"

	"github.com/SridarDhandapani/go-ipro/pkg/enum"
	"gopkg.in/validator.v2"
)

func encrypt(strInput string, strKey string, iSize int) string {
	var strCode string
	var maskCode int
	var asciiCode int
	var asciiKey int

	for i := 0; i < iSize; i++ {
		if i < len(strInput) {
			asciiCode = int(strInput[i])
		} else {
			asciiCode = 0
		}
		if i != 0 {
			asciiCode ^= maskCode
		}
		for j := 0; j < len(strKey); j++ {
			asciiKey = int(strKey[j])
			asciiCode = (asciiCode ^ asciiKey) - asciiKey
		}
		maskCode = asciiCode & 0xFF
		hexCode := strconv.FormatInt(int64(maskCode), 16)
		if maskCode < 16 {
			strCode += "0"
		}
		strCode += hexCode
	}
	return strCode
}

type User struct {
	Name        string   `url:"name" validate:"nonzero,min=1,max=32"`
	Password    string   `url:"password" validate:"nonzero,min=8,max=32"`
	AccessLevel enum.ACL `url:"access_level,omitempty" validate:"min=0,max=3"`
}

func (u *User) Encrypt(token string) error {
	if err := validator.Validate(u); err != nil {
		return err
	}

	u.Name = encrypt(u.Name, token, 32)
	u.Password = encrypt(u.Password, token, 32)

	return nil
}
