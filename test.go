package main
import (
	"./parser"
	"./utils"
	"fmt"
)

func main(){
	task,err := parser.ReadAndParseJSON("./test/taskQueue.json")
	if(err != nil){
		panic(err)
	}
	switch task["type"] {
		case "get":
			fmt.Println(utils.SimpleGet(task["url"].(string),task["params"].([]interface{})))
	}
}