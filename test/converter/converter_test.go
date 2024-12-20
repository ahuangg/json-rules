package test

import (
	"testing"

	"github.com/ahuangg/json-rules/converter"
)


func TestJSONToExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{
			name:     "equal test",
			input:    `{"eq": [{"var": "x"}, 1]}`,
			expected: "(x eq 1)",
		},
		{
			name:     "equal test 2",
			input:    `{"==": [{"var": "x"}, 2]}`,
			expected: "(x == 2)",
		},
		{
			name:     "less than test",
			input:    `{"<": [{"var": "x"}, 2]}`,
			expected: "(x < 2)",
		},
		{
			name:     "less than test 2",
			input:    `{"<": [{"var": "x"}, 1]}`,
			expected: "(x < 1)",
		},
		{
			name:     "greater than test",
			input:    `{">": [{"var": "x"}, 0]}`,
			expected: "(x > 0)",
		},
		{
			name:     "greater than test 2",
			input:    `{">": [{"var": "x"}, 2]}`,
			expected: "(x > 2)",
		},
		{
			name:     "equal and less than equal test",
			input:    `{"and": [{"==": [{"var": "x.a"}, 1]}, {"<=": [{"var": "x.b.c"}, 2]}]}`,
			expected: "((x.a == 1) and (x.b.c <= 2))",
		},
		{
			name:     "equal and greater than test",
			input:    `{"and": [{"==": [{"var": "y"}, 4]}, {">": [{"var": "x"}, 1]}]}`,
			expected: "((y == 4) and (x > 1))",
		},
		{
			name:     "in test",
			input:    `{"and": [{"==": [{"var": "y"}, 4]}, {"in": [{"var": "x"}, [1, 2, 3]]}]}`,
			expected: "((y == 4) and (x in [1 2 3]))",
		},
		{
			name:     "equal string test",
			input:    `{"and": [{"==": [{"var": "y"}, 4]}, {"eq": [{"var": "x"}, "1.2.3"]}]}`,
			expected: `((y == 4) and (x eq 1.2.3))`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := converter.JSONToExpression(tt.input)
			if err != nil {
				return
			} 

			if got != tt.expected {
				t.Errorf("JSONToExpression() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestExpressionToJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{
			name:     "equal test",
			input:    "x eq 1",
			expected: `{"eq":[{"var":"x"},1]}`,
		},
		{
			name:     "equal test 2", 
			input:    "x == 2",
			expected: `{"==":[{"var":"x"},2]}`,
		},
		{
			name:     "less than test",
			input:    "x < 2",
			expected: `{"<":[{"var":"x"},2]}`,
		},
		{
			name:     "less than test 2",
			input:    "x < 1",
			expected: `{"<":[{"var":"x"},1]}`,
		},
		{
			name:     "greater than test",
			input:    "x > 0",
			expected: `{">":[{"var":"x"},0]}`,
		},
		{
			name:     "greater than test 2",
			input:    "x > 2",
			expected: `{">":[{"var":"x"},2]}`,
		},
		{
			name:     "equal and less than equal test",
			input:    "((x.a == 1) and (x.b.c <= 2))",
			expected: `{"and":[{"==":[{"var":"x.a"},1]},{"<=":[{"var":"x.b.c"},2]}]}`,
		},
		{
			name:     "equal and greater than test",
			input:    "y == 4 and (x > 1)",
			expected: `{"and":[{"==":[{"var":"y"},4]},{">":[{"var":"x"},1]}]}`,
		},
		{
			name:     "in test",
			input:    "y == 4 and (x in [1 2 3])",
			expected: `{"and":[{"==":[{"var":"y"},4]},{"in":[{"var":"x"},[1,2,3]]}]}`,
		},
		{
			name:     "equal string test",
			input:    `y == 4 and (x eq 1.2.3)`,
			expected: `{"and":[{"==":[{"var":"y"},4]},{"eq":[{"var":"x"},"1.2.3"]}]}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := converter.ExpressionToJSON(tt.input)
			if err != nil {
				return
			}

			if got != tt.expected {
				t.Errorf("ExpressionToJSON() = %v, want %v", got, tt.expected)
			}
		})
	}
}