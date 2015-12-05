# CloudC By Go

一个Golang写的多协程网络爬虫

## 结构

1. 任务队列
2. 分发器
3. 工作协程
4. 数据库模板

### 任务队列

>其中一条

`{
	"type":"get",		
	"cookie":{
		"need":"no"
		},
	"range":"url",		
	"url":"http://ke.qq.com/course/list",
	"params":[
		{
			"name":"mt",
			"type":"int",	
			"range":[1001,1009]	
		},
		{
			"name":"st",
			"type":"int",
			"range":[2004,2008]
		},
		{
			"name":"tt",
			"type":"int",
			"range":[3026,3030]
		}],
	"pattern":"<li class=\"course-card-item\"><a .*? href=\"(.*?)\".*?>.*?<img src=\"(.*?)\".*?title=\"(.*?)\"",
	"databaseTemplate":"dt/tecentKt.dt"
}`

### 数据库模板

>dt/tencentKt.dt

`{
	"database":"CloudKt",
	"table":"videos",
	"bind":["href","img","title"]
}`