//bin/true  ;  go build -o day "$0"  &&  exec -a "$0" ./day "$@"  ;  exec /bin/false

package main

import "fmt"
import "os"
import "bufio"
import "strings"



func checkPport(pport map[string]string) bool {
    //Check if Valid:
    _, byrOk := pport["byr"];
    _, iyrOk := pport["iyr"];
    _, eyrOk := pport["eyr"];
    _, hgtOk := pport["hgt"];
    _, hclOk := pport["hcl"];
    _, eclOk := pport["ecl"];
    _, pidOk := pport["pid"];
    //_, cidOk := pport["cid"];
    if byrOk && iyrOk && eyrOk && hgtOk && hclOk && eclOk && pidOk {
        return true;
    }
    return false;
}

func checkPportStrict(pport map[string]string) bool {
    //Check if Valid:
    byr, byrOk := pport["byr"];
    iyr, iyrOk := pport["iyr"];
    eyr, eyrOk := pport["eyr"];
    hgt, hgtOk := pport["hgt"];
    hcl, hclOk := pport["hcl"];
    ecl, eclOk := pport["ecl"];
    pid, pidOk := pport["pid"];
    //_, cidOk := pport["cid"];
    if byrOk && iyrOk && eyrOk && hgtOk && hclOk && eclOk && pidOk {
        byrInt := 0;
        n, ok := fmt.Sscanf(byr, "%d", &byrInt);
        if n != 1 || ok != nil || byrInt < 1920 || byrInt > 2002 {
            return false;
        }

        iyrInt := 0;
        n, ok = fmt.Sscanf(iyr, "%d", &iyrInt);
        if n != 1 || ok != nil || iyrInt < 2010 || iyrInt > 2020 {
            return false;
        }

        eyrInt := 0;
        n, ok = fmt.Sscanf(eyr, "%d", &eyrInt);
        if n != 1 || ok != nil || eyrInt < 2020 || eyrInt > 2030 {
            return false;
        }

        hgtNumber := 0
        hgtU := ""
        n, ok = fmt.Sscanf(hgt, "%d%s", &hgtNumber, &hgtU);
        if n != 2 || ok != nil || (hgtU != "cm" && hgtU != "in") || (hgtU == "cm" && (hgtNumber < 150 || hgtNumber > 193)) || (hgtU == "in" && (hgtNumber < 59 || hgtNumber > 76)) {
            return false;
        }

        hclNumber := 0
        n, ok = fmt.Sscanf(hcl, "#%6x", &hclNumber)
        if n != 1 || ok != nil || len(hcl) != 7 {
            return false;
        }

        if ecl != "amb" && ecl != "blu" && ecl != "brn" && ecl != "gry" && ecl != "grn" && ecl != "hzl" && ecl != "oth" {
            return false;
        }

        pidNumber := 0
        n, ok = fmt.Sscanf(pid, "%9d", &pidNumber)
        if n != 1 || ok != nil || len(pid) != 9 {
            return false;
        }

        return true;
    }
    return false;
}


func main() {
    fmt.Println("Advent of Code - Day4");

    //Read input file:
    file, _ := os.Open("day4_input.txt");
    scanner := bufio.NewScanner(file);
    lines := make([]string, 0);
    for scanner.Scan() {
        line := scanner.Text();
        lines = append(lines, line)
    }

    fmt.Printf("Line count: %d\n", len(lines));


    totalCount := 0
    validCount := 0
    validStrictCount := 0
    pport := make(map[string]string, 0);
    for i:=0; i < len(lines); i++ {
        if (lines[i] == "") {
            //End of pport
            totalCount++;

            if (checkPport(pport)) {
                validCount++;
            }

            if (checkPportStrict(pport)) {
                validStrictCount++;
            }

            //Start again with new:
            pport = make(map[string]string, 0);
        }

        keyValuePairs := strings.Split(lines[i], " ");
        for j:=0; j < len(keyValuePairs); j++ {
            if keyValuePairs[j] == "" {
                continue;
            }
            //fmt.Printf("kvp: %s\n", keyValuePairs[j]);
            pair := strings.Split(keyValuePairs[j], ":");
            pairName := pair[0];
            pairValue := pair[1];
            pport[pairName] = pairValue;
        }
    }

    //End of last pport:
    if (checkPport(pport)) {
        validCount++;
    }
    if (checkPportStrict(pport)) {
        validStrictCount++;
    }

    fmt.Printf("Total pport count: %d\n", totalCount);
    fmt.Printf("Valid pport count (answer): %d\n", validCount);
    fmt.Printf("Valid strict pport count (answer 2): %d\n", validStrictCount);
}




