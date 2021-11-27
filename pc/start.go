package pc

func makeFunc(name string, Names []string, Tags []string ,Types []string,RealTypes []string){
	makeDaoFunc(name , Names , Tags  ,Types ,RealTypes)
	makeDaoFuncR(name , Names , Tags  ,Types ,RealTypes)
	makeBizFunc(name , Names , Tags  ,Types)
}
