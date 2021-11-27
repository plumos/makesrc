package pc

import (
	"fmt"
	"github.com/plumos/makesrc/cfile"
	"strings"
)

var importst =  `import (
 "fmt"
 "github.com/jmoiron/sqlx"
 "github.com/plumos/dinghuo/model"
)`
const pre = "dh_"
const common_create = `_,err :=d.db.Exec(sql)
return err
}`
func gettable(name string) string{
	return fmt.Sprintf("%s_table",strings.ToLower(name))
}
func makeDaoFunc(name string, Names []string, Tags []string ,Types []string,RealTypes []string){
	all:=""
	dao:= name+"Dao"
	all+="package dao"
	all += "\n"+importst
	tb:= "const "+gettable(name) +`="`+ pre+strings.ToLower(name)+`"`
	all+= "\n"+tb
	st:= "type "+dao +" struct{"+"\n" + " db *sqlx.DB"+"\n" +"}"
	all+= "\n"+st

	all+= "\n" + MakeGSql(name,Tags)
	all+= "\n" + makeCreate(dao,name,Names,Tags,Types)

	all += "\n" + makeDelete(dao,name,Tags,Types,RealTypes)
	all += "\n" + makeUpdate(dao,name)
	all+="\n"+makeGet(dao,name,Names,Tags,Types,RealTypes)
	all+="\n"+makeGetList(dao,name)
	cfile.WriteToFile(all,"output\\dao"+"\\"+strings.ToLower(name)+"_db.go")
}
func MakeGSql(name string,tags []string)string{
	var tag = ""
	for _,item:= range tags{
		tag += item+","
	}
	tag = tag[:len(tag)-1]

	line1 := "const "+ name+"GetSql= `select "+tag+" from %s where %s`"
	return line1
}
func makeCreate(dao,name string, Names []string, Tags []string ,Types []string)string{
	cstr := makeCsql(name,Tags,Types)
	line1:= `func (d *` +dao+`)create(m *model.`+ name + ")error{"
	line2pre:= `sql := fmt.Sprintf(`+name+"CreateSql,"+gettable(name)+","
	for _,item:=range Names{
		line2pre += "m."+item+","
	}
	line2 := line2pre[:len(line2pre)-1]
	line2+=")"
	content:= cstr+"\n"+line1+"\n" + line2 + "\n" +common_create
	return content
}
func makeCsql(name string,tags []string,Types []string)string{
	var tag = ""
	for _,item:= range tags{
		tag += item+","
	}
	tag = tag[:len(tag)-1]
	var stype = ""
	for _,item:= range Types{
		stype += item+","
	}
	stype = stype[:len(stype)-1]
	line1 := "const "+ name+"CreateSql= `insert into %s ("+tag+") values` +"
	line2:= `"(` + stype + `)"`
	return line1 + "\n" + line2
}

func makeDelete(dao,name string , Tags []string,Types []string,RealTypes []string)string{
	key,item:=0,""
	for key,item=range Tags{
		if item =="wx_id" || item =="id"{
			break
		}
	}

	line1:= `func (d *` +dao+`)delete(`+ Tags[key] + " " + RealTypes[key] + ")error{"
	line2:= `sql := fmt.Sprintf(CommonDelete,`+gettable(name)+`, fmt.Sprintf(" ` +Tags[key] + "= "+Types[key]+`",` + Tags[key] + "))"
	line3:=" _, err := d.db.Exec(sql) \n  return err \n}"
	return line1 + "\n"+line2 + "\n"+ line3

}

func makeUpdate(dao,name string )string{
	line1:= `func (d *` +dao+`) update(setinfo, condition string) error {`
	line2:= `sql := fmt.Sprintf(CommonUpdate,`+gettable(name)+`, setinfo, condition) `
	line3:=" _, err := d.db.Exec(sql) \n  return err \n}"
	return line1 + "\n"+line2 + "\n"+ line3
}

func makeGet(dao,name string,Names []string, Tags []string,Types []string, RealTypes []string )string{
	key,item:=0,""
	for key,item=range Tags{
		if item =="wx_id" || item =="id"{
			break
		}
	}

	line1:=  `func (d *` +dao+`)get(`+ Tags[key] + " " + RealTypes[key] + ")(*model."+name+",error){"
	linec:= `condition := fmt.Sprintf(" `+Tags[key]+` = `+ Types[key] +`,", ` +Tags[key]+")"

	line2:= `sql := fmt.Sprintf(`+name+"GetSql,"+gettable(name)+",condition)"
	line3:= `var ` + strings.ToLower(name) + " model."+name
	line4:= "err := d.db.Get(&"+strings.ToLower(name)+",sql)"
	line5:=" if err!=nil{ \n  return nil,err \n}"
	line6:= " return &"+strings.ToLower(name) +", nil \n}"
	return line1 + "\n"+linec+"\n"+line2 + "\n"+ line3+"\n"+line4 +"\n"+line5+"\n"+line6
}

func makeGetList(dao,name string )string{

	line1:=  `func (d *` +dao+`) getList(condition string) (*[]model.`+name+",error){"
	line2:= `sql := fmt.Sprintf(`+name+"GetSql,"+gettable(name)+",condition)"
	line3:= `var ` + strings.ToLower(name) + " []model."+name
	line4:= "err := d.db.Get(&"+strings.ToLower(name)+",sql)"
	line5:=" if err!=nil{ \n  return nil,err \n}"
	line6:= " return &"+strings.ToLower(name) +", nil \n}"
	return line1 + "\n"+"\n"+line2 + "\n"+ line3+"\n"+line4 +"\n"+line5+"\n"+line6
}