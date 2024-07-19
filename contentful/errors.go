package contentful

import (
	"errors"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/labd/contentful-go"
)

func parseError(err error) diag.Diagnostics {
	if err == nil {
		return nil
	}

	var contentfulErr contentful.ErrorResponse
	if !errors.As(err, &contentfulErr) {
		return diag.FromErr(err)
	}

	var diags diag.Diagnostics
	for _, e := range contentfulErr.Details.Errors {
		var path []string
		if e.Path != nil {
			for _, p := range e.Path.([]interface{}) {
				path = append(path, fmt.Sprintf("%v", p))
			}
		}
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  fmt.Sprintf("%s (%s)", e.Details, strings.Join(path, ".")),
		})
	}

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  contentfulErr.Message,
	})

	return diags
}
