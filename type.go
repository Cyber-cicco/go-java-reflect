package reflect

type Type struct {
    Identifier string
    OptClass *Class
}

func (t *Type) GetDeclaredName() string {
    if t.OptClass != nil {
        return t.OptClass.GetDeclaredName()
    }
    return t.Identifier
}
