/**
    @author: PDK
    @date: 2022/9/13
**/
package main

import (
	"github.com/pengdengke/tools_box/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
