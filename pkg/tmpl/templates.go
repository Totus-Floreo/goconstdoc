/*
Copyright © 2024 Timur Kulakov totusfloreodev@proton.me
*/
package tmpl

import "html/template"

const (
	HTMLTemplate  = "html"
	TableTemplate = "table"
)

var BaseTemplate = template.Must(template.New("base").Parse(`
		{{define "html"}}
		<!DOCTYPE html>
		<html lang="ru">
		<meta charset="UTF-8">
		{{template "table" .}}
		</html>
		{{end}}
		
		{{define "table"}}
		<table>
		<thead>
		<tr>
			{{range .Columns}}
				<th>{{.Name}}</th>
			{{end}}
		</tr>
		</thead>
		<tbody>
			{{$Columns := .Columns}}
			{{range $Row := .Rows}}
				<tr>
					{{range $Column := $Columns}}
						<td>{{index $Row $Column.Value}}</td>
					{{end}}
				</tr>
			{{end}}
		</tbody>
		</table>
		{{end}}
	`))
