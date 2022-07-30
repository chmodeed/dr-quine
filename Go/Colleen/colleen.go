package main

import "fmt"

func outsideFunc() string {
	/*
	* Outside comment
	*/
	a := "package main\n\nimport \"fmt\"\n\nfunc outsideFunc() string {\n\t/*\n\t* Outside comment\n\t*/\n\ta := %q\n\treturn a\n}\n\nfunc main() {\n\ta := outsideFunc()\n\tfmt.Printf(a, a)\n}\n"
	return a
}

func main() {
	a := outsideFunc()
	fmt.Printf(a, a)
}
