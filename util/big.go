package util

import (
	"math/big"
)

func BigPow(a, b int) *big.Int {
	c := new(big.Int)
	c.Exp(big.NewInt(int64(a)), big.NewInt(int64(b)), big.NewInt(0))
	return c
}

func Big(num string) *big.Int {
	n := new(big.Int)
	n.SetString(num, 0)
	return n
}

func BigD(data []byte) *big.Int {
	n := new(big.Int)
	n.SetBytes(data)
	return n
}

func BigToBytes(num *big.Int, base int) []byte {
	ret := make([]byte, base/8)
	if len(num.Bytes()) > base/8 {
		return num.Bytes()
	}
	return append(ret[:len(ret)-len(num.Bytes())], num.Bytes()...)
}

func BigCopy(src *big.Int) *big.Int {
	return new(big.Int).Set(src)
}

func BigMax(x, y *big.Int) *big.Int {
	if x.Cmp(y) <= 0 {
		return y
	}
	return x
}

func BigMin(x, y *big.Int) *big.Int {
	if x.Cmp(y) >= 0 {
		return y
	}
	return x
}
