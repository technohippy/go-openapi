package gen

import (
	//"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
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

func (o *OpenAPIObject) TemplateVariables() map[string]interface{} {
	variables := map[string]interface{}{
		"packageName":    "swagger", // TODO
		"hasImport":      false,     // TODO
		"modelPackage":   "TODO:modelPackage",
		"package":        "TODO:package",
		"clientPackage":  "TODO:clientPackage",
		"version":        "TODO:version",
		"classVarName":   "TODO:classVarName",
		"basePath":       "TODO:basePath",
		"packageVersion": o.Info.Version,       // TODO
		"infoEmail":      o.Info.Contact.Email, // TODO
		"importPath":     "TODO:importPath",
		"licenceInfo":    o.Info.License.Name,
		"hasMore":        false, // TODO
		"generatedDate":  "TODO:generatedDate",
		"classname":      "TODO:classname",
		"imports": []map[string]string{{
			"import": "TODO:import",
		}},
		"appName":        o.Info.Title,
		"appVersion":     o.Info.Version,
		"generatorClass": "gen/generator.go",
		"baseName":       "TODO:baseName",
		"contextPath":    "TODO:contextPath",
		"apiInfo": map[string]interface{}{
			"apis": map[string]interface{}{
				"operations": o.OperationsTemplateVariables(),
				"hasMore":    false,
			},
		},
		"models":      o.ModelsTemplateVariables(),
		"authMethods": []interface{}{},
	}
	return variables
}

func (o *OpenAPIObject) OperationsTemplateVariables() map[string]interface{} {
	operationVariables := []map[string]interface{}{}
	for path, pathItem := range o.Paths {
		operationVariables = append(operationVariables, pathItem.OperationTemplateVariables(path))
	}
	operationsVariables := map[string]interface{}{"operation": operationVariables}
	return operationsVariables
}

func (o *OpenAPIObject) ModelsTemplateVariables() map[string]interface{} {
	return map[string]interface{}{}
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

// https://github.com/swagger-api/swagger-codegen/wiki/Mustache-Template-Variables
func (o PathItemObject) OperationTemplateVariables(path string) map[string]interface{} {
	operationValue := map[string]interface{}{}
	operationValue["responseHeaders"] = []interface{}{}
	opType, op := o.Operation()
	if op != nil {
		produces := []map[string]interface{}{}
		for _, produce := range op.Produces() {
			produces = append(produces, map[string]interface{}{
				"hasMore":   true,
				"mediaType": produce})
		}
		produces[len(produces)-1]["hasMore"] = false

		operationValue["hasProduces"] = true
		if 0 < len(op.Parameters) {
			operationValue["hasParams"] = true
		}
		operationValue["hasMore"] = true // last one must be false
		mediaType := produces[0]["mediaType"].(string)
		operationValue["isResponseBinary"] = strings.HasPrefix(mediaType, "image/") ||
			strings.HasPrefix(mediaType, "audio/") ||
			strings.HasPrefix(mediaType, "video/") ||
			strings.HasPrefix(mediaType, "application/octet-stream")
		operationValue["path"] = path
		operationValue["operationId"] = op.OperationId
		operationValue["httpMethod"] = opType
		operationValue["summary"] = op.Summary
		operationValue["notes"] = op.Description
		operationValue["baseName"] = "" // TODO
		operationValue["produces"] = produces
		bodyParams := []map[string]interface{}{}
		formParams := []map[string]interface{}{}
		pathParams := []map[string]interface{}{}
		queryParams := []map[string]interface{}{}
		headerParams := []map[string]interface{}{}
		for _, param := range op.Parameters {
			var params []map[string]interface{}
			isBodyParam := false
			switch param.In {
			case "body":
				params = bodyParams
				isBodyParam = true
			case "form":
				params = formParams
			case "path":
				params = pathParams
			case "query":
				params = queryParams
			case "header":
				params = headerParams
			}
			prm := map[string]interface{}{
				"isBodyParam":      isBodyParam,
				"baseName":         param.Name,
				"paramName":        param.Name,
				"dataType":         opType, // TODO: is it ok?
				"description":      param.Description,
				"jsonSchema":       "",
				"isEnum":           false,
				"vendorExtensions": nil,
				"required":         param.Required,
			}
			params = append(params, prm)
		}
		if 0 < len(bodyParams) {
			operationValue["bodyParam"] = bodyParams[0]
			operationValue["bodyParams"] = bodyParams
			operationValue["hasBodyParam"] = true
			operationValue["hasBodyParams"] = true
		}
		if 0 < len(formParams) {
			operationValue["formParams"] = formParams
			operationValue["hasFormParams"] = true
		}
		if 0 < len(pathParams) {
			operationValue["pathParams"] = pathParams
			operationValue["hasPathParams"] = true
		}
		if 0 < len(queryParams) {
			operationValue["queryParams"] = queryParams
			operationValue["hasQueryParams"] = true
		}
		if 0 < len(headerParams) {
			operationValue["headerParams"] = headerParams
			operationValue["hasHeaderParams"] = true
		}
		allParams := []map[string]interface{}{}
		allParams = append(allParams, bodyParams...)
		allParams = append(allParams, formParams...)
		allParams = append(allParams, pathParams...)
		allParams = append(allParams, queryParams...)
		allParams = append(allParams, headerParams...)
		operationValue["allParams"] = allParams
		operationValue["tags"] = op.Tags
		responses := []map[string]interface{}{}
		for code, response := range op.Responses {
			resp := map[string]interface{}{
				"headers":         response.Headers,
				"code":            code,
				"message":         response.Description,
				"hasMore":         true,
				"isDefault":       code == "default",
				"simpleType":      true,  // TODO
				"primitiveType":   true,  // TODO
				"isMapContainer":  false, // TODO
				"isListContainer": false, // TODO
				"isBinary":        false, // TODO
				"jsonSchema":      "{}",  // TODO
				"wildcard":        true,  // TODO
			}
			responses = append(responses, resp)
		}
		responses[len(responses)-1]["hasMore"] = false
		operationValue["responses"] = responses
		operationValue["imports"] = ""           // TODO
		operationValue["vendorExtensions"] = nil // TODO
		operationValue["nickname"] = ""          // TODO
	}

	return operationValue // TODO
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
			if !found {
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
	Ref string `yaml:"$ref,omitempty"`

	PathItemObject
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
