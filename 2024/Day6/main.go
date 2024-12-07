package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    length:=130

    var posx,posy int
    dirx,diry:=-1,0

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
    var nextx,nexty int
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


        if array[nextx][nexty]=="#" {
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
        if array[posx][posy]=="." {
            array[posx][posy]="X" 
        }
    }
    for _,v:=range array{
        for _,k:=range v{
            if k == "X" {
                sum++
            }
        }
    }
    fmt.Println("\nThe final output is ",sum+1)//output part 1 (need to add 1 to account for the starting position)
}


