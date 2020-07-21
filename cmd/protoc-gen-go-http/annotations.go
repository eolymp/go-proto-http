package main

import (
	"errors"
	"github.com/eolymp/go-proto-http/internal/annotations"
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
	"strings"
)

func getHTTPRuleAnnotation(method *protogen.Method) (*annotations.HttpRule, bool) {
	opts, ok := method.Desc.Options().(*descriptorpb.MethodOptions)
	if !ok {
		return nil, false
	}

	ext, err := proto.GetExtension(opts, annotations.E_Http)
	if err != nil {
		panic(err)
	}

	rule, ok := ext.(*annotations.HttpRule)
	if !ok {
		panic(errors.New("http annotation is not HttpRule"))
	}

	return rule, true
}

type httpBinding struct {
	Method      string
	Path        string
	Parameters  []string
	RequestBody string
}

func getHTTPBindings(method *protogen.Method) []httpBinding {
	annotation, ok := getHTTPRuleAnnotation(method)
	if !ok {
		return nil
	}

	switch p := annotation.Pattern.(type) {
	case *annotations.HttpRule_Get:
		return []httpBinding{{
			Method:      "GET",
			Path:        p.Get,
			RequestBody: annotation.Body,
			Parameters:  getURLParameters(p.Get),
		}}
	case *annotations.HttpRule_Put:
		return []httpBinding{{
			Method:      "PUT",
			Path:        p.Put,
			RequestBody: annotation.Body,
			Parameters:  getURLParameters(p.Put),
		}}
	case *annotations.HttpRule_Post:
		return []httpBinding{{
			Method:      "POST",
			Path:        p.Post,
			RequestBody: annotation.Body,
			Parameters:  getURLParameters(p.Post),
		}}
	case *annotations.HttpRule_Patch:
		return []httpBinding{{
			Method:      "PATCH",
			Path:        p.Patch,
			RequestBody: annotation.Body,
			Parameters:  getURLParameters(p.Patch),
		}}
	case *annotations.HttpRule_Delete:
		return []httpBinding{{
			Method:      "DELETE",
			Path:        p.Delete,
			RequestBody: annotation.Body,
			Parameters:  getURLParameters(p.Delete),
		}}
	case *annotations.HttpRule_Custom:
		return []httpBinding{{
			Method:      p.Custom.GetKind(),
			Path:        p.Custom.GetPath(),
			RequestBody: annotation.Body,
			Parameters:  getURLParameters(p.Custom.GetPath()),
		}}
	}

	return nil
}

func getURLParameters(path string) []string {
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
