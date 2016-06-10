package main

import (
	"flag"
	"fmt"
)

const (
	UsageString = `iumac is a tool to compile and package an application.
  
Usage:
    
  iumac build [flags]
    
The flags are:
 `
)

func Usage() {
	fmt.Println(UsageString)
	flag.PrintDefaults()
	fmt.Println()
}
