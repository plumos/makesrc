package pc

import (
	"github.com/plumos/makesrc/cfile"
	"strings"
)
var importbz =  `import (
 "fmt"
 "github.com/plumos/dinghuo/dao"
 "github.com/plumos/dinghuo/model"
)`
func makeBizFunc(name string ,Names []string,Tags []string, Types []string){
	all:=""
	biz:= name+"Biz"

	all+="package biz"
	all += "\n"+importbz
	st:= "type "+biz +" struct{}"
	all+= "\n"+st
	all+= "\n"+"var My" + biz + " "+ biz
	all+="\n"+makebizAdd(biz,name)
	all+="\n"+makebizSet(biz,name,Names,Tags,Types)

	cfile.WriteToFile(all,"output\\biz"+"\\"+strings.ToLower(name)+".go")
}
func makebizAdd(biz,name string )string{
	lname:= strings.ToLower(name)
	content:= `func (d *` +biz+`)Create(`+lname+` *model.`+ name + ")error{"
	content+= `return dao.Def`+name+"Dao.Create("+lname+")"
	content+="}"
	return content
}
func makebizSet(biz,name string ,Names []string,Tags []string, Types []string)string{
	lname:= strings.ToLower(name)
	content:=""
	content+= `func (d *` +biz+`)Update(`+lname+` *model.`+ name + "Api)error{"

	content+="\n" + `setinfo:=""`
	for key,item:=range Names{
		if key ==0{
			continue
		}
		if item =="ctime" ||item == "mtime"||item=="creator" || item=="modifier"{
			continue
		}
		content += "\n"+ "if "+lname+"."+item +" != nil{ "
		content += "\n"+`setinfo += fmt.Sprintf(" `+Tags[key]+` = `+ Types[key] +`,", ` +getname(lname,item,Types[key])+")"
		content += "\n"+"}"
	}

	content+= "\n"+ `condition := fmt.Sprintf(" `+Tags[0]+` = `+ Types[0] +`,", ` +getid(lname,Names[0],Types[0])+")"
	content+= "\n"+`return dao.Def`+name+"Dao.Update(setinfo,condition)"
	content+="\n" + "}"
	return content
}

func getid(stname ,name,tag string )string{
	if tag == "%d"{
		return stname+"."+name
	}else{
		return "escape(" + stname+"."+name+")"
	}
}

func getname(stname ,name,tag string )string{
	if tag == "%d"{
		return "*"+stname+"."+name
	}else{
		return "escape(*" + stname+"."+name+")"
	}
}