//bin/true  ;  go build -o day "$0"  &&  exec -a "$0" ./day "$@"  ;  exec /bin/false

package main

import "fmt"
//import "time"
import "os"
import "bufio"
//import "path"



func runProgramList(programList []string) (int, int) {

    var stepCount []int = make([]int, len(programList));
    var instPointer int = 0;
    var accumulator int = 0;


    for instPointer=0; instPointer < len(programList);  {
        var instructionName string = "";
        var argument int = 0;
        fmt.Sscanf(programList[instPointer], "%s %d", &instructionName, &argument);
        stepCount[instPointer] += 1;
        if stepCount[instPointer] > 1 {
            //Infinite loop found.
            break;
        }

        //Run instruction:
        switch instructionName {
            case "acc":
                accumulator += argument;
                instPointer++;
                break;
            case "jmp":
                instPointer += argument;
                if instPointer < 0 || instPointer > len(programList) + 1 {
                    //jmp out of range
                    break;
                }
                break;
            case "nop":
                instPointer++;
                break;
        }

        //fmt.Printf("%s %d -> %d \n", instructionName, argument, accumulator);
    }

    return instPointer, accumulator;
}


func main() {
    fmt.Println("Advent of Code - Day8");

    //Read input file:
    file, _ := os.Open("day8_input.txt");
    scanner := bufio.NewScanner(file);
    programList := make([]string, 0);
    for scanner.Scan() {
        line := scanner.Text();
        programList = append(programList, line)
    }

    //Part 1:
    instPointer, accumulator := runProgramList(programList);
    if instPointer != len(programList) {
        fmt.Printf("Instruction %d to be run a second time, with accumulator value of %d\n", instPointer, accumulator);
    }

    //Part 2:
    var modProgramList []string = make([]string, len(programList));
    for i:=0; i < len(programList); i++ {
        copy(modProgramList, programList);

        var instructionName string = "";
        var argument int = 0;
        fmt.Sscanf(modProgramList[i], "%s %d", &instructionName, &argument);

        if instructionName == "jmp" {
            instructionName = "nop";
        } else if instructionName == "nop" {
            instructionName = "jmp";
        } else {
            continue;
        }

        modProgramList[i] = fmt.Sprintf("%s %d", instructionName, argument);

        instPointer, accumulator = runProgramList(modProgramList);

        if instPointer == len(modProgramList) {
            fmt.Printf("Found Fix!, changing instruction at line %d, the program ends without infinite loop with accumulator = %d\n", i, accumulator);
            break;
        }
    }

}




