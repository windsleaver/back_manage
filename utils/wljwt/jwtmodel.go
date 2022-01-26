package wljwt

import "github.com/dgrijalva/jwt-go"

/**
 * @Time   : 2022/1/25 23:53
 * @Author : WindsLeaver
 * @File   : jwtmodel
 * @Version: 1.0.0
 * @Description:
 **/

var (
	SmAuth    *SysManageCustomClaims           //后台管理
	RdAuth    *RiderCustomClaims               //骑手端
	MeAuth    *MerchantCustomClaims            //商户端
	MuaAuth   *MainUserAppCustomClaims         //用户app端
	MumAuth   *MainUserMiniProgramCustomClaims //用户微信小程序端口
	JType     JWTTYPE                          //用户类别
	IpAddress string                           //用户请求IP地址
)

// SysManageCustomClaims 后台管理CustomClaims
type SysManageCustomClaims struct {
	TenantId     string `json:"tenant_id"`      //所属租户(或校区)ID
	UserId       string `json:"user_id"`        //用户ID
	SchoolId     string `json:"school_id"`      //校区ID
	IsTenant     bool   `json:"is_tenant"`      //是否是租主[true:是; false:否]
	IsAdmin      bool   `json:"is_admin"`       //是否是管理员[true:是; false:否]
	IsBindWechat bool   `json:"is_bind_wechat"` //是否绑定微信[true:是; false:否]
	jwt.StandardClaims
}

// MerchantCustomClaims 商户端CustomClaims
type MerchantCustomClaims struct {
	TenantId   string `json:"tenant_id"`   //所属租户(或校区)ID
	SchoolId   string `json:"school_id"`   //校区ID
	MerchantId string `json:"merchant_id"` //商户ID
	StoreId    string `json:"store_id"`    //店铺ID
	UnionId    string `json:"union_id"`    //微信UnionId
	OpenId     string `json:"open_id"`     //微信OpenId
	AppleId    string `json:"apple_id"`    //苹果AppleId
	jwt.StandardClaims
}

// RiderCustomClaims 骑手APP端CustomClaims
type RiderCustomClaims struct {
	TenantId string `json:"tenant_id"` //所属租户ID
	SchoolId string `json:"school_id"` //校区ID
	RiderId  string `json:"rider_id"`  //骑手用户ID
	UnionId  string `json:"union_id"`  //微信UnionId
	OpenId   string `json:"open_id"`   //微信OpenId
	AppleId  string `json:"apple_id"`  //苹果AppleId
	jwt.StandardClaims
}

// MainUserAppCustomClaims 用户APP端CustomClaims
type MainUserAppCustomClaims struct {
	SchoolId   string `json:"school_id"`    //校区ID
	MainUserId string `json:"main_user_id"` //用户ID
	UnionId    string `json:"union_id"`     //微信UnionId
	OpenId     string `json:"open_id"`      //微信OpenId
	AppleId    string `json:"apple_id"`     //苹果AppleId
	jwt.StandardClaims
}

// MainUserMiniProgramCustomClaims 用户微信小程序端CustomClaims
type MainUserMiniProgramCustomClaims struct {
	SchoolId   string `json:"school_id"`    //校区ID
	MainUserId string `json:"main_user_id"` //用户ID
	UnionId    string `json:"union_id"`     //微信UnionId
	SmOpenId   string `json:"sm_open_id"`   //微信OpenId
	jwt.StandardClaims
}

type JwtCommonRes struct {
	Success bool        `json:"success"` //[成功:true;失败:false]
	Msg     string      `json:"msg"`     //成功或失败消息
	Data    interface{} `json:"data"`    //数据
}
