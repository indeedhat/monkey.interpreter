package evaluator

import (
	"testing"

	"github.com/indeedhat/monkey-lang/evaluator/object"
)

var ifElseTests = []struct {
	input string
	value any
}{
	{"if (true) { 10 }", 10},
	{"if (false) { 10 }", nil},
	{"if (1) { 10 }", 10},
	{"if (1 < 2) { 10 }", 10},
	{"if (1 > 2) { 10 }", nil},
	{"if (1 > 2) { 10 } else { 20 }", 20},
	{"if (1 < 2) { 10 } else { 20 }", 10},
	{"if (1 > 2) { 10 } else if (false) { 20 } else { 30 }", 30},
	{"if (1 > 2) { 10 } else if (true) { 20 } else { 30 }", 20},
}

func TestIfElseExpressions(t *testing.T) {
	for _, tCase := range ifElseTests {
		t.Run(tCase.input, func(t *testing.T) {
			evald := testEval(tCase.input)

			switch val := tCase.value.(type) {
			case int:
				testIntegerObject(t, evald, int64(val))
			case nil:
				testNullObject(t, evald)
			}
		})
	}
}

func testNullObject(t *testing.T, obj object.Object) {
	if _, ok := obj.(*object.Null); !ok {
		t.Fatalf("bad object type: expect(Null) found(%T)", obj)
	}
}
