package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main(){
    file,_:=os.Open("day13")
    defer file.Close()
    sc := bufio.NewScanner(file)
    var arr []int
    var sum int
    i:=0
    for sc.Scan(){
        if sc.Text() == ""{
            continue
        }
        line := strings.Split(sc.Text(), ":")[1]
        sym := regexp.MustCompile(`[,=+]`)
        parts := sym.Split(line, -1)
        number1,_ := strconv.Atoi(parts[1])
        number2,_ := strconv.Atoi(parts[3])
        if i == 2{
            number1 += 10000000000000 
            number2 += 10000000000000 
        }
        //fmt.Println(number1,number2)
        arr = append(arr,number1)
        arr = append(arr,number2)
        i++
        if len(arr) == 6{
            i=0
            a1,a2,b1,b2,c1,c2 := arr[0],arr[1],arr[2],arr[3],arr[4],arr[5]
            a,b:=deter(a1,a2,b1,b2,c1,c2)
            sum += (a*3)+b
            //fmt.Println((a*3)+b)
            arr = nil
        }
    }
    fmt.Println("The final sum is: ",sum)
}
func deter(a1,a2,b1,b2,c1,c2 int) (int,int){
    d := (a1*b2 - a2*b1)
    da := (c1*b2 - c2*b1)
    db := (a1*c2 - a2*c1)
    if da%d==0 &&  db%d == 0{
        return da/d,db/d
    }
    return 0,0
}
