package calculator

import (
	"math"

	"github.com/shopspring/decimal"
)

// Calculator 提供基本的数学计算功能
type Calculator struct {
	precision int32 // 计算精度
}

// NewCalculator 创建一个新的计算器实例，指定计算精度
func NewCalculator(precision int32) *Calculator {
	if precision < 0 {
		precision = 10 // 默认精度为10位小数
	}
	return &Calculator{precision: precision}
}

// Add 执行加法运算
func (c *Calculator) Add(left, right string) string {
	l, _ := decimal.NewFromString(left)
	r, _ := decimal.NewFromString(right)
	return l.Add(r).Round(c.precision).String()
}

// Subtract 执行减法运算
func (c *Calculator) Subtract(left, right string) string {
	l, _ := decimal.NewFromString(left)
	r, _ := decimal.NewFromString(right)
	return l.Sub(r).Round(c.precision).String()
}

// Multiply 执行乘法运算
func (c *Calculator) Multiply(left, right string) string {
	l, _ := decimal.NewFromString(left)
	r, _ := decimal.NewFromString(right)
	return l.Mul(r).Round(c.precision).String()
}

// Divide 执行除法运算
func (c *Calculator) Divide(left, right string) string {
	l, _ := decimal.NewFromString(left)
	r, _ := decimal.NewFromString(right)
	if r.IsZero() {
		panic("除数不能为零")
	}
	decimal.DivisionPrecision = int(c.precision)
	return l.Div(r).Round(c.precision).String()
}

// Power 执行乘方运算
func (c *Calculator) Power(base, exponent string) string {
	b, _ := decimal.NewFromString(base)
	e, _ := decimal.NewFromString(exponent)
	res, _ := b.PowWithPrecision(e, c.precision)
	return res.Round(c.precision).String()
}

// Sqrt 执行开方运算
func (c *Calculator) Sqrt(value string) string {
	v, _ := decimal.NewFromString(value)
	if v.IsNegative() {
		panic("不能对负数进行开方")
	}
	if v.IsZero() {
		return "0"
	}

	// 使用牛顿迭代法计算平方根
	z := v.Div(decimal.NewFromInt(2))
	decimal.DivisionPrecision = int(c.precision)

	// 迭代计算，直到达到指定精度
	tolerance := decimal.NewFromFloat(math.Pow(10, -float64(c.precision)))
	for i := 0; i < int(c.precision)*2; i++ {
		prev := z
		z = z.Add(v.Div(z)).Div(decimal.NewFromInt(2))

		if z.Sub(prev).Abs().LessThan(tolerance) {
			break
		}
	}

	return z.Round(c.precision).String()
}

// Sin 执行正弦运算
func (c *Calculator) Sin(value string) string {
	v, _ := decimal.NewFromString(value)
	floatVal := v.InexactFloat64()
	sinVal := math.Sin(floatVal)
	return decimal.NewFromFloat(sinVal).Round(c.precision).String()
}

// Cos 执行余弦运算
func (c *Calculator) Cos(value string) string {
	v, _ := decimal.NewFromString(value)
	floatVal := v.InexactFloat64()
	cosVal := math.Cos(floatVal)
	return decimal.NewFromFloat(cosVal).Round(c.precision).String()
}

// Tan 执行正切运算
func (c *Calculator) Tan(value string) string {
    v, _ := decimal.NewFromString(value)
    // 由于 decimal 包不直接支持三角函数，我们需要先转换为 float64
    floatVal := v.InexactFloat64()
    tanVal := math.Tan(floatVal)
    return decimal.NewFromFloat(tanVal).Round(c.precision).String()
}

// Asin 执行反正弦运算
func (c *Calculator) Asin(value string) string {
    v, _ := decimal.NewFromString(value)
    floatVal := v.InexactFloat64()
    if floatVal < -1 || floatVal > 1 {
        panic("反正弦函数的输入必须在 [-1,1] 范围内")
    }
    asinVal := math.Asin(floatVal)
    return decimal.NewFromFloat(asinVal).Round(c.precision).String()
}

// Acos 执行反余弦运算
func (c *Calculator) Acos(value string) string {
    v, _ := decimal.NewFromString(value)
    floatVal := v.InexactFloat64()
    if floatVal < -1 || floatVal > 1 {
        panic("反余弦函数的输入必须在 [-1,1] 范围内")
    }
    acosVal := math.Acos(floatVal)
    return decimal.NewFromFloat(acosVal).Round(c.precision).String()
}

// Atan 执行反正切运算
func (c *Calculator) Atan(value string) string {
    v, _ := decimal.NewFromString(value)
    floatVal := v.InexactFloat64()
    atanVal := math.Atan(floatVal)
    return decimal.NewFromFloat(atanVal).Round(c.precision).String()
}

// PI 返回π常量
func (c *Calculator) PI() string {
	pi, err := decimal.NewFromString("3.141592653589793238462643383279502884197169399375105820974944592307816406286")
	if err != nil {
		return "3.141592653589793"
	}
	return pi.Round(c.precision).String()
}
