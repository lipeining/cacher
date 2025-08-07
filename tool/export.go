package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"export/maps"

	"go/ast"
	"go/parser"
	"go/token"
)

func inspectFile(file *ast.File) map[string][]string {
	structs := map[string][]string{}

	ast.Inspect(file, func(n ast.Node) bool {
		// 检查当前节点是否为类型声明
		if typeSpec, ok := n.(*ast.TypeSpec); ok {
			// 检查类型声明是否为结构体类型
			if structType, ok := typeSpec.Type.(*ast.StructType); ok {
				fields := inspectStruct(structType)
				structs[typeSpec.Name.Name] = fields
			}
		}
		return true
	})

	return structs
}

func inspectStruct(s *ast.StructType) []string {
	fields := []string{}
	for _, field := range s.Fields.List {
		for _, name := range field.Names {
			// 是否需要管 exported？
			fields = append(fields, name.Name)
		}
	}
	return fields
}

func sortKeys(origins []string) []string {
	keys := origins
	sort.Strings(keys)
	return keys
}

func sortMap(origins map[string][]string) map[string][]string {
	keys := maps.Keys(origins)
	sort.Strings(keys)

	structs := map[string][]string{}

	for _, key := range keys {
		structs[key] = sortKeys(origins[key])
	}

	return structs
}

func getFileStructs(filename string) map[string][]string {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, parser.AllErrors)
	if err != nil {
		fmt.Printf("Error parsing file: %v\n", err)
		return nil
	}

	origins := inspectFile(f)
	structs := sortMap(origins)

	return structs
}

func writeStructs(filename string, structs map[string][]string) {
	jsonData, err := json.Marshal(structs)
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}

	os.WriteFile(filename, jsonData, 0644)
}

func main() {
	leftStructs := getFileStructs("left.go")
	rightStructs := getFileStructs("right.go")

	writeStructs("left.json", leftStructs)
	writeStructs("right.json", rightStructs)
}
