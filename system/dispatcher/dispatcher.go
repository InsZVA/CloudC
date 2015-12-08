package dispatcher
import (
	"../../system"
	"../../parser"
)

var TaskQueue []*system.Task
var MinNum int

func StartWorker(taskQueue []*system.Task){
	go system.Worker(taskQueue)
}

func Run(){
	LoadTaskQueue()
	taskNum := len(TaskQueue)
	workersNum := system.Config.Workers
	taskQueue := make([][]*system.Task,workersNum)
	i := 0;w := 0
	for i < taskNum{
		taskQueue[w % workersNum] = append(taskQueue[w % workersNum],TaskQueue[i])
		i++;w++
	}
	MinNum = workersNum
	if taskNum < MinNum{MinNum = taskNum}
	system.WorkerEnd = make(chan bool,10)
	for w = 0;w < MinNum;w++{
		StartWorker(taskQueue[w])
	}
}

func LoadTaskQueue(){
	//Initialize task queue
	tasks,err := parser.ReadAndParseJSONArray("./test/taskQueue.json")
	if(err != nil){
			panic(err)
	}
	for _,taskNode := range tasks{
		switch taskNode["type"] {
			case "get":
				cookie := taskNode["cookie"].(map[string]interface{})
				needCookie := cookie["need"].(string)
				if(needCookie == "no"){
					
					produceFunction := "SimpleGet"
					pattern := taskNode["pattern"].(string)
					params := taskNode["params"].([]interface{})
					url := taskNode["url"].(string)
					databaseTemplate := taskNode["databaseTemplate"].(string)
					newTask := new(system.Task)
					newTask.ProduceFunction = produceFunction
					newTask.Pattern = pattern
					newTask.Url = url
					for _,param := range params{
						newTask.Params = append(newTask.Params,param)
					}
					newTask.DatabaseTemplate=databaseTemplate
					TaskQueue = append(TaskQueue,newTask)
				}
			case "phantomGet":
				produceFunction := "PhantomGet"
				pattern := taskNode["pattern"].(string)
				params := taskNode["params"].([]interface{})
				url := taskNode["url"].(string)
				databaseTemplate := taskNode["databaseTemplate"].(string)
				newTask := new(system.Task)
				newTask.ProduceFunction = produceFunction
				newTask.Pattern = pattern
				newTask.Url = url
				for _,param := range params{
					newTask.Params = append(newTask.Params,param)
				}
				newTask.DatabaseTemplate=databaseTemplate
				TaskQueue = append(TaskQueue,newTask)
		}
	}

}