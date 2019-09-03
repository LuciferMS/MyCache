package cache

//数据
type dataMap struct {
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
func InitDataMap()(instance *dataMap) {
	return &dataMap{
		data:make(map[string]string),
	}
}

//添加数据
func (instance*dataMap)set(key string, value string) (result *resultData){
	instance.data[key] = value
	return &resultData{
		code:status{
			code:SUCCESS,
		},
	}
}

//根据key获取数据
func (instance *dataMap)get(key string)(result *resultData){
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
func (instance *dataMap)del(key string)(result *resultData){
	delete(instance.data, key)
	return &resultData{
		code:status{
			code:SUCCESS,
		},
	}
}