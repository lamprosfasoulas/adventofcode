package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
    file, _ := os.Open("input")
    defer file.Close()
    sc := bufio.NewScanner(file)
    read:=false
    i := 0
    for sc.Scan(){
        //fmt.Println(sc.Text())
        for j,k := range sc.Text(){
            if k == '@'{
                fmt.Printf("@ is at i: %d j: %d",i-1,j-1)
            }
        }
        if sc.Text() == ""{
            read = !read
        }
        if read{
            for _,k := range sc.Text(){
                if string(k) == "^"{
                    fmt.Println("")
                }
                if string(k) == "v"{
                    fmt.Println("")
                }
                fmt.Print(string(k))
            }
        }
        i++
    }
}
