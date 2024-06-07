package contentful

import (
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/labd/contentful-go"
	"strings"
)

func parseError(err error) diag.Diagnostics {
	if !errors.As(err, &contentful.ErrorResponse{}) {
		return diag.FromErr(err)
	}

	var warnings []diag.Diagnostic
	for _, e := range err.(contentful.ErrorResponse).Details.Errors {
		var path []string
		if e.Path != nil {
			for _, p := range e.Path.([]interface{}) {
				path = append(path, fmt.Sprintf("%v", p))
			}
		}
		warnings = append(warnings, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  fmt.Sprintf("%s (%s)", e.Details, strings.Join(path, ".")),
		})
	}

	return append(warnings, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  err.(contentful.ErrorResponse).Message,
	})
}
