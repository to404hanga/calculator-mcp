# Calculator-MCP

## High-Precision Scientific Calculator

A high-precision scientific calculator implemented in Go, supporting complex mathematical operations with customizable precision.

### Features

1. Basic Operations
   - Addition(+), Subtraction(-), Multiplication(*), Division(/)
   - Nested parentheses support: (1 + 2) * 3
   - Arbitrary precision decimal calculations

2. Mathematical Constants
   - PI (π): Mathematical constant pi
   - E (e): Base of natural logarithm

3. Mathematical Functions
   - sqrt(x): Square root calculation
   - pow(x, y): Exponentiation, e.g., 2 ^ 3
   
4. Trigonometric Functions
   - sin(x): Sine function
   - cos(x): Cosine function
   - tan(x): Tangent function
   - asin(x): Arcsine function (input range [-1,1])
   - acos(x): Arccosine function (input range [-1,1])
   - atan(x): Arctangent function

5. Precision Control
   - Customizable calculation precision
   - Default: 10 decimal places
   - Maximum: 75 decimal places

### Usage Examples:
1. Basic operation: 1 + 2 * 3
2. Constant operation: 2 * PI
3. Function calculation: sqrt(16)
4. Compound operation: (1 + sqrt(16)) * 2
5. Trigonometric function: sin(PI/2)
6. High precision: E ^ 2

### Important Notes:
1. Division by zero is not allowed
2. Square root of negative numbers is not allowed
3. Input values for inverse trigonometric functions must be within valid range`
