package evaluator

import (
	"github.com/indeedhat/monkey-lang/evaluator/object"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Name: "len",
		Fn:   builtinLen,
	},
}

func builtinLen(args ...object.Object) object.Object {
	if len(args) != 1 {
		return error("bad arg count: expect(1) found(%d)", len(args))
	}

	switch subject := args[0].(type) {
	case *object.String:
		return &object.Integer{Value: int64(len(subject.Value))}
	}

	return error("invalid arg type `%s`", args[0].Type())
}
