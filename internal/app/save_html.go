/*
Copyright © 2024 Timur Kulakov totusfloreodev@proton.me
*/
package app

import (
	"bytes"
	"fmt"
	"github.com/Totus-Floreo/goconstdoc/internal/domain"
	"github.com/Totus-Floreo/goconstdoc/pkg/tmpl"
	"os"
)

func SaveHTMLToFile(table *domain.Table, filename string, overwrite bool) error {
	var buffer bytes.Buffer
	if err := tmpl.BaseTemplate.ExecuteTemplate(&buffer, tmpl.HTMLTemplate, table); err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	flag := os.O_CREATE | os.O_EXCL | os.O_WRONLY
	if overwrite {
		flag = os.O_CREATE | os.O_WRONLY
	}

	file, err := os.OpenFile(filename, flag, 0666)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	if _, err := file.Write(buffer.Bytes()); err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}
