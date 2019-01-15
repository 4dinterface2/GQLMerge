package schema

type Field struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	IsDeprecated bool             `json:isDeprecated`
	DeprecationReason interface{} `json:deprecationReason` //непонятнокакой тип
	Args        []*Argument       `json:args`
	Type OfType                   `json:"type"`
	ArgumenetsMap map[string] *Argument
}

func (o *Field) Init(){
	o.ArgumenetsMap = make(map[string] *Argument)
	for _,arg:=range o.Args{            //input fields это аргументы
		o.ArgumenetsMap[arg.Name] = arg
		//field.Init()
	}
}

func (o *Field) Clone() *Field{
	result:= *o
	return &result
}