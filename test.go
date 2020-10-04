package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 声明一个空结构体
	type cat struct {
		Name string

		// 带有结构体tag的字段
		User float32 `valid:"min=5"`
	}

	// 创建cat的实例
	ins := cat{Name: "mimi", User: 15.0}

	// 获取结构体实例的反射类型对象
	fields := reflect.ValueOf(ins)

	// 遍历结构体所有成员
	for i := 0; i < fields.NumField(); i++ {

		// 获取每个成员的结构体字段类型
		field := fields.Type().Field(i)
		valid := field.Tag.Get("valid")
		fmt.Println("1-->", field)
		if valid == "" {
			continue
		}
		fmt.Println("2--->", valid)
		value := fields.FieldByName(field.Name)

		// 输出成员名和tag
		fmt.Println("3----->", int(value.Float()))

		fmt.Println("4----->", value.Type().Kind())

	}
}
