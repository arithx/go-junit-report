package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/arithx/go-junit-report/parser"
)

var (
	noXMLHeader   bool
	packageName   string
	goVersionFlag string
	setExitCode   bool
)

func init() {
	flag.BoolVar(&noXMLHeader, "no-xml-header", false, "do not print xml header")
	flag.StringVar(&packageName, "package-name", "", "specify a package name (compiled test have no package name in output)")
	flag.StringVar(&goVersionFlag, "go-version", "", "specify the value to use for the go.version property in the generated XML")
	flag.BoolVar(&setExitCode, "set-exit-code", false, "set exit code to 1 if tests failed")
}

func main() {
	flag.Parse()

	// Read input
	report, err := parser.Parse(os.Stdin, packageName)
	if err != nil {
		fmt.Printf("Error reading input: %s\n", err)
		os.Exit(1)
	}

	// Write JSON
	out, err := json.Marshal(report)
	if err != nil {
		fmt.Printf("Erorr writing JSON: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(string(out))

	// Write xml
	/*
	err = JUnitReportXML(report, noXMLHeader, goVersionFlag, os.Stdout)
	if err != nil {
		fmt.Printf("Error writing XML: %s\n", err)
		os.Exit(1)
	}*/

	if setExitCode && report.Failures() > 0 {
		os.Exit(1)
	}
}
