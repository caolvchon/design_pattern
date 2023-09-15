package processInitial

import (
	"github.com/caolvchon/design_pattern/abstract_factory/processor/interactive"
	"github.com/caolvchon/design_pattern/abstract_factory/processor/movie"
)

func InitProcess(){
	interactive.Register()
	movie.Register()
}
