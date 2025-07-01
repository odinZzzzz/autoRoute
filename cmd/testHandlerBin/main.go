package main

import (
	"fmt"
	"github.com/odinZzzzz/autoRoute/tool"
	"reflect"
)

func main() {
	// 定义一个函数
	myFunc := func(a int, b string) bool {
		fmt.Printf("a = %d, b = %s\n", a, b)
		return true
	}

	// 获取函数的元数据
	funcType := reflect.TypeOf(myFunc)
	metadata := tool.FunctionMetadata{
		Name:       "myFunc",
		ParamTypes: []string{funcType.In(0).String(), funcType.In(1).String()},
		ReturnType: funcType.Out(0).String(),
	}

	// 导出函数元数据
	if err := tool.ExportFunctionMetadata("func_metadata.bin", metadata); err != nil {
		fmt.Println("Error exporting function metadata:", err)
	}

	// 导入函数元数据
	importedMetadata, err := tool.ImportFunctionMetadata("func_metadata.bin")
	if err != nil {
		fmt.Println("Error importing function metadata:", err)
	} else {
		fmt.Printf("Imported Function Metadata: %+v\n", importedMetadata)
	}

}
