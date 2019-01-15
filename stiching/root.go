package stiching

import (
	"awesomeProject/service"
	"github.com/graphql-go/graphql/language/ast"
	"awesomeProject/filter"
	"awesomeProject/schema"

	"encoding/json"
	"fmt"
)

type Service struct {
	services  []*service.Service
	rootShema *schema.Schema
}

func NewRootService() *Service{
	return &Service{
		rootShema:schema.NewSchema(),
	}
}

//Добавить сервис в список
func (s *Service) AddService(service *service.Service){
	Merge(s.rootShema, service.Schema);
	s.services=append(s.services, service)
}

//Добавить сервис в список
func (s *Service) Do(ast *ast.Document) string {
	res:= ResultJson{}
	for _, service := range s.services {
		filtredAST:= filter.Filter(service.Schema, ast)
		r:=service.Do(filtredAST)
		res.AddMap(r)
	}

	r, _:=json.Marshal(&res.result)
	fmt.Println("res = ",  string(r) )
	return "res"
}