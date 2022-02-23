package security

import (
	"back_manage/utils/commonutil"
	"golang.org/x/crypto/bcrypt"
)

/**
 * @Time   : 2022/2/24 00:43
 * @Author : WindsLeaver
 * @File   : password
 * @Version: 1.0.0
 * @Description:
 **/

//GenderPassword 生成密码
func GenderPassword(password, salt string) (string, error) {
	gPassword := commonutil.StitchingBufStr(password, salt)
	hash, err := bcrypt.GenerateFromPassword([]byte(gPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	encodePwd := string(hash)
	return encodePwd, nil
}

//VerifyPassword 密码验证
func VerifyPassword(loginPassword, salt, password string) bool {
	gPassword := commonutil.StitchingBufStr(loginPassword, salt)
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(gPassword))
	if err != nil {
		return false
	} else {
		return true
	}
}
