package main

import(
    "fmt"
    "os"
    "bufio"
    "log"
    "strings"
    "strconv"
)

func contains(line string) int {
        parts := strings.Split(line,",")
        first := strings.Split(parts[0],"-")
        second := strings.Split(parts[1],"-")
        
        //fmt.Printf("%s-%s, %s-%s\n",first[0],first[1],second[0],second[1])
        iStart, err0 := strconv.Atoi(first[0])
        iEnd, err1 := strconv.Atoi(first[1])
        jStart, err2 := strconv.Atoi(second[0])
        jEnd, err3 := strconv.Atoi(second[1])
    
        if err0 != nil || err1 != nil || err2 != nil || err3 != nil {
            log.Fatal("Conversion error. Goodbye")
        }
    
        if iStart <= jStart && iEnd >= jEnd {
            return 1
        }
        if iStart >= jStart && iEnd <= jEnd {
            return 1
        }
        return 0
}

func partTwoContains(line string) int {
        parts := strings.Split(line,",")
        first := strings.Split(parts[0],"-")
        second := strings.Split(parts[1],"-")
        
        //fmt.Printf("%s-%s, %s-%s\n",first[0],first[1],second[0],second[1])
        iStart, err0 := strconv.Atoi(first[0])
        iEnd, err1 := strconv.Atoi(first[1])
        jStart, err2 := strconv.Atoi(second[0])
        jEnd, err3 := strconv.Atoi(second[1])
    
        if err0 != nil || err1 != nil || err2 != nil || err3 != nil {
            log.Fatal("Conversion error. Goodbye")
        }
        
        if jStart <= iEnd && jEnd >= iStart {
            return 1
        }
        if iStart <= jEnd && iEnd >= jStart{
            return 1
        }
        return 0    
}

func main(){
    inputFile, err := os.Open("inputs/day4.txt")
    if err != nil {
        log.Fatal("Input file didn't work.")
    }
    
    fileScanner := bufio.NewScanner(inputFile)
    fileScanner.Split(bufio.ScanLines)
    
    var count int
    var countPartTwo int
    for fileScanner.Scan() {
        line := fileScanner.Text()
        count += contains(line)
        countPartTwo += partTwoContains(line)
        //fmt.Println(line, partTwoContains(line))
    }
    fmt.Println("Part One:", count)
    fmt.Println("Part Two:", countPartTwo)
    inputFile.Close() 

}
