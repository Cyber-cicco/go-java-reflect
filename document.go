package reflect

import (
	"errors"

	"github.com/Cyber-cicco/tree-sitter-query-builder/querier"
	sitter "github.com/smacker/go-tree-sitter"
)

type Document struct {
    root *sitter.Node
    content []byte
}

//Returns all imports object of a given java document
func (d *Document) GetImports() []Type {
    return []Type{}
}

func (d *Document) GetMainClass() *Class {
    return &Class{}
}

func (d *Document) GetAbsolutePath() string {
    return ""
}

func (d *Document) GetPackage() (string, error){
    match := querier.GetFirstMatch(d.root, func(node *sitter.Node) bool {
        return node.Type() == "package_declaration"
    })

    if match == nil {
        return "", errors.New("Class doesn't have a package declaration")
    }

    identifier := querier.GetFirstMatch(match, func(node *sitter.Node) bool {
        return node.Type() == "scoped_identifier"
    })

    if identifier == nil {
        return "", errors.New("found package declaration without scoped identifier")
    }

    return identifier.Content(d.content), nil

}
