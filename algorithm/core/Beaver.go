package core

import (
	"eos/algorithm/utils"
	"math/big"
)

type BeaverTriple struct {
	A *big.Int
	B *big.Int
	C *big.Int
}
type Beaver interface {
	GenerateBeaverTriple() BeaverTriple
	BeaverAdd(x, y, triple_A, triple_B Share) (Share, Share)
	BeaverMul(x, y *big.Int, triple_A, triple_B, triple_C Share) Share
}

// GenerateBeaverTriple 生成Beaver三元组
func GenerateBeaverTriple() BeaverTriple {
	a := utils.RandomField()
	b := utils.RandomField()
	c := utils.ModMul(a, b)
	return BeaverTriple{A: a, B: b, C: c}
}

// BeaverMultiplication 使用Beaver三元组实现乘法同态
// x,y是恢复的明文x-a和y-b
func BeaverAdd(x, y, triple_A, triple_B Share) (Share, Share) {
	x.Value = utils.ModSub(x.Value, triple_A.Value)
	y.Value = utils.ModSub(y.Value, triple_B.Value)
	return x, y
}
func BeaverMul(x, y *big.Int, triple_A, triple_B, triple_C Share) Share {

	// 计算结果
	xy := utils.ModMul(x, y)
	result := utils.ModAdd(utils.ModAdd(triple_C.Value, utils.ModMul(x, triple_B.Value)), utils.ModMul(y, triple_A.Value))
	result = utils.ModAdd(result, xy)

	return Share{Value: result, Party: triple_A.Party}
}
