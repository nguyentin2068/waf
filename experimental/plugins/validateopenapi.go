// Copyright 2023 Juan Pablo Tosso and the OWASP Coraza contributors
// SPDX-License-Identifier: Apache-2.0

package plugins

import (
	"github.com/nguyentin2068/validateopenapi"

	"github.com/corazawaf/coraza/v3/experimental/plugins/plugintypes"
)

func ValidateOpenAPI(name string, validateAPI func() plugintypes.ValidateOpenAPI) {
	validateopenapi.Register(name, validateAPI)
}
