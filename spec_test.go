package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	validator "github.com/rayterrill/httptest-openapi/openapi3"
)

//this test should fail - response body is incorrect (missing id)
func Test_Handler_Contract(t *testing.T) {
	doc, err := openapi3.NewLoader().LoadFromFile("openapi.yaml")
	if err != nil {
		t.Error("unexpected test setup error reading OpenAPI contract", err)
	}

	t.Run("Validate Spec", func(t *testing.T) {
		//validate spec
		err = doc.Validate(context.Background())
		if err != nil {
			t.Error("OpenAPI contract failed to compile", err)
		}
	})

	t.Run("GET PETS FAIL", func(t *testing.T) {
		handler := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("content-type", "application/json")
			w.WriteHeader(200)
			w.Write([]byte(`{"name": "test"}`))
		}

		rr := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/pets/123", nil)

		v := validator.Validator{Openapi: doc}

		handler(rr, req)

		err = v.Validate(rr, req)

		//this fails - intentionally. if the message matches, dont fail the test
		if err != nil {
			if !strings.Contains(err.Error(), "property \"id\" is missing") {
				t.Error(err)
			}
		}
	})

	t.Run("GET PETS SUCCESS", func(t *testing.T) {
		handler := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("content-type", "application/json")
			w.WriteHeader(200)
			w.Write([]byte(`{"id": 2, "name": "test"}`))
		}

		rr := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/pets/123", nil)

		v := validator.Validator{Openapi: doc}

		handler(rr, req)

		err = v.Validate(rr, req)
		if err != nil {
			t.Error(err)
		}
	})
}
