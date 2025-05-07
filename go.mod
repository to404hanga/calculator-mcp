module github.com/to404hanga/calculator-mcp

go 1.23.4

replace github.com/to404hanga/calculator-mcp => ./calculator-mcp

require (
	github.com/mark3labs/mcp-go v0.5.1
	github.com/shopspring/decimal v1.4.0
)

require github.com/google/uuid v1.6.0 // indirect
