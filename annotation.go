package reflect

import sitter "github.com/smacker/go-tree-sitter"

type Annotation struct {
    root *sitter.Node
    parent *Annotated
}
