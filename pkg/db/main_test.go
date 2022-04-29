package db

import (
	"fmt"
	"strconv"
	"testing"
)

/**
普通json
*/
//type Actress struct {
//	Name       string
//	Birthday   string
//	BirthPlace string
//	Opus       []string
//}

/*
内嵌普通json
*/

//type Actress struct {
//	Name       string
//	Birthday   string
//	BirthPlace string
//	Opus       Opus
//}

/**
json内嵌数组json
*/

//type Actress struct {
//	Name       string
//	Birthday   string
//	BirthPlace string
//	Opus       []Opus
//}

/*
内嵌动态key
*/
type Actress struct {
	Name       string
	Birthday   string
	BirthPlace string
	Opus       map[string]Opus
}

type Opus struct {
	Date  string
	Title string
}

func TestDoInsert(t *testing.T) {
	//第一种
	//jsonData := []byte(`{
	//	"name":"汪峰",
	//	"birthday":"1998-04-16",
	//	"birthPlace":"湖南省长沙市",
	//	"opus":[
	//		"存在",
	//		"春天里",
	//		"活着"
	//	]
	//}`)

	//第二种
	//jsonData := []byte(`{
	//	"name":"汪峰",
	//	"birthday":"1998-04-16",
	//	"birthPlace":"湖南省长沙市",
	//	"opus":{
	//		"Date":"2022",
	//		"Title":"存在"
	//	}
	//}`)

	//第三种
	//jsonData := []byte(`{
	//	"name":"汪峰",
	//	"birthday":"1998-04-16",
	//	"birthPlace":"湖南省长沙市",
	//	"opus":[{
	//		"Date":"2022",
	//		"Title":"存在"
	//	},
	//	{
	//		"Date":"2222",
	//		"Title":"不存在"
	//	},
	//	{
	//		"Date":"2399",
	//		"Title":"无谓存在"
	//	}]
	//}`)

	//第四种
	//jsonData := []byte(`{
	//	"name":"汪峰",
	//	"birthday":"1998-04-16",
	//	"birthPlace":"湖南省长沙市",
	//	"opus":{
	//		"2013":{
	//		"Date":"2022",
	//		"Title":"存在"
	//	},
	//		"2014":{
	//		"Date":"2222",
	//		"Title":"不存在"
	//	},
	//		"2015":{
	//		"Date":"2399",
	//		"Title":"无谓存在"
	//	}
	//	}
	//}`)
	//var actress Actress
	//err := json.Unmarshal(jsonData, &actress)
	//if err != nil {
	//	fmt.Println("error", err)
	//	return
	//}
	//fmt.Printf(actress.Name)
	//fmt.Printf(actress.Birthday)
	//fmt.Printf(actress.BirthPlace)
	//fmt.Printf("作品")
	//for index, val := range actress.Opus {
	//	fmt.Printf("\t标签:%s\n", index)
	//	fmt.Printf("\t\t日期：%s\n", val.Date)
	//	fmt.Printf("\t\t标题：%s\n", val.Title)
	//}
	itemIDStr := "1233"
	itemID, err := strconv.ParseInt(itemIDStr, 10, 64)
	if err == nil {
		fmt.Println("值为：", itemID)
	}
}
