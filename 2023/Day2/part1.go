package Day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)


var sumEr int
func Test(){
	file, err := os.Open("day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
        value := scanner.Text()
        parts := strings.Split(value, ": ")
        indices,_ := strconv.Atoi(strings.Split(parts[0]," ")[1])
        subset := strings.Split(parts[1], "; ")
        
        sumEr += findNums(subset, indices)

        fmt.Println(value)
        fmt.Println(indices)
		fmt.Println("------------------------------")

	}
    if err := scanner.Err(); err != nil {
		log.Fatal(err)
    }
	file.Close()
    fmt.Println(sumEr)

}

func findNums(pieces []string, index int)int{
        for _,v := range pieces{
            pieces := strings.Split(v,", ")
            fmt.Println(pieces)
        for _,v := range pieces{
            one,_ := strconv.Atoi(strings.Split(v," ")[0])
            two := strings.Split(v," ")[1]
            if(string(two) == "red" && one > 12){
                return 0
            }
            if (string(two) == "green" && one > 13)  {
                return 0
            }
            if (string(two) == "blue" && one > 14)  {
                return 0
            }
            fmt.Println(one,string(two)) 
        }
        }
    return index
}
