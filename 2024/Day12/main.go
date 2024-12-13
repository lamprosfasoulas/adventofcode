package main

import (
	"fmt"
	"os"
)

func main(){
    file,_:=os.Open("input")
    defer file.Close()
    fmt.Println(string(file))
}
