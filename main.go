package main

import (
	"bytes"
	"fmt"
)

type APIMethod struct {
	Description string
	Method      string
	Path        string
	Title       string
}

func main() {
	methods := map[string][]*APIMethod{
		"Product": []*APIMethod{
			&APIMethod{
				Description: "Allows a user to list the products under their account.",
				Method:      "GET",
				Path:        "/products",
				Title:       "list_products",
			},
			&APIMethod{
				Description: "Allow a user to create a new product under their account.",
				Method:      "POST",
				Path:        "/products",
				Title:       "create_product",
			},
		},
	}

	node := buildAst(methods)
	code := buildCode(node)

	fmt.Println("emitted code:\n")
	fmt.Println(code.String())
}

func buildAst(methods map[string][]*APIMethod) Node {
	resourcesNode := &StatementListNode{}

	parentNode := &StatementListNode{
		[]Node{
			&RequireNode{&StringScalarNode{"excon"}},
			&RequireNode{&StringScalarNode{"stripe"}},
			&ModuleNode{
				&SymbolNode{"Stripe"},
				resourcesNode,
			},
		},
	}

	for resource, endpoints := range methods {
		methodsNode := &StatementListNode{}

		resourceNode := &ClassNode{
			&SymbolNode{resource},
			methodsNode,
		}

		resourcesNode.Children = append(resourcesNode.Children, resourceNode)

		for _, endpoint := range endpoints {
			methodNode := &MethodNode{
				&SymbolNode{endpoint.Title},
				&ArgumentListNode{
					[]*SymbolNode{&SymbolNode{"params"}},
					[]*SymbolNode{},
					[]*DefaultArgumentNode{},
				},
				&StatementListNode{},
			}

			methodsNode.Children =
				append(methodsNode.Children, &CommentNode{endpoint.Description})
			methodsNode.Children = append(methodsNode.Children, methodNode)
		}
	}

	return parentNode
}

func buildCode(node Node) bytes.Buffer {
	emitter := &CodeEmitter{}
	node.Emit(emitter)
	return emitter.Buffer
}
