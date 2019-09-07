package Cache

import (
	"fmt"
	"strconv"
)

//数据
type DataMap struct {
	data map[string]string
}

//返回状态码
type status struct {
	code int
}

type resultData struct {
	code status
	data string
}

const(
	SUCCESS = 200
	FAIL    = 400
)

//初始化map
func InitDataMap()(instance *DataMap) {
	return &DataMap{
		data:make(map[string]string),
	}
}

//添加数据
func (instance *DataMap)Set(key string, value string) (result *resultData){
	instance.data[key] = value
	return &resultData{
		code:status{
			code:SUCCESS,
		},
	}
}

//根据key获取数据
func (instance *DataMap)Get(key string)(result *resultData){
	value := instance.data[key]
	if value == "" {
		return &resultData{
			code:status{
				code:FAIL,
			},
			data: "没有这个key",
		}
	}
	return &resultData{
		code:status{
			code:SUCCESS,
		},
		data:value,
	}
}

//删除key
func (instance *DataMap)Del(key string)(result *resultData){
	delete(instance.data, key)
	return &resultData{
		code:status{
			code:SUCCESS,
		},
	}
}

func (instance *DataMap)Print(){
	for key, value := range instance.data{
		fmt.Printf("key : %s, value : %s", key, value)
		fmt.Println()
	}
	fmt.Println("===============美丽的分割线====================")
}

func (resultData *resultData)GetData()(result []byte){
	str := "{"
	str += "code : " + strconv.Itoa(resultData.code.code) + ","
	str += "result : " + resultData.data + "}"
	return []byte(str)
}
