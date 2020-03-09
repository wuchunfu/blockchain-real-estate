/**
 * @Author: 夜央 Oh oh oh oh oh oh (https://github.com/togettoyou)
 * @Email: zoujh99@qq.com
 * @Date: 2020/3/4 5:14 下午
 * @Description: sdk的测试
 */
package blockchain

import (
	"fmt"
	"testing"
)

func TestInvoke_QueryAccountList(t *testing.T) {
	Init()
	response, e := ChannelQuery("queryAccountList", [][]byte{})
	if e != nil {
		fmt.Println(e.Error())
		t.FailNow()
	}
	fmt.Println(string(response.Payload))
}
