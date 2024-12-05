package Day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var sum int
func Test1() {
	// open the file using Open() function from os library
	file, err := os.Open("values1.txt")
	if err != nil {
		log.Fatal(err)
	}

	// read the file line by line using a scanner
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		//Print the line
        value := scanner.Bytes()
        GetNum(value)
        fmt.Println(string(value))
		fmt.Println("------------")
        // len := len(value)
        // num,_ := strconv.Atoi(string(value))
        // fmt.Println(value, num, "length: ",len, "number: ", string(value))

	}
    // check for the error that occurred during the scanning
    
    if err := scanner.Err(); err != nil {
		log.Fatal(err) }
    
	// Close the file at the end of the program
	file.Close()
    fmt.Println(sum)

}
func GetNum(value []byte){
    var charset string
    for _,v := range value {
        if(v >= 48 && v <= 57){
            charset += string(v)
        }
    }
    // Call AddNums that keeps only first and last number
    // and add the to the sum
    fmt.Println(charset)
    res,_ := strconv.Atoi(AddNums(charset))
    sum += res
}
func AddNums(test string) string{
    len := len(test)
    var stbc string
    if (len == 1){
        stbc += test
        stbc += test
    }else if(len > 2){
        stbc += string(test[0])
        stbc += string(test[len -1])
    }else{
        stbc = test
    }
    return stbc
}
