package reflect

import sitter "github.com/smacker/go-tree-sitter"

type Method struct {
    root *sitter.Node
    document *Document
}

//Argument of a given method
type Argument struct {
    root *sitter.Node
}


//Search an argument of a method by it's type and it's name
//
//If you pass nil in either of the args, it doesn't try to match it.
//
//Returns an argument and true if there is match, and an argument with
//an nil Tree Sitter Node and false if it wasn't found
func (m *Method) ArgumentSelector(searchedType *Type, name *string) (Argument, bool) {
    return Argument{}, false
}

//Gets the type of an argument
func (a *Argument) GetType() Type {
    return Type{}
}

//Gets the name of an argument
func (a *Argument) GetName() string {
    return ""
}
