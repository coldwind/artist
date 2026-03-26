package iutils

import (
	"crypto/md5"
	crand "crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"net"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

func Md5(content string) string {
	h := md5.New()
	h.Write([]byte(content))
	return hex.EncodeToString(h.Sum(nil))
}

func GenSalt() string {
	rand.Seed(time.Now().Unix())
	saltByte := make([]byte, 0)
	for i := 0; i < 10; i++ {
		asciiCode := rand.Intn(26) + 65
		saltByte = append(saltByte, byte(asciiCode))
	}

	return string(saltByte)
}

func GenVerifyCode() string {

	rand.Seed(time.Now().Unix())
	code := rand.Intn(9000)
	code += 1000

	return strconv.Itoa(code)
}

func Seed() {
	rand.Seed(time.Now().UnixNano())
}

// Rand [min, max]
func Rand(min int32, max int32) int32 {
	return RealRandInt32(min, max)
}

// [min, max)
func RealRand(min, max int64) int64 {
	num := max - min
	randNum, _ := crand.Int(crand.Reader, big.NewInt(num))
	return randNum.Int64() + min
}

func RealRandInt(min, max int) int {
	num := max - min
	randNum, _ := crand.Int(crand.Reader, big.NewInt(int64(num)))
	return int(randNum.Int64()) + min
}

// RandByWeightInt32 带权重随机返回key
// weight: key=结果值，value=权重
func RandByWeightInt32(weight map[int32]int32) (int32, error) {
	// 1. 空 map 判断
	if len(weight) == 0 {
		return 0, errors.New("weight map is empty")
	}

	// 2. 计算总权重
	totalWeight := int64(0)
	for _, w := range weight {
		totalWeight += int64(w)
	}

	// 3. 总权重为0判断
	if totalWeight <= 0 {
		return 0, errors.New("total weight is zero or negative")
	}

	// 4. 生成安全随机数 [0, totalWeight-1]
	randNum, err := crand.Int(crand.Reader, big.NewInt(totalWeight))
	if err != nil {
		return 0, fmt.Errorf("generate random failed: %w", err)
	}
	target := randNum.Int64()

	// 5. map 转为有序切片（解决遍历无序问题）
	type pair struct {
		key    int32
		weight int64
	}
	var list []pair
	for k, v := range weight {
		list = append(list, pair{key: k, weight: int64(v)})
	}

	// 按 key 排序，保证每次遍历顺序一致（可选，也可按权重排序）
	sort.Slice(list, func(i, j int) bool {
		return list[i].key < list[j].key
	})

	// 6. 权重累加匹配
	var weightSum int64 = 0
	for _, item := range list {
		weightSum += item.weight
		if weightSum > target { // 正确判断
			return item.key, nil
		}
	}

	return 0, errors.New("no item selected, unexpected error")
}

func RandByWeightInt64(weight map[int64]int64) (int64, error) {
	totalWeight := int64(0)
	for _, w := range weight {
		totalWeight += w
	}
	randNum, _ := crand.Int(crand.Reader, big.NewInt(int64(totalWeight)))
	var weightSum int64 = 0
	for k, v := range weight {
		weightSum += v
		if weightSum > randNum.Int64() {
			return k, nil
		}
	}
	return 0, fmt.Errorf("failed to select weighted random int64")
}

func RealRandInt32(min, max int32) int32 {
	num := max - min
	randNum, _ := crand.Int(crand.Reader, big.NewInt(int64(num)))
	return int32(randNum.Int64()) + min
}

// BytesToInt32Array bytes转为int32数组
func BytesToInt32Array(byteData []byte) []uint32 {
	byteLen := len(byteData)
	int32Array := make([]uint32, 0, 1024)
	max := int(math.Ceil(float64(byteLen) / 4))
	for i := 0; i < max; i++ {
		begin := i * 4
		end := begin + 4
		if end > byteLen {
			end = byteLen
		}

		soleNum := binary.LittleEndian.Uint32(byteData[begin:end])
		int32Array = append(int32Array, soleNum)
	}

	return int32Array
}

// GenRandomNumber 生成count个start到end之间的随机数[start,end)
func GenRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}

// 获取IP地址
func LocalIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}
	var ip string = "localhost"
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
			}
		}
	}
	return ip
}

func ClientIpFromFasthttp(ctx *fasthttp.RequestCtx) string {
	clientIP := string(ctx.Request.Header.Peek("X-Forwarded-For"))
	if index := strings.IndexByte(clientIP, ','); index >= 0 {
		clientIP = clientIP[0:index]
	}
	clientIP = strings.TrimSpace(clientIP)
	if len(clientIP) > 0 {
		return clientIP
	}
	clientIP = strings.TrimSpace(string(ctx.Request.Header.Peek("X-Real-Ip")))
	if len(clientIP) > 0 {
		return clientIP
	}
	return ctx.RemoteIP().String()
}

// 乱序一个数组的key
func ConfusionArrayKey(arrayLen int) []int {

	array := make([]int, 0)
	for i := 0; i < arrayLen; i++ {
		array = append(array, 0)
	}

	rand.Seed(time.Now().UnixNano())
	keys := make([]int, 0)
	for i := 0; i < arrayLen; i++ {
		tLen := len(array)
		index := rand.Intn(tLen)
		array = append(array[0:index], array[index+1:]...)
		keys = append(keys, index)
	}

	return keys
}

func IsNil(data interface{}) bool {
	if data == nil {
		return true
	}

	return reflect.ValueOf(data).IsNil()
}
