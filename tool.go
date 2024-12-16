package autoRoute

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"unicode"
	"unicode/utf8"
)

type handlerData struct {
	Name    string
	Define  reflect.Type
	Handler interface{}
}

func (a *handlerData) Func(args []reflect.Value) interface{} {
	// 创建结构体实例的反射值
	val := reflect.ValueOf(a.Handler)
	// 调用方法
	resultValues := val.MethodByName(a.Name).Call(args)
	return resultValues[0].Interface()
}
func (a *AutoRoute) registerHandler(group string, handler interface{}) map[string]handlerData {
	// 获取结构体的反射值
	val := reflect.ValueOf(handler)

	// 如果不是结构体，则返回空 map
	if val.Kind() != reflect.Struct {
		return nil
	}

	// 创建一个空 map，用于存储结构体字段名和字段值
	result := make(map[string]handlerData)
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
	DebugLog(fmt.Sprintf("开始注册 %s 到 AutoHandler", group))

	// 遍历结构体的字段
	for i := 0; i < val.NumMethod(); i++ {
		// 获取字段名
		method := val.Type().Method(i)
		fieldName := method.Name
		hData := handlerData{
			Name:    fieldName,
			Handler: handler,
		}
		DebugLog(fmt.Sprintf("注册 %s接口  AutoHandler 成功", fieldName))
		mType := method.Type
		if mType.NumIn() > 1 {
			paramDefine := mType.In(1)
			hData.Define = paramDefine.Elem()
		}
		result[fieldName] = hData

	}
	autoHandlerMap[group] = result
	return result

}

func (a *AutoHandler) Suc(data map[string]any) map[string]any {
	return gin.H{
		"code": 200,
		"data": data,
	}
}

func DebugLog(msg string) {
	if !RouteOpt.Debug {
		return
	}
	fmt.Printf("\033[32m[AutoRoute-Debug]:%s \033[0m\r\n", msg)
}
func WarnLog(msg string) {
	fmt.Printf("\033[33m[AutoRoute-Warn]:%s \033[0m\r\n", msg)
}
func ErrorLog(msg string) {
	fmt.Printf("\033[41m[AutoRoute-Warn]:%s \033[0m\r\n", msg)
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

			fieldValue := reflect.ValueOf(value)
			if fieldValue.CanConvert(fieldType) {
				field.Set(fieldValue.Convert(fieldType))
			} else {
				fmt.Printf(" FormatParam 【%s】 err：[%s ] value [%s]format 失败！！\r\n", structType.Name(), Capitalize(key), fieldValue)
			}

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
