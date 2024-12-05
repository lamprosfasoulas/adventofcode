package Day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
func init(){
    matrix = make([][]byte,140)
    for i:= range matrix{
        matrix[i] = make([]byte, 140)
    }
    //numset = make([][]int,200)
    //for i:= range matrix{
    //    numset[i] = make([]int, 140)
    //}
}
func Test2(){
    file,_ := os.Open("day3.txt")
    scanner := bufio.NewScanner(file)
    j := 0
    for scanner.Scan(){
        value := scanner.Bytes()
        //value2 := scanner.Text()
        //fmt.Println(value2)
        //fmt.Println(string(value))
        symChange2(value)
        for i,v:= range value{
            matrix[j][i] = v
        }
        //numLine(value)
        //synLine(value)
       //if j == 2 {break}
        j++
        //for i,v:= range value{
        //    matrix[j][i] = v
        //}
        //fmt.Println("Num:", j,string(value),"Len: ",len(value))
        //numS2(value)
        //sym(value)
    }
    numS2(matrix)
    file.Close()
}
func symChange2 (arr []byte){
    for j := range arr{
        p := &arr[j]
        if (*p == 46){//here we define spaces
            //*p = 126 
            *p = 45
        }else if *p != 42 && (*p >= 32 && *p < 48)||(*p > 57 && *p < 65){//here we define the symbols
            //*p = 35
            *p = 45
        }
    }
}
func numS2(arr [][]byte){
    var flag bool
    var sum int64
    var sumsum int64
    for i,v := range arr{
        var charset string 
        for j,k:= range v{
            if (k >= 48 && k <= 57) { 
                //fmt.Println(k)
                if (condition2(i,j)){
                    flag = true
                }
                charset += string(k)
                if (j == 139 && flag == true){
                    casualmidvar,_:=strconv.Atoi(charset)
                    sum += int64(casualmidvar)
                    flag = false
                }
                //fmt.Println("@@",charset,"@@")
                //fmt.Printf("Number %c at index %d\n",k,j)
            }else{
                if charset != "" && flag == true{
                    casualmidvar,_:=strconv.Atoi(charset)
                    sum += int64(casualmidvar)
                    flag = false
                }
                charset = ""
            }
        }
        fmt.Print(string(v))
        fmt.Println(positions)
        positions = positions[:0]
    }
    mapsym()
    mapnum()
    fmt.Println("\nSum of all is ",sumsum," result should be ",sum)
}
func condition2(x int, y int)bool{
    for i:=-1;i<=1;i++{
        for j:=-1;j<=1;j++{
            if (x+i<0 || y+j<0)||(x+i>139|| y+j > 139){
                continue
            }else if matrix[x+i][y+j] == 42{
                return true
            }
        }
    }
    return false
}
var npositions []string
func mapnum(){
    for i,v:=range matrix{
        strind:=0
        endind:=0
        for j,k:=range v{
            if k >= 48 && k <= 57 && strind == 0{

            }
            if j == 139{
                npositions = append(npositions, number)
                number = ""
            }else{
                npositions = append(npositions, number)
                number = ""
            }
        }
        fmt.Println(npositions)
        npositions = npositions[:0]
    }

}
var positions []string
func mapsym(){
    for i,v:=range matrix{
        for j,k:=range v{
            if k == 42{
                positions = append(positions, fmt.Sprint(i)+":"+fmt.Sprint(j))
            }
        }
        fmt.Println(positions)
        positions = positions[:0]
    }
}

