package stiching

import (
	"awesomeProject/schema"
)

func Merge(dst *schema.Schema, src *schema.Schema){
	resultType :=  dst.Clone()
	MergeTypes(resultType, src)
	MergeDirective(resultType, src)
}

func MergeTypes(destination *schema.Schema, source *schema.Schema)  {
	log:= NewLog()
	for _, src := range source.Types  {
		//fmt.Println(src.Name)
		dst, ok := destination.TypesMap[src.Name]
		if !ok {
			dst = schema.NewObject()
			dst.Name = src.Name
			dst.Kind = src.Kind
			//log.Conflict(src.Name, src, dst)
		} else {
			dst = dst.Clone()
			//fmt.Println("dst",dst)
			//log.Merge(dst.Name, src, dst)
		}

		//dst.Description = dst.Description + src.Description  //Тут обьект а должно быть описание
		strictMergeString("kind", &dst.Kind, &src.Kind, log )

		MergeFields(dst, src.FieldsMap)
		MergeInputFields(dst.InputFieldMap, src.InputFieldMap)
		// TODO src.Interfaces  - смерджить интерфейсы
		// TODO src.PossibleTypes  - смерджить PossibleTypes

		//сохраним результат в map и массив
		destination.TypesMap[src.Name] = dst
		destination.Types = append(destination.Types, dst)
	}
}

func MergeFields(destination *schema.Object, source map[string]*schema.Field) {
	for _, src := range source  {
		if val, ok:=destination.FieldsMap[src.Name]; ok {
			destination.ReplaceField(val.Clone()) // если есть то клонируем
		} else {
			destination.AddField(src.Clone())     // если нет то добавляем новый
		}
		//MergeInputFields(src.)
	}
}

func MergeInputFields(destination map[string]*schema.Argument, source map[string]*schema.Argument) {
	//resultMap:=make(map[string]*schema.Argument)
	//resultArr:=make([]*schema.Argument,)
	/*for _, src := range source  {
		fmt.Println("directive=", src.Name)
	}*/
}

//мерджим дерективы
func MergeDirective(destination *schema.Schema, source *schema.Schema)  {
	/*for _, src := range source.Directives  {
		fmt.Println("directive=", src.Name);
	}*/
}

//=================================== utils =========================================//
func strictMergeString(fieldName string, dst *string, src *string, log *Log ){
	if dst != src {
		//log.Conflict(fieldName, dst, src)
	} else {
		dst = src
	}
}