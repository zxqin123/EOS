package core

import (
	"math/big"
)

type BeaverTriple struct {
	A *big.Int
	B *big.Int
	C *big.Int
}

// GenerateBeaverTriple 生成Beaver三元组
func GenerateBeaverTriple() BeaverTriple {
	a := RandomField()
	b := RandomField()
	c := ModMul(a, b)
	return BeaverTriple{A: a, B: b, C: c}
}

// BeaverMultiplication 使用Beaver三元组实现乘法同态
// x,y是恢复的明文x-a和y-b
func BeaverAdd(x, y, triple_A, triple_B *big.Int) (*big.Int, *big.Int) {
	x = ModSub(x, triple_A)
	y = ModSub(y, triple_B)
	return x, y
}
func BeaverMul(x, y, triple_A, triple_B, triple_C *big.Int) *big.Int {

	// 计算结果
	xy := ModMul(x, y)
	result := ModAdd(ModAdd(triple_C, ModMul(x, triple_B)), ModMul(y, triple_A))
	result = ModAdd(result, xy)

	return result
}
