package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    file,_:= os.Open("input.txt")
    defer file.Close()
    var final int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line:= scanner.Text()
        sum,_:=strconv.Atoi(strings.Split(line,":")[0])//this is the sum value
        var nums []int//these will be the numbers in question
        numsAsString:=strings.Fields(strings.Split(line,":")[1])//this is numbs as strings
        for _,v:=range numsAsString{
            num,_:=strconv.Atoi(v)
            nums=append(nums,num)
        }
        final += isOkay(sum,nums)
    }
    fmt.Println("Final result:",final)
}

func isOkay(sum int, nums []int)int{
    var actions []string
    symbols:=[]string{"+","*","c"}
    posComb("",len(nums)-1,&actions,symbols)
    for _,v:=range actions{
        partSum:= nums[0]
        for i,k:=range v{
            if string(k)=="+"{
                partSum+=nums[i+1]
            }else if string(k)=="*"{
                partSum*=nums[i+1]
            }else if string(k)=="c"{
                s1:=strconv.Itoa(partSum)
                s2:=strconv.Itoa(nums[i+1])
                partSum,_=strconv.Atoi(s1+s2)
            }
        }
        if partSum==sum{
            return sum
        }
    }
    return 0
}
func posComb(current string,length int,result *[]string,symbols []string){
    if len(current)==length{
        *result=append(*result,current)
        return
    }
    for _,v:=range symbols{
        posComb(current+v,length,result,symbols)
    }

}
