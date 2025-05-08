package ast

import (
	"strings"

	"github.com/to404hanga/calculator-mcp/calculator"
)

// NodeType 定义节点类型
type NodeType int

const (
	NumberNode NodeType = iota
	BinaryOpNode
	UnaryOpNode
	PINode
	SqrtNode
	PowNode
	SinNode
	CosNode
	TanNode
	AsinNode
	AcosNode
	AtanNode
	ENode    // 自然对数e常量
	LogNode  // 对数运算
)

// Node 接口定义了所有 AST 节点必须实现的方法
type Node interface {
	Evaluate() string
	Type() NodeType
}

// NumberNode 表示数字常量
type NumberLiteral struct {
	Value string
}

func (n *NumberLiteral) Evaluate() string {
	return n.Value
}

func (n *NumberLiteral) Type() NodeType {
	return NumberNode
}

// BinaryOperator 表示二元运算符节点
type BinaryOperator struct {
	Left     Node
	Right    Node
	Operator string
	calc     *calculator.Calculator
}

func (b *BinaryOperator) Evaluate() string {
	left := b.Left.Evaluate()
	right := b.Right.Evaluate()

	switch b.Operator {
	case "+":
		return b.calc.Add(left, right)
	case "-":
		return b.calc.Subtract(left, right)
	case "*":
		return b.calc.Multiply(left, right)
	case "/":
		return b.calc.Divide(left, right)
	}
	panic("未知的运算符: " + b.Operator)
}

func (b *BinaryOperator) Type() NodeType {
	return BinaryOpNode
}

// PIConstant 表示π常量
type PIConstant struct {
	calc *calculator.Calculator
}

func (p *PIConstant) Evaluate() string {
	return p.calc.PI()
}

func (p *PIConstant) Type() NodeType {
	return PINode
}

// SqrtOperation 表示开方操作
type SqrtOperation struct {
	Operand Node
	calc    *calculator.Calculator
}

func (s *SqrtOperation) Evaluate() string {
	return s.calc.Sqrt(s.Operand.Evaluate())
}

func (s *SqrtOperation) Type() NodeType {
	return SqrtNode
}

// PowOperation 表示乘方操作
type PowOperation struct {
	Base     Node
	Exponent Node
	calc     *calculator.Calculator
}

func (p *PowOperation) Evaluate() string {
	return p.calc.Power(p.Base.Evaluate(), p.Exponent.Evaluate())
}

func (p *PowOperation) Type() NodeType {
	return PowNode
}

// SinOperation 表示正弦操作
type SinOperation struct {
	Operand Node
	calc    *calculator.Calculator
}

func (s *SinOperation) Evaluate() string {
	return s.calc.Sin(s.Operand.Evaluate())
}

func (s *SinOperation) Type() NodeType {
	return SinNode
}

// CosOperation 表示余弦操作
type CosOperation struct {
	Operand Node
	calc    *calculator.Calculator
}

func (c *CosOperation) Evaluate() string {
	return c.calc.Cos(c.Operand.Evaluate())
}

func (c *CosOperation) Type() NodeType {
	return CosNode
}

// Parse 解析整个表达式
func (p *Parser) Parse() Node {
	return p.parseExpression()
}

// parseExpression 解析表达式
func (p *Parser) parseExpression() Node {
	left := p.parseTerm()

	for p.pos < len(p.tokens) {
		if p.tokens[p.pos] == "+" || p.tokens[p.pos] == "-" {
			operator := p.tokens[p.pos]
			p.pos++
			right := p.parseTerm()
			left = &BinaryOperator{Left: left, Right: right, Operator: operator, calc: p.calc}
		} else {
			break
		}
	}

	return left
}

// parseTerm 解析项
func (p *Parser) parseTerm() Node {
	left := p.parseFactor()

	for p.pos < len(p.tokens) {
		if p.tokens[p.pos] == "*" || p.tokens[p.pos] == "/" || p.tokens[p.pos] == "^" {
			operator := p.tokens[p.pos]
			p.pos++
			right := p.parseFactor()
			if operator == "^" {
				left = &PowOperation{Base: left, Exponent: right, calc: p.calc}
			} else {
				left = &BinaryOperator{Left: left, Right: right, Operator: operator, calc: p.calc}
			}
		} else {
			break
		}
	}

	return left
}

