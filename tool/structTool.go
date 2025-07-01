package tool

import (
	"encoding/binary"
	"fmt"
	"os"
)

// 定义函数元数据结构体
type FunctionMetadata struct {
	Name       string
	ParamTypes []string
	ReturnType string
}

// 导出函数元数据为二进制文件
func ExportFunctionMetadata(filename string, metadata FunctionMetadata) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// 写入函数名称长度和内容
	nameLen := uint32(len(metadata.Name))
	if err := binary.Write(file, binary.LittleEndian, nameLen); err != nil {
		return err
	}
	if err := binary.Write(file, binary.LittleEndian, []byte(metadata.Name)); err != nil {
		return err
	}

	// 写入参数类型数量
	paramCount := uint32(len(metadata.ParamTypes))
	if err := binary.Write(file, binary.LittleEndian, paramCount); err != nil {
		return err
	}

	// 写入每个参数类型
	for _, paramType := range metadata.ParamTypes {
		paramTypeLen := uint32(len(paramType))
		if err := binary.Write(file, binary.LittleEndian, paramTypeLen); err != nil {
			return err
		}
		if err := binary.Write(file, binary.LittleEndian, []byte(paramType)); err != nil {
			return err
		}
	}

	// 写入返回值类型长度和内容
	returnTypeLen := uint32(len(metadata.ReturnType))
	if err := binary.Write(file, binary.LittleEndian, returnTypeLen); err != nil {
		return err
	}
	if err := binary.Write(file, binary.LittleEndian, []byte(metadata.ReturnType)); err != nil {
		return err
	}

	fmt.Printf("Function metadata exported to %s\n", filename)
	return nil
}

// 从二进制文件导入函数元数据
func ImportFunctionMetadata(filename string) (FunctionMetadata, error) {
	file, err := os.Open(filename)
	if err != nil {
		return FunctionMetadata{}, err
	}
	defer file.Close()

	var metadata FunctionMetadata

	// 读取函数名称长度和内容
	var nameLen uint32
	if err := binary.Read(file, binary.LittleEndian, &nameLen); err != nil {
		return FunctionMetadata{}, err
	}
	nameBytes := make([]byte, nameLen)
	if err := binary.Read(file, binary.LittleEndian, &nameBytes); err != nil {
		return FunctionMetadata{}, err
	}
	metadata.Name = string(nameBytes)

	// 读取参数类型数量
	var paramCount uint32
	if err := binary.Read(file, binary.LittleEndian, &paramCount); err != nil {
		return FunctionMetadata{}, err
	}

	// 读取每个参数类型
	metadata.ParamTypes = make([]string, paramCount)
	for i := 0; i < int(paramCount); i++ {
		var paramTypeLen uint32
		if err := binary.Read(file, binary.LittleEndian, &paramTypeLen); err != nil {
			return FunctionMetadata{}, err
		}
		paramTypeBytes := make([]byte, paramTypeLen)
		if err := binary.Read(file, binary.LittleEndian, &paramTypeBytes); err != nil {
			return FunctionMetadata{}, err
		}
		metadata.ParamTypes[i] = string(paramTypeBytes)
	}

	// 读取返回值类型长度和内容
	var returnTypeLen uint32
	if err := binary.Read(file, binary.LittleEndian, &returnTypeLen); err != nil {
		return FunctionMetadata{}, err
	}
	returnTypeBytes := make([]byte, returnTypeLen)
	if err := binary.Read(file, binary.LittleEndian, &returnTypeBytes); err != nil {
		return FunctionMetadata{}, err
	}
	metadata.ReturnType = string(returnTypeBytes)

	return metadata, nil
}
