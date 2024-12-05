package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const filename = "input.txt"

func main(){
    var sum int
    file, _ := os.Open(filename)
    defer file.Close()
    scanner := bufio.NewScanner(file)
    scan := false
    for scanner.Scan() {
        line := scanner.Text()
        //here we skip the first part of the file
        if !scan{
            if line == ""{
                scan = true
                fmt.Println("Update lists")
            }
            continue
        }
        //here we print the lists
        //we will use grep to find the rules corresponding to each update list item
        //fmt.Println("New Line : ",line)
        numbers := strings.Split(line,",")
        //count := false

        //this is for part one
        //for i,v := range numbers{//this will check if the order is right
        //    if (i == len(numbers)-1) && count{sum += findMiddle(numbers);continue}
        //    count = order(v,numbers[i+1])
        //    if !count{break}
        //}

        //this is for part two
        isCorrect := true
        for i:=0;i<len(numbers)-1;i++{
            isCorrect = order(numbers[i],numbers[i+1])
            if !isCorrect{break}
        }
        if !isCorrect{
            sum += makeCorrect(numbers)
        }
    }
    fmt.Println(sum)
}

func order(searchTerm string, check string)bool{
    cmd := exec.Command("bash","-c",fmt.Sprintf("grep '%s|' %s | cut -d '|' -f 2",searchTerm,filename))
    out,_ := cmd.Output()
    rules := strings.Split(string(out),"\n")
    for _,v := range rules{
        if v == check {
            return true
        }
    }
    return false
}

func findMiddle(numbs []string)int{
    //we use this function to find the middle number of the correct ordered list
    var middle int
    length := len(numbs)
    middle,_ = strconv.Atoi(numbs[length/2])
    return middle
}

func makeCorrect(numbs []string)int{
    //newArr := make([]string,0,len(numbs))
    var middle int
    var nextPos[]int

    for _,v:=range numbs{
        cmd := exec.Command("bash","-c",fmt.Sprintf("grep '|%s' %s | cut -d '|' -f 1",v,filename))
        out,_ := cmd.Output()
        rules := strings.Split(string(out),"\n")
        gravity := 0
        for _,v := range rules{
            if IndexOf(v,numbs){
                gravity++
            }
        }
        if gravity == len(numbs)/2{
            middle,_ = strconv.Atoi(v)
        }
        nextPos= append(nextPos,gravity) 
    }
    //here we find the middle fo the newly created array adn return it 
    //tho be added to the new sum
    //length := len(newArr)
    //middle,_ := strconv.Atoi(newArr[length/2])
    return middle
}


func IndexOf(v string, arr []string)bool{
    for _,val := range arr{
        if v == val{
            return true
        }
    }
    return false
}
