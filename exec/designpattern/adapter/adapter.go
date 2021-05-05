package main

import (
	dp "github.com/Reljod/Reljod-Portfolio-Backend/app/designpattern"
)

func main() {
	var processor dp.IProcess = dp.Adapter{}
	processor.Process()
}
