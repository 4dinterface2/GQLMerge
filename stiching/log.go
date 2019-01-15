package stiching

import "fmt"

type LogStatus int
const (
	Conflicted    LogStatus = 0
	Added         LogStatus = 1
	Replace       LogStatus = 2
	Merged        LogStatus = 3
)

type LogItem struct {
	Status LogStatus
	FieldName string
	val1 interface{}
	val2 interface{}
}

type Log struct{
	items []*LogItem
	conflictCouner int
}

func NewLog () *Log {
	return &Log{
		items: make([]*LogItem, 0),
		conflictCouner:0,
	}
}

func (l *Log) Conflict(fieldName string, va1 *interface{}, val2 *interface{}) {
	l.items = append(l.items, &LogItem{
		Status: Conflicted,
		FieldName:fieldName,
		val1:va1,
		val2:val2,
	})
	fmt.Println("conflict", fieldName, va1, val2)
}

func (l *Log) Merge(fieldName string, va1, val2 *interface{} ) {
	l.items = append(l.items, &LogItem{
		Status: Merged,
		FieldName:fieldName,
		val1:va1,
		val2:val2,
	})
	fmt.Println("merge ", fieldName, va1, val2)
}