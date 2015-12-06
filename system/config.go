package system
import (
	"../parser"
)

type serversConfig struct{
	Num int
	ServerPool []string
}

type config struct{
	Workers int
	Servers serversConfig
}

var Config config

func LoadConfig(){
	config,err := parser.ReadAndParseJSON("./test/config.json")
	if(err != nil){panic(err)}
	workers := int(config["workers"].(float64))
	serversJSON := config["servers"].(map[string]interface{})
	num := int(serversJSON["num"].(float64))
	serverPools := serversJSON["serverPool"].([]interface{})
	var serverPool []string
	for _,serverPoolNode := range serverPools{
		serverPool = append(serverPool,serverPoolNode.(string))
	}
	Config.Workers = workers
	Config.Servers.Num = num
	Config.Servers.ServerPool = serverPool
}