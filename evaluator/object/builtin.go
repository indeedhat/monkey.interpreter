package object

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Name string
	Fn   BuiltinFunction
}

// Inspect implements Object
func (b *Builtin) Inspect() string {
	return "builtin<" + b.Name + ">"
}

// Type implements Object
func (*Builtin) Type() ObjectType {
	return BuiltinObj
}

var _ Object = (*Builtin)(nil)
