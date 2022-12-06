package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/tealeg/xlsx"
)

func Run(src, dst string) {
	fileIO, err := os.ReadDir(src)
	if err != nil {
		fmt.Println("get dir:", err)
		return
	}

	for _, v := range fileIO {
		if v.IsDir() {
			continue
		}

		splitName := strings.Split(v.Name(), ".")

		if len(splitName) != 2 || splitName[1] != "xlsx" {
			continue
		}

		fstChar := splitName[0][0:1]
		if fstChar == "~" {
			continue
		}

		getExcel(src+"/"+v.Name(), dst)
	}
}

func getExcel(filename string, dst string) {

	fmt.Println("开始转换", filename)

	names := strings.Split(filename, "/")
	splitFileName := strings.Split(names[len(names)-1], ".")
	splitFileName = strings.Split(splitFileName[0], "-")
	if len(splitFileName) < 3 {
		fmt.Println("error:" + filename)
		return
	}

	file, _ := xlsx.OpenFile(filename)
	sheet := file.Sheet["Sheet1"]

	var jsonString []byte

	if splitFileName[1] == "r" {
		titles := make([]string, 0, 16)
		// 获取title
		for _, v := range sheet.Rows[0].Cells {
			titles = append(titles, v.Value)
		}

		types := make([]string, 0, 16)
		for _, v := range sheet.Rows[1].Cells {
			types = append(types, v.Value)
		}

		records := make([]map[string]interface{}, 0, 128)

		// 获取类型
		for k, v := range sheet.Rows {
			if k <= 2 {
				continue
			}
			res := make(map[string]interface{})
			for rk, rv := range v.Cells {

				rv.Value = strings.Replace(rv.Value, " ", "", -1)

				switch types[rk] {
				case "int":
					if rv.Value == "nil" {
						res[titles[rk]] = 0
						continue
					}

					if rv.Value != "" {

						pv, err := strconv.ParseInt(rv.Value, 10, 64)
						if err != nil {
							fmt.Println(err)
						}
						res[titles[rk]] = pv
					} else {
						res[titles[rk]] = 0
					}
				case "string":
					res[titles[rk]] = rv.Value
				case "float":
					if rv.Value == "nil" || rv.Value == "" {
						res[titles[rk]] = 0
						continue
					}

					pv, err := strconv.ParseFloat(rv.Value, 64)
					if err != nil {
						fmt.Println(err)
						return
					}
					res[titles[rk]] = pv
				case "[int]":
					disposedData := make([]int64, 0, 16)

					if rv.Value == "nil" || rv.Value == "" {
						res[titles[rk]] = disposedData
						continue
					}

					splitVal := strings.Split(rv.Value, ",")
					for _, sv := range splitVal {
						pv, err := strconv.ParseInt(sv, 10, 64)
						if err != nil {
							fmt.Println("[int] parse error:", sv, rv.Value)
							fmt.Println(err)
						}
						disposedData = append(disposedData, pv)
					}
					res[titles[rk]] = disposedData
				case "[string]":
					res[titles[rk]] = strings.Split(rv.Value, ",")
				case "[float]":
					disposedData := make([]float64, 0, 16)
					if rv.Value == "nil" {
						res[titles[rk]] = disposedData
						continue
					}

					splitVal := strings.Split(rv.Value, ",")
					for _, sv := range splitVal {
						pv, err := strconv.ParseFloat(sv, 64)
						if err != nil {
							fmt.Println(err)
						}
						disposedData = append(disposedData, pv)
					}
					res[titles[rk]] = disposedData
				case "json":
					jsonData := make([]map[string]interface{}, 0, 16)
					if rv.Value == "nil" {
						res[titles[rk]] = jsonData
						continue
					}

					json.Unmarshal([]byte(rv.Value), &jsonData)
					res[titles[rk]] = jsonData
				case "map":
					mapData := make(map[string]interface{})
					mapStrings := strings.Split(rv.Value, ",")
					for _, ms := range mapStrings {
						if ms == "" {
							continue
						}
						msData := strings.Split(ms, ":")
						if len(msData) != 2 {
							continue
						}
						mapData[msData[0]], _ = strconv.Atoi(msData[1])
					}
					res[titles[rk]] = mapData
				}
			}
			records = append(records, res)

		}

		var err error
		// 判断是否需要转换为object格式的json文件
		if len(splitFileName) == 4 {
			// 判断第4位是否有中括号转为数组
			reg := regexp.MustCompile(`\[(\w+)\]`)
			regStrings := reg.FindStringSubmatch(splitFileName[3])
			if len(regStrings) > 0 {
				objects := make(map[int64][]interface{})
				for _, v := range records {

					key := v[regStrings[1]].(int64)

					if _, ok := objects[key]; !ok {
						objects[key] = make([]interface{}, 0)
					}
					objects[key] = append(objects[key], v)
				}
				jsonString, err = json.MarshalIndent(objects, "", "\t")
			} else {
				objects := make(map[int64]interface{})
				for _, v := range records {
					objects[v[splitFileName[3]].(int64)] = v
				}

				jsonString, err = json.MarshalIndent(objects, "", "\t")
			}

		} else if len(splitFileName) == 5 {
			// 需要被转为object的配置必须有id字段
			objects := make(map[int64]map[int64]interface{})
			for _, v := range records {
				fstKey := v[splitFileName[3]].(int64)
				secKey := v[splitFileName[4]].(int64)
				if _, ok := objects[fstKey]; !ok {
					objects[fstKey] = make(map[int64]interface{})
				}
				objects[fstKey][secKey] = v
			}

			jsonString, err = json.MarshalIndent(objects, "", "\t")
		} else {
			jsonString, err = json.MarshalIndent(records, "", "\t")
		}

		if err != nil {
			fmt.Println(err)
			return
		}
	} else if splitFileName[1] == "c" {
		records := make(map[string]interface{})

		for _, v := range sheet.Rows {
			key := v.Cells[0].Value
			value := v.Cells[1].Value
			pv, perr := strconv.ParseInt(value, 10, 64)
			if perr == nil {
				records[key] = pv
				continue
			}

			fv, ferr := strconv.ParseFloat(value, 64)
			if ferr == nil {
				records[key] = fv
				continue
			}

			records[key] = value
		}

		jsonString, _ = json.MarshalIndent(records, "", "\t")
	}

	outFilename := fmt.Sprintf("%s%s.json", dst, splitFileName[2])
	err := os.WriteFile(outFilename, jsonString, 0666)
	if err != nil {
		fmt.Println(outFilename, "ERROR")
		fmt.Println(err)
	} else {
		fmt.Println("转换完成", outFilename)
	}
}
