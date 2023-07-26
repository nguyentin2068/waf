// Copyright 2023 Juan Pablo Tosso and the OWASP Coraza contributors
// SPDX-License-Identifier: Apache-2.0

package operators

import (
	"net/http"
	"strings"
	"log"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers/gorillamux"
	"github.com/nguyentin2068/waf/experimental/plugins/plugintypes"
)
const (
	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH" // RFC 5789
	MethodDelete  = "DELETE"
	MethodConnect = "CONNECT"
	MethodOptions = "OPTIONS"
	MethodTrace   = "TRACE"
)

type validateOpenAPI struct{}

var _ plugintypes.Operator = (*validateOpenAPI)(nil)

func newValidateOpenAPI(plugintypes.OperatorOptions) (plugintypes.Operator, error) {
	return &validateOpenAPI{}, nil
}

func (o *validateOpenAPI) Evaluate(_ plugintypes.TransactionState, value string) bool {
	// schemaFile := "/home/tinnt2/FSOFT/github/waf/APISchema/api.json"
	schemaFile := "/opt/APISchema/api.json"
	print(value)
	reqe := strings.Split(value, " ")
	methd :=  getMethod(reqe[0])
	uri := reqe[1]
	req, err := http.NewRequest(methd, uri, nil)
	if err != nil {
		log.Fatal("Error create request:", err)
	}
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

func getMethod(methodStr string) string {
	switch methodStr {
	case "GET":
		return http.MethodGet
	case "HEAD":
		return http.MethodHead
	case "POST":
		return http.MethodPost
	case "PUT":
		return http.MethodPut
	case "PATCH":
		return http.MethodPatch
	case "DELETE":
		return http.MethodDelete
	case "CONNECT":
		return http.MethodConnect
	case "OPTIONS":
		return http.MethodOptions
	case "TRACE":
		return http.MethodTrace
	default:
		return "" // Return an empty string for an unsupported method
	}
}