// TanOperation 表示正切操作
type TanOperation struct {
	Operand Node
	calc    *calculator.Calculator
}

func (t *TanOperation) Evaluate() string {
	return t.calc.Tan(t.Operand.Evaluate())
}

func (t *TanOperation) Type() NodeType {
	return TanNode
}

// AsinOperation 表示反正弦操作
type AsinOperation struct {
	Operand Node
	calc    *calculator.Calculator
}

func (a *AsinOperation) Evaluate() string {
	return a.calc.Asin(a.Operand.Evaluate())
}

func (a *AsinOperation) Type() NodeType {
	return AsinNode
}

// AcosOperation 表示反余弦操作
type AcosOperation struct {
	Operand Node
	calc    *calculator.Calculator
}

func (a *AcosOperation) Evaluate() string {
	return a.calc.Acos(a.Operand.Evaluate())
}

func (a *AcosOperation) Type() NodeType {
	return AcosNode
}

// AtanOperation 表示反正切操作
type AtanOperation struct {
	Operand Node
	calc    *calculator.Calculator
}

func (a *AtanOperation) Evaluate() string {
	return a.calc.Atan(a.Operand.Evaluate())
}

func (a *AtanOperation) Type() NodeType {
	return AtanNode
}

// EConstant 表示自然对数e常量
type EConstant struct {
	calc *calculator.Calculator
}

func (e *EConstant) Evaluate() string {
	return e.calc.E()
}

func (e *EConstant) Type() NodeType {
	return ENode
}

// LogOperation 表示自然对数操作
type LogOperation struct {
	Value Node
	Base  Node
	calc  *calculator.Calculator
}

func (l *LogOperation) Evaluate() string {
	return l.calc.Log(l.Value.Evaluate(), l.Base.Evaluate())
}

func (l *LogOperation) Type() NodeType {
	return LogNode
}

