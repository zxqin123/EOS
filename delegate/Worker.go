package delegete

import (
	"fmt"
	"math/big"
)

type WorkerImpl interface {
	new() Worker
	Work() []*big.Int
}

// 定义一个结构体
type Worker struct {
	ID      int
	Address string
}

// 为结构体实现接口方法
func (w WorkerImpl) Work() []*big.Int {
	fmt.Println("Method1 called")
	return
}
