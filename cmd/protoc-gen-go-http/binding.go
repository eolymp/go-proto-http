package main

import (
	"errors"
	"github.com/eolymp/go-proto-http/internal/annotations"
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
	"net/http"
	"strings"
)

func getRuleForMethod(method *protogen.Method) (*annotations.HttpRule, bool) {
	opts, ok := method.Desc.Options().(*descriptorpb.MethodOptions)
	if !ok {
		return nil, false
	}

	ext, err := proto.GetExtension(opts, annotations.E_Http)
	if err == proto.ErrMissingExtension {
		return nil, false
	}

	if err != nil {
		panic(err)
	}

	rule, ok := ext.(*annotations.HttpRule)
	if !ok {
		panic(errors.New("http annotation is not HttpRule"))
	}

	return rule, true
}

type Binding struct {
	Method          string   // HTTP method to use
	Path            string   // Path template (following mux package template format
	PathParameters  []string // List of Path parameters (these coming from Path template)
	QueryParameters []string // List of Query parameters
	RequestBody     string   // Name of the field used for request body, or empty if no body is expected, or "*" if whole input is in body
	ResponseBody    string   // Name of the field used in response body, or "*" if whole output is in body (empty defaults to "*")
}

// getBindings for a given Mathod
// @todo: should also extract additionalBindings
func getBindings(desc *protogen.Method) []Binding {
	rule, ok := getRuleForMethod(desc)
	if !ok {
		return nil
	}

	return getBindingsByRule(desc, rule)
}

func getBindingsByRule(desc *protogen.Method, rule *annotations.HttpRule) []Binding {
	var method, path string

	switch p := rule.Pattern.(type) {
	case *annotations.HttpRule_Get:
		method, path = http.MethodGet, p.Get
	case *annotations.HttpRule_Put:
		method, path = http.MethodPut, p.Put
	case *annotations.HttpRule_Post:
		method, path = http.MethodPost, p.Post
	case *annotations.HttpRule_Patch:
		method, path = http.MethodPatch, p.Patch
	case *annotations.HttpRule_Delete:
		method, path = http.MethodDelete, p.Delete
	case *annotations.HttpRule_Custom:
		method, path = p.Custom.GetKind(), p.Custom.GetPath()
	default:
		panic(errors.New("unexpected Pattern type"))
	}

	bindings := []Binding{{
		Method:          method,
		Path:            path, // todo: path from google.api.http annotation is not the same as mux path template, we should convert it correctly
		RequestBody:     rule.Body,
		ResponseBody:    rule.ResponseBody,
		PathParameters:  getPathParameters(path),
		QueryParameters: getQueryParameters(path, desc),
	}}

	for _, r := range rule.AdditionalBindings {
		bindings = append(bindings, getBindingsByRule(desc, r)...)
	}

	return bindings
}

// getPathParameters returns list of "variables" in path template
func getPathParameters(path string) []string {
	var params []string

	parts := strings.Split(path, "/")
	for _, part := range parts {
		if part == "" || part[0] != '{' || part[len(part)-1] != '}' {
			continue
		}

		// todo: resolve patterns such as `{field_name=*}` and `{field_name=**}`
		params = append(params, part[1:len(part)-1])
	}

	return params
}

// getQueryParameters returns list of fields from method.Input which should be populated from query string
// By definition query parameters are used for all fields which are not coming from path nor body
func getQueryParameters(path string, desc *protogen.Method) (params []string) {
	annotation, ok := getRuleForMethod(desc)
	if !ok {
		return
	}

	if annotation.Body == "*" { // body includes everything, query parameters are not needed
		return
	}

	exclude := map[string]string{}

	// exclude body param
	if annotation.Body != "" {
		exclude[annotation.Body] = annotation.Body
	}

	// exclude path params
	for _, param := range getPathParameters(path) {
		exclude[param] = param
	}

	for _, field := range desc.Input.Fields {
		name := string(field.Desc.Name())
		if _, ok := exclude[name]; ok {
			continue
		}

		params = append(params, name)
	}

	return
}
