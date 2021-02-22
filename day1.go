//bin/true  ;  go build -o day "$0"  &&  exec -a "$0" ./day "$@"  ;  exec /bin/false

package main

import "fmt"
import "strconv"
//import "time"
import "os"
import "bufio"
//import "path"


func main() {
    fmt.Println("Advent of Code - Day1");

    //Read input file:
    file, _ := os.Open("day1_input.txt");
    scanner := bufio.NewScanner(file);
    list := make([]int64, 0);
    for scanner.Scan() {
        line := scanner.Text();
        v, _ := strconv.ParseInt(line, 10, 32);
        list = append(list, v)
    }

    var found bool = false;
    for i:=0; i < len(list); i++ {
        for j:=i+1; j < len(list); j++ {
            if list[i] + list[j] == 2020 {
                fmt.Printf("Found pair that sums exactly 2020 at pos %d and %d. Values %d and %d\n", i, j, list[i], list[j]);
                fmt.Printf("Part 1: Multiplied: %d\n", list[i] * list[j]);
                found = true
                break;
            }
        }
        if found {
            break;
        }
    }

    found = false;
    for i:=0; i < len(list); i++ {
        for j:=i+1; j < len(list); j++ {
            for k:=j+1; k < len(list); k++ {

                if list[i] + list[j] + list[k] == 2020 {
                    fmt.Printf("Found the three entries that sum 2020\n");
                    fmt.Printf("Part 2: Multiplied: %d\n", list[i] * list[j] * list[k]);
                    found = true;
                    break;
                }
            }
            if found {
                break;
            }
        }
        if found {
            break;
        }
    }


    //file.Close();
}

