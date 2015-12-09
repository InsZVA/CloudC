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
var Commit bool	//最后一项事务是否处理完毕

func IO(){
	//Initialize
	Query = make(chan InsertTemplate,MAX_CHAN)
	for{
		query := <- Query
		Commit = false
		utils.DBWork(query.Template,query.Data)
		Commit = true
	}
}