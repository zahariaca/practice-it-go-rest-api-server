module example.com/tester

go 1.18

replace example.com/practice => ../practice

require (
	example.com/library v0.0.0-00010101000000-000000000000
	example.com/practice v0.0.0-00010101000000-000000000000
)

replace example.com/library => ../library
