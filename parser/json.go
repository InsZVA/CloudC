package parser
import (
	"encoding/json"
	"io/ioutil"
	"os"
)

/*
	Function: readAndParseJSON
	Read a json object from a file
	params:
	filePath, the path of file to parse
*/
func ReadAndParseJSON(filePath string) (map[string]interface{},error){
	fi,err := os.Open(filePath)
	js := make(map[string]interface{})
	if err != nil{
		return js,err
	}
	defer fi.Close()
	fd,err := ioutil.ReadAll(fi)
	err = json.Unmarshal(fd,&js)
	if err != nil{
		return js,err
	}
	return js,nil
}

func ReadAndParseJSONArray(filePath string) ([]map[string]interface{},error){
	fi,err := os.Open(filePath)
	var js []map[string]interface{}
	if err != nil{
		return js,err
	}
	defer fi.Close()
	fd,err := ioutil.ReadAll(fi)
	err = json.Unmarshal(fd,&js)
	if err != nil{
		return js,err
	}
	return js,nil
}