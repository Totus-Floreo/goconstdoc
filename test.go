package main

/*
	goconstdoc:column:ColumnName:Name
	goconstdoc:column:ColumnValue:Value
	goconstdoc:column:ColumnComment:Comment
*/

const (
	TestOne   = 100 // goconstdoc:Comment:TestValue1
	TestTwo   = 200 // goconstdoc:Comment:TestValue2
	TestThree = 300 // goconstdoc:Comment:TestValue3
)

const UndocumentedConst = "Undocumented constant"
