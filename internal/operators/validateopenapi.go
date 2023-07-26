// Copyright 2023 Juan Pablo Tosso and the OWASP Coraza contributors
// SPDX-License-Identifier: Apache-2.0

package operators

import (
	"fmt"
	"net/http"
	"strings"
	"log"

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
	print(value)
	reqe := strings.Split(value, " ")
	methd := reqe[0]
	uri := reqe[1]
	req, _ := http.NewRequest(methd, uri, nil)
	// Load the OpenAPI document
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromFile(schemaFile)
	if err != nil {
		log.Fatal("Error load schema file:", err)
	}

	// Find the operation (HTTP method + path) that matches the request
	router, err := gorillamux.NewRouter(doc)
	if err != nil {
		log.Fatal("Error creating router:", err)
	}
	route, pathParams, err := router.FindRoute(req)
	if err != nil {
		log.Fatal("Load PathParams router error:", err)
	}
	fmt.Printf(route.Path)
	fmt.Printf(route.Method)
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
