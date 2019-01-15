package service

import (
	"awesomeProject/service/drivers"
	"awesomeProject/schema"
	"github.com/graphql-go/graphql/language/ast"
	"fmt"
	"github.com/graphql-go/graphql/language/printer"

	"regexp"
	"encoding/json"
)

type Service struct {
	driver *drivers.HTTPService
	Schema *schema.Schema
}

func (s *Service) init() {
	req,_:=s.driver.FetchByte(&introspection)
	s.Schema = schema.FromJson(&req)
}

func (s *Service) fetchQuery (){

}

type QueryResult struct {
	data interface{}
	errors  interface{}
}

//Исполняет запрос
func (s *Service) Do(doc *ast.Document) *interface{} {
	query := printer.Print( doc )
	re:= regexp.MustCompile(`[[:space:]]`)
	str45 := re.ReplaceAllString(fmt.Sprint(query), " ")

	//IntrospectionQuery - имя операции для интроспекции
	str := `{
		"operationName": "",
		"variables": {},
		"query": "` + str45 + `"` +
	`}`

	var ff []byte = []byte(str)
	data, _ := s.driver.FetchByte( &ff )

	var v interface{}
	json.Unmarshal(data, &v)
	//fmt.Println(str, v)

	//fmt.Println("отправляю=", str ,  string(data) )
	//req,_:=s.driver.FetchByte(&introspection)
	//s.Schema = schema.FromJson(&req)
	return  &v //string(data)
}