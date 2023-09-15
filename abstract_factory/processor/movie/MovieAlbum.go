package movie

import "github.com/caolvchon/design_pattern/abstract_factory/processor/factory"

const name="movie"

type processorFactory struct{

}


func Register(){
	factory.Register(name,processorFactory{})
}

func (f processorFactory) Create() factory.ProcessorInf{
	return &MovieProcessor{}
}

type MovieProcessor struct {

}

func (processor MovieProcessor)Process() string {
	return "我是MovieProcessor，我从数据库中读movie数据，并进行format"
}