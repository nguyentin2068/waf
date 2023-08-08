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
	reqe := strings.Split(value, " ")
	methd :=  getMethod(reqe[0])
	uri := reqe[1]
	req, err := http.NewRequest(methd, uri, nil)
	if err != nil {
		log.Fatal("Error create request:", err)
	}
	// Load the OpenAPI document
	loader := openapi3.NewLoader()
	domain := strings.Split(uri,"/")[2]
	schemaFile := "/etc/coraza-spoa/apischema/"+ domain +".json"
	doc, err := loader.LoadFromFile(schemaFile)
	if err != nil {
		log.Fatal("Error load schema file:", err)
		return true
	}

	// Find the operation (HTTP method + path) that matches the request
	router, err := gorillamux.NewRouter(doc)
	if err != nil {
		log.Fatal("Error creating router:", err)
	}
	// Validate Method, Endpoint from request
	route, pathParams, err := router.FindRoute(req)
	if err != nil {
		return true
	}

	requestValidationInput := &openapi3filter.RequestValidationInput{
		Request:    req,
		PathParams: pathParams,
		Route:      route,
	}
	httpreq := req.Context()
	// Validate the request Paramester required
	if er := openapi3filter.ValidateRequest(httpreq, requestValidationInput); er != nil {
		return true
	}
	return false
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