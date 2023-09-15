package factory


var processorFactories= make(map[string]ProcessorFactoryInf)

func Register(name string,factory ProcessorFactoryInf){
	processorFactories[name]=factory
}

type ProcessorFactoryInf interface {
	 Create() ProcessorInf
}


type ProcessorInf interface {
	 Process() (string)
}

func GetProcessor(name string) ProcessorInf{
	f,ok:=processorFactories[name]
	if !ok{
		panic("not match processor")
	}
	return f.Create()
}

//策略模式
//模板方法模式

