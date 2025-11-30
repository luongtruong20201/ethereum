package util

import (
	"encoding/json"
	"reflect"
)

type List struct {
	list   reflect.Value
	Length int
}

func NewList(t interface{}) *List {
	list := reflect.ValueOf(t)
	if list.Kind() != reflect.Slice {
		panic("list container initialized with a non-slice type")
	}
	return &List{
		list:   list,
		Length: list.Len(),
	}
}

func EmptyList() *List {
	return NewList([]interface{}{})
}

func (l *List) Get(i int) interface{} {
	if l.list.Len() > i {
		return l.list.Index(i).Interface()
	}
	return nil
}

func (l *List) Append(v interface{}) {
	l.list = reflect.Append(l.list, reflect.ValueOf(v))
	l.Length = l.list.Len()
}

func (l *List) Interface() interface{} {
	return l.list.Interface()
}

func (l *List) ToJSON() string {
	var list []interface{}
	for i := 0; i < l.Length; i++ {
		list = append(list, l.Get(i))
	}
	data, _ := json.Marshal(list)
	return string(data)
}
