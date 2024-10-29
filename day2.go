package main

import(
    "fmt"
    "os"
    "bufio"
    "log"
)
func calculateScorePartTwo(line string) int {
    // X lose, Y draw, Z win
    switch line{
    case "A X":
        return 3 + 0 
    case "A Y":
        return 1 + 3
    case "A Z":
        return 2 + 6
    case "B X":
        return 1 + 0
    case "B Y":
        return 2 + 3
    case "B Z":
        return 3 + 6
    case "C X":
        return 2 + 0
    case "C Y":
        return 3 + 3
    case "C Z":
        return 1 + 6
    default:
        log.Fatal("Error with input", line)
        return -1
    }
}


func calculateScore(line string) int {
    switch line{
    case "A X":
        return 1 + 3 
    case "A Y":
        return 2 + 6
    case "A Z":
        return 3 + 0
    case "B X":
        return 1 + 0
    case "B Y":
        return 2 + 3
    case "B Z":
        return 3 + 6
    case "C X":
        return 1 + 6
    case "C Y":
        return 2 + 0
    case "C Z":
        return 3 + 3
    default:
        log.Fatal("Error with input", line)
        return -1
    }
}

func main() {
    /*
    A,X = Rock      1pt
    B,Y = Paper     2pt
    C,Z = Scissors  3pt
    0 lost, 3 draw, 6 win
    */
    inputFile, err := os.Open("inputs/day2.txt")
    if err != nil {
        log.Fatal("Input file didn't work.")
    }
    
    fileScanner := bufio.NewScanner(inputFile)
    fileScanner.Split(bufio.ScanLines)
    
    var total int  
    var totalPartTwo int
    for fileScanner.Scan() {
        line := fileScanner.Text()
        total += calculateScore(line)
        totalPartTwo += calculateScorePartTwo(line)
    }
    fmt.Println("Part One:", total)
    fmt.Println("Part Two:", totalPartTwo)
    inputFile.Close() 
}
