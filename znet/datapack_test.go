package znet

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println(11111111111)

}

// 负责测试拆包和封包
func TestDataPack(t *testing.T) {

	pack := NewDataPack()
	message := NewMessage(12, []byte("cahng"))
	bytes, err := pack.Pack(message)
	fmt.Println(bytes, "\n", err)

}
