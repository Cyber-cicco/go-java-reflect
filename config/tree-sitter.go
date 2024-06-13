package config

import (
    sitter "github.com/smacker/go-tree-sitter"
    "github.com/smacker/go-tree-sitter/java"
)

var JavaLang *sitter.Language
var JavaParser *sitter.Parser

func init() {
    JavaParser = sitter.NewParser()
    JavaLang = java.GetLanguage()
    JavaParser.SetLanguage(JavaLang)
}
