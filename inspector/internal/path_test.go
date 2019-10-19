package internal

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go/ast"
	"go/parser"
	"go/token"

	"testing"
)

var funcs = map[string]*ast.FieldList{}

func TestMain(m *testing.M) {
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, "./testdata/sample.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	for _, stmt := range file.Decls {
		if funcDecl, ok := stmt.(*ast.FuncDecl); ok {
			funcs[funcDecl.Name.Name] = funcDecl.Type.Params
		}
	}
	m.Run()
}

func Test_pathForFieldList(t *testing.T) {
	cases := []struct {
		name          string
		expected      string
		expectedForce string
	}{
		{
			name:          "noParams",
			expected:      "",
			expectedForce: "()",
		},
		{
			name:          "singleParamString",
			expected:      "string",
			expectedForce: "(string)",
		},
		{
			name:          "singleParamInt32",
			expected:      "int32",
			expectedForce: "(int32)",
		},
		{
			name:          "singleParamInt32Pointer",
			expected:      "*int32",
			expectedForce: "(*int32)",
		},
		{
			name:          "singleParamPerson",
			expected:      "person",
			expectedForce: "(person)",
		},
		{
			name:          "singleParamExternalPerson",
			expected:      "_package.Person",
			expectedForce: "(_package.Person)",
		},
		{
			name:          "singleParamPersonPointer",
			expected:      "*person",
			expectedForce: "(*person)",
		},
		{
			name:          "singleParamExternalPersonPointer",
			expected:      "*_package.Person",
			expectedForce: "(*_package.Person)",
		},
		{
			name:          "singleParamInterface",
			expected:      "interface{}",
			expectedForce: "(interface{})",
		},
		{
			name:          "singleParamStruct",
			expected:      "struct{}",
			expectedForce: "(struct{})",
		},
		{
			name:          "singleParamArrayOfString",
			expected:      "[]string",
			expectedForce: "([]string)",
		},
		{
			name:          "singleParamArrayOfInterface",
			expected:      "[]interface{}",
			expectedForce: "([]interface{})",
		},
		{
			name:          "singleParamArrayOfStruct",
			expected:      "[]struct{}",
			expectedForce: "([]struct{})",
		},
		{
			name:          "singleParamArrayOfStringPointer",
			expected:      "[]*string",
			expectedForce: "([]*string)",
		},
		{
			name:          "singleParamMapStringString",
			expected:      "map[string]string",
			expectedForce: "(map[string]string)",
		},
		{
			name:          "singleParamMapStringPerson",
			expected:      "map[string]_package.Person",
			expectedForce: "(map[string]_package.Person)",
		},
		{
			name:          "singleParamMapStringPersonPointer",
			expected:      "map[string]*_package.Person",
			expectedForce: "(map[string]*_package.Person)",
		},
		{
			name:          "singleParamFuncEmpty",
			expected:      "func()",
			expectedForce: "(func()())",
		},
		{
			name:          "singleParamFuncArg",
			expected:      "func(*_package.Person)",
			expectedForce: "(func(*_package.Person)())",
		},
		{
			name:          "singleParamFuncArgStringInt",
			expected:      "func(string,int)",
			expectedForce: "(func(string,int)())",
		},
		{
			name:          "singleParamFuncArgStringPersonPointer",
			expected:      "func(string,*_package.Person)",
			expectedForce: "(func(string,*_package.Person)())",
		},
	}
	for _, c := range cases {
		fmt.Printf("[TEST] %s: %s \n",c.name, c.expected)
		fieldList := funcs[c.name]
		result := pathForFieldList(fieldList, true)
		if !assert.EqualValues(t, c.expectedForce, result) {
			t.FailNow()
		}
		result = pathForFieldList(fieldList, false)
		if !assert.EqualValues(t, c.expected, result) {
			t.FailNow()
		}
	}

}