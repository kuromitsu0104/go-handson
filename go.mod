module example.com/go-handson

go 1.16

require (
	local.packages/hello v0.0.0
)

replace local.packages/hello => ./hello