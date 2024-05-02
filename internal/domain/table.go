/*
Copyright © 2024 Timur Kulakov totusfloreodev@proton.me
*/
package domain

type Table struct {
	Name    string
	Columns []Column
	Rows    []Row
}

type Row map[string]string

func (r *Row) SetName(name string) {
	(*r)["Name"] = name
}

func (r *Row) SetValue(value string) {
	(*r)["Value"] = value
}

type Column struct {
	Name  string
	Value string
}
