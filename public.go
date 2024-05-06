package reflect

import (
	"github.com/Cyber-cicco/java-reflect/structs"
	sitter "github.com/smacker/go-tree-sitter"
)

func GetrootClass(n *sitter.Node) *structs.Document {
    return &structs.Document{Root: n}
}
