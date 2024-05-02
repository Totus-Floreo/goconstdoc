# GoConstDoc

GoConstDoc is a command-line tool written in Go. It parses constants from a Go file and generates a documentation HTML table for them.

## Features

- Parse constants from a Go file.
- Generate a documentation HTML table for the parsed constants.
- Output the generated table to the terminal or save it to a file.
- Control overwriting of the output file.

## Defining Columns and Values with Comments

In GoConstDoc, you can use comments in your Go files to define the columns and their values for the generated documentation table. This is done by using a specific comment syntax.

Here's an example:

```go
/*
    constant has default value fields Name and Value
    Do not override the default value fields
*/

/*
	goconstdoc:column:ColumnName:Name 
	goconstdoc:column:ColumnValue:Value 
	goconstdoc:column:ColumnComment:Comment 
*/

const (
    TestOne   = 100 // goconstdoc:Comment:Test1 
    TestTwo   = 200 // goconstdoc:Comment:Test2 
    TestThree = 300 // goconstdoc:Comment:Test3 
)
```

In the above example:

- `ColumnName` is the name of the column and `Name` (builtin) is the value of the column.
- `ColumnValue` is the name of the column and `Value` (builtin) is the value of the column.
- `ColumnComment` is the name of the column and `Comment` (your_comment_value) is the value of the column.

For the constants:

- `Test1` is value of the column `Comment` for the constant `TestOne`.
- `Test2` is value of the column `Comment` for the constant `TestTwo`.
- `Test3` is value of the column `Comment` for the constant `TestThree`.

## Usage

The main command for the tool is `parse`. Here's how you can use it:

```bash
goconstdoc parse -p /path/to/your/file.go -o /path/to/output.html
```

### Flags

- `-p, --path`: Path to the Go file to parse. This flag is required.
- `-o, --output`: Output file for the documentation. If not provided, the output will be written to the terminal.
- `--cmd`: If set, the output will be written to the terminal. Default is `true`.
- `--overwrite`: If set, the output file will be overwritten if it exists. Default is `false`.

## Installation

To install the tool, you need to have Go installed on your machine. Then, you can clone the repository and build the tool:

```bash
go install github.com/Totus-Floreo/goconstdoc@main
```

Now, you can use the tool with `./goconstdoc`.

## Contributing

Contributions are welcome. Please open an issue or submit a pull request on GitHub.

## License

This project is licensed under the terms of the MIT license.
