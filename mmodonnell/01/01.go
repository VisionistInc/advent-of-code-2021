package main

import "io/ioutil"
import "fmt"
import "strconv"
import "strings"

func main() {
    filename := "input.txt"
    content, err := ioutil.ReadFile(filename)
    if err != nil {
        //Do something
    }
    lines := strings.Split(string(content), "\n")
    count := 0 
    value01, _ := strconv.Atoi(lines[0]) 
    value02, _ := strconv.Atoi(lines[1])
    value03, _ := strconv.Atoi(lines[2])
    previous_sum := value01 + value02 + value03
    for i :=2; i < (len(lines)); i++ {
        value1, _ := strconv.Atoi(lines[i]) 
        value2, _ := strconv.Atoi(lines[i-1])
        value3, _ := strconv.Atoi(lines[i-2])
        cur_sum := value1+value2+value3
        if  cur_sum > previous_sum {
            count++
        }
        previous_sum = cur_sum // THIS IS NOT THE SAME AS PREVIOUS_SUM := CUR_SUM ADFSJKADJKHFDAJHFD
    }
    fmt.Println("test")
    fmt.Println(count)
}
