package iredis

import "github.com/shopspring/decimal"

func StringSliceToUInt64Slice(slice []string) []uint64 {
	iSlice := make([]uint64, 0, len(slice))
	for _, v := range slice {
		if dec, err := decimal.NewFromString(v); err == nil {
			iSlice = append(iSlice, dec.BigInt().Uint64())
		}
	}
	return iSlice
}

func StringSliceToInt64Slice(slice []string) []int64 {
	iSlice := make([]int64, 0, len(slice))
	for _, v := range slice {
		if dec, err := decimal.NewFromString(v); err == nil {
			iSlice = append(iSlice, dec.IntPart())
		}
	}
	return iSlice
}

func StringSliceToFloat64Slice(slice []string) []float64 {
	iSlice := make([]float64, 0, len(slice))
	for _, v := range slice {
		if dec, err := decimal.NewFromString(v); err == nil {
			fnum, _ := dec.Float64()
			iSlice = append(iSlice, fnum)
		}
	}
	return iSlice
}
