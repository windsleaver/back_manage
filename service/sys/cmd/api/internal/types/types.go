// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type LoginResp struct {
	ID       string `json:"id"`
	UserName string `json:"userName"`
}
