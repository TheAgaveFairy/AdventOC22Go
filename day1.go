package main

import(
    "fmt"
    "os"
    "bufio"
    "log"
    "strconv"
)

func updateTopThree(num int, topThree *[3]int){
    for i, val := range topThree {
        if num > val {
            copy(topThree[i+1:], topThree[i:])
            topThree[i] = num
            break
        }
    }
}

func sumTopThree(topThree []int) (sum int) {
    for _, val := range topThree{
        sum += val
    }
    return
}

func main(){
    inputFile, err := os.Open("inputs/day1.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer inputFile.Close()

    fileScanner := bufio.NewScanner(inputFile)
    fileScanner.Split(bufio.ScanLines)
    
    var elves []int
    var tempSum int
    var maxCalories int
    var part2 [3]int

    for fileScanner.Scan() {
        line := fileScanner.Text()
        //fmt.Println(line)
        if line != ""{
            tempInt, err := strconv.Atoi(line)
            if err != nil{
                log.Fatal(err)
            }
            tempSum += tempInt
        } else {
            elves = append(elves, tempSum)
            maxCalories = max(maxCalories, tempSum)
            updateTopThree(tempSum, &part2)
            tempSum = 0
        }
    }
    /*
    for i, val := range elves{
        fmt.Printf("Elf %d has %d calories.\n", i, val)
    }
    */
    fmt.Println("The Max Number of Calories:", maxCalories)
    fmt.Println("Total of Top Three:", sumTopThree(part2[:]))
}
