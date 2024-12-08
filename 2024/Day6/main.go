package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    length:=130
    var sum int
    //open input file
    file,_:=os.Open("input.txt")
    defer file.Close()
    //create 2D array
    array:=make([][]string,length)
    for i:=range array{
        array[i]=make([]string,length)
    }
    //read input file
    scanner := bufio.NewScanner(file)
    i:=0
    var posx,posy int
    for scanner.Scan() {
        line := scanner.Text()
        for j,v := range line {
            if string(v)=="^" {
                posx=i
                posy=j
            }
            array[i][j]=string(v)
        }
        i++
    }
    final2:=0
    for i:=range array{
        for j:=range array[i]{
            if mov(posx,posy,&array,i,j) {
                final2++
            }
        }
    }
    fmt.Println("\nThe final output is ",sum+1)//output part 1 (need to add 1 to account for the starting position)
    fmt.Println("\nThe final output 2 is ",final2)//output part 2
    //this should be 1915. I will solve it tomorrow
}


func mov(posx,posy int, array *[][]string,x,y int) bool {
    moment:=(*array)[x][y]
    if moment=="#"{return false}
    if moment=="^"{return false}
    (*array)[x][y]="#"
    //for _,v:=range *array{
    //    fmt.Println(v)
    //}
    var nextx,nexty int
    var dirx,diry int = -1,0
    var turns int
        for true{
            nextx=posx+dirx
            nexty=posy+diry
            if nextx == -1 || nextx == 130 || nexty == -1 || nexty == 130 {break}

            //if dirx==-1 && diry==0 && array[posx][posy+1]=="X"{
            //    part2++
            //}else if dirx==0 && diry==1 && array[posx+1][posy]=="X"{
            //    part2++
            //}else if dirx==1 && diry==0 && array[posx][posy-1]=="X"{
            //    part2++
            //}else if dirx==0 && diry==-1 && array[posx-1][posy]=="X"{
            //    part2++
            //}
            if (*array)[nextx][nexty]=="#" {
                turns++
                if dirx==-1 && diry==0 {
                    dirx=0
                    diry=1
                }else if dirx==0 && diry==1 {
                    dirx=1
                    diry=0
                }else if dirx==1 && diry==0 {
                    dirx=0
                    diry=-1
                }else if dirx==0 && diry==-1 {
                    dirx=-1
                    diry=0
                }
                nextx=posx+dirx
                nexty=posy+diry
            }
            posx = nextx
            posy = nexty
            //if (*array)[posx][posy]=="." {
            //    (*array)[posx][posy]="X" 
            //}
        if turns == 40000 {(*array)[x][y]=moment;return true}
        }
    (*array)[x][y]=moment
    //sum:=0
    //for _,v:=range *array{
    //    for _,k:=range v{
    //        if k == "X" {
    //            sum++
    //        }
    //    }
    //}
    return false    
}
