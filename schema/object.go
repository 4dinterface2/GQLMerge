package schema

//import "fmt"

/*type ObjectConfig struct {
	Name        string      `json:"name"`
	Interfaces  interface{} `json:"interfaces"`
	Fields      interface{} `json:"fields"`
	IsTypeOf    IsTypeOfFn  `json:"isTypeOf"`
	Description string      `json:"description"`
}*/

//важно чтобы мы не храним ссылку на массив, так как это используется при клонировании
type Object struct {
	InputFields   []*Argument    `json:"inputFields"`
	Name          string        `json:"name"`
	Description   interface{}   `json:"description"`
	Interfaces    []interface{} `json:"interfaces"`         // попробывал добавить массив
	EnumValues    interface{}   `json:"enumValues"`
	Fields      []*Field       `json:"fields"`
	//Fields        map[string]*Field       `json:"fields"`   // вроде как можно распарсить в map, непонятно как map в array Обратно

	Kind          string        `json:"kind"`
	PossibleTypes interface{}   `json:"possibleTypes"`
	//IsTypeOf    IsTypeOfFn    `json:"isTypeOf"`

	//map
	FieldsMap map[string] *Field
	InputFieldMap map[string] *Argument
}

func (o *Object) Init(){
	//field map
	o.FieldsMap = make(map[string] *Field)
	for index:=range o.Fields{
		field:=o.Fields[index];
		//fmt.Println("  ", field.Name)
		o.FieldsMap[field.Name] = field
		field.Init()
	}


	//input field map
	o.InputFieldMap = make(map[string] *Argument)
	for index:=range o.InputFields{            //input fields это аргументы
		field := o.InputFields[index];
		//fmt.Println("  input=", field.Name)
		o.InputFieldMap[field.Name] = field
		//field.Init()
	}
}

func NewObject() *Object{
	instance:=&Object{}
	instance.Init()
	return instance
}

func (o *Object) AddField(field *Field){
	o.FieldsMap[field.Name] = field
	//TODO управлять массивом
}

func (o *Object) ReplaceField(field *Field){
	o.FieldsMap[field.Name] = field
	//TODO управлять массивом
}


func (o *Object) AddArgument(arg *Argument) {
	o.InputFieldMap[arg.Name] = arg
	//TODO управлять массивом
}

func (o *Object) ReplaceArgument(arg *Argument) {
	o.InputFieldMap[arg.Name] = arg
	//TODO управлять массивом
}

func (o *Object) Clone() *Object {
	ret:=*o
	ret.Init()
	return  &ret
}