package main

import (
	"fmt"
	"strings"
)

func GenerateMethod(g Group) (codes []string) {
	name := upperFirstLetter(g.Name)
	hasParams := len(g.Fields) != 0

	if hasParams {
		paramsName := name + "Params"

		codes = append(codes, GenerateType(Group{
			Name:        paramsName,
			Description: "contains the method's parameters",
			Fields:      g.Fields,
		}))

		hasUploadable, isGenerated := generateHasUploadable(paramsName, g.Fields)
		if isGenerated {
			codes = append(codes, hasUploadable)
		}
	}

	codes = append(codes, generateMethod(name, hasParams, g))

	return codes
}

func generateHasUploadable(name string, fields []Field) (text string, generated bool) {
	var inputFileFields []string
	for _, f := range fields {
		if typeOf(f.Name, f.Type) == "InputFile" {
			inputFileFields = append(inputFileFields, fmt.Sprintf("d.%s.NeedsUpload()", snakeToPascal(f.Name)))
		}
	}

	generated = inputFileFields != nil
	if generated {
		text = fmt.Sprintf("func (d %s) HasUploadable() bool {\nreturn %s\n}", name, strings.Join(inputFileFields, " || "))
	}
	return
}

func generateMethod(name string, hasParams bool, g Group) string {
	var params string
	paramsParam := "nil"
	if hasParams {
		params = "params " + name + "Params"
		paramsParam = "params"
	}

	returnType := "json.RawMessage"
	if types := extractReturnType(g.Description); len(types) == 1 {
		returnType = typeOf(g.Name, types[0])
	}

	return fmt.Sprintf(`// %s %s
func (c *API) %s(%s) (data %s,err error) {
	return doHTTP[%s](c.client, c.url, "%s", %s)
}`, name, g.Description, name, params, returnType, returnType, g.Name, paramsParam)
}