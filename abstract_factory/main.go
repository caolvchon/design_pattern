package main

import (
	"fmt"
	"github.com/caolvchon/design_pattern/abstract_factory/processor/factory"
	"github.com/caolvchon/design_pattern/abstract_factory/processor/processInitial"
	"os"
	"time"
)


func initial(){
	processInitial.InitProcess()
}

func main(){
	initial()


	names:=[]string{
		"movie","interactive","comic","comic","interactive","comic",
	}
	for _,v:=range names{
		time.Sleep(time.Second)
		go func(s string) {
			defer func() {
				if err:=recover();err!=nil{
					fmt.Printf("fatal no match name panic=%v \n" ,err)
					return
				}
			}()
			processor:=factory.GetProcessor(s)
			res:=processor.Process()
			fmt.Println(res)
		}(v)
	}
	time.Sleep(time.Second*10)
	os.Exit(1)
}