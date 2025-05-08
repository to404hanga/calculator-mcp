package mcp

import (
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/to404hanga/calculator-mcp/ast"
	"github.com/to404hanga/calculator-mcp/calculator"
)

const toolDescriptionEN = `High-Precision Scientific Calculator
This is a scientific calculator supporting high-precision calculations with the following features:

1. Basic Operations
   - Addition(+), Subtraction(-), Multiplication(*), Division(/)
   - Supports nested parentheses, e.g., (1 + 2) * 3
   - Supports arbitrary precision decimal calculations

2. Mathematical Constants
   - PI (π): Mathematical constant pi
   - E (e): Base of natural logarithm

3. Mathematical Functions
   - sqrt(x): Square root calculation
   - pow(x, y): Exponentiation, e.g., 2 ^ 3
   - log(x,b): Logarithm with base b, e.g., log(8,2) = 3
   - ln(x): Natural logarithm (base e), e.g., ln(e) = 1
   - lg(x): Common logarithm (base 10), e.g., lg(100) = 2
   
4. Trigonometric Functions
   - sin(x): Sine function
   - cos(x): Cosine function
   - tan(x): Tangent function
   - asin(x): Arcsine function, input range [-1,1]
   - acos(x): Arccosine function, input range [-1,1]
   - atan(x): Arctangent function

5. Precision Control
   - Supports custom calculation precision
   - Default precision of 10 decimal places
   - Maximum precision up to 75 decimal places

Usage Examples:
1. Basic operation: 1 + 2 * 3
2. Constant operation: 2 * PI
3. Function calculation: sqrt(16)
4. Compound operation: (1 + sqrt(16)) * 2
5. Trigonometric function: sin(PI/2)
6. High precision: E ^ 2
7. Logarithm: log(1000,10) = 3
8. Natural logarithm: ln(E^2) = 2
9. Common logarithm: lg(1000) = 3

Important Notes:
1. Division by zero is not allowed
2. Square root of negative numbers is not allowed
3. Input values for inverse trigonometric functions must be within valid range
4. Logarithm input and base must be positive, base cannot be 1
5. Natural logarithm input must be positive`

var calcInputSchema = mcp.ToolInputSchema{
	Type: "object",
	Properties: map[string]any{
		"expression": map[string]any{
			"type":        "string",
			"description": "The expression to evaluate",
		},
		"precision": map[string]any{
			"type":        "number",
			"description": "The precision of the result",
		},
	},
	Required: []string{"expression"},
}

type CalcServer struct {
	server *server.MCPServer
}

func (s *CalcServer) runCalc(expression string, precision int32) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			result, err = "Internal Error", r.(error)
		}
	}()

	calc := calculator.NewCalculator(precision)
	parser := ast.NewParser(expression, calc)
	result = parser.Parse().Evaluate()
	return result, nil
}

func (s *CalcServer) handleToolCall(arguments map[string]any) (*mcp.CallToolResult, error) {
	log.Printf("handleToolCall called with arguments: %+v", arguments)

	expression, ok := arguments["expression"].(string)
	if !ok {
		return nil, fmt.Errorf("expression is required")
	}

	var precision int
	if exist, ok := arguments["precision"]; ok {
		if precision, ok = exist.(int); !ok {
			precision = 10
		}
	}

	result, err := s.runCalc(expression, int32(precision))
	if err != nil {
		log.Printf("Error running calc: %v", err)
		return nil, err
	}

	return &mcp.CallToolResult{
		Content: []any{
			map[string]any{
				"type": "text",
				"text": result,
			},
		},
	}, nil
}

func NewCalcServer() *server.MCPServer {
	calcServer := &CalcServer{}

	s := server.NewMCPServer(
		"calculator-mcp",
		"1.0.0",
		server.WithToolCapabilities(true),
		server.WithLogging(),
	)
	calcServer.server = s

	log.Printf("Adding calc tool...")
	tool := mcp.Tool{
		Name:        "calc",
		Description: toolDescriptionEN,
		InputSchema: calcInputSchema,
	}
	s.AddTool(tool, calcServer.handleToolCall)

	return s
}
