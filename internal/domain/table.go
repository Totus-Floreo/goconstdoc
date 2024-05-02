/*
Copyright © 2024 Timur Kulakov totusfloreodev@proton.me
*/
package domain

import "github.com/Totus-Floreo/goconstdoc/internal/constant"

type Table struct {
	Name    string
	Columns []Column
	Rows    []Row
}

type Row map[string]string

func (r *Row) SetName(name string) {
	(*r)[constant.BuiltinName] = name
}

func (r *Row) SetValue(value string) {
	(*r)[constant.BuiltinValue] = value
}

func (r *Row) HasName() bool {
	_, ok := (*r)[constant.BuiltinName]
	return ok
}

func (r *Row) HasValue() bool {
	_, ok := (*r)[constant.BuiltinValue]
	return ok
}

type Column struct {
	Name  string
	Value string
}
