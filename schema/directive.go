package schema

type Directive struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Locations   []string   `json:"locations"`
	Args        []*Argument `json:"args"`

	//аргументы
	ArgMap       map[string] *Argument
}

func (d *Directive) Init(){
	d.ArgMap = make(map[string] *Argument)
	for _,arg:= range d.Args {
		d.ArgMap[arg.Name] = arg
	}
}


func (d *Directive) AddArg(arg *Argument){
	d.ArgMap[arg.Name] = arg
	_ = append(d.Args, arg)
}