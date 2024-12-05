package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
    file, _ := os.Open("input1.txt")
    defer file.Close()
    puzzle2(file)
}

func puzzle1(file *os.File) {
    sum := 0
    left,right := []int{}, []int{}
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        leftToBe ,_:= strconv.Atoi(strings.Fields(scanner.Text())[0])
        left = append(left, leftToBe)
        rightToBe,_:= strconv.Atoi(strings.Fields(scanner.Text())[1])
        right = append(right, rightToBe)
    }
    sort.Ints(left)
    sort.Ints(right)
    var p1 *int
    var p2 *int
    for i := range left{
        p1 = &left[i]
        p2 = &right[i]
        diff := (*p1 - *p2)
        if diff < 0 {
            diff *= -1
        }
        sum += diff
    }
    fmt.Println(sum)
}

func puzzle2(file *os.File){
    sum := 0
    left,right := []int{}, []int{}
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        leftToBe ,_:= strconv.Atoi(strings.Fields(scanner.Text())[0])
        left = append(left, leftToBe)
        rightToBe,_:= strconv.Atoi(strings.Fields(scanner.Text())[1])
        right = append(right, rightToBe)
    }
    for i := range left{
        p1 := &left[i]
        appear := 0
        for j := range right{
            p2 := &right[j]
            if *p1 == *p2 {
                appear += 1
            }
        }
        sum += (*p1 * appear)
    }
    fmt.Println(sum)
}
