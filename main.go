package main

import (
	"github.com/plumos/makesrc/model"
	"github.com/plumos/makesrc/pc"
)
/*
	1 dao (增 删(condition) 改(setinfo,condition)查(id) 、 获取列表(condition))
2 biz（增（st）删（id）查（id）改（api），setstate（id,state）。  getlist （by a, by b））

3 url（对应biz）

*/
func main(){
	pc.Parse(model.Tenant{})
}


