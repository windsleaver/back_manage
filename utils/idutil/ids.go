package idutil

/**
 * @Time   : 2022/2/24 00:19
 * @Author : WindsLeaver
 * @File   : ids
 * @Version: 1.0.0
 * @Description:
 **/

import (
	"back_manage/utils/snowflake"
	"math/rand"
	"sync"
)

var sfNode *snowflake.Node
var once sync.Once

func NextId() string {
	once.Do(func() {
		for retry := 0; retry < 10; retry++ {
			node, err := snowflake.NewNode(int64(rand.Int() % 1024))
			if err != nil {
				continue
			}
			sfNode = node
			break
		}
		if sfNode == nil {
			panic("snowflake id generate error")
		}
	})
	return string(sfNode.Generate().Bytes())
}
