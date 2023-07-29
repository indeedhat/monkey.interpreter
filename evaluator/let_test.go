package evaluator

import "testing"

var letTests = []struct {
	input string
	value int64
}{
	{"let a = 5; a;", 5},
	{"let a = 5 * 5; a;", 25},
	{"let a = 5; let b = a; b;", 5},
	{"let a = 5; let b = a; let c = a + b + 5; c;", 15},
}

func TestEvalLetStatements(t *testing.T) {
	for _, tCase := range letTests {
		t.Run(tCase.input, func(t *testing.T) {
			evald := testEval(tCase.input)
			testIntegerObject(t, evald, tCase.value)
		})
	}
}
