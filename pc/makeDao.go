package pc

import (
	"github.com/plumos/makesrc/cfile"
	"strings"
)

var importstd =  `import (
 "fmt"
 "github.com/plumos/dinghuo/model"
)`

func makeDaoFuncR(name string, Names []string, Tags []string ,Types []string,RealTypes []string){
	all:=""
	//dao:= name+"Dao"
	all+="package dao"
	all += "\n"+importstd
	//tb:= "const "+gettable(name) +`="`+ pre+strings.ToLower(name)+`"`
	//all+= "\n"+tb
	//st:= "type "+dao +" struct{"+"\n" + " db *sqlx.DB"+"\n" +"}"
	//all+= "\n"+st

	all+= "\n" + makeDCreate(name)

	_,_,rtype:=getKey(Tags,Types,RealTypes)
	all += "\n" + makeDDelete(name,rtype)
	all += "\n" + makeDUpdate(name)
	all+="\n"+makeDGet(name,rtype)
	all+="\n"+makeDGetList(name)
	cfile.WriteToFile(all,"output\\dao"+"\\"+strings.ToLower(name)+".go")
}

func makeDCreate(name string)string{
	line:= `func (d *` + getDao(name) +")Create(m *"+getValue(name)+")(error){"
	line+= "\nreturn d.create(m) \n}"
	return line
}

func makeDDelete(name string ,realtype string)string{
	line:= `func (d *` + getDao(name) +")Delete(key "+ realtype+")(error){"
	line+= "\nreturn d.delete(key) \n}"
	return line
}

func makeDUpdate(name string )string{
	line:= `func (d *` + getDao(name) +")Update(setinfo,condition string)(error){"
	line+= "\nreturn d.update(setinfo,condition) \n}"
	return line
}

func makeDGet(name string,realtype string )string{
	line:= `func (d *` + getDao(name) +")Get(key "+ realtype+")(*"+getValue(name)+",error){"
	line+= "\nreturn d.get(key) \n}"
	return line
}

func makeDGetList(name string )string{
	line:= `func (d *` + getDao(name) +")GetList(condition string)("+getValues(name)+",error){"
	line+= "\nreturn d.getList(condition) \n}"
	return line
}

func getValue(name string)string{
	return "model."+name
}

func getValues(name string)string{
	return "*[]model."+name
}

func getDao(name string)string{
	return name +"Dao"
}

func getlname(name string) string{
	return strings.ToLower(name)
}

func getKey(Tags, Types , RealTypes []string )(string,string,string){
	key,item:=0,""
	for key,item=range Tags{
		if item =="wx_id" || item =="id"{
			break
		}
	}
	return Tags[key],Types[key],RealTypes[key]
}