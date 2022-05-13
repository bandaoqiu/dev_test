package entity

import (
	"errors"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type ValidErr error

var (
	Required ValidErr = errors.New("不能为空")
)

type User struct {
	FirstName string `json:"first_name" valid:"required=true,min=2"`
	LastName  string `json:"last_name" valid:"required=true,min=2"`
	Email     string `json:"email" valid:"required=true,type=email,unique=true"`
	Password  string `json:"password" valid:"required=true,type=alphanumber,min=6,max=16"`
}

func (t *User) Validate() (err error) {
	v := reflect.ValueOf(*t)
	numField := v.NumField()
	for i := 0; i < numField; i++ {
		fieldTag := v.Type().Field(i).Tag.Get("valid")

		fieldName := v.Type().Field(i).Name
		//fieldType := v.Field(i).Type()
		fieldVal := v.Field(i).Interface().(string)
		if fieldTag == "" || fieldTag == "-" {
			continue
		}
		tags := strings.Split(fieldTag, ",")

		for _,v := range tags {

			spv := strings.Split(v,"=")

			if spv[0] == "required" && spv[1] == "true"{
				if fieldVal == ""{
					err = errors.New("字段：" + fieldName + "不能为空")
					return err
				}
			}

			if spv[0] == "min"{
				spv1 := spv[1]
				spvInt,err := strconv.Atoi(spv1)
				if err != nil{
					return err
				}
				if len(fieldVal) < spvInt{
					err = errors.New("字段：" + fieldName + "长度最少为" + spv1)
					return err
				}
			}
			if spv[0] == "max"{
				spv1 := spv[1]
				spvInt,err := strconv.Atoi(spv1)
				if err != nil{
					return err
				}
				if len(fieldVal) > spvInt{
					err = errors.New("字段：" + fieldName + "长度最大为" + spv1)
					return err
				}
			}
			if spv[0] == "type" && spv[1] == "alphanumber"{
				reg := regexp.MustCompile(`[a-zA-Z0-9]{6,16}`)
				b := reg.MatchString(fieldVal)
				if !b{
					err = errors.New("字段:"+ fieldName + "不匹配")
					return err
				}
			}
			if spv[0] == "type" && spv[1] == "email"{
				reg := regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)
				b := reg.MatchString(fieldVal)
				if !b {
					err = errors.New("字段:" + fieldName + "不是正确邮箱")
					return err
				}
			}
		}
	}
	return nil
}
