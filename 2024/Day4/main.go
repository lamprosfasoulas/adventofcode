package main

import (
	"bufio"
	"fmt"
	"os"
)
var sum int
func main(){
    var chars [140][140]string
    file,_ := os.Open("input.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)
    j:=0
    for scanner.Scan(){
        line := scanner.Text()
        for i,v:=range line{
            chars[j][i] = string(v)
        }
        j++
    }
    for i,v:=range chars{
        for j,k:=range v{
            if k == "A"{
                check(i,j,chars)
            }
        }
    }
    fmt.Println(sum)
}

func check(i int, j int, chars [140][140]string){
    amM:=2
    amS:=2
    if !(i+1 >= 140 || j+1 >= 140 || i-1 < 0 || j-1 < 0){
        p1 := &chars[i+1][j+1]
        p2 := &chars[i+1][j-1]
        p3 := &chars[i-1][j+1]
        p4 := &chars[i-1][j-1]
        variables := []string{*p1,*p2,*p3,*p4}
        fmt.Println(*p1,*p2,*p3,*p4)
        if *p1 == *p4 || *p2 == *p3{
            return
        }
        for _,v:=range variables{
            if v == "M"{
                amM--
            }
            if v == "S"{
                amS--
            }
        }
    }
    if amM == 0 && amS==0{
        sum ++
    }
}
