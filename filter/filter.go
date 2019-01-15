  package filter
//import "github.com/graphql-go/graphql/language/visitor"
import "github.com/graphql-go/graphql/language/ast"

import (
	"awesomeProject/schema"
)

func Filter(schema *schema.Schema, node *ast.Document) *ast.Document{
	doc:=filterOperation(schema, node);
	return doc
}

func filterOperation(schema *schema.Schema, doc *ast.Document) *ast.Document{
	//fmt.Println("schema", ast)
	newDoc:=ast.NewDocument(nil)
	for _, definition:= range doc.Definitions {
		operation := definition.(*ast.OperationDefinition)
		newOperation := ast.NewOperationDefinition(nil)

		query := schema.GetOperation(operation.GetOperation()) //TypesMap["RootQuery"]

		newOperation.Loc = operation.Loc
		newOperation.SelectionSet = FilterSelectionSet(schema, query, operation.SelectionSet)
		newOperation.Name = operation.Name
		newOperation.Directives = operation.Directives
		newOperation.Kind = operation.Kind
		newOperation.Operation = operation.Operation
		newOperation.VariableDefinitions = operation.VariableDefinitions

		newDoc.Definitions = append(newDoc.Definitions, newOperation) //operation defenition
	}
	return newDoc
}


func FilterSelectionSet(schema *schema.Schema, currentType *schema.Object, selectionSet *ast.SelectionSet) *ast.SelectionSet {
	if(selectionSet==nil) {
		return nil
	}
	newSet:=ast.NewSelectionSet(nil)

	for _, selection:= range selectionSet.Selections {
		field := selection.(*ast.Field)                             // selection
		fieldDefinition := currentType.FieldsMap[field.Name.Value]; // описание поля в схеме

		if fieldDefinition == nil {
			continue;
		}
		resultType := schema.TypesMap[fieldDefinition.Type.GetName()] // TODO интроспекция возвращает цепочку типо (списки, notNull)
		newField := ast.NewField(nil)
		//newField.Arguments = FilterArgs(schema, fieldDefinition, field.Arguments)    // аргументы пока не фильтруем
		newField.Alias = field.Alias
		newField.Name = field.Name
		newField.Kind = field.Kind
		newField.Directives = field.Directives  // директивы пока не фильтруем
		newField.Loc = field.Loc
		newField.SelectionSet = FilterSelectionSet(schema, resultType, field.SelectionSet)
		newField.Arguments=filterArgs(schema, fieldDefinition, field.Arguments)
		newSet.Selections = append(newSet.Selections, newField)
	}
	return newSet
}

//фильтруем аргументы
func filterArgs(schema *schema.Schema, fieldDefinition *schema.Field, args []*ast.Argument)[]*ast.Argument {
	result:= make( []*ast.Argument,0)

	for _,arg:= range args {
		def:=fieldDefinition.ArgumenetsMap[arg.Name.Value]
		if (def!=nil && def.Name == arg.Name.Value) {
			newArg:=ast.NewArgument(nil);
			newArg.Loc = arg.Loc
			newArg.Name = arg.Name
			newArg.Kind = arg.Kind
			newArg.Value = makeValue(schema, def.Type.GetName(), arg.Value)
			result=append(result, newArg)
		}
	}
	return result
}

func makeValue(schema *schema.Schema, itemType string, value ast.Value ) ast.Value {
	argType:=schema.TypesMap[itemType]

	switch value.(type) {
	case *ast.ObjectValue:
		objVal:=value.(*ast.ObjectValue)
		newObjVal:= ast.NewObjectValue(nil)
		newObjVal.Kind   = objVal.Kind
		newObjVal.Fields = FilterObjectValueFields(schema, argType, objVal.Fields)
		newObjVal.Loc = objVal.Loc
		return objVal
		break;

	case *ast.ListValue: //TODO
		objVal:=value.(*ast.ListValue)
		newObjVal:= ast.NewListValue(nil)
		newObjVal.Kind   = objVal.Kind
		newObjVal.Values  =   objVal.Values//FilterObjectValueFields(schema, argType, objVal.Fields)
		newObjVal.Loc = objVal.Loc
		return value
		break;
	case *ast.EnumValue: //TODO (какая логика у enum ?)
		return value
		break;
	default:
		return value
	}
	return value
}


func FilterObjectValueFields(schema *schema.Schema, typeDef *schema.Object, fields []*ast.ObjectField) []*ast.ObjectField{
	newFields:=make([]*ast.ObjectField,0)
	//for _, item:= range fields {
		//newFields:=makeValue()
	//}
	return newFields;
}