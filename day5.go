package main

import(
    "fmt"
    "bufio"
    "log"
    "os"
)

func main() {
    inputFile, err := os.Open("inputs/day5.txt")
    if err != nil {
        log.Fatal("Input file didn't work.")
    }
    
    fileScanner := bufio.NewScanner(inputFile)
    fileScanner.Split(bufio.ScanLines)
    
    for fileScanner.Scan() {
        line := fileScanner.Text()
    }
    fmt.Println("Part One:", count)
    inputFile.Close() 

}
