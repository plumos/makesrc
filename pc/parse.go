package pc

import (
	"reflect"
	"strings"
)

func Parse(st interface{}){
	typeOfCat := reflect.TypeOf(st)
	var names []string
	var tags []string
	var types []string
	var realtypes []string
	// 遍历结构体所有成员
	for i := 0; i < typeOfCat.NumField(); i++ {

		// 获取每个成员的结构体字段类型
		fd := typeOfCat.Field(i)
		names = append(names,fd.Name)
		tags = append(tags, fd.Tag.Get("json"))
		types = append(types,getType(fd.Type.String()))

		realtypes = append(realtypes,fd.Type.String())
	}
	makeFunc(typeOfCat.Name(),names,tags,types,realtypes)

}

func getType(stype string)string{
	if strings.Contains(stype,"int"){
		return "%d"
	}else if stype=="string"{
		return "'%s'"
	}
	return "%d"
}
