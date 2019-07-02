package lib

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// APIDefinitions holds all information about SORACOM API
type APIDefinitions struct {
	Host       string
	Schemes    []string
	BasePath   string
	Methods    []APIMethod
	StructDefs map[string]StructDef
}

// APIMethod holds information about an API method
type APIMethod struct {
	Path        string
	Method      string
	Use         string
	Tags        []string                 `yaml:"tags"`
	Summary     string                   `yaml:"summary"`
	Description string                   `yaml:"description"`
	OperationID string                   `yaml:"operationId"`
	Security    []map[string]interface{} `yaml:"security"`
	CLI         []string                 `yaml:"x-soracom-cli"`
	Parameters  []APIParam               `yaml:"parameters"`
	Responses   map[string]APIResponse   `yaml:"responses"`
}

// APIParam holds information about an API parameter
type APIParam struct {
	Name        string             `yaml:"name"`
	In          string             `yaml:"in"`
	Required    bool               `yaml:"required"`
	Description string             `yaml:"description"`
	Type        string             `yaml:"type"`
	Enum        []string           `yaml:"enum"`
	Schema      APIParamSchema     `yaml:"schema"`
	Items       APIParamArrayItems `yaml:"items"`
}

// APIParamSchema holds information about a Schema in an API parameter
type APIParamSchema struct {
	Type   string          `yaml:"type"`
	Format string          `yaml:"format"`
	Items  *APIParamSchema `yaml:"items"`
	Ref    string          `yaml:"$ref"`
}

// APIParamArrayItems holds information about array type of parameters
type APIParamArrayItems struct {
	Type   string `yaml:"type"`
	Format string `yaml:"format"`
}

// APIResponse holds information about an API response
type APIResponse struct {
	Schema *APIParamSchema
}

// StructDef holds all information under /definitions in API definition file
type StructDef struct {
	Properties []StructProperty
	References []StructReference
}

// StructProperty holds information about an entry under /definitions/{StructName}/properties in API definition file
type StructProperty struct {
	Name        string
	Type        string
	Format      string
	Description string
	Required    bool
}

// StructReference holds information about a reference from a property in API definition file
type StructReference struct {
	Name    string
	RefType string
}

// LoadAPIDef loads API definitions from the specified file
func LoadAPIDef(apiDefYAMLFile string) (*APIDefinitions, error) {
	apiDefYAML, err := loadAPIDefYAML(apiDefYAMLFile)
	if err != nil {
		return nil, err
	}

	apiDefMap := make(map[interface{}]interface{})
	err = yaml.Unmarshal(bytes.NewBufferString(apiDefYAML).Bytes(), &apiDefMap)
	if err != nil {
		return nil, err
	}

	structDefs := loadStructDefs(apiDefMap["definitions"].(map[interface{}]interface{}))
	methods, err := loadMethods(apiDefMap)
	if err != nil {
		return nil, err
	}

	return &APIDefinitions{
		Host:       apiDefMap["host"].(string),
		Schemes:    loadSchemes(apiDefMap),
		BasePath:   apiDefMap["basePath"].(string),
		Methods:    methods,
		StructDefs: structDefs,
	}, nil
}

func loadSchemes(apiDefMap map[interface{}]interface{}) []string {
	result := []string{}
	for _, v := range apiDefMap["schemes"].([]interface{}) {
		result = append(result, v.(string))
	}
	return result
}

func loadAPIDefYAML(inputFile string) (string, error) {
	f, err := os.Open(inputFile)
	if err != nil {
		return "", err
	}
	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Printf("warning: unable to close file %s", f.Name())
		}
	}()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return bytes.NewBuffer(data).String(), nil
}

func loadStructDefs(data map[interface{}]interface{}) map[string]StructDef {
	//fmt.Println("loading struct definitions")
	result := map[string]StructDef{}
	for structName, structDef := range data {
		//fmt.Printf("  name: %s\n", structName)
		s := structDef.(map[interface{}]interface{})
		if s["properties"] == nil {
			continue
		}
		var required = map[string]bool{}
		if s["required"] != nil {
			for _, v := range s["required"].([]interface{}) {
				name := v.(string)
				required[name] = true
			}
		}
		var d StructDef
		for propName, propDef := range s["properties"].(map[interface{}]interface{}) {
			//fmt.Printf("    %s: ", propName)
			var sp StructProperty
			sp.Name = propName.(string)
			p := propDef.(map[interface{}]interface{})
			if p["type"] == nil {
				continue
			}
			t := p["type"].(string)
			sp.Type = t
			//fmt.Printf("%s", sp.Type)

			if p["format"] != nil {
				sp.Format = p["format"].(string)
				//fmt.Printf(" (%s)", sp.Format)
			}

			if required[sp.Name] {
				sp.Required = true
			}
			//fmt.Println()
			d.Properties = append(d.Properties, sp)
		}
		result[structName.(string)] = d
	}
	return result
}

func loadMethods(apiDefMap map[interface{}]interface{}) ([]APIMethod, error) {
	result := make([]APIMethod, 0, len(apiDefMap))
	paths := apiDefMap["paths"].(map[interface{}]interface{})
	for path, p := range paths {
		methods := p.(map[interface{}]interface{})
		for method, m := range methods {
			mi, err := decodeAPIMethod(m.(map[interface{}]interface{}))
			if err != nil {
				return nil, err
			}
			mi.Path = path.(string)
			mi.Method = method.(string)
			result = append(result, *mi)
		}
	}
	return result, nil
}

func decodeAPIMethod(data map[interface{}]interface{}) (*APIMethod, error) {
	y, err := yaml.Marshal(&data)
	if err != nil {
		return nil, err
	}
	var result APIMethod
	err = yaml.Unmarshal(y, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
