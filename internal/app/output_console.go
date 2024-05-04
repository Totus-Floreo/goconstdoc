/*
Copyright © 2024 Timur Kulakov totusfloreodev@proton.me
*/
package app

import (
	"fmt"
	"github.com/Totus-Floreo/goconstdoc/internal/domain"
	"github.com/Totus-Floreo/goconstdoc/pkg/tmpl"
	"github.com/Totus-Floreo/goconstdoc/pkg/util"
	"strings"
)

func WriteToConsole(table *domain.Table, pretty bool) error {
	var buf strings.Builder
	if err := tmpl.BaseTemplate.ExecuteTemplate(&buf, tmpl.TableTemplate, table); err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	if !pretty {
		fmt.Println(util.TrimSpaceCharacters(buf.String()))
		return nil
	}

	fmt.Println(buf.String())
	return nil
}
