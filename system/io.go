package system
/* 多路IO复用协程 */
import (
	"../utils"
)

const (
	MAX_CHAN = 10000
)

type InsertTemplate struct{
	Data [][]string
	Template string
}

var Query chan InsertTemplate

func IO(){
	//Initialize
	Query = make(chan InsertTemplate,MAX_CHAN)
	for{
		query := <- Query
		utils.DBWork(query.Template,query.Data)
	}
}