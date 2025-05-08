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
		{"E", "2.718281828459045235360287471352662497757247093699959574966967627724076630353"},
		{"2 * E", "5.436563656918090470720574942705324995514494187399919149933935255448153260706"},
		{"E ^ 2", "7.389056098930650227230427460575007813180315570551847324087127822522573796079"},
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
		{"tan(PI/4)", "1"},                // tan(π/4) = 1
		{"asin(0)", "0"},                  // arcsin(0) = 0
		{"acos(1)", "0"},                  // arccos(1) = 0
		{"atan(1)", "0.7853981633974483"}, // arctan(1) = π/4
		{"tan(0)", "0"},                   // tan(0) = 0
		{"asin(1)", "1.5707963267948966"}, // arcsin(1) = π/2
		{"acos(0)", "1.5707963267948966"}, // arccos(0) = π/2
		{"atan(0)", "0"},                  // arctan(0) = 0
		{"ln(E)", "1"},                     // ln(e) = 1
		{"ln(1)", "0"},                     // ln(1) = 0
		{"ln(E^2)", "2"},                   // ln(e²) = 2
		{"2 * ln(E)", "2"},                 // 2ln(e) = 2
		{"log(8,2)", "3"},                  // log₂(8) = 3
		{"log(1000,10)", "3"},              // log₁₀(1000) = 3
		{"log(E^2,E)", "2"},                // log_e(e²) = 2
		{"log(16,2)", "4"},                 // log₂(16) = 4
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

func TestParserErrors(t *testing.T) {
    tests := []struct {
        input       string
        expectPanic bool
        panicMsg    string
    }{
        {"", true, "表达式不完整"},
        {"(1 + 2", true, "缺少右括号"},
        {"1 + ", true, "表达式不完整"},
        {"sqrt", true, "sqrt后需要括号"},  // 修改期望的错误消息
        {"sqrt(", true, "表达式不完整"},
        {"sqrt)", true, "sqrt后需要括号"},
        {"log(10)", true, "log函数需要两个参数，用逗号分隔"},
        {"log(10,)", true, "log缺少右括号"},  // 修改期望的错误消息
        {"log(,2)", true, "log函数需要两个参数，用逗号分隔"},  // 修改期望的错误消息
        {"log(10,2", true, "log缺少右括号"},
        {"ln(", true, "表达式不完整"},
        {"ln)", true, "ln后需要括号"},
        {"lg(", true, "表达式不完整"},
        {"lg)", true, "lg后需要括号"},
        {"1 +", true, "表达式不完整"},
        {"*", true, "无效的表达式"},      // 修改错误消息
        {"+", true, "无效的表达式"},      // 修改错误消息
        {"-", true, "无效的表达式"},      // 修改错误消息
        {"/", true, "无效的表达式"},      // 修改错误消息
        {"^", true, "无效的表达式"},      // 修改错误消息
        {"+ 2", true, "无效的表达式"},    // 修改错误消息
        {"* 2", true, "无效的表达式"},    // 修改错误消息
        {"/ 2", true, "无效的表达式"},    // 修改错误消息
        {"2 *", true, "表达式不完整"},    // 保持原有错误消息
        {"2 /", true, "表达式不完整"},    // 保持原有错误消息
    }

    calc := calculator.NewCalculator(10)

    for _, test := range tests {
        func() {
            defer func() {
                r := recover()
                if test.expectPanic {
                    if r == nil {
                        t.Errorf("对于输入 %s: 期望发生panic，但没有", test.input)
                    } else if r.(string) != test.panicMsg {
                        t.Errorf("对于输入 %s: 期望panic消息为 %s, 得到 %s", test.input, test.panicMsg, r)
                    }
                } else if r != nil {
                    t.Errorf("对于输入 %s: 不期望发生panic，但发生了: %v", test.input, r)
                }
            }()
            parser := NewParser(test.input, calc)
            parser.Parse()
        }()
    }
}

func TestParserPrecision(t *testing.T) {
    tests := []struct {
        input     string
        precision int32
        expected  string
    }{
        {"PI", 5, "3.14159"},
        {"E", 4, "2.7183"},
        {"1.23456789", 4, "1.2346"},
        {"sin(PI/6)", 4, "0.5000"},
    }

    for _, test := range tests {
        calc := calculator.NewCalculator(test.precision)
        parser := NewParser(test.input, calc)
        result := parser.Parse().Evaluate()
        if result != test.expected {
            t.Errorf("对于输入 %s (精度 %d): 期望 %s, 得到 %s", 
                test.input, test.precision, test.expected, result)
        }
    }
}

func TestComplexExpressions(t *testing.T) {
    tests := []struct {
        input    string
        expected string
    }{
        {"2 * PI + sin(PI/2) * cos(PI/3)", "7.062177826491071"},
        {"log(E^2, E) + ln(E^3)", "5"},
        {"sqrt(PI^2 + E^2)", "4.180754863754716"},
        {"sin(PI/4)^2 + cos(PI/4)^2", "1"},
        {"log(1000,10) + ln(E) + lg(100)", "6"},
    }

    calc := calculator.NewCalculator(12)

    for _, test := range tests {
        parser := NewParser(test.input, calc)
        result := parser.Parse().Evaluate()
        expected, _ := decimal.NewFromString(test.expected)
        actual, _ := decimal.NewFromString(result)
        if !expected.Sub(actual).Abs().LessThan(decimal.NewFromFloat(1e-10)) {
            t.Errorf("对于输入 %s: 期望 %s, 得到 %s", test.input, test.expected, result)
        }
    }
}
