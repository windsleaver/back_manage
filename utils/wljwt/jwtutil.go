package wljwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

/**
 * @Time   : 2022/1/25 23:13
 * @Author : WindsLeaver
 * @File   : jwtutil
 * @Version: 1.0.0
 * @Description:
 **/

type JWTTYPE int32

const (
	_                   JWTTYPE = iota
	SYSMANAGE                   //后台管理
	MERCHANT                    //商户
	RIDER                       //骑手
	MAINUSERAPP                 //用户APP
	MAINUSERMINIPROGRAM         //用户微信小程序
)

var (
	SysManage_SecretKey           = []byte("C41638FCE3C1DD49D2BECF186EE57EF2")
	Merchant_SecretKey            = []byte("10B3F45148A98C6FB1CC59D4E6CDED08")
	Rider_SecretKey               = []byte("4EC9280256BE3760363AFEC015B309F0")
	MainUserApp_SecretKey         = []byte("4F512FA2CF415B23ED750451B8CD40AE")
	MainUserMiniProgram_SecretKey = []byte("41C02C48FA13E952D2F001F018DB22E4")
)

// CreateJwtToken 生成Token
func CreateJwtToken(typeArg JWTTYPE, data interface{}) (string, error) {
	var (
		key   []byte
		token *jwt.Token
	)
	switch typeArg {
	case SYSMANAGE:
		key = SysManage_SecretKey
		claims := data.(SysManageCustomClaims)
		claims.StandardClaims = generateStandardClaims()
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	case MERCHANT:
		key = Merchant_SecretKey
		claims := data.(MerchantCustomClaims)
		claims.StandardClaims = generateStandardClaims()
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	case RIDER:
		key = Rider_SecretKey
		claims := data.(RiderCustomClaims)
		claims.StandardClaims = generateStandardClaims()
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	case MAINUSERAPP:
		key = MainUserApp_SecretKey
		claims := data.(MainUserAppCustomClaims)
		claims.StandardClaims = generateStandardClaims()
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	case MAINUSERMINIPROGRAM:
		key = MainUserMiniProgram_SecretKey
		claims := data.(MainUserMiniProgramCustomClaims)
		claims.StandardClaims = generateStandardClaims()
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	default:
		return "", nil
	}
	return token.SignedString(key)
}

// ParseToken 解析Token
func ParseToken(typeArg JWTTYPE, token string) (interface{}, error) {
	var (
		jwtToken *jwt.Token
		err      error
	)
	switch typeArg {
	case SYSMANAGE:
		jwtToken, err = jwt.ParseWithClaims(token, &SysManageCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return SysManage_SecretKey, nil
		})
		if err != nil || jwtToken == nil {
			return nil, err
		}
		claim, ok := jwtToken.Claims.(*SysManageCustomClaims)
		if ok && jwtToken.Valid {
			return claim, nil
		}
	case MERCHANT:
		jwtToken, err = jwt.ParseWithClaims(token, &MerchantCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return Merchant_SecretKey, nil
		})
		if err != nil || jwtToken == nil {
			return nil, err
		}
		claim, ok := jwtToken.Claims.(*MerchantCustomClaims)
		if ok && jwtToken.Valid {
			return claim, nil
		}
	case RIDER:
		jwtToken, err = jwt.ParseWithClaims(token, &RiderCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return Rider_SecretKey, nil
		})
		if err != nil || jwtToken == nil {
			return nil, err
		}
		claim, ok := jwtToken.Claims.(*RiderCustomClaims)
		if ok && jwtToken.Valid {
			return claim, nil
		}
	case MAINUSERAPP:
		jwtToken, err = jwt.ParseWithClaims(token, &MainUserAppCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return MainUserApp_SecretKey, nil
		})
		if err != nil || jwtToken == nil {
			return nil, err
		}
		claim, ok := jwtToken.Claims.(*MainUserAppCustomClaims)
		if ok && jwtToken.Valid {
			return claim, nil
		}
	case MAINUSERMINIPROGRAM:
		jwtToken, err = jwt.ParseWithClaims(token, &MainUserMiniProgramCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return MainUserMiniProgram_SecretKey, nil
		})
		if err != nil || jwtToken == nil {
			return nil, err
		}
		claim, ok := jwtToken.Claims.(*MainUserMiniProgramCustomClaims)
		if ok && jwtToken.Valid {
			return claim, nil
		}
	default:
		return nil, nil
	}
	return nil, nil
}

func ConvertTokenArg(typeArg string) JWTTYPE {
	var r JWTTYPE
	switch typeArg {
	case "SYSMANAGE":
		r = SYSMANAGE
	case "MERCHANT":
		r = MERCHANT
	case "RIDER":
		r = RIDER
	case "MAINUSERAPP":
		r = MAINUSERAPP
	case "MAINUSERMINIPROGRAM":
		r = MAINUSERMINIPROGRAM
	}
	return r
}

// generateStandardClaims 生成StandardClaims
func generateStandardClaims() jwt.StandardClaims {
	return jwt.StandardClaims{
		Audience:  "yt_cloud",
		ExpiresAt: time.Now().Add(48 * time.Hour).Unix(),
		Id:        "12156120115412",
		IssuedAt:  time.Now().Unix(),
		Issuer:    "hs_lynnss",
		NotBefore: time.Now().Unix(),
		Subject:   "yt",
	}
}
