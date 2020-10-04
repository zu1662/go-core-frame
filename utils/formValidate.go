package utils

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

//自定义验证规则
// 组合使用
// valid:"required,oneof=10 20 30"
// valid:"required,min=5,max=10"
const (
	REQUIRED = "required" //字符串不能为空 ---> required
	MAX      = "max"      //最大值 ---> max=10
	MIN      = "min"      //最小值 ---> min=0
	EQUAL    = "equal"    // 等于 ---> eq=5
	LEN      = "len"      //长度等于 ---> len=11
	ONEOF    = "oneof"    // 数值内的其中一个 ---> oneof=37 39 41
	EMAIL    = "email"    // email 类型 ---> email
)

// StructValidate 对外暴露结构体验证函数
func StructValidate(bean interface{}) error {
	fields := reflect.ValueOf(bean)
	for i := 0; i < fields.NumField(); i++ {
		field := fields.Type().Field(i)
		valid := field.Tag.Get("valid")
		if valid == "" {
			continue
		}
		value := fields.FieldByName(field.Name)
		err := fieldValidate(field.Name, valid, value)
		if err != nil {
			return err
		}
	}
	return nil
}

//属性验证
func fieldValidate(fieldName, validStr string, value reflect.Value) error {
	valids := strings.Split(validStr, ",")
	for index, valid := range valids {

		if strings.Contains(valid, REQUIRED) {
			str := value.String()
			if str == "" {
				return errors.New(fieldName + " value can't empty")
			}
		}

		if strings.Contains(valid, MIN) || strings.Contains(valid, MAX) || strings.Contains(valid, EQUAL) {
			compareErr := compare(fieldName, valid, value)
			if compareErr == nil {
				if index < len(valids)-1 {
					continue
				}
				return nil
			}
			return compareErr
		}
	}
	return nil
}

// 针对数字或者字符串长度进行比较
func compare(fieldName string, valid string, value reflect.Value) error {
	validStrArr := strings.Split(valid, "=")
	valueType := value.Type().Kind()
	switch valueType {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		valueInt := int(value.Int())
		validInt, err := strconv.Atoi(validStrArr[1]) // 字符串转 整数
		if err != nil {
			return errors.New(fieldName + " " + validStrArr[0] + "value is not valid")
		}
		switch validStrArr[0] {
		case MIN:
			if valueInt >= validInt {
				return nil
			}
			return errors.New(fieldName + " value must >= " + strconv.Itoa(validInt))
		case MAX:
			if valueInt <= validInt {
				return nil
			}
			return errors.New(fieldName + " value must <= " + strconv.Itoa(validInt))
		case EQUAL:
			if valueInt == validInt {
				return nil
			}
			return errors.New(fieldName + " value must equal " + strconv.Itoa(validInt))
		}
	case reflect.Float32, reflect.Float64:
		valueInt := int(value.Float())
		validInt, err := strconv.Atoi(validStrArr[1]) // 字符串转 整数
		if err != nil {
			return errors.New(fieldName + " " + validStrArr[0] + "value is not valid")
		}
		switch validStrArr[0] {
		case MIN:
			if valueInt >= validInt {
				return nil
			}
			return errors.New(fieldName + " value must >= " + strconv.Itoa(validInt))
		case MAX:
			if valueInt <= validInt {
				return nil
			}
			return errors.New(fieldName + " value must <= " + strconv.Itoa(validInt))
		case EQUAL:
			if valueInt == validInt {
				return nil
			}
			return errors.New(fieldName + " value must equal " + strconv.Itoa(validInt))
		}
	case reflect.String:
		valueInt := len(value.String())
		validInt, err := strconv.Atoi(validStrArr[1]) // 字符串转 整数
		if err != nil {
			return errors.New(fieldName + " " + validStrArr[0] + "value is not valid")
		}
		switch validStrArr[0] {
		case MIN:
			if valueInt >= validInt {
				return nil
			}
			return errors.New(fieldName + " value length must >= " + strconv.Itoa(validInt))
		case MAX:
			if valueInt <= validInt {
				return nil
			}
			return errors.New(fieldName + " value length must <= " + strconv.Itoa(validInt))
		case EQUAL:
			if valueInt == validInt {
				return nil
			}
			return errors.New(fieldName + " value length must equal " + strconv.Itoa(validInt))
		}
	default:
		return errors.New(fieldName + " " + validStrArr[0] + "type is not valid")
	}
	return nil
}
