package token

import (
	"back_manage/utils/commonutil"
	"back_manage/utils/randomutil"
	"back_manage/utils/wljwt"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strings"
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
	markInfo, err := GenerateMark(wljwt.SYSMANAGE)
	if err != nil {
		return "", err
	}
	loginIn := wljwt.SysManageUserClaims{
		UserId:         userId,
		MarkInfo:       markInfo,
		StandardClaims: jwt.StandardClaims{},
	}
	token, err := wljwt.CreateJwtToken(wljwt.SYSMANAGE, loginIn)
	if err != nil {
		return "", errors.New("创建Token错误")
	}
	return token, nil
}

// GenerateMark 生成token判断标识
func GenerateMark(typeArg wljwt.JWTTYPE) (string, error) {
	typeInfo := wljwt.ConvertTokenString(typeArg)
	if typeInfo == "" {
		return "", errors.New("创建Token错误")
	}
	markTop := randomutil.GetRandomBase32String(6)
	markTail := randomutil.GetRandomBase32String(6)
	mark := commonutil.StitchingBufStr(markTop, ".", typeInfo, ".", markTail)
	mark = commonutil.Encode(mark)
	return mark, nil
}

// ParseMark 解析token判断标识
func ParseMark(markInfo, jwtType string) error {
	mark := commonutil.Decode(markInfo)
	info := strings.Split(mark, ".")
	if info[1] != jwtType {
		return errors.New("ERR_4012")
	}
	return nil
}
