package evaluator

import (
	"fmt"

	"github.com/indeedhat/monkey-lang/evaluator/object"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Name: "len",
		Fn:   builtinLen,
	},
	"append": {
		Name: "append",
		Fn:   builtinAppend,
	},
	"print": {
		Name: "print",
		Fn:   builtinPrint,
	},
	"printf": {
		Name: "print",
		Fn:   builtinPrintf,
	},
	"sprintf": {
		Name: "print",
		Fn:   builtinSprintf,
	},
}

func builtinLen(args ...object.Object) object.Object {
	if len(args) != 1 {
		return error("bad arg count: expect(1) found(%d)", len(args))
	}

	switch subject := args[0].(type) {
	case *object.String:
		return &object.Integer{Value: int64(len(subject.Value))}
	case *object.Array:
		return &object.Integer{Value: int64(len(subject.Elements))}
	}

	return error("invalid arg type `%s`", args[0].Type())
}

func builtinAppend(args ...object.Object) object.Object {
	if len(args) < 2 {
		return error("append requires at least 2 arguments")
	}

	array, ok := args[0].(*object.Array)
	if !ok {
		return error("bad subject type: expected(*object.Array) found(%T)", args[0])
	}

	subject := *array

	subject.Elements = append(subject.Elements, args[1:]...)

	return &subject
}

func builtinPrint(args ...object.Object) object.Object {
	printables := make([]any, 0, len(args))

	for _, ob := range args {
		printables = append(printables, ob.Inspect())
	}

	fmt.Print(printables...)
	return Void
}

func builtinPrintf(args ...object.Object) object.Object {
	printables := make([]any, 0, len(args))

	for _, ob := range args {
		printables = append(printables, ob.Inspect())
	}

	fmt.Printf(args[0].Inspect(), printables[1:]...)
	return Void
}

func builtinSprintf(args ...object.Object) object.Object {
	printables := make([]any, 0, len(args))

	for _, ob := range args {
		printables = append(printables, ob.Inspect())
	}

	return &object.String{Value: fmt.Sprintf(args[0].Inspect(), printables[1:]...)}
}
