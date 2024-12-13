package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
    var letter string
    file,_:=os.Open("input2")
    defer file.Close()
    var chart []string
    //var chart = make([][]string,20)
    /*
    for i:=range chart{
        chart[i] = make([]string,20)
    }*/

    sc := bufio.NewScanner(file)
    i:=0
    for sc.Scan(){
        line := ""
        for _,v := range sc.Text(){
            fmt.Println(string(v))
            if string(v) == letter{
                 line += string(v)
            }else{
                if letter == ""{
                    letter = string(v)
                    line += "|"
                    line += string(v)
                    fmt.Println("haha we are here")
                }else{
                    letter = string(v)
                    line += "|"
                    line += string(v)
                }
            }

        }
        line += "|"
        chart = append(chart,line)
        line = strings.Repeat("-",len(line))
        chart = append(chart,line)
        line = ""
        i++
    }
    for _,v := range chart{
        fmt.Println(v)
    }
}
