package commonutil

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"math"
	"net/url"
	"reflect"
	"strconv"
	"unsafe"
)

/**
 * @Time   : 2022/2/24 00:45
 * @Author : WindsLeaver
 * @File   : commonutil
 * @Version: 1.0.0
 * @Description:
 **/

// StitchingBufStr 字符串拼接
func StitchingBufStr(args ...string) (str string) {
	var buf bytes.Buffer
	for _, v := range args {
		buf.WriteString(v)
	}
	str = buf.String()
	return
}

// EncodeURL URL编码
func EncodeURL(api string, params map[string]string) (string, error) {
	reqUrl, err := url.Parse(api)
	if err != nil {
		return "", err
	}
	query := reqUrl.Query()
	for k, v := range params {
		query.Set(k, v)
	}
	reqUrl.RawQuery = query.Encode()
	result, _ := url.QueryUnescape(reqUrl.String())
	return result, nil
}

// MathCeil 向上取整数
func MathCeil(value float64) int {
	rst := math.Ceil(value)
	return int(rst)
}

// MathFloor 向下取整数
func MathFloor(value float64) int {
	rst := math.Floor(value)
	return int(rst)
}

// ConvertFloat 转换
func ConvertFloat(value float64) float64 {
	if math.IsNaN(value) {
		return 0.0
	}
	v, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return v
}

func OrderNumber(tag string, number int) string {
	var numberStr string
	if len(tag) > 0 {
		numberStr = StitchingBufStr("#", tag, "_", fmt.Sprintf("%04d", number+1))
	} else {
		numberStr = fmt.Sprintf("#%04d", number+1)
	}
	return numberStr
}

func HidePhone(phone string) string {
	newPhone := phone[:3] + "****" + phone[6:]
	return newPhone
}

func HideValue(value string) string {
	var newValue string
	if len(value) > 3 {
		newValue = string([]rune(value)[:1]) + "***"
	}
	return newValue
}

func HideFoodValue(value string) string {
	var newValue string
	if len(value) > 3 {
		newValue = string([]rune(value)[:3]) + "@@@"
	}
	return newValue
}

//转化为弧度(rad)
func rad(d float64) (r float64) {
	r = d * math.Pi / 180.0
	return
}

func DistanceLatOrLng(lon1, lat1, lon2, lat2 float64) (distance float64) {
	//赤道半径(单位m)
	const EARTH_RADIUS = 6378137
	rad_lat1 := rad(lat1)
	rad_lon1 := rad(lon1)
	rad_lat2 := rad(lat2)
	rad_lon2 := rad(lon2)
	if rad_lat1 < 0 {
		rad_lat1 = math.Pi/2 + math.Abs(rad_lat1)
	}
	if rad_lat1 > 0 {
		rad_lat1 = math.Pi/2 - math.Abs(rad_lat1)
	}
	if rad_lon1 < 0 {
		rad_lon1 = math.Pi*2 - math.Abs(rad_lon1)
	}
	if rad_lat2 < 0 {
		rad_lat2 = math.Pi/2 + math.Abs(rad_lat2)
	}
	if rad_lat2 > 0 {
		rad_lat2 = math.Pi/2 - math.Abs(rad_lat2)
	}
	if rad_lon2 < 0 {
		rad_lon2 = math.Pi*2 - math.Abs(rad_lon2)
	}
	x1 := EARTH_RADIUS * math.Cos(rad_lon1) * math.Sin(rad_lat1)
	y1 := EARTH_RADIUS * math.Sin(rad_lon1) * math.Sin(rad_lat1)
	z1 := EARTH_RADIUS * math.Cos(rad_lat1)

	x2 := EARTH_RADIUS * math.Cos(rad_lon2) * math.Sin(rad_lat2)
	y2 := EARTH_RADIUS * math.Sin(rad_lon2) * math.Sin(rad_lat2)
	z2 := EARTH_RADIUS * math.Cos(rad_lat2)
	d := math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2) + (z1-z2)*(z1-z2))
	theta := math.Acos((EARTH_RADIUS*EARTH_RADIUS + EARTH_RADIUS*EARTH_RADIUS - d*d) / (2 * EARTH_RADIUS * EARTH_RADIUS))
	distance = theta * EARTH_RADIUS
	return
}

// getBase64Table Base64表
func getBase64Table() string {
	str := "IJjkKLMNO567PQX12RVW3YZaDEFGbcdefghiABCHlSTUmnopqrxyz04stuvw89+/"
	return str
}

// Encode Base64加密
func Encode(data string) string {
	content := *(*[]byte)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&data))))
	coder := base64.NewEncoding(getBase64Table())
	return coder.EncodeToString(content)
}

// Decode Base64解密
func Decode(data string) string {
	coder := base64.NewEncoding(getBase64Table())
	result, _ := coder.DecodeString(data)
	return *(*string)(unsafe.Pointer(&result))
}
