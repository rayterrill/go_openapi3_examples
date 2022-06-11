package main

import (
	"context"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func Test_Spec(t *testing.T) {
	doc, err := openapi3.NewLoader().LoadFromFile("openapi.yaml")
	if err != nil {
		t.Fatal("unable to load yaml file", err)
	}

	//validate spec
	err = doc.Validate(context.Background())
	if err != nil {
		t.Error("OpenAPI contract failed to compile", err)
	}
}
