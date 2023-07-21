// Copyright 2023 Juan Pablo Tosso and the OWASP Coraza contributors
// SPDX-License-Identifier: Apache-2.0

package operators

import (
	"net/http"
	"strings"
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers/gorillamux"
	"github.com/nguyentin2068/waf/experimental/plugins/plugintypes"
)

type validateOpenAPI struct{}

var _ plugintypes.Operator = (*validateOpenAPI)(nil)

func newValidateOpenAPI(plugintypes.OperatorOptions) (plugintypes.Operator, error) {
	return &validateOpenAPI{}, nil
}

func (o *validateOpenAPI) Evaluate(_ plugintypes.TransactionState, value string) bool {
	schemaFile := "/opt/APISchema/api.json"
	reqe := strings.Split(value, " ")
	methd := reqe[0]
	uri := reqe[1]
	req, _ := http.NewRequest(methd, uri, nil)
	// Load the OpenAPI document
	loader := openapi3.NewLoader()
	doc, _ := loader.LoadFromFile(schemaFile)

	// Find the operation (HTTP method + path) that matches the request
	router, _ := gorillamux.NewRouter(doc)
	route, pathParams, _ := router.FindRoute(req)

	// Create a RequestValidationInput
	requestValidationInput := &openapi3filter.RequestValidationInput{
		Request:    req,
		PathParams: pathParams,
		Route:      route,
	}
	httpreq := req.Context()
	// Validate the request
	if er := openapi3filter.ValidateRequest(httpreq, requestValidationInput); er != nil {
		return false
	}
	return true
}

func init() {
	Register("validateOpenAPI", newValidateOpenAPI)
}
