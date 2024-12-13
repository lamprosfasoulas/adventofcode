package main

import (
	"bufio"
	"fmt"
	"os"
)

var moves = [][2]int {
    {0,-1},
    {0,1},
    {-1,0},
    {1,0},
}

var nines int
func main(){
    var board = make([][]int,53)
    for i:=range board{
        board[i] = make([]int,53)
    }
    file,_:=os.Open("day10")
    defer file.Close()
    var sum int
    sc := bufio.NewScanner(file)
    i:=0
    for sc.Scan(){
        for j,v := range sc.Bytes(){
            board[i][j] = int(v-byte('0'))
        }
        fmt.Println(board[i])
        i++
    }
    for i:=range board{
        for j:=range board[i]{
            if board[i][j] == 0{
                nines = 0
                score(board,i,j)
                sum += nines
                fmt.Println(sum)
            }
        }
    }
}
func score(board [][]int,i,j int,/*found map[[2]int]bool*/){
    found := make(map[[2]int]bool)
    if board[i][j] == 9 && found[[2]int{i,j}] == false{
        nines++
        found[[2]int{i,j}] = true
    }
    fmt.Println(i,j,"Value:",board[i][j])
    for _,m := range moves{
        x,y := i+m[0],j+m[1]
        if x>=0 && x<len(board) && y>=0 && y<len(board){
            if board[x][y] == board[i][j]+1{
                score(board,x,y)
            }else{
                continue
            }
        }
    }
    fmt.Println("Nines:",nines)
}
