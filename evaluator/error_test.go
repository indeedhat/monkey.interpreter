package evaluator

import (
	"testing"

	"github.com/indeedhat/monkey-lang/evaluator/object"
	"github.com/stretchr/testify/require"
)

var errTests = []struct {
	input   string
	message string
}{
	{
		"5 + true;",
		"type mismatch: int + bool",
	},
	{
		"true + 5;",
		"type mismatch: bool + int",
	},
	{
		"5 + true; 5;",
		"type mismatch: int + bool",
	},
	{
		"-true",
		"unknown operator: -bool",
	},
	{
		"true + false;",
		"unknown operator: bool + bool",
	},
	{
		"5; true + false; 5",
		"unknown operator: bool + bool",
	},
	{
		"if (10 > 1) { true + false; }",
		"unknown operator: bool + bool",
	},
	{
		`
if (10 > 1) {
if (10 > 1) {
return true + false;
}
return 1;
}
`,
		"unknown operator: bool + bool",
	},
	{
		"nope",
		"undefined identifier: nope",
	},
}

func TestErrorHandling(t *testing.T) {
	for _, tCase := range errTests {
		t.Run(tCase.input, func(t *testing.T) {
			evaluated := testEval(tCase.input)

			err, ok := evaluated.(*object.Error)
			if !ok {
				t.Fatalf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			}

			require.Equal(t, tCase.message, err.Message)
		})
	}
}
