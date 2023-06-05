package core

type Field struct {
	Name string
	Type string
}

type FieldList []*Field

type Type struct {
	Name string
	*FieldList
	Funcs []*Func
}

type Func struct {
	Name   string
	Return *FieldList
	Params *FieldList
}

type Module struct {
	Name  string
	Deps  []*Module
	Types []*Type
	Funcs []*Func
}
