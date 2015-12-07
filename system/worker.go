package system

import (
	"../utils"
)

var WorkerEnd chan bool

type Task struct{
	ProduceFunction string
	Pattern string
	Params []interface{}
	DatabaseTemplate string
	Url string
}

func Worker(taskQueue []*Task){
	for _,task := range taskQueue{
		switch task.ProduceFunction{
			case "SimpleGet":
			response,err := utils.SimpleGet(task.Url,task.Params)
			if(err != nil){panic(nil)}
			matches,err := utils.Peek(response,task.Pattern)
			if(err != nil){panic(nil)}
			//utils.DBWork("tecentKt",matches)
			Query <- InsertTemplate{matches,"tecentKt"}
		}
	}
	WorkerEnd <- true
}