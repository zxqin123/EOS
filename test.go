package main

import (
	"eos/core"
	"fmt"
	"log"
	"math/big"
)

//测试shamir
/* func main() {
	secret := big.NewInt(12)
	threshold := 5
	numShares := 5
	vector := make([]int, numShares)
	for i := range vector {
		vector[i] = i + 1
	}
	coefficients, err := core.PrecomputeLagrangeCoefficients(vector)
	shares, err := core.ShamirShare(secret, threshold, numShares)
	if err != nil {
		log.Fatalf("Failed to generate shares: %v", err)
	}
	fmt.Println("secret:", secret)
	fmt.Println("Shares:", shares)
	share, err := core.Recover(shares, coefficients)
	if err != nil {
		log.Fatalf("Failed to recover shares: %v", err)
	}
	fmt.Println(share)
} */

// 测试beaver
func main() {
	triple := core.GenerateBeaverTriple()
	secret_x := big.NewInt(12)
	secret_y := big.NewInt(4)
	threshold := 5
	numShares := 5
	vector := make([]int, numShares)
	for i := range vector {
		vector[i] = i + 1
	}
	coefficients, err := core.PrecomputeLagrangeCoefficients(vector)
	shares_x, err := core.ShamirShare(secret_x, threshold, numShares)
	shares_y, err := core.ShamirShare(secret_y, threshold, numShares)
	triple_A, err := core.ShamirShare(triple.A, threshold, numShares)
	triple_B, err := core.ShamirShare(triple.B, threshold, numShares)
	triple_C, err := core.ShamirShare(triple.C, threshold, numShares)
	shares_e := make([]*big.Int, numShares)
	shares_d := make([]*big.Int, numShares)
	for i := range shares_x {
		shares_e[i], shares_d[i] = core.BeaverAdd(shares_x[i], shares_y[i], triple_A[i], triple_B[i])
	}
	secret_e, err := core.Recover(shares_e, coefficients)
	secret_d, err := core.Recover(shares_d, coefficients)
	fmt.Println("x,y:", secret_x, secret_y)

	shares_xy := make([]*big.Int, numShares)
	for i := range shares_x {
		shares_xy[i] = core.BeaverMul(secret_e, secret_d, triple_A[i], triple_B[i], triple_C[i])
	}
	secret_xy, err := core.Recover(shares_xy, coefficients)
	fmt.Println("恢复的xy:", secret_xy)
	if err != nil {
		log.Fatalf("Failed to recover shares: %v", err)
	}
}
