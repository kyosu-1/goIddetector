package main

import (
	"github.com/kyosu-1/goIddetector"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(goIddetector.Analyzer) }
