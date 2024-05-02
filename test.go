package main

/*
	goconstdoc:column:ColumnName:Name
	goconstdoc:column:ColumnValue:Value
	goconstdoc:column:ColumnComment:Comment
*/

const (
	TestString          = "Hello, " // goconstdoc:Comment:It Merging Mode:Value:goconstdoc!
	TestInt             = 24        // goconstdoc:Comment:command like:Value:24
	TestBool            = true      // goconstdoc:Comment:goconstdoc:Value:true
	TestFloat           = 3.14      // goconstdoc:Comment:parse:Value:3.14
	TestRune            = 'G'       // goconstdoc:Comment:-p test.go:Value:o
	TestByte            = 'B'       // goconstdoc:Comment:-o index.html:Value:66
	TestIntFloat        = 15        // goconstdoc:Comment:-i merge:Value:10.5
	TestUndoc           = "Undocumented constant"
	TestIntWithoutValue = 42              // goconstdoc:Comment:--nocmd
	TestStrWithoutValue = "No More Merge" // goconstdoc:Comment:--overwrite
)

const UndocumentedConst = "Undocumented constant"
