package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3gen"
	"github.com/godpeny/goserv/ent"
	"sigs.k8s.io/yaml"
)

type Swagger struct {
	Components openapi3.Components `json:"components,omitempty" yaml:"components,omitempty"`
}

func main() {

	components := openapi3.NewComponents()

	// Generate Components Schema
	components.Schemas = make(map[string]*openapi3.SchemaRef)

	UserSchema, _, err := openapi3gen.NewSchemaRefForValue(&ent.User{})
	checkErr(err)
	APIResponseSchema, _, err := openapi3gen.NewSchemaRefForValue(&ent.APIResponse{})
	checkErr(err)

	// Generate Components RequestBodies
	components.RequestBodies = make(map[string]*openapi3.RequestBodyRef)

	UserRequestBodies := createRequestBodiesSchema(UserSchema)

	// Set Schema
	components.Schemas["ApiResponse"] = APIResponseSchema
	components.Schemas["User"] = UserSchema

	// Set RequestBodies
	components.RequestBodies["User"] = UserRequestBodies

	// Generate Swagger YAML
	gen(&components)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func createRequestBodiesSchema(schema *openapi3.SchemaRef) *openapi3.RequestBodyRef {
	consume := []string{"application/json"}
	reqBody := openapi3.NewRequestBody().WithSchemaRef(schema, consume)

	return &openapi3.RequestBodyRef{
		Value: reqBody,
	}
}

func gen(components *openapi3.Components) {
	swagger := Swagger{}
	swagger.Components = *components

	b := &bytes.Buffer{}
	json.NewEncoder(b).Encode(swagger)

	schema, err := yaml.JSONToYAML(b.Bytes())
	checkErr(err)

	b = &bytes.Buffer{}
	b.Write(schema)

	err = ioutil.WriteFile("./api/server/components-server.yaml", b.Bytes(), 0666)
	checkErr(err)

	paths, err := ioutil.ReadFile("./api/server/path-server.yaml")

	b = &bytes.Buffer{}
	b.Write(paths)
	b.Write(schema)

	err = ioutil.WriteFile("./api/swagger-server.yaml", b.Bytes(), 0666)
	checkErr(err)

	// Check if generated YAML is valid
	_, err = openapi3.NewSwaggerLoader().LoadSwaggerFromData(b.Bytes())
	checkErr(err)
}
