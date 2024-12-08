package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    const length = 50
    const filename = "input.txt"
    file,_ := os.Open(filename)
    defer file.Close()
    sum := 0
    array := make([][]string,length)
    for i:=range array{
        array[i] = make([]string,length)
    }
    var list []string
    scanner := bufio.NewScanner(file)
    j := 0
    for scanner.Scan(){
        line := scanner.Text()
        for i,v:=range line{
            array[j][i] = string(v)
            if string(v)!="."{
                str := string(v) + fmt.Sprintf(":%d:%d",j,i)
                list = append(list,str)
            }
        }
        j++
    }
    //fmt.Println(len(list))
    check(&list,&sum,&array)
    //fmt.Println(list)
    for _,v:=range array{
        fmt.Println(v)
    }
    fmt.Println(sum)
}
func check(arr *[]string,sum *int,horse *[][]string){
    var found []string
    for _,v:=range *arr{
        fmt.Println("Checking ",v)
        sym := strings.Split(string(v), ":")[0]
        x,_ := strconv.Atoi(strings.Split(v, ":")[1])
        y,_ := strconv.Atoi(strings.Split(v, ":")[2])
        fmt.Println("Value of X: ",x," Value of Y: ",y)
        for i,k:=range *horse{
            for j,v:=range k{
                diffv := x-i
                diffh := y-j
                if diffv == 0 && diffh == 0{continue}
                if sym == string(v)&& x+diffv>=0 && x+diffv<len(*horse) && y+diffh>=0 && y+diffh<len(*horse) {
                    fmt.Println("Found antinode #"," at ",x+diffv,y+diffh," with ",diffv," ",diffh)
                    if !isThere(&found,x+diffv,y+diffh){
                        found = append(found,fmt.Sprintf("%d:%d",x+diffv,y+diffh))
                        *sum ++
                    }
                }
            }
        }
    }
    fmt.Println(*arr)
    fmt.Println("List found ",found)
    fmt.Println("Length of found: ",len(found))
    //fmt.Println("The final result is: ",*sum-errors)
}
func isThere (list *[]string, x int, y int) bool{
    for _,v:=range *list{
        if v == fmt.Sprintf("%d:%d",x,y){
            return true
        }
    }
    return false
}
