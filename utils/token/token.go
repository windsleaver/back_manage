package token

import (
	"back_manage/utils/wljwt"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

/**
 * @Time   : 2022/2/17 00:16
 * @Author : WindsLeaver
 * @File   : token
 * @Version: 1.0.0
 * @Description:
 **/

// GenerateUserBackToken 生成管理后台的token
func GenerateUserBackToken(userId string) (string, error) {
	loginIn := wljwt.SysManageUserClaims{
		UserId:         userId,
		StandardClaims: jwt.StandardClaims{},
	}
	token, errs := wljwt.CreateJwtToken(wljwt.SYSMANAGE, loginIn)
	if errs != nil {
		return "", errors.New("创建Token错误")
	}
	return token, nil
}
