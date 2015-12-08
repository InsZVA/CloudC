# CLoudC v0.03

## 系统结构

CloudC 0.03由 持久化IO总线，任务分发器，工作者协程 组成，由分发器分发任务给工作者协程，工作者协程抓取信息并多路复用持久化协程来达到高性能的抓取效果

## 配置文件

taskQueue.json

任务队列配置文件，为一个JSONArray，描述了所有要分发的任务，片段如下：
`[{
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
			"range":[{
				"start":1001,
				"end":1001}]	
		},
		{
			"name":"st",
			"type":"int",
			"range":[{
				"start":2001,
				"end":2004}]
		},
		{
			"name":"tt",
			"type":"int",
			"range":[{
				"start":3026,
				"end":3030}]
		}],
	"pattern":"<li class=\"course-card-item\"><a .*? href=\"(.*?)\".*?>.*?<img src=\"(.*?)\".*?title=\"(.*?)\"",
	"databaseTemplate":"tecentKt"
},`
对于每个节点，type目前只支持get，cookie只支持如例参数，range只支持url，url为目标url，params数组描述了多个参数，name为参数名，type目前只支持int（实际上是字符串，但范围按照数字比较），range数组描述了多个范围段，每个段由start和end描述了开始和结束数字，参数取值范围在此范围
pattern描述了正则匹配规则，抓取的内容会通过databaseTemplate所指文件的规定持久化到数据库

tecentKt.json:

数据库持久化模板，JSONObject，模板绑定字段必须与正则抓取数量一致
`{
	"database":"CloudKt",
	"table":"videos",
	"bind":["title","href","img"]
}`
database描述了数据库名，table描述了表名，bind描述了绑定的字段（顺序与正则抓取顺序一致）

config.json:

`{
	"workers":6,
	"servers":{
		"num":1,
		"serverPool":[
			"127.0.0.1"
		]
	}
}`
workers描述了工作者协程的最大值（但不会超过任务队列中任务数量），servers字段内容暂时不允许更改

## 优化策略

1. 使workers等于CPU数量（多核核心数）-1，部分时候-1得不到高效率时可以-2
2. 分割任务队列，使其为workers的整数倍，并尽量使每个任务工作量均匀


## 性能

在多核CPU及小带宽下性能优于等效php7代码10~50倍，本产品不适用于核心较少的服务器或单机使用