package hsresult

import (
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
)

/**
 * @Time   : 2022/1/26 23:53
 * @Author : WindsLeaver
 * @File   : hsresult
 * @Version: 1.0.0
 * @Description:
 **/

type HsResult struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func HsResponse(w http.ResponseWriter, resp interface{}, err error) {
	var result HsResult
	if err != nil {
		result.Success = false
		result.Msg = err.Error()
	} else {
		result.Success = true
		result.Data = resp
		result.Msg = "OK"
	}
	httpx.OkJson(w, result)
}
