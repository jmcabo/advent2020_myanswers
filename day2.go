//bin/true  ;  go build -o day "$0"  &&  exec -a "$0" ./day "$@"  ;  exec /bin/false

package main

import "fmt"
//import "time"
import "os"
import "bufio"
//import "path"


func isValidLine(min int, max int, ch string, p string) (bool) {

    var appearancesOfChar = 0;
    for i:=0; i < len(p); i++ {
        if p[i] == ch[0] {
            appearancesOfChar++;
        }
    }

    if appearancesOfChar >= min && appearancesOfChar <= max {
        return true;
    }
    return false;
}


func isValidLinePart2(firstPos int, secondPos int, ch string, p string) (bool) {

    var appearancesOfChar = 0;
    if firstPos >= 1 && firstPos <= len(p) && secondPos >= 1 && secondPos <= len(p) {
        if p[firstPos-1] == ch[0] {
            appearancesOfChar++;
        }
        if p[secondPos-1] == ch[0] {
            appearancesOfChar++;
        }
    }

    if appearancesOfChar == 1 {
        return true;
    }
    return false;
}


func main() {
    fmt.Println("Advent of Code - Day2");

    //Read input file:
    file, _ := os.Open("day2_input.txt");
    scanner := bufio.NewScanner(file);
    list := make([]string, 0);
    for scanner.Scan() {
        line := scanner.Text();
        list = append(list, line)
    }

    var validLines = 0;
    for i:=0; i < len(list); i++ {
        var min int;
        var max int;
        var ch string;
        var p string;
        fmt.Sscanf(list[i], "%d-%d %1s: %s", &min, &max, &ch, &p);
        //fmt.Printf("%d - %d char %s, w: %s\n", min, max, ch, p);

        if isValidLine(min, max, ch, p) {
            validLines++;
        }
    }

    fmt.Printf("Valid lines: %d\n", validLines);


    //Part 2
    validLines = 0;
    for i:=0; i < len(list); i++ {
        var firstPos int;
        var secondPos int;
        var ch string;
        var p string;
        fmt.Sscanf(list[i], "%d-%d %1s: %s", &firstPos, &secondPos, &ch, &p);

        if isValidLinePart2(firstPos, secondPos, ch, p) {
            validLines++;
        }
    }

    fmt.Printf("Valid lines: %d\n", validLines);

}
