/*
Copyright © 2024 Timur Kulakov totusfloreodev@proton.me
*/
package app

import (
	"fmt"
	"github.com/Totus-Floreo/goconstdoc/internal/domain"
	"github.com/Totus-Floreo/goconstdoc/pkg/tmpl"
	"os"
)

func WriteToConsole(table *domain.Table) error {
	if err := tmpl.BaseTemplate.ExecuteTemplate(os.Stdout, tmpl.TableTemplate, table); err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	return nil
}
