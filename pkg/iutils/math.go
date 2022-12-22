package iutils

import "github.com/shopspring/decimal"

func StringToInt64(num string) (int64, error) {
	dec, err := decimal.NewFromString(num)
	if err != nil {
		return 0, err
	}

	return dec.IntPart(), nil
}

func StringToFloat64(num string) float64 {
	dec, err := decimal.NewFromString(num)
	if err != nil {
		return 0
	}

	return dec.InexactFloat64()
}
