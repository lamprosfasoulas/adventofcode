package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    file, _ := os.Open("input.txt")
    defer file.Close()
    part1(file)
}
func part1(file *os.File) {
    var sum int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        var nums []int
        for _,v := range strings.Fields(line){
            p1,_:= strconv.Atoi(v)
            nums = append(nums, p1)
        }
        //Here we run the functions for each line
        if check(nums){
            sum += 1
        }else{
            for i:=range nums{
                var newArr []int
                for j:=range nums{
                    if i != j {newArr = append(newArr, nums[j])}
                }
                if check(newArr){sum +=1;break}
            }
        }
    }
    fmt.Println(sum)
}
func check(nums []int) bool {
    if (isIncreasing(nums) || isDecreasing(nums)) && hasDiff(nums){
        return true 
    }
    return false
}
func isIncreasing(nums []int) bool {
    for i:= range nums[:len(nums)-1]{
        p1 := &nums[i]
        p2 := &nums[i+1]
        if *p1 > *p2{
            return false
        }
    }
    return true
}
func isDecreasing(nums []int) bool {
    for i:= range nums[:len(nums)-1]{
        p1 := &nums[i]
        p2 := &nums[i+1]
        if *p1 < *p2{
            return false
        }
    }
    return true
}
func hasDiff(nums []int) bool {
    for i:= range nums[:len(nums)-1]{
        p1 := &nums[i]
        p2 := &nums[i+1]
        diff := *p1 - *p2
        if diff < 0 {diff *= -1}
        if diff > 3 || diff < 1{
            return false
        }
    }
    return true

}
