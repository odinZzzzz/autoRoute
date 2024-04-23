package autoRoute

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"unicode"
	"unicode/utf8"
)

func (a *AutoRoute) registerHandler(group string, handler interface{}) map[string]interface{} {
	// 获取结构体的反射值
	val := reflect.ValueOf(handler)

	// 如果不是结构体，则返回空 map
	if val.Kind() != reflect.Struct {
		return nil
	}

	// 创建一个空 map，用于存储结构体字段名和字段值
	result := make(map[string]interface{})
	opt := make(map[string]interface{})

	// 遍历结构体的字段
	for i := 0; i < val.NumField(); i++ {
		// 获取字段名
		fieldName := val.Type().Field(i).Name
		// 获取字段值
		fieldValue := val.Field(i).Interface()
		// 将字段名和字段值存储到 map 中
		opt[fieldName] = fieldValue
	}
	a.log(fmt.Sprintf("开始注册 %s 到 AutoHandler", group))

	// 遍历结构体的字段
	for i := 0; i < val.NumMethod(); i++ {
		// 获取字段名
		fieldName := val.Type().Method(i).Name
		// 将方法存储在 map 中
		result[fieldName] = func(args []reflect.Value) interface{} {
			// 创建结构体实例的反射值
			val := reflect.ValueOf(handler)
			// 调用方法
			resultValues := val.MethodByName(fieldName).Call(args)
			return resultValues[0].Interface()
		}
	}
	result["opt"] = opt
	autoHandlerMap[group] = result
	a.log(fmt.Sprintf("注册 %s  AutoHandler 成功", opt["HandlerName"]))
	return result

}

func (a *AutoHandler) Suc(data gin.H) gin.H {
	return gin.H{
		"code": 200,
		"data": data,
	}
}
func (a *AutoRoute) log(msg string) {
	if !a.Debug {
		return
	}
	fmt.Printf("[AutoRoute-debug]:%s \r\n", msg)
}

// 合并raw 参数 和 query参数
func mergePara(c *gin.Context) map[string]interface{} {
	// 获取所有的POST参数
	var mergedData map[string]interface{}
	if err := c.ShouldBind(&mergedData); err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if mergedData == nil {
		mergedData = make(map[string]interface{})
	}
	// 获取查询参数
	queryData := c.Request.URL.Query()
	for key, value := range queryData {
		mergedData[key] = value[0]
	}

	return mergedData
}
func FormatParam[T any](m map[string]interface{}, s T) T {
	// 获取结构体类型
	structType := reflect.TypeOf(s)
	// 创建结构体实例
	structValue := reflect.New(structType).Elem()

	// 遍历 map 中的键值对
	for key, value := range m {
		// 获取结构体字段 首字母主动大写匹配struct
		field := structValue.FieldByName(Capitalize(key))
		// 如果字段存在且可设置
		if field.IsValid() && field.CanSet() {
			// 获取字段类型
			fieldType := field.Type()
			// 将 interface{} 类型的值转换为字段类型并设置到结构体中
			fieldValue := reflect.ValueOf(value).Convert(fieldType)
			field.Set(fieldValue)
		}
	}

	return structValue.Interface().(T)
}
func Capitalize(s string) string {
	if s == "" {
		return s
	}
	r, size := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[size:]
}
