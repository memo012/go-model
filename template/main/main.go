package main

import "designMode/template"

func main() {
	localDac := template.NewLocalDoc()
	netDac   := template.NewNetDoc()
	localDac.DoOperator()
	netDac.DoOperator()
}