// parseFactor 解析因子
func (p *Parser) parseFactor() Node {
	if p.pos >= len(p.tokens) {
		panic("表达式不完整")
	}

	token := p.tokens[p.pos]

	// 检查单个运算符和前缀运算符的情况
	if token == "+" || token == "-" || token == "*" || token == "/" || token == "^" {
		if p.pos == 0 || p.pos == len(p.tokens)-1 {
			panic("无效的表达式")
		}
		// 如果运算符在开头，也是无效的表达式
		if p.pos == 0 {
			panic("无效的表达式")
		}
	}

	p.pos++

	switch {
	case token == "(":
		node := p.parseExpression()
		if p.pos >= len(p.tokens) || p.tokens[p.pos] != ")" {
			panic("缺少右括号")
		}
		p.pos++
		return node

	case token == "PI":
		return &PIConstant{calc: p.calc}

	case token == "E": // 新增
		return &EConstant{calc: p.calc}

	case token == "sqrt":
		if p.pos >= len(p.tokens) || p.tokens[p.pos] != "(" {
			panic("sqrt后需要括号")
		}
		p.pos++
		operand := p.parseExpression()
		if p.pos >= len(p.tokens) || p.tokens[p.pos] != ")" {
			panic("sqrt缺少右括号")
		}
		p.pos++
		return &SqrtOperation{Operand: operand, calc: p.calc}

	case token == "sin":
		if p.pos >= len(p.tokens) || p.tokens[p.pos] != "(" {
			panic("sin后需要括号")
		}
		p.pos++
		operand := p.parseExpression()
		if p.pos >= len(p.tokens) || p.tokens[p.pos] != ")" {
			panic("sin缺少右括号")
		}
		p.pos++
		return &SinOperation{Operand: operand, calc: p.calc}

	case token == "cos":
		if p.pos >= len(p.tokens) || p.tokens[p.pos] != "(" {
			panic("cos后需要括号")
		}
		p.pos++
		operand := p.parseExpression()
		if p.pos >= len(p.tokens) || p.tokens[p.pos] != ")" {
			panic("cos缺少右括号")
		}
		p.pos++
		return &CosOperation{Operand: operand, calc: p.calc}

	case token == "tan":
		if p.pos >= len(p.tokens) || p.tokens[p.pos] != "(" {
			panic("tan后需要括号")
		}
		p.pos++
		operand := p.parseExpression()
		if p.pos >= len(p.tokens) || p.tokens[p.pos] != ")" {
			panic("tan缺少右括号")
		}
		p.pos++
		return &TanOperation{Operand: operand, calc: p.calc}

	case token == "asin":
		if p.pos >= len(p.tokens) || p.tokens[p.pos] != "(" {
			panic("asin后需要括号")
		}
		p.pos++
		operand := p.parseExpression()
		if p.pos >= len(p.tokens) || p.tokens[p.pos] != ")" {
			panic("asin缺少右括号")
		}
		p.pos++
		return &AsinOperation{Operand: operand, calc: p.calc}

	case token == "acos":
		if p.pos >= len(p.tokens) || p.tokens[p.pos] != "(" {
			panic("acos后需要括号")
		}
		p.pos++
		operand := p.parseExpression()
		if p.pos >= len(p.tokens) || p.tokens[p.pos] != ")" {
			panic("acos缺少右括号")
		}
		p.pos++
		return &AcosOperation{Operand: operand, calc: p.calc}

	case token == "atan":
		if p.pos >= len(p.tokens) || p.tokens[p.pos] != "(" {
			panic("atan后需要括号")
		}
		p.pos++
		operand := p.parseExpression()
		if p.pos >= len(p.tokens) || p.tokens[p.pos] != ")" {
			panic("atan缺少右括号")
		}
		p.pos++
		return &AtanOperation{Operand: operand, calc: p.calc}

	case token == "log":
		if p.pos >= len(p.tokens) || p.tokens[p.pos] != "(" {
			panic("log后需要括号")
		}
		p.pos++
		value := p.parseExpression()
		if p.pos >= len(p.tokens) || p.tokens[p.pos] != "," {
			panic("log函数需要两个参数，用逗号分隔")
		}
		p.pos++
		base := p.parseExpression()
		if p.pos >= len(p.tokens) || p.tokens[p.pos] != ")" {
			panic("log缺少右括号")
		}
		p.pos++
		return &LogOperation{Value: value, Base: base, calc: p.calc}

	case token == "lg":
		if p.pos >= len(p.tokens) || p.tokens[p.pos] != "(" {
			panic("lg后需要括号")
		}
		p.pos++
		operand := p.parseExpression()
		if p.pos >= len(p.tokens) || p.tokens[p.pos] != ")" {
			panic("lg缺少右括号")
		}
		p.pos++
		// 使用 LogOperation，将 10 作为底数
		return &LogOperation{
			Value: operand,
			Base:  &NumberLiteral{Value: "10"},
			calc:  p.calc,
		}

	case token == "ln":
		if p.pos >= len(p.tokens) || p.tokens[p.pos] != "(" {
			panic("ln后需要括号")
		}
		p.pos++
		operand := p.parseExpression()
		if p.pos >= len(p.tokens) || p.tokens[p.pos] != ")" {
			panic("ln缺少右括号")
		}
		p.pos++
		// 使用 LogOperation，将 E 作为底数
		return &LogOperation{
			Value: operand,
			Base:  &EConstant{calc: p.calc},
			calc:  p.calc,
		}

	default:
		// 直接将数字作为字符串存储
		return &NumberLiteral{Value: token}
	}
}

// Parser 结构体用于解析表达式
type Parser struct {
	tokens []string
	pos    int
	calc   *calculator.Calculator
}

// NewParser 创建新的解析器
func NewParser(expression string, calc *calculator.Calculator) *Parser {
	// 将表达式转换为标记序列
	expression = strings.ReplaceAll(expression, "(", " ( ")
	expression = strings.ReplaceAll(expression, ")", " ) ")
	expression = strings.ReplaceAll(expression, ",", " , ") // 添加对逗号的处理
	expression = strings.ReplaceAll(expression, "+", " + ")
	expression = strings.ReplaceAll(expression, "-", " - ")
	expression = strings.ReplaceAll(expression, "*", " * ")
	expression = strings.ReplaceAll(expression, "/", " / ")
	expression = strings.ReplaceAll(expression, "^", " ^ ")
	tokens := strings.Fields(expression)
	return &Parser{
		tokens: tokens,
		pos:    0,
		calc:   calc,
	}
}
