package schema

type OfType struct {
	Kind   string      `json:"kind"`
	Name   string      `json:"name"`
	OfType *OfType      `json:"ofType"` // попробывать ofType
}

func (o *OfType) GetName() string {

	typer:=o
	for {
		if typer.Name != ""{
			//fmt.Println("result=",  typer.Name)
			return typer.Name
		}
		if typer.OfType==nil {
			break
		}
		typer = typer.OfType
	}
	return ""
}
/*package schema

type ArgType struct{
	Kind   string      `json:"kind"`
	Name   interface{} `json:"name"`
	OfType OfType      `json:"ofType"`
}*/
