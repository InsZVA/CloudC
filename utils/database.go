package utils
import (
	"../parser"
	"fmt"
	"database/sql"
	_"github.com/ziutek/mymysql/godrv"
)

const (
	DB_USER = "root"
	DB_PASS = "thefirstgeek"
	MERGE_NUM = 10		//每10条insert事务合并
)

type DatabaseTemplate struct{
	database string
	table string
	bind []string
}

func insert(data [][]string,template *DatabaseTemplate) ([]string){
	ret := make([]string,len(data))
	j := 0
	for _,dataItem := range data{
		i := 0
		sql := "insert into " + template.table + "("
		for ;i < len(template.bind);i++{
			if i != 0{sql += ","}
			sql += "`" + template.bind[i] + "`"
		}
		sql += ") values ("
		for i = 1;i < len(dataItem);i++{
			if i != 1{sql += ","}
			sql += "'" + dataItem[i] + "'"
		}
		ret[j] = sql + ")"
		j++
	}
	return ret
}

func loadTemplate(filePath string) (*DatabaseTemplate){
	templateJSON,err := parser.ReadAndParseJSON(filePath)
	if(err != nil){panic(err)}
	template := new(DatabaseTemplate)
	template.database = templateJSON["database"].(string)
	template.table = templateJSON["table"].(string)
	binds := templateJSON["bind"].([]interface{})
	template.bind = make([]string,len(binds))
	for i := 0;i < len(binds);i++{
		template.bind[i] = binds[i].(string)
	}
	return template
}

func DBWork(tempName string,data [][]string){
	template := loadTemplate("./test/" + tempName + ".json")
	sqls := insert(data,template)
	db,err := sql.Open("mymysql",fmt.Sprintf("%s/%s/%s",template.database,DB_USER,DB_PASS))
	if(err != nil){panic(err)}
	defer db.Close()
	i := 0 //事务合并
	query := ""
	for _,sql := range sqls{
		if i == 0 {
			query = "START TRANSACTION;\n"
		}else if i > MERGE_NUM {
			i = 0
			query += "COMMIT;"
			_,err := db.Query(query)
			query = ""
			if(err != nil){}
			continue
		}else {
			query += sql + ";\n"
		}
		i++
	}
	if i != 0{
		query += "COMMIT;"
		_,err := db.Query(query)
		if(err != nil){}
	}
}

