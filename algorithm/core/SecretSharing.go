package core

import (
	"eos/algorithm/utils"
	"errors"
	"math/big"
)

type Share struct {
	Value *big.Int
	Party Party
}
type Party struct {
	ID      int
	Address string
}
type SecretSharing interface {
	PrecomputeLagrangeCoefficients(partyIDs []int) ([]*big.Int, error)
	ShamirShare(secret *big.Int, threshold, numShares int) ([]Share, error)
	Recover(shares []Share, precomputedCoefficients []*big.Int) (*big.Int, error)
}

// PrecomputeLagrangeCoefficients calculates Lagrange basis polynomials at x=0 for given party IDs
func PrecomputeLagrangeCoefficients(partyIDs []int) ([]*big.Int, error) {
	coefficients := make([]*big.Int, len(partyIDs))

	for i, id := range partyIDs {
		coefficient := big.NewInt(1)
		xi := big.NewInt(int64(id))

		for j, otherId := range partyIDs {
			if i != j {
				xj := big.NewInt(int64(otherId))

				// Calculate Li(0) = Π(j≠i) (-xj)/(xi-xj)
				numerator := new(big.Int).Neg(xj)
				denominator := new(big.Int).Sub(xi, xj)
				denominatorInverse, err := utils.ModInverse(denominator)
				if err != nil {
					return nil, errors.New("failed to compute modular inverse")
				}

				coefficient = utils.ModMul(coefficient,
					utils.ModMul(numerator, denominatorInverse))
			}
		}
		coefficients[i] = coefficient
	}

	return coefficients, nil
}

// SecretShare 使用Shamir的秘密共享生成分享
func ShamirShare(secret *big.Int, threshold, numShares int) ([]Share, error) {
	if threshold < 1 || numShares < 1 {
		return nil, errors.New("threshold and numShares must be greater than 0")
	}

	// 生成Shamir秘密共享的系数
	coefficients := make([]*big.Int, threshold)
	coefficients[0] = secret
	for i := 1; i < threshold; i++ {
		coefficients[i] = utils.RandomField()
	}

	// 生成分享
	shares := make([]Share, numShares)
	for i := 1; i <= numShares; i++ {
		x := big.NewInt(int64(i))
		y := new(big.Int).Set(coefficients[threshold-1])
		for j := threshold - 2; j >= 0; j-- {
			temp := utils.ModMul(y, x)
			y = utils.ModAdd(temp, coefficients[j])
		}
		shares[i-1] = Share{Value: y, Party: Party{ID: i}}
	}
	return shares, nil
}

// Recover reconstructs the secret using precomputed Lagrange coefficients
func Recover(shares []Share, precomputedCoefficients []*big.Int) (*big.Int, error) {
	if len(shares) == 0 {
		return nil, errors.New("no shares provided")
	}
	if len(shares) != len(precomputedCoefficients) {
		return nil, errors.New("number of shares and coefficients must match")
	}

	secret := big.NewInt(0)
	for i, share := range shares {
		term := utils.ModMul(share.Value, precomputedCoefficients[i])
		secret = utils.ModAdd(secret, term)
	}

	return secret, nil
}
