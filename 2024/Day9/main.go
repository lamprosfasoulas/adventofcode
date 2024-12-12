package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var FILENAME string = os.Args[1]
var real int = 6460170593016
func main(){
    file,_ := os.Open(FILENAME)
    defer file.Close()
    scanner := bufio.NewScanner(file)
    files := make([]string,0)
    //line := []int{}
    for scanner.Scan(){
        file := true
        fileID := 0
        for _,b := range scanner.Text(){
            value,_ := strconv.Atoi(string(b))
            symbol := strconv.Itoa(fileID)
            if file {
                printLine(symbol,value,&files)
                fileID++
                file = false
            }else{
                printLine(".",value,&files)
                file = true
            }
        }
    }
    //order(&files)
    //fmt.Println("Step 1 ###")
    order2(&files)
    //fmt.Println(files)
    mul(&files)
    //right answer should be 6460170593016

}

func  printLine(symbol string,mul int, str *[]string){
    if mul == 0{return}
    for range mul{
        *str = append(*str,symbol)
    }
}
func order(str *[]string){
    //part 1
    //newstring := []string{}
    //for j,v:=range *str{
    //    char := string(v)
    //    if char == "."{
    //        for i:=len(*str)-1;i>0;i--{
    //            if string((*str)[i]) != "."{
    //                newstring =append(newstring, string((*str)[i]))
    //                (*str)[i] = "."
    //                break
    //            }
    //        }
    //    }else{
    //        newstring = append(newstring,char)
    //        (*str)[j] = "."
    //    }
    //}
    //mul(&newstring)
}
func order2(str *[]string){
    //fmt.Println("Step 2 ###")
    total := len(*str)
    barLength := 80
    tempSpace := 0
    temp := []string{}
    for i:=len((*str))-1;i>0;i--{
        progress := float64(total-i)/float64(total)
        bar := strings.Repeat("=",int(progress*float64(barLength)))
        fmt.Printf("\rProgress %v : %s",int(progress*100),bar)
        //fmt.Println("Index is ",i," and value is ",string((*str)[i]))
        length := len(temp)
        //fmt.Println("Length is ",length)
        if string((*str)[i]) != "."{
            if length == 0{
                length = len(temp)
                temp = append(temp,string((*str)[i]))
                if string((*str)[i]) != string((*str)[i-1]){
                    trial(&i,&tempSpace,&length,&temp,str)
                }
            }
            if length > 0 && string((*str)[i]) == string((*str)[i-1]){temp = append(temp,string((*str)[i]))}
            if length>0 && string((*str)[i]) != string((*str)[i-1]) && string((*str)[i]) == temp[0]{
                temp = append(temp,string((*str)[i]))
                trial(&i,&tempSpace,&length,&temp,str)

            }
        }
        //if string((*str)[i][0]) != "."{
        //    chars :=len((*str)[i])
        //    for j:=range spaces{
        //        if spaces[j] == 0{continue}
        //        diff := spaces[j] - chars
        //        if diff >= 0{
        //            final[j] += string((*str)[i])
        //            spaces[j]-=chars
        //            test := ""
        //            for range chars{
        //                test+=string(".")
        //            }
        //            (*str)[i] = test
        //            break
        //        }
        //    }

        //}
    }
    //fmt.Printf("\nStep 3 ###")
    //part 2
}
func trial(i,tempSpace,length *int,temp *[]string,str *[]string){
    //fmt.Printf("Adding to temp number %s at index %d\n",string((*str)[*i]),*i)
    //fmt.Println("temp is ",*temp)
    for j:=0;j<*i;j++{
        v := string((*str)[j])
        if *tempSpace>= len(*temp)&& string(v) != "."{
            if *tempSpace > *length{
                for p:=0;p<=len(*temp)-1;p++{
                    //if ((j-*tempSpace)+p <0){
                    //    panic(fmt.Sprintf("We have tempSpace %d and temp %v, with length %d , and index j:%d, that has entry %v",*tempSpace, *temp, *length,j,string((*str)[j])))
                    //}
                    if (j-*tempSpace)+p>=0{
                        (*str)[*i+p] = "."
                        (*str)[(j-*tempSpace)+p] = (*temp)[p]
                        //fmt.Println("We did")
                    }else{
                        (*str)[*i+p] = "."
                        (*str)[(j)+p] = (*temp)[p]
                        //fmt.Println("We did")
                    }
                }
                //fmt.Println(*str)
                *tempSpace = 0
                *temp = []string{}
                //fmt.Println("We replaced the str and now temp is",*temp, "and tempSpace is ",*tempSpace)
                break
            }
        }
        if string(v) != "."{
            *tempSpace = 0
            //fmt.Println("Reseting tempSpace ... ")
            //fmt.Println("TempSpace is ",*tempSpace," index ",j)
        }
        if string(v) == "."{
            *tempSpace ++
            //fmt.Println("Found a space ... ")
            //fmt.Println("TempSpace is ",*tempSpace," index ",j)
        }
        if j == *i-1{
            *temp = []string{}
            //fmt.Println("Reseting temp... ",*temp)
        }
    }
    //fmt.Println(*str,*i,len(*str),*tempSpace,*temp)
}
func mul(s *[]string){
    sum := 0
    for i,v:=range *s{
        num,_ := strconv.Atoi(string(v))
        sum += (i*num)
    }
    fmt.Printf("\nThe real sum :%d The final sum is %d The diff %d\n",real,sum,real-sum)
}
