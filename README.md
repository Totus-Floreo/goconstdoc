# GoConstDoc

GoConstDoc is a command-line tool written in Go. It parses constants from a Go file and generates a documentation HTML table for them.

## Features

- Parse constants from a Go file.
- Generate a documentation HTML table for the parsed constants.
- Output the generated table to the terminal or save it to a file.
- Control overwriting of the output file.
- Interact with built-in values of the program.

## Defining Columns and Values with Comments

In GoConstDoc, you can use comments in your Go files to define the columns and their values for the generated documentation table. This is done by using a specific comment syntax.

Here's an example:

```go
/*
    constant has default value fields Name and Value
    You can override the default values by using -i, --interaction flag, 
    more in Interacting with Built-in Values section
*/

/*
	goconstdoc:column:ColumnName:Name
	goconstdoc:column:ColumnValue:Value
	goconstdoc:column:ColumnComment:Comment
*/

const (
TestString          = "Hello, " // goconstdoc:Comment:TestValue1:Value:goconstdoc!
TestInt             = 24        // goconstdoc:Comment:TestValue2:Value:24
TestBool            = true      // goconstdoc:Comment:TestValue3:Value:true
TestFloat           = 3.14      // goconstdoc:Comment:TestValue4:Value:3.14
TestRune            = 'G'       // goconstdoc:Comment:TestValue5:Value:o
TestByte            = 'B'       // goconstdoc:Comment:TestValue6:Value:66
TestIntFloat        = 15        // goconstdoc:Comment:TestValue7:Value:10.5
TestUndoc           = "Undocumented constant"
TestIntWithoutValue = 42              // goconstdoc:Comment:TestValue8
TestStrWithoutValue = "No More Merge" // goconstdoc:Comment:TestValue9
)

const UndocumentedConst = "Undocumented constant"
```

In the above example:

- `ColumnName` is the name of the column and `Name` (builtin) is the value of the column.
- `ColumnValue` is the name of the column and `Value` (builtin) is the value of the column.
- `ColumnComment` is the name of the column and `Comment` (your_comment_value) is the value of the column.

For the first constant:

- `TestString` is value of the column `ColumnName`, 
- `"Hello, "` is value of the column `ColumnValue`,
- `TestValue1` is value of the column `ColumnComment`,
- `goconstdoc!` is value of the column `ColumnValue`.

Check the [example](https://totus-floreo.github.io/goconstdoc/) for the output of the above example.

## Interacting with Built-in Values

In GoConstDoc, you can interact with built-in values of the program using the `interaction` flag. This flag accepts three values: `builtin`, `merge`, and `overwrite`.

- `builtin`: This is the default value. When this value is set, the built-in values are used.
- `merge`: When this value is set, the custom values are merged with the built-in values.
- `overwrite`: When this value is set, the custom values overwrite the built-in values.

About interacting:
- If custom values are not present, the built-in values are used.

About merging:
- If custom values are present, they are merged with the built-in values.
- if built-in values and the custom values are numeric, the custom values are added to the built-in values.
- if built-in values and the custom values are strings, the custom values are concatenated to the built-in values.
- if built-in values and the custom values are boolean, the custom values are ORed with the built-in values.
- if built-in values and the custom values are runes, the custom values are concatenated to the built-in values like strings.
- if built-in values and the custom values are bytes, the custom values are concatenated to the built-in values like strings.
- if built-in values are numeric and the custom values are strings, the custom values are concatenated to the built-in values.
- All of the above rules are applied to the custom values.

Check the [example](https://totus-floreo.github.io/goconstdoc/).

## Usage

The main command for the tool is `parse`. Here's how you can use it:

```bash
goconstdoc parse -p /path/to/your/file.go -o /path/to/output.html
```

### Flags

- `-h, --help`: Show help message.
- `-p, --path`: Path to the Go file to parse. This flag is required.
- `-o, --output`: Output file for the documentation. If not provided, the output will be written to the terminal.
- `--nocmd`: If set, the output will not be written to the terminal. Default is `false`.
- `--overwrite`: If set, the output file will be overwritten if it exists. Default is `false`.
- `-i, --interaction`: Type of interaction with built-in values. Allowed values are `builtin`, `merge`, `overwrite`. Default is `builtin`.

## Installation

To install the tool, you need to have Go installed on your machine. Then, you can clone the repository and build the tool:

```bash
go install github.com/Totus-Floreo/goconstdoc@main
```

Now, you can use the tool with `./goconstdoc` or just `goconstdoc`.

## Contributing

Contributions are welcome. Please open an issue or submit a pull request on GitHub.

## License

This project is licensed under the terms of the MIT license.
