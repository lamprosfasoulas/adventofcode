package Day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)


func Test2(){
	file, err := os.Open("day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
        value := scanner.Text()
        parts := strings.Split(value, ": ")
        //indices,_ := strconv.Atoi(strings.Split(parts[0]," ")[1])
        subset := strings.Split(parts[1], "; ")
        
        sumEr += findNums2(subset)

		//fmt.Println("------------------------------")

	}
    if err := scanner.Err(); err != nil {
		log.Fatal(err)
    }
	file.Close()
    fmt.Println(sumEr)
}

func findNums2(pieces []string)int{
        blue,red,green := 0,0,0
        for _,v := range pieces{
            pieces := strings.Split(v,", ")
            //fmt.Println(pieces)
        for _,v := range pieces{
            one,_ := strconv.Atoi(strings.Split(v," ")[0])
            two := strings.Split(v," ")[1]
            if(string(two) == "red" && one > red){
                red = one 
            }
            if (string(two) == "green" && one > green)  {
                green = one
            }
            if (string(two) == "blue" && one > blue)  {
                blue = one
            }
            //fmt.Println(one,string(two)) 
        }
        //fmt.Printf("High red: %d,High blue: %d, High green: %d\n",red,blue,green)
        //fmt.Println(blue*red*green)

        }
    return blue*red*green
}
