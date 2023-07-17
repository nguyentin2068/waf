module github.com/nguyentin2068/waf

go 1.18

// Testing dependencies:
// - go-mockdns
// - go-modsecurity (optional)

// Development dependencies:
// - mage

// Build dependencies:
// - libinjection-go
// - aho-corasick
// - gjson
// - binaryregexp

require (
	github.com/anuraaga/go-modsecurity v0.0.0-20220824035035-b9a4099778df
	github.com/corazawaf/coraza/v3 v3.0.2
	github.com/corazawaf/libinjection-go v0.1.2
	github.com/foxcpp/go-mockdns v1.0.0
	github.com/getkin/kin-openapi v0.118.0
	github.com/magefile/mage v1.15.0
	github.com/mccutchen/go-httpbin/v2 v2.9.0
	github.com/nguyentin2068/validateopenapi v0.0.0-20230705040532-37661eed72af
	github.com/petar-dambovaliev/aho-corasick v0.0.0-20211021192214-5ab2d9280aa9
	github.com/tidwall/gjson v1.14.4
	golang.org/x/net v0.11.0
	rsc.io/binaryregexp v0.2.0
)

require (
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/swag v0.19.5 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/invopop/yaml v0.1.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/miekg/dns v1.1.50 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/perimeterx/marshmallow v1.1.4 // indirect
	github.com/stretchr/testify v1.8.2 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	golang.org/x/mod v0.8.0 // indirect
	golang.org/x/sys v0.9.0 // indirect
	golang.org/x/tools v0.6.0 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
