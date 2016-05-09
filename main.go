package main

import "github.com/andrewarrow/spire-go/spire"
import "os"
import "fmt"

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Usage:")
		fmt.Println("./spire-go 2016-05-03")
		return
	}
	spire.GetDate(os.Args[1])
}
