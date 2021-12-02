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
    aim := 0
    depth := 0
    forward := 0
    count := 0
    for i :=0; i < (len(lines)); i++ {
        count++
        
        fmt.Println(lines[i])
        line := strings.Split(string(lines[i]), " ")
        // if(count==17){
        //     // println(count)
        //     break
        // }
        if(len(line)==1){
            fmt.Println(line)
            println(count)
            break
        }
        movement_type := line[0]
        movement_value, _ := strconv.Atoi(line[1]) 
        if(strings.Contains(movement_type, "f")){ // forward
            // fmt.Println("fo")
            forward +=  movement_value
            depth += movement_value * aim
        } else if (strings.Contains(movement_type, "d")){ // down
            // fmt.Println("do")
            aim +=  movement_value
        } else { // up
            // fmt.Println("u")
            aim -=  movement_value
        }
        fmt.Println(forward)
        fmt.Println(aim)
    }
    fmt.Println(forward)
    fmt.Println(depth)
    fmt.Println(depth*forward)
}
