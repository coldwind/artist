package iutils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

func StringToInt64(num string) (int64, error) {
	dec, err := decimal.NewFromString(num)
	if err != nil {
		return 0, err
	}

	return dec.IntPart(), nil
}

func StringToInt32(num string) (int32, error) {
	dec, err := decimal.NewFromString(num)
	if err != nil {
		return 0, err
	}

	return int32(dec.IntPart()), nil
}

func StringToInt(num string) (int, error) {
	dec, err := decimal.NewFromString(num)
	if err != nil {
		return 0, err
	}

	return int(dec.IntPart()), nil
}

func StringToUint64(num string) (uint64, error) {
	dec, err := decimal.NewFromString(num)
	if err != nil {
		return 0, err
	}

	return uint64(dec.IntPart()), nil
}

func StringToFloat64(num string) float64 {
	dec, err := decimal.NewFromString(num)
	if err != nil {
		return 0
	}

	return dec.InexactFloat64()
}

func StringToIntArrBySep(num string, sep string) ([]int, error) {
	arr := strings.Split(num, sep)
	vals := make([]int, 0, len(arr))
	for _, v := range arr {
		if val, err := StringToInt(v); err == nil {
			vals = append(vals, val)
		} else {
			fmt.Println(err)
			return nil, errors.New("convert error")
		}
	}

	return vals, nil
}

func StringToInt32ArrBySep(num string, sep string) ([]int32, error) {
	arr := strings.Split(num, sep)
	vals := make([]int32, 0, len(arr))
	for _, v := range arr {
		if val, err := StringToInt32(v); err == nil {
			vals = append(vals, val)
		} else {
			fmt.Println(err)
			return nil, errors.New("convert error")
		}
	}

	return vals, nil
}

func StringToInt64ArrBySep(num string, sep string) ([]int64, error) {
	arr := strings.Split(num, sep)
	vals := make([]int64, 0, len(arr))
	for _, v := range arr {
		if val, err := StringToInt64(v); err == nil {
			vals = append(vals, val)
		} else {
			fmt.Println(err)
			return nil, errors.New("convert error")
		}
	}

	return vals, nil
}
