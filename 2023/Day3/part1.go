package Day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var matrix [][]byte
var numbers []string 
func init(){
    matrix = make([][]byte,140)
    for i:= range matrix{
        matrix[i] = make([]byte, 140)
    }
}
func Test(){
    file,_ := os.Open("day3.txt")
    scanner := bufio.NewScanner(file)
    j := 0
    for scanner.Scan(){
        value := scanner.Bytes()
        //value2 := scanner.Text()
        //fmt.Println(value2)
        //fmt.Println(string(value))
        symChange(value)
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
        //numS(value)
        //sym(value)
    }
    numS(matrix)
    file.Close()
}
func symChange (arr []byte){
    for i := range arr{
        p := &arr[i]
        if (*p == 46){//here we define spaces
            //*p = 126 
            *p = 32
        }else if (*p >= 32 && *p < 48)||(*p > 57 && *p < 65){//here we define the symbols
            //*p = 35
            *p = 35
        }
    }
}
func numS(arr [][]byte){
    var flag bool
    var sum int64
    var sumsum int64
    for i,v := range arr{
        var charset string 
        for j,k:= range v{
            if (k >= 48 && k <= 57) { 
                //fmt.Println(k)
                if (condition(i,j)){
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
    }
    fmt.Println("\nSum of all is ",sumsum," result should be ",sum)
}
func condition(x int, y int)bool{
    for i:=-1;i<=1;i++{
        for j:=-1;j<=1;j++{
            if (x+i<0 || y+j<0)||(x+i>139|| y+j > 139){
                continue
            }else if matrix[x+i][y+j] == 35{
                return true
            }
        }
    }
    return false
}
