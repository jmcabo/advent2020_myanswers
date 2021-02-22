//bin/true  ;  go build -o day "$0"  &&  exec -a "$0" ./day "$@"  ;  exec /bin/false

package main

import "fmt"
//import "strconv"
//import "time"
//import "os"
//import "bufio"
//import "path"


func computeLastPos(allValues []int) int {

    n := allValues[len(allValues)-1];
    lastPos := 0;
    for i:=len(allValues)-2; i >= 0; i-- {
        if allValues[i] == n {
            lastPos = len(allValues) - 1 - i;
            break;
        }
    }
    return lastPos;
}


func main() {
    fmt.Println("Advent of Code - Day15");

    last3 := make([]int, 0);
    //last3 = append(last3, 3);
    //last3 = append(last3, 1);
    //last3 = append(last3, 2);

    last3 = append(last3, 0);
    last3 = append(last3, 6);
    last3 = append(last3, 1);
    last3 = append(last3, 7);
    last3 = append(last3, 2);
    last3 = append(last3, 19);
    last3 = append(last3, 20);


    allValues := last3;

    //for i:=len(allValues); i < 2020; i++ {
    for i:=len(allValues); i < 3000000; i++ {
        n := computeLastPos(allValues);
        allValues = append(allValues, n);
    }

    fmt.Printf("Last number in 2020th position is: %d\n", allValues[len(allValues)-1]);
    //fmt.Printf("list: %#v \n", allValues); //testing
}

