package core

import (
	"errors"
	"math/big"
)

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
				denominatorInverse, err := ModInverse(denominator)
				if err != nil {
					// 处理 err
					return nil, err
				}
				coefficient = ModMul(coefficient,
					ModMul(numerator, denominatorInverse))
			}
		}
		coefficients[i] = coefficient
	}

	return coefficients, nil
}

// SecretShare 使用Shamir的秘密共享生成分享
func ShamirShare(secret *big.Int, threshold, numShares int) ([]*big.Int, error) {
	if threshold < 1 || numShares < 1 {
		return nil, errors.New("threshold and numShares must be greater than 0")
	}

	// 生成Shamir秘密共享的系数
	coefficients := make([]*big.Int, threshold)
	coefficients[0] = secret
	for i := 1; i < threshold; i++ {
		coefficients[i] = RandomField()
	}

	// 生成分享
	shares := make([]*big.Int, numShares)
	for i := 1; i <= numShares; i++ {
		x := big.NewInt(int64(i))
		y := new(big.Int).Set(coefficients[threshold-1])
		for j := threshold - 2; j >= 0; j-- {
			temp := ModMul(y, x)
			y = ModAdd(temp, coefficients[j])
		}
		shares[i-1] = y
	}
	return shares, nil
}

// Recover reconstructs the secret using precomputed Lagrange coefficients
func Recover(shares, precomputedCoefficients []*big.Int) (*big.Int, error) {
	if len(shares) == 0 {
		return nil, errors.New("no shares provided")
	}
	if len(shares) != len(precomputedCoefficients) {
		return nil, errors.New("number of shares and coefficients must match")
	}

	secret := big.NewInt(0)
	for i, share := range shares {
		term := ModMul(share, precomputedCoefficients[i])
		secret = ModAdd(secret, term)
	}

	return secret, nil
}
