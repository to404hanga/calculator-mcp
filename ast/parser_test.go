package ast

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/to404hanga/calculator-mcp/calculator"
)

func TestParser(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1 + 2", "3"},
		{"2 * 3", "6"},
		{"(1 + 2) * 3", "9"},
		{"PI", "3.141592653589793238462643383279502884197169399375105820974944592307816406286"},
		{"2 * PI", "6.283185307179586476925286766559005768394338798750211641949889184615632812572"},
		{"sqrt(16)", "4"},
		{"sqrt(2 * 8)", "4"},
		{"sqrt(PI)", "1.772453850905516027298167483341145182797549456122387128213807789852911284591"},
		{"(1 + sqrt(16)) * 2", "10"},
	}

	calc := calculator.NewCalculator(75)

	for _, test := range tests {
		parser := NewParser(test.input, calc)
		result := parser.Parse().Evaluate()
		// Compare decimal values instead of float64
		expected, _ := decimal.NewFromString(test.expected)
		actual, _ := decimal.NewFromString(result)
		if !expected.Equal(actual) {
			t.Errorf("对于输入 %s: 期望 %s, 得到 %s", test.input, test.expected, result)
		}
	}
}

func TestExtendedOperations(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"2 ^ 3", "8"},
		{"sin(PI / 2)", "1"},
		{"cos(PI)", "-1"},
		{"2 ^ 2 + sin(PI/2)", "5"},
		{"cos(0) ^ 2", "1"},
		{"sin(PI/6) ^ 2 + cos(PI/6) ^ 2", "1"}, // sin²θ + cos²θ = 1
		{"sin(0)", "0"},                        // 添加边界值测试
		{"cos(PI/2)", "0"},                     // 添加边界值测试
		{"sin(PI)", "0"},                       // 添加边界值测试
		{"2 ^ 0", "1"},                         // 添加指数为0的测试
		{"2 ^ 0.5", "1.4142135623730950488016887242096980785696718753769480731766797379907324784621"}, // 添加小数指数测试
		{"tan(PI/4)", "1"},                          // tan(π/4) = 1
		{"asin(0)", "0"},                           // arcsin(0) = 0
		{"acos(1)", "0"},                           // arccos(1) = 0
		{"atan(1)", "0.7853981633974483"},         // arctan(1) = π/4
		{"tan(0)", "0"},                            // tan(0) = 0
		{"asin(1)", "1.5707963267948966"},         // arcsin(1) = π/2
		{"acos(0)", "1.5707963267948966"},         // arccos(0) = π/2
		{"atan(0)", "0"},                           // arctan(0) = 0
	}

	calc := calculator.NewCalculator(75)

	for _, test := range tests {
		parser := NewParser(test.input, calc)
		result := parser.Parse().Evaluate()
		// Compare decimal values instead of float64
		expected, _ := decimal.NewFromString(test.expected)
		actual, _ := decimal.NewFromString(result)
		// Use a higher precision comparison for trigonometric functions
		if !expected.Sub(actual).Abs().LessThan(decimal.NewFromFloat(1e-10)) {
			t.Errorf("对于输入 %s: 期望 %s, 得到 %s", test.input, test.expected, result)
		}
	}
}
