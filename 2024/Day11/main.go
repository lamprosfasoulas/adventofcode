package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type numAndBlinks struct {
    num     int
    blinks  int
}
func process(stone int, blinks int, found map[numAndBlinks]int) int {
    if blinks == 0 {
        return 1
    }
    prev := found[numAndBlinks{num:stone,blinks:blinks}]
    if prev != 0 {
        return prev
    }
    var total int
    if stone == 0 {
        total = process(1,blinks-1,found)
    }else if even,half1,half2 := isEven(stone); even {
        total = process(half1,blinks-1,found) + process(half2,blinks-1,found)
    }else{
        total = process(stone*2024,blinks-1,found)
    }
    found[numAndBlinks{num:stone,blinks:blinks}] = total
    return total
}
func solve(stones []int, blinks int) {
    start := time.Now()
    var found = make(map[numAndBlinks]int)

    var total = 0
    for _, v := range stones {
        total += process(v,blinks,found)
    }
    duration := time.Since(start)
    fmt.Println("\nResult: ",total, " With duration: ",duration)
}

func isEven(num int) (bool,int,int){
    numStr := strconv.Itoa(num)
    if len(numStr)%2 != 0 {
        return false,0,0
    }
    mid := len(numStr)/2
    num1,_ := strconv.Atoi(numStr[:mid])
    num2,_ := strconv.Atoi(numStr[mid:])
    return true,num1,num2
}

func main() {
    var BLINKS int
    fmt.Scanln(&BLINKS)

    file, _:= os.Open("day11")
    defer file.Close()
    sc := bufio.NewScanner(file)
    var stones []int
    for sc.Scan(){
        line := strings.Fields(sc.Text())
        for _,v := range line{
            num,_ := strconv.Atoi(v) 
            stones = append(stones, num)
        }
    }
    solve(stones,BLINKS)
}

