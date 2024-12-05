package Day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Test(){
    var result int = 211
    var repeats []int
	file, err := os.Open("day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
        value := scanner.Text()
        parts := strings.Split(value, ": ")
        //indices := strings.Fields(parts[0])[1]
        subset := strings.Split(parts[1], "| ")
        winNums := strings.Fields(subset[0])
        regNums := strings.Fields(subset[1])
        //result += getPoints(winNums,regNums)

        additor := getNums(winNums,regNums)
        for i:=range repeats{
            p := &repeats[i]
            if *p == 0{continue}
            *p --
            result ++
            if (additor !=0){
                repeats = append(repeats,additor)
                fmt.Printf("Executing: %s \r",parts[0])
            }
        }
        repeats = append(repeats, additor)

        fmt.Printf("Executing: %s \r",parts[0])

	}
    if err := scanner.Err(); err != nil {
		log.Fatal(err)
    }
	file.Close()
    fmt.Println("The result is : ",result)
}

func getNums(winning []string,numbers []string)int{
    sum := 0
    for _,v:=range winning{
        for _,k:=range numbers{
            if v==k{
                sum ++
            }
        }
    }
    return sum
}
func getPoints(winning []string,numbers []string)int{
    sum := 0
    for _,v:=range winning{
        for _,k:=range numbers{
            if v==k && sum == 0{
                sum ++
            }else if v==k{
                sum *=2
            }
        }
    }
    return sum
}

