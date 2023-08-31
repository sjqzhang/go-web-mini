package listener

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"testing"
)

func Test_boolean_expression_with_equal_evaluates(t *testing.T) {
	variables := map[string]interface{}{
		"A": "",
	}

	result, err := evaluateExpression("A == \"\"", variables)

	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, result, is.True())

	variables = map[string]interface{}{
		"A": "a value",
	}

	result, err = evaluateExpression("A != \"\"", variables)

	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, result, is.True())
}
