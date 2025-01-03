package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type FiniteField interface {
	RandomField() *big.Int
	ModAdd(a, b *big.Int) *big.Int
	ModSub(a, b *big.Int) *big.Int
	ModMul(a, b *big.Int) *big.Int
	ModInverse(a *big.Int) (*big.Int, error)
}

// 定义质数 p，作为有限域的模
var p = big.NewInt(101) // 这里我们选择 p = 101，作为有限域的模

// ModAdd 执行有限域上的加法 (a + b) % p
func ModAdd(a, b *big.Int) *big.Int {
	return new(big.Int).Mod(new(big.Int).Add(a, b), p)
}

func ModSub(a, b *big.Int) *big.Int {
	// 在有限域中，减法等价于 a + (-b)
	// 首先计算 b 的加法逆元 (p - b)
	negB := new(big.Int).Sub(p, b)
	// 然后执行 (a + (-b)) % p
	return ModAdd(a, negB)
}

// ModMul 执行有限域上的乘法 (a * b) % p
func ModMul(a, b *big.Int) *big.Int {
	return new(big.Int).Mod(new(big.Int).Mul(a, b), p)
}

// ModInverse 执行有限域上的求逆运算 a^(-1) mod p
func ModInverse(a *big.Int) (*big.Int, error) {
	// 使用扩展欧几里得算法计算模逆
	g, x, _ := extendedGCD(a, p)
	if g.Cmp(big.NewInt(1)) != 0 {
		return nil, fmt.Errorf("modular inverse does not exist")
	}
	return new(big.Int).Mod(x, p), nil
}

// extendedGCD 扩展欧几里得算法，计算 a 和 p 的最大公约数，以及 a 的逆元
func extendedGCD(a, b *big.Int) (*big.Int, *big.Int, *big.Int) {
	zero := big.NewInt(0)
	x0, x1 := big.NewInt(1), big.NewInt(0)
	y0, y1 := big.NewInt(0), big.NewInt(1)

	for b.Cmp(zero) != 0 {
		q := new(big.Int).Div(a, b)
		r := new(big.Int).Mod(a, b)
		a, b = b, r

		tmpX := new(big.Int).Sub(x0, new(big.Int).Mul(q, x1))
		x0, x1 = x1, tmpX

		tmpY := new(big.Int).Sub(y0, new(big.Int).Mul(q, y1))
		y0, y1 = y1, tmpY
	}

	return a, x0, y0
}

// RandomField 生成有限域上的随机数
func RandomField() *big.Int {
	randomElement, _ := rand.Int(rand.Reader, p)
	return randomElement
}

func main() {
	a := big.NewInt(45) // 第一个数
	b := big.NewInt(25) // 第二个数

	// 加法运算 (a + b) % p
	sum := ModAdd(a, b)
	fmt.Printf("%d + %d mod %d = %d\n", a, b, p, sum)

	// 乘法运算 (a * b) % p
	product := ModMul(a, b)
	fmt.Printf("%d * %d mod %d = %d\n", a, b, p, product)

	// 求逆运算 a^(-1) mod p
	inv, err := ModInverse(a)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Inverse of %d mod %d = %d\n", a, p, inv)
	}
}
