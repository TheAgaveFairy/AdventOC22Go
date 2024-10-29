package main

import(
    "fmt"
    "os"
    "bufio"
    "log"
//    "strconv"
)

func partTwo(line1 string, line2 string, line3 string) int {
    for _, c1 := range line1 {
        for _, c2 := range line2 {
            for _, c3 := range line3 {
                if c1 == c2 && c2 == c3 {
                    return getCharValue(c1)
                }
            }
        }
    }
    return -1
}

func getCharValue(char rune) int {
    if char >= 'a' && char <= 'z' {
        return int(char - 'a') + 1
    } else if char >= 'A' && char <= 'Z' {
        return int(char - 'A') + 1 + 26
    } else {
        log.Fatal("Error with rune.")
        return 0
    }
}

func process(line string) int {
    length := len(line)
    half := length / 2
    front := line[:half]
    back := line[half:]
    //fmt.Printf("%d, %d. ", len(front), len(back))
    for _, f := range front {
        for _, b := range back {
            if f == b {
                //fmt.Print(f,string(b),getCharValue(f)," ")
                return getCharValue(f)
            }
        }
    }
    return -1
}

func main() {
    inputFile, err := os.Open("inputs/day3.txt")
    if err != nil {
        log.Fatal("Input file didn't work.")
    }
    
    fileScanner := bufio.NewScanner(inputFile)
    fileScanner.Split(bufio.ScanLines)
    
    var total int
    var groupsTotal int
    lineGroup := make([]string, 0, 3)
    for fileScanner.Scan() {
        line := fileScanner.Text()
        lineGroup = append(lineGroup, line)
        if len(lineGroup) == 3 {
            groupsTotal += partTwo(lineGroup[0], lineGroup[1], lineGroup[2])
            lineGroup = lineGroup[:0]
        }
        total += process(line)
    }
    fmt.Println("Part One:", total)
    fmt.Println("Part Two:", groupsTotal)
    inputFile.Close() 
}
