package main

import (
	"awesomeProject/service"
	"awesomeProject/stiching"
	"fmt"
	"net/http"
	"log"
	"github.com/graphql-go/graphql/language/source"
	"github.com/graphql-go/graphql/language/parser"
)

func Do(schema *stiching.Service , query string) string {
	source2 := source.NewSource(&source.Source{
		Body: []byte(query),
		Name: "GraphQL request",
	})

	AST, _ := parser.Parse(parser.ParseParams{Source: source2})
	return schema.Do(AST)

	//t1:=time.Now()
	//t2:=time.Now()
	//fmt.Println("execution time=", time.Now().Sub(t), "filterTime=",t2.Sub(t1))/
}

func main() {
	ser:=stiching.NewRootService()
	ser.AddService(service.NewHTTPService("https://hivdb.stanford.edu/graphql"))
	ser.AddService(service.NewHTTPService("https://graphql-demo.azurewebsites.net/?"))
	ser.AddService(service.NewJSONService())

	//t:=time.Now()
	query := `
		{
  			me {
    			id
  			}

            allComments (filter:{name:{demo:10}}){
				id 
				name
            }
            viewer {
                currentVersion { text, publishDate }
            }
		}
	`

	//===========================================================================//
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		result := Do(ser, query)
		fmt.Fprintf(w, "%q", result)

		//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	fmt.Println("server on 9000")

	err := http.ListenAndServe(":9000", nil) // задаем слушать порт
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	//===========================================================================//
}