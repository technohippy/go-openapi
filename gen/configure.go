package gen

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

func ReadConfig(filepath string) (*OpenAPIObject, error) {
	buf, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var conf OpenAPIObject
	yaml.Unmarshal(buf, &conf)
	return &conf, nil
}

type OpenAPIObject struct {
	OpenAPI      string                      `yaml:"openapi,omitempty"`
	Info         InfoObject                  `yaml:"info,omitempty"`
	Servers      []ServerObject              `yaml:"servers,omitempty"`
	Paths        map[string]PathItemObject   `yaml:"paths,omitempty"`
	Components   ComponentsObject            `yaml:"components,omitempty"`
	Security     []SecurityRequirementObject `yaml:"security,omitempty"`
	Tags         []TagObject                 `yaml:"tags,omitempty"`
	ExternalDocs ExternalDocumentationObject `yaml:"externalDocs,omitempty"`
}

func (o *OpenAPIObject) OperationTemplateVariable() map[string]interface{} {
	operationValue := []map[string]interface{}{}
	for path, pathItem := range o.Paths {
		operationValue = append(operationValue, pathItem.OperationTemplateVariable(path))
	}
	operationsValue := map[string]interface{}{"operation": operationValue}
	return operationsValue // TODO
}

func (o *OpenAPIObject) Yaml() (string, error) {
	buf, err := yaml.Marshal(o)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

type InfoObject struct {
	Title          string        `yaml:"title,omitempty"`
	Description    string        `yaml:"description,omitempty"`
	TermsOfService string        `yaml:"termsOfService,omitempty"`
	Contact        ContactObject `yaml:"contact,omitempty"`
	License        LicenseObject `yaml:"license,omitempty"`
	Version        string        `yaml:"version,omitempty"`
}

type ContactObject struct {
	Name  string `yaml:"name,omitempty"`
	Url   string `yaml:"url,omitempty"`
	Email string `yaml:"email,omitempty"`
}

type LicenseObject struct {
	Name string `yaml:"name,omitempty"`
	Url  string `yaml:"url,omitempty"`
}

type ServerObject struct {
	Url         string                          `yaml:"url,omitempty"`
	Description string                          `yaml:"description,omitempty"`
	Variables   map[string]ServerVariableObject `yaml:"variables,omitempty"`
}

type ServerVariableObject struct {
	Enum        []string `yaml:"enum,omitempty"`
	Default     string   `yaml:"default,omitempty"`
	Description string   `yaml:"description,omitempty"`
}

type ComponentsObject struct {
	Schemas         map[string]SchemaOrRefObject         `yaml:"schemas,omitempty"`
	Responses       map[string]ResponseOrRefObject       `yaml:"responses,omitempty"`
	Parameters      map[string]ParameterOrRefObject      `yaml:"parameters,omitempty"`
	Examples        map[string]ExampleOrRefObject        `yaml:"examples,omitempty"`
	RequestBodies   map[string]RequestBodyOrRefObject    `yaml:"requestBodies,omitempty"`
	Headers         map[string]HeaderOrRefObject         `yaml:"headers,omitempty"`
	SecuritySchemes map[string]SecuritySchemeOrRefObject `yaml:"securitySchemes,omitempty"`
	Links           map[string]LinkOrRefObject           `yaml:"links,omitempty"`
	Callbacks       map[string]CallbackOrRefObject       `yaml:"callbacks,omitempty"`
}

type PathItemObject struct {
	Ref         string                 `yaml:"$ref,omitempty"`
	Summary     string                 `yaml:"summary,omitempty"`
	Description string                 `yaml:"description,omitempty"`
	Get         OperationObject        `yaml:"get,omitempty"`
	Put         OperationObject        `yaml:"put,omitempty"`
	Post        OperationObject        `yaml:"post,omitempty"`
	Delete      OperationObject        `yaml:"delete,omitempty"`
	Options     OperationObject        `yaml:"options,omitempty"`
	Head        OperationObject        `yaml:"head,omitempty"`
	Patch       OperationObject        `yaml:"patch,omitempty"`
	Trace       OperationObject        `yaml:"trace,omitempty"`
	Servers     []ServerObject         `yaml:"servers,omitempty"`
	Parameters  []ParameterOrRefObject `yaml:"parameters,omitempty"`
}

func (o PathItemObject) Operation() (string, *OperationObject) {
	if o.Get.HasValue() {
		return "get", &o.Get
	}
	if o.Put.HasValue() {
		return "put", &o.Put
	}
	if o.Post.HasValue() {
		return "post", &o.Post
	}
	if o.Delete.HasValue() {
		return "delete", &o.Delete
	}
	if o.Options.HasValue() {
		return "options", &o.Options
	}
	if o.Head.HasValue() {
		return "head", &o.Head
	}
	if o.Patch.HasValue() {
		return "patch", &o.Patch
	}
	if o.Trace.HasValue() {
		return "trace", &o.Trace
	}
	return "error", nil
}

func (o PathItemObject) OperationTemplateVariable(path string) map[string]interface{} {
	operationValue := map[string]interface{}{}
	operationValue["responseHeaders"] = []interface{}{}
	opType, op := o.Operation()
	fmt.Sprintf("%v", opType)
	if op != nil {
		operationValue["hasProduces"] = true
		produces := []map[string]interface{}{}
		for _, produce := range op.Produces() {
			operationValue["produces"] = append(produces, map[string]interface{}{
				"hasMore":   true,
				"mediaType": produce})
		}
		produces[len(produces)-1]["hasMore"] = false
		operationValue["produces"] = produces
	}

	// TODO

	return nil
}

type OperationObject struct {
	Tags         []string                       `yaml:"tags,omitempty"`
	Summary      string                         `yaml:"summary,omitempty"`
	Description  string                         `yaml:"description,omitempty"`
	ExternalDocs ExternalDocumentationObject    `yaml:"externalDocs,omitempty"`
	OperationId  string                         `yaml:"operationId,omitempty"`
	Parameters   []ParameterOrRefObject         `yaml:"parameters,omitempty"`
	RequestBody  RequestBodyOrRefObject         `yaml:"requestBody,omitempty"`
	Responses    map[string]ResponseOrRefObject `yaml:"responses,omitempty"`
	Callbacks    map[string]CallbackOrRefObject `yaml:"callbacks,omitempty"`
	Deprecated   bool                           `yaml:"deprecated,omitempty"`
	Security     []SecurityRequirementObject    `yaml:"securty,omitempty"`
	Servers      []ServerObject                 `yaml:"servers,omitempty"`
}

func (o *OperationObject) HasValue() bool {
	return 0 < len(o.Responses)
}

func (o *OperationObject) Produces() []string {
	produces := []string{}
	for _, response := range o.Responses {
		produce := response.Produce()

		if produce != "" {
			found := false
			for _, p := range produces {
				if p == produce {
					found = true
					break
				}
			}
			if found {
				produces = append(produces, produce)
			}
		}
	}
	return produces
}

type ExternalDocumentationObject struct {
	Description string `yaml:"description,omitempty"`
	Url         string `yaml:"url,omitempty"`
}

type ParameterOrRefObject struct {
	Ref string `yaml:"$ref,omitempty"`

	Name            string `yaml:"name,omitempty"`
	In              string `yaml:"in,omitempty"`
	Description     string `yaml:"description,omitempty"`
	Required        bool   `yaml:"required,omitempty"`
	Deprecated      bool   `yaml:"deprecated,omitempty"`
	AllowEmptyValue bool   `yaml:"allowEmptyValue,omitempty"`

	Style         string                        `yaml:"style,omitempty"`
	Explode       bool                          `yaml:"explode,omitempty"`
	AllowReserved bool                          `yaml:"allowReserved,omitempty"`
	Schema        SchemaOrRefObject             `yaml:"schema,omitempty"`
	Example       interface{}                   `yaml:"example,omitempty"`
	Examples      map[string]ExampleOrRefObject `yaml:"examples,omitempty"`
}

type RequestBodyOrRefObject struct {
	Ref string `yaml:"$ref,omitempty"`

	Description string                     `yaml:"description,omitempty"`
	Content     map[string]MediaTypeObject `yaml:"content,omitempty"`
	Required    bool                       `yaml:"required1,omitempty"`
}

type MediaTypeObject struct {
	Schema   SchemaOrRefObject             `yaml:"schema,omitempty"`
	Example  interface{}                   `yaml:"example,omitempty"`
	Examples map[string]ExampleOrRefObject `yaml:"examples,omitempty"`
	Encodng  map[string]EncodingObject     `yaml:"encoding,omitempty"`
}

type EncodingObject struct {
	ContentType   string                       `yaml:"contentType,omitempty"`
	Headers       map[string]HeaderOrRefObject `yaml:"headers,omitempty"`
	Style         string                       `yaml:"style,omitempty"`
	Explode       bool                         `yaml:"explode,omitempty"`
	AllowReserved bool                         `yaml:"allowReserved,omitempty"`
}

type ResponseOrRefObject struct {
	Ref string `yaml:$ref,omitempty"`

	Description string                       `yaml:description,omitempty"`
	Headers     map[string]HeaderOrRefObject `yaml:headers,omitempty"`
	Content     map[string]MediaTypeObject   `yaml:content,omitempty"`
	Links       map[string]LinkOrRefObject   `yaml:links,omitempty"`
}

func (r ResponseOrRefObject) Produce() string {
	for produce, _ := range r.Content {
		return produce
	}
	return ""
}

type CallbackOrRefObject struct {
	PathItemObject
	Ref string `yaml:"$ref,omitempty"`
}

type ExampleOrRefObject struct {
	Ref string `yaml:"$ref,omitempty"`

	Summary       string      `yaml:"summary,omitempty"`
	Description   string      `yaml:"description,omitempty"`
	Value         interface{} `yaml:"value,omitempty"`
	ExternalValue string      `yaml:"externalValue,omitempty"`
}

type LinkOrRefObject struct {
	Ref string `yaml:"$ref,omitempty"`

	OperationRef string                 `yaml:"operationRef,omitempty"`
	OperationId  string                 `yaml:"operationId,omitempty"`
	Parameters   map[string]interface{} `yaml:"parameters,omitempty"`
	RequestBody  interface{}            `yaml:"requestBody,omitempty"`
	Description  string                 `yaml:"description,omitempty"`
	Server       ServerObject           `yaml:"server,omitempty"`
}

type HeaderOrRefObject struct {
	ParameterOrRefObject
}

type TagObject struct {
	Name         string                      `yaml:"name,omitempty"`
	Description  string                      `yaml:"description,omitempty"`
	ExternalDocs ExternalDocumentationObject `yaml:"externalDocs,omitempty"`
}

// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.1.md#schemaObject
type SchemaOrRefObject struct {
	Ref string `yaml:"$ref,omitempty"`

	Type string `yaml:"type,omitempty"`
	// TODO:
}

type SecuritySchemeOrRefObject struct {
	Ref string `yaml:"$ref,omitempty"`

	Type             string           `yaml:"type,omitempty"`
	Description      string           `yaml:"description,omitempty"`
	Name             string           `yaml:"name,omitempty"`
	In               string           `yaml:"in,omitempty"`
	Scheme           string           `yaml:"scheme,omitempty"`
	BearerFormat     string           `yaml:"bearerFormat,omitempty"`
	Flows            OAuthFlowsObject `yaml:"flows,omitempty"`
	OpenIdConnectUrl string           `yaml:"openIdConnectUrl,omitempty"`
}

type OAuthFlowsObject struct {
	Implicit          OAuthFlowObject `yaml:"implicit,omitempty"`
	Password          OAuthFlowObject `yaml:"password,omitempty"`
	ClientCredentials OAuthFlowObject `yaml:"clientCredentials,omitempty"`
	AuthorizationCode OAuthFlowObject `yaml:"authorizationCode,omitempty"`
}

type OAuthFlowObject struct {
	AuthorizationUrl string            `yaml:"authorizationUrl,omitempty"`
	TokenUrl         string            `yaml:"tokenUrl,omitempty"`
	RefreshUrl       string            `yaml:"refreshUrl,omitempty"`
	Scopes           map[string]string `yaml:"scopes,omitempty"`
}

type SecurityRequirementObject struct {
	Requirements map[string][]string
}
