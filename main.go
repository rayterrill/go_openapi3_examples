package main

import (
	"os"

	"github.com/getkin/kin-openapi/openapi3"
	"gopkg.in/yaml.v3"
)

// function showing how to use openapi3 to generate an openapi spec
// using golang. intentionally contains a number of contrived examples
// and different ways of doing the same things to illustrate a number
// of options
func main() {
	//build base of spec
	s := openapi3.T{
		OpenAPI: "3.0.0",
	}

	//build info section
	info := &openapi3.Info{
		Title: "Swagger Petstore",
		License: &openapi3.License{
			Name: "MIT",
		},
		Version: "1.0.0",
	}
	//add info to spec
	s.Info = info

	//tags
	tags := openapi3.Tags{
		{
			Name:        "test1",
			Description: "something",
		},
		{
			Name:        "test2",
			Description: "something2",
		},
	}
	//add the tags to our doc
	s.Tags = tags

	//build a basic server
	server1 := openapi3.Server{
		URL: "http://petstore.swagger.io/v1",
	}
	//add the server to our spec
	s.AddServer(&server1)

	//build a server with a variable and description
	server2 := openapi3.Server{
		URL:         "http://petstore.{env}.swagger.io/v1",
		Description: "another server",
		Variables: map[string]*openapi3.ServerVariable{
			"env": {
				Default: "dev",
				Enum: []string{
					"dev",
					"staging",
				},
			},
		},
	}
	//add another server to our spec
	s.AddServer(&server2)

	//build /pets operation
	//build descriptions because we need to attach them to our response as *string
	var description200 = "a paged array of pets"
	var descriptionDefault = "unexpected error"
	getPets := openapi3.Operation{
		Tags:        []string{"pets"},
		Summary:     "List all pets",
		OperationID: "listPets",
		Parameters: openapi3.Parameters{
			&openapi3.ParameterRef{
				Ref: "",
				Value: &openapi3.Parameter{
					Name:        "limit",
					In:          "query",
					Description: "How many items to return at one time (max 100)",
					Required:    false,
					Schema: &openapi3.SchemaRef{
						Ref: "",
						Value: &openapi3.Schema{
							Type:   "integer",
							Format: "int32",
						},
					},
				},
			},
		},
		Responses: openapi3.Responses{
			"200": &openapi3.ResponseRef{
				Ref: "",
				Value: &openapi3.Response{
					Description: &description200,
					Headers: openapi3.Headers{
						"x-next": &openapi3.HeaderRef{
							Ref: "",
							Value: &openapi3.Header{
								Parameter: openapi3.Parameter{
									Description: "A link to the next page of responses",
									Schema: &openapi3.SchemaRef{
										Ref: "",
										Value: &openapi3.Schema{
											Type: "string",
										},
									},
								},
							},
						},
					},
					Content: openapi3.Content{
						"application/json": &openapi3.MediaType{
							Schema: &openapi3.SchemaRef{
								Ref: "#/components/schemas/Pets",
							},
						},
					},
				},
			},
			"default": &openapi3.ResponseRef{
				Ref: "",
				Value: &openapi3.Response{
					Headers: nil,
					Content: openapi3.Content{
						"application/json": &openapi3.MediaType{
							Schema: &openapi3.SchemaRef{
								Ref: "#/components/schemas/Error",
							},
						},
					},
					Description: &descriptionDefault,
					Links:       nil,
				},
			},
		},
	}

	//build pathitem mapping in our getPets method
	pets := openapi3.PathItem{
		Get: &getPets,
	}

	//build paths object and add out pathitem(s)
	paths := openapi3.Paths{
		"/pets": &pets,
	}
	s.Paths = paths

	//build components
	components := openapi3.Components{
		Schemas: openapi3.Schemas{
			"Pet": &openapi3.SchemaRef{
				Ref: "",
				Value: &openapi3.Schema{
					Required: []string{"id", "name"},
					//Enum:         []interface{}{"VALUE1", "VALUE2"},
					Properties: openapi3.Schemas{
						"id": &openapi3.SchemaRef{
							Ref: "",
							Value: &openapi3.Schema{
								Type:   "integer",
								Format: "int64",
							},
						},
						"name": &openapi3.SchemaRef{
							Ref: "",
							Value: &openapi3.Schema{
								Type: "string",
							},
						},
						"tag": &openapi3.SchemaRef{
							Ref: "",
							Value: &openapi3.Schema{
								Type: "string",
							},
						},
					},
				},
			},
			"Pets": &openapi3.SchemaRef{
				Ref: "",
				Value: &openapi3.Schema{
					Type: "array",
					Items: &openapi3.SchemaRef{
						Ref: "#/components/schemas/Pet",
					},
				},
			},
			"Error": &openapi3.SchemaRef{
				Ref: "",
				Value: &openapi3.Schema{
					Required: []string{"code", "message"},
					Properties: openapi3.Schemas{
						"code": &openapi3.SchemaRef{
							Ref: "",
							Value: &openapi3.Schema{
								Type:   "integer",
								Format: "int32",
							},
						},
						"message": &openapi3.SchemaRef{
							Ref: "",
							Value: &openapi3.Schema{
								Type: "string",
							},
						},
					},
				},
			},
		},
	}
	s.Components = components

	//marshall spec into yaml
	d, err := yaml.Marshal(&s)
	if err != nil {
		panic(err)
	}

	//write the yaml to a file
	err = os.WriteFile("output.yaml", d, 0644)
	if err != nil {
		panic(err)
	}
}
