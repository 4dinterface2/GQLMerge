package service

import (
	"awesomeProject/service/drivers"
	"awesomeProject/schema"
)

func NewHTTPService(url string) (*Service){
	service:=&Service{
		driver : &drivers.HTTPService{URL:url},
	}
	service.init();
	return service;
}

func NewNatsService() (*Service){
	return &Service{
		driver : &drivers.HTTPService{},
	}
}

func NewJSONService() (*Service){
	s:=&Service{
		driver : &drivers.HTTPService{},
	}
	s.Schema = schema.ReadJsonFile()
	return s
}