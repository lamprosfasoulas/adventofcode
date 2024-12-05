package Day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Test(){
    //variables
    var seeds []int


    file,_ := os.Open("day5.txt")
    scanner:=bufio.NewScanner(file)
    if scanner.Scan(){
        value := scanner.Text()
        seedlings := strings.Fields(strings.Split(value, ":")[1])
        for _,v:=range seedlings{
            apVal,_ := strconv.Atoi(v)
            seeds = append(seeds, apVal)
        }
        fmt.Println(seeds)
    }
    file.Close()
    for i:=range seeds{
        file,_ := os.Open("day5.txt")
        scanner:=bufio.NewScanner(file)
        fmt.Println("Seed ", seeds[i])
        for scanner.Scan(){
            line := scanner.Text()
            if strings.Contains(line,"seeds:"){
                continue
            }
            if strings.Contains(line,":"){
                blocks(scanner,line)
            }
        }
    }

}

func blocks(scanner *bufio.Scanner,startline string){
    fmt.Println("Starting At @@",startline)
    for scanner.Scan(){
        line := scanner.Text()
        if line == ""{break}

    }
}

