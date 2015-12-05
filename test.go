package main
import (
	"./parser"
	"./utils"
)

func main(){
	task,err := parser.ReadAndParseJSON("./test/taskQueue.json")
	if(err != nil){
		panic(err)
	}
	cookie := task["cookie"].(map[string]interface{})
	needCookie := cookie["need"].(string)
	switch task["type"] {
		case "get":
			if(needCookie == "no"){
				response,err := utils.SimpleGet(task["url"].(string),task["params"].([]interface{}))
				if(err != nil){panic(nil)}
				matches,err := utils.Peek(response,task["pattern"].(string))
				if(err != nil){panic(nil)}
				utils.TestDB("tecentKt",matches)
			}
	}
}