package evaluator

import (
	"monkey/m/v2/object"
)

var builtins = map[string]*object.BuiltIn{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to `len` not supported, got=%v", arg.Type())
			}
		},
	},
	"first": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			arr, ok := args[0].(*object.Array)
			if !ok {
				return newError("argument to `first` not supported, got=%s", args[0].Type())
			}
			if len(arr.Elements) >= 0 {
				return arr.Elements[0]
			}
			return NULL
		},
	},
	"last": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			arr, ok := args[0].(*object.Array)
			if !ok {
				return newError("argument to `last` not supported, got=%s", args[0].Type())
			}
			length := len(arr.Elements)

			if length >= 0 {
				return arr.Elements[length-1]
			}
			return NULL
		},
	},
	"rest": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			arr, ok := args[0].(*object.Array)
			if !ok {
				return newError("argument to `rest` not supported, got=%s", args[0].Type())
			}

			length := len(arr.Elements)

			if length >= 0 {
				newElements := make([]object.Object, length-1, length-1)
				copy(newElements, arr.Elements[1:])
				return &object.Array{Elements: newElements}
			}
			return NULL
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			arr, ok := args[0].(*object.Array)
			if !ok {
				return newError("argument to `push` not supported, got=%s", args[0].Type())
			}

			length := len(arr.Elements)
			newElements := make([]object.Object, length, length+1)
			copy(newElements, arr.Elements)
			newElements = append(newElements, args[1])
			return &object.Array{Elements: newElements}
		},
	},
}
