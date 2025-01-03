package core

import "math/big"

type FiniteField interface {
	init()
	RandomField() *big.Int
	ModAdd(a, b *big.Int) *big.Int
	ModSub(a, b *big.Int) *big.Int
	ModMul(a, b *big.Int) *big.Int
	ModInverse(a *big.Int) *big.Int
	extendedGCD(a, b *big.Int) (*big.Int, *big.Int, *big.Int)
}

type SecretSharing interface {
	PrecomputeLagrangeCoefficients(partyIDs []int) ([]*big.Int, error)
	ShamirShare(secret *big.Int, threshold, numShares int) ([]*big.Int, error)
	Recover(shares, precomputedCoefficients []*big.Int) (*big.Int, error)
}
type Beaver interface {
	GenerateBeaverTriple() BeaverTriple
	BeaverAdd(x, y, triple_A, triple_B *big.Int) (*big.Int, *big.Int)
	BeaverMul(x, y, triple_A, triple_B, triple_C *big.Int) *big.Int
}
