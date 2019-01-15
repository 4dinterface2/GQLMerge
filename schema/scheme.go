package schema

type Schema struct {
	Directives []*Directive `json:"directives"`
	MutationType struct {
		Name string `json:"name"`
	} `json:"mutationType"`
	SubscriptionType struct {
		Name string `json:"name"`
	} `json:"subscriptionType"`
	QueryType struct {
		Name string `json:"name"`
	} `json:"queryType"`
	Types [] *Object `json:"types"`

	TypesMap map[string] *Object; //карта типов для быстрого доступа
	DirectiveMap map[string] *Directive
}



func NewSchema() *Schema{
	return &Schema{
		Types: make([]*Object,0),
		Directives:make([]*Directive,0),
	}
}

func (s *Schema) Init(){
	s.TypesMap = make(map[string] *Object)

	for index:=range s.Types {
		obj:=s.Types[index]
		//fmt.Println(obj.Name)
		s.TypesMap[obj.Name] = obj
		obj.Init()
	}

	s.DirectiveMap = make(map[string] *Directive)
	for _, dir:=range s.Directives {
		//fmt.Println(dir)
		s.DirectiveMap[dir.Name] = dir
		dir.Init()
	}
}

func (s *Schema) Clone() *Schema{
	result:= *s
	result.Init()
	return &result
}

func (s *Schema) GetOperation(opName string) *Object{
	if(opName == "query"){
		return s.TypesMap[s.QueryType.Name]
	}
	return nil
}
