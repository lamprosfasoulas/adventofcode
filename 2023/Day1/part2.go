package Day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	//"strconv"
)

var wordMap = map[string]string{
"one":    "1",
"two":    "2",
"three":  "3",
"four":   "4",
"five":   "5",
"six":    "6",
"seven":  "7",
"eight":  "8",
"nine":   "9",
"1":  "1",
"2":  "2",
"3":  "3",
"4":  "4",
"5":  "5",
"6":  "6",
"7":  "7",
"8":  "8",
"9":  "9",
}

func Test2() {
	file, err := os.Open("./2023/Day1/values1.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//Print the line
        value := scanner.Text()
        test(string(value),wordMap) 
        //getNewNum(value)
        fmt.Println(value)
		fmt.Println("------------")

	}
    if err := scanner.Err(); err != nil {
		log.Fatal(err)
    }
	file.Close()
    fmt.Println(sum)
}
func getNewNum(testvalue string){
    var charset string
    for word := range wordMap{
        if strings.Contains(testvalue,word){
            charset += wordMap[word]
            fmt.Println("Contains word: ",word," Which is the number:  ", wordMap[word])
        }
    }
    res,_ := strconv.Atoi(AddNums(charset))
    sum += res
    fmt.Println(testvalue,AddNums(charset))
}

func test(testString string, wordMap map[string]string) {
    firstIndex := -1    // Overall first occurrence index
    lastIndex := -1     // Overall last occurrence index
    firstWord := ""     // Word for the first occurrence
    lastWord := ""      // Word for the last occurrence
    charset := ""
    // Loop through each word in the map
    for word := range wordMap {
        // Find the first occurrence of the current word
        index := strings.Index(testString, word)
        if index != -1 {
            // Check if it's the first occurrence overall
            if firstIndex == -1 || index < firstIndex {
                firstIndex = index
                firstWord = word
            }
        }

        // Find the last occurrence of the current word
        lastIndexFound := strings.LastIndex(testString, word)
        if lastIndexFound != -1 {
            // Check if it's the last occurrence overall
            if lastIndex == -1 || lastIndexFound > lastIndex {
                lastIndex = lastIndexFound
                lastWord = word
            }
        }
    }

    // Print results
    charset = wordMap[firstWord]+wordMap[lastWord]
    res,_ := strconv.Atoi(AddNums(charset))
    sum += res
    if firstIndex != -1 {
        fmt.Printf("First occurrence: Word: '%s'\n", firstWord)
    } else {
        fmt.Println("No words from the map found in the test string.")
    }

    if lastIndex != -1 {
        fmt.Printf("Last occurrence: Word: '%s'\n", lastWord)
    } else {
        fmt.Println("No words from the map found in the test string.")
    }
}
