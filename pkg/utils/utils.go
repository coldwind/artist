package utils

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"crypto/md5"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"net"
	"strconv"
	"time"
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
	n := max - min + 1
	randNum := rand.Int31n(n) + min

	return randNum
}

// DoZlibCompress 进行zlib压缩
func DoZlibCompress(src []byte) []byte {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	w.Write(src)
	w.Close()
	return in.Bytes()
}

// DoZlibUnCompress 进行zlib解压缩
func DoZlibUnCompress(compressSrc []byte) []byte {
	b := bytes.NewReader(compressSrc)
	var out bytes.Buffer
	r, _ := zlib.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
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

func GZipData(data []byte) string {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	defer w.Close()

	w.Write(data)
	w.Flush()

	str := base64.StdEncoding.EncodeToString(b.Bytes())

	return str
}

func GZipParse(baseData string) string {

	decodeData, _ := base64.StdEncoding.DecodeString(baseData)
	rdata := bytes.NewReader(decodeData)

	r, _ := gzip.NewReader(rdata)

	byteData, _ := ioutil.ReadAll(r)

	return string(byteData)
}

// Int64ToBytes 字节与int互换--大端法
func Int64ToBytes(n int64) []byte {
	data := int64(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

// BytesToInt64 字节与int互换--大端法
func BytesToInt64(bys []byte) int64 {
	bytebuff := bytes.NewBuffer(bys)
	var data int64
	binary.Read(bytebuff, binary.BigEndian, &data)
	return data
}

// Int32ToBytes 字节与int互换--大端法
func Int32ToBytes(n int32) []byte {
	data := n
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

// BytesToInt32 字节与int互换--大端法
func BytesToInt32(bys []byte) int32 {
	bytebuff := bytes.NewBuffer(bys)
	var data int32
	binary.Read(bytebuff, binary.BigEndian, &data)
	return data
}

// 字节与int互换--大端法
func UInt16ToBytes(n int) []byte {
	data := uint16(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

// 字节与int互换--大端法
func BytesToUInt16(bys []byte) int {
	bytebuff := bytes.NewBuffer(bys)
	var data uint16
	binary.Read(bytebuff, binary.BigEndian, &data)
	return int(data)
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

// BASE64
func Base64Encode(bytes []byte) string {
	str := base64.StdEncoding.EncodeToString(bytes)

	return str
}
func Base64Decode(data string) []byte {
	decodeData, _ := base64.StdEncoding.DecodeString(data)

	return decodeData
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

func JsonMarshal(data interface{}) []byte {
	b, err := json.Marshal(data)
	if err != nil {
		return nil
	}

	return b
}

func JsonUnMarshal(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)

	return err
}
