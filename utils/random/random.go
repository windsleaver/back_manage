package random

import "github.com/gofrs/uuid"

/**
 * @Time   : 2022/2/4 00:04
 * @Author : WindsLeaver
 * @File   : random
 * @Version: 1.0.0
 * @Description:
 **/

// GenerateUuid uuid 随机生成
func GenerateUuid() (string, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return uid.String(), nil
}
