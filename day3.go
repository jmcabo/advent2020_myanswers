//bin/true  ;  go build -o day "$0"  &&  exec -a "$0" ./day "$@"  ;  exec /bin/false

package main

import "fmt"
//import "time"
import "os"
import "bufio"
//import "path"


func isTree(mapOfTrees []string, x int, y int) (bool) {
    if y >= len(mapOfTrees) {
        return false;
    }
    row := mapOfTrees[y];
    if row[x % len(row)] == "#"[0] {
        return true;
    }
    return false;
}

func countTreesForSlope(slopeX int, slopeY int, mapOfTrees []string) (int) {

    x := 0
    y := 0
    treesFound := 0;

    for ; y < len(mapOfTrees); {
        x = x + slopeX;
        y = y + slopeY;
        if isTree(mapOfTrees, x, y) {
            treesFound++;
        }
    }

    return treesFound;
}


func main() {
    fmt.Println("Advent of Code - Day3");

    //Read input file:
    file, _ := os.Open("day3_input.txt");
    scanner := bufio.NewScanner(file);
    mapOfTrees := make([]string, 0);
    for scanner.Scan() {
        line := scanner.Text();
        mapOfTrees = append(mapOfTrees, line)
    }

    fmt.Printf("Part1 Trees found: %d\n", countTreesForSlope(3, 1, mapOfTrees));

    var mult int = 1;
    mult *= countTreesForSlope(1, 1, mapOfTrees);
    mult *= countTreesForSlope(3, 1, mapOfTrees);
    mult *= countTreesForSlope(5, 1, mapOfTrees);
    mult *= countTreesForSlope(7, 1, mapOfTrees);
    mult *= countTreesForSlope(1, 2, mapOfTrees);

    fmt.Printf("Multiplication of trees found in each slope: %d\n", mult);
}




