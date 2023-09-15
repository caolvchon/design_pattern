package interactive

import "github.com/caolvchon/design_pattern/abstract_factory/processor/factory"

const name="interactive"

type Processorfactory struct{

}

func Register(){
	factory.Register(name,Processorfactory{})
}

func (f Processorfactory) Create() factory.ProcessorInf{
	return &InteractiveProcessor{}
}

type InteractiveProcessor struct {

}

func (processor InteractiveProcessor)Process() string {
	return "InteractiveProcessor，我从数据库中读 互动剧 数据，并进行format"
}





