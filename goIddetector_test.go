package goIddetector_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/kyosu-1/goIddetector"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, goIddetector.Analyzer, "a")
}
