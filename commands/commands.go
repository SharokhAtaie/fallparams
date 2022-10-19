package commands

import (
	"bytes"
	"fmt"
	"github.com/SharokhAtaie/fallparams/requests"
	"os/exec"
	"strings"
)

func Regex(url string) string {
	data := requests.MakeRequest(url)
	cut := SingleAndDouble(data)
	let := LetVariable(data)
	json := JsonVar(data)
	jsonVar := JsonVariable(data)
	final := cut + let + json + jsonVar
	sort := Sort(final)
	return sort
}

func SingleAndDouble(str string) string {
	grep := exec.Command("grep", "-Eo", `("|')(\w+)("|')`)
	grep.Stdin = strings.NewReader(str)

	var grepout bytes.Buffer
	grep.Stdout = &grepout

	err := grep.Run()
	if err != nil {
		fmt.Println("grep not working")
	}

	tr := exec.Command("tr", "-d", `'"`)
	tr.Stdin = strings.NewReader(grepout.String())

	var trout bytes.Buffer
	tr.Stdout = &trout

	err1 := tr.Run()
	if err1 != nil {
		fmt.Println("tr not working")
	}
	return trout.String()
}

func LetVariable(str string) string {
	grep := exec.Command("grep", "-Eo", `(var|let|const) (\w+)`)
	grep.Stdin = strings.NewReader(str)

	var grepout bytes.Buffer
	grep.Stdout = &grepout

	err := grep.Run()
	if err != nil {
		fmt.Println("grep let function not working")
	}

	awk := exec.Command("awk", `{print $2 }`)
	awk.Stdin = strings.NewReader(grepout.String())

	var awkout bytes.Buffer
	awk.Stdout = &awkout

	err1 := awk.Run()
	if err1 != nil {
		fmt.Println("awk not working")
	}
	return awkout.String()
}

func JsonVar(str string) string {
	grep := exec.Command("grep", "-Eo", `(\w+):`)
	grep.Stdin = strings.NewReader(str)

	var grepout bytes.Buffer
	grep.Stdout = &grepout

	err := grep.Run()
	if err != nil {
		fmt.Println("jsonvar grep not working")
	}

	tr := exec.Command("tr", "-d", ":")
	tr.Stdin = strings.NewReader(grepout.String())

	var trout bytes.Buffer
	tr.Stdout = &trout

	err1 := tr.Run()
	if err1 != nil {
		fmt.Println("jsonvar tr not working")
	}
	return trout.String()
}

func JsonVariable(str string) string {
	grep := exec.Command("grep", "-Eo", `(\w+) =`)
	grep.Stdin = strings.NewReader(str)

	var grepout bytes.Buffer
	grep.Stdout = &grepout

	err := grep.Run()
	if err != nil {
		fmt.Println("jsonvariable grep not working")
	}

	tr := exec.Command("tr", "-d", "=")
	tr.Stdin = strings.NewReader(grepout.String())

	var trout bytes.Buffer
	tr.Stdout = &trout

	err1 := tr.Run()
	if err1 != nil {
		fmt.Println("jsonvariable tr not working")
	}
	return trout.String()
}

func Sort(str string) string {
	sort := exec.Command("sort", "-u")
	sort.Stdin = strings.NewReader(str)

	var sortout bytes.Buffer
	sort.Stdout = &sortout

	err := sort.Run()
	if err != nil {
		fmt.Println("sort not working")
	}
	return sortout.String()
}
