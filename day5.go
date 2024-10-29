package main

import(
    "fmt"
    "bufio"
    "log"
    "os"
    "strings"
    "strconv"
    "slices"
)

func processInstr(line string) (amount int, from int, to int) {
    n, err := fmt.Sscanf(line, "move %d from %d to %d", &amount, &from, &to)
    if err != nil && n > 0{
        log.Fatal("Bad line of instructions")
    }
    return // remember that these are 1 indexed stacks, annoying
}

func findStackCount(filename string) (int, int) {
    inputFile, err := os.Open(filename)
    if err != nil {
        log.Fatal("Error opening file to get counts.")
    }
    fileScanner := bufio.NewScanner(inputFile)
    fileScanner.Split(bufio.ScanLines)

    var lineNumber int
    for fileScanner.Scan() {
        lineNumber += 1
        line := fileScanner.Text()
        trimmed := strings.TrimSpace(line)
        splits := strings.Split(trimmed," ")
        
        _, err := strconv.Atoi(splits[0])
        if err == nil {
            tempLen := len(splits)
            answer, err := strconv.Atoi(splits[tempLen-1])
            if err != nil {
                log.Fatal("Couldn't find range.")
            } else {
                inputFile.Close()
                return answer, lineNumber
            }
        }
    }
    inputFile.Close()
    return 0, 0
}

func processingHeader(filename string) [][]rune  {
    stackCount, maxHeight := findStackCount(filename)
    stacks := make([][]rune, stackCount)
    for i := range stacks {
        stacks[i] = make([]rune, 0, maxHeight * 2)
    }

    inputFile, err := os.Open(filename)
    if err != nil {
        log.Fatal("Error opening file to get counts.")
    }
    fileScanner := bufio.NewScanner(inputFile)
    for fileScanner.Scan() {
        line := fileScanner.Text()
        runes := []rune(line)
        for i := 0; i < len(runes); i += 4 {
            stackNum := i / 4
            if runes[i] == '[' {
                tempRune := runes[i+1]
                stacks[stackNum] = append(stacks[stackNum], tempRune)
            }
        }
    }
    for row := range stacks {
        slices.Reverse(stacks[row])
    }
    inputFile.Close()
    return stacks // they're in reverse order for now
}

func partOneMoving(filename string, stacks [][]rune) {
    inputFile, err := os.Open(filename)
    if err != nil {
        log.Fatal("Error opening file to start movements.")
    }
    fileScanner := bufio.NewScanner(inputFile)

    for fileScanner.Scan() {
        line := fileScanner.Text()
        amount, from, to := processInstr(line)
        //fmt.Printf("%d, %d, %d. ", amount, from, to)
        if amount > 0 { // valid line
            for i := 0; i < amount; i++ {
                endOfFrom := len(stacks[from-1]) - 1
                x := stacks[from-1][endOfFrom] // get last element
                stacks[from-1] = stacks[from-1][:endOfFrom] // remove last, finish pop()
                stacks[to-1] = append(stacks[to-1], x) // push ()
            }
        }
    }
    inputFile.Close()
    return

}

func main() {
    filename := "inputs/day5.txt"
    stacks := processingHeader(filename)

    for r := range stacks { // row by row
        fmt.Println(string(stacks[r]))
    }
    fmt.Println("Answer is the top of each of these stacks:")
    partOneMoving(filename,stacks)    
     for r := range stacks { // row by row
        fmt.Println(string(stacks[r]))
    }
}
