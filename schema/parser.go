package schema

import (
	"io/ioutil"
	"encoding/json"
)

//TODO допилить скаляры
type IntrospectSchema struct {
	Data struct {
		Schema Schema `json:"__schema"`
	} `json:"data"`
}

func FromJson(b *[]byte) *Schema{
	var v2 IntrospectSchema
	json.Unmarshal(*b, &v2)
	//fmt.Println(v2.Data.Schema.Types[17].Fields[2])
	v2.Data.Schema.Init()
	return &v2.Data.Schema
}

func ReadJsonFile() *Schema{
	b, err := ioutil.ReadFile("./config.json") // just pass the file name
	if err != nil {
		//fmt.Print(err)
	}
	return FromJson(&b);
}