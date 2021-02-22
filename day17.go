//bin/true  ;  go build -o day "$0"  &&  exec -a "$0" ./day "$@"  ;  exec /bin/false

package main

import "fmt"
//import "time"
import "os"
import "bufio"
//import "path"


func cubesSet(cubes map[int] map[int] map[int] byte, x int, y int, z int, v byte) {

    m, ok := cubes[x];
    if !ok {
        m = make(map[int] map[int] byte);
        cubes[x] = m;
    }

    m2, ok := cubes[x][y];
    if !ok {
        m2 = make(map[int] byte);
        cubes[x][y] = m2;
    }

    cubes[x][y][z] = v;
}

func cubes4Set(cubes map[int] map[int] map[int] map[int] byte, x int, y int, z int, w int, v byte) {

    m, ok := cubes[x];
    if !ok {
        m = make(map[int] map[int] map[int] byte);
        cubes[x] = m;
    }

    m2, ok := cubes[x][y];
    if !ok {
        m2 = make(map[int] map[int] byte);
        cubes[x][y] = m2;
    }

    m3, ok := cubes[x][y][z];
    if !ok {
        m3 = make(map[int] byte);
        cubes[x][y][z] = m3;
    }

    cubes[x][y][z][w] = v;
}


//Returns empty '.' if not set before.
func cubesGet(cubes map[int] map[int] map[int] byte, x int, y int, z int) (byte) {
    _, ok := cubes[x];
    if !ok {
        return '.';
    }

    _, ok = cubes[x][y];
    if !ok {
        return '.';
    }

    v, ok := cubes[x][y][z];
    if !ok {
        return '.';
    }
    return v;
}

//Returns empty '.' if not set before.
func cubes4Get(cubes map[int] map[int] map[int] map[int] byte, x int, y int, z int, w int) (byte) {
    _, ok := cubes[x];
    if !ok {
        return '.';
    }

    _, ok = cubes[x][y];
    if !ok {
        return '.';
    }

    _, ok = cubes[x][y][z];
    if !ok {
        return '.';
    }

    v, ok := cubes[x][y][z][w];
    if !ok {
        return '.';
    }
    return v;
}

func copyCubes(cubes map[int] map[int] map[int] byte) (map[int] map[int] map[int] byte) {
    var destCubes map[int] map[int] map[int] byte = make(map[int] map[int] map[int] byte);

    for x, m := range cubes {
        for y, m2 := range m {
            for z, v := range m2 {
                cubesSet(destCubes, x, y, z, v);
            }
        }
    }

    return destCubes;
}

func copyCubes4(cubes map[int] map[int] map[int] map[int] byte) (map[int] map[int] map[int] map[int] byte) {
    var destCubes map[int] map[int] map[int] map[int] byte = make(map[int] map[int] map[int] map[int] byte);

    for x, m := range cubes {
        for y, m2 := range m {
            for z, m3 := range m2 {
                for w, v := range m3 {
                    cubes4Set(destCubes, x, y, z, w, v);
                }
            }
        }
    }

    return destCubes;
}

func countActiveNear(cubes map[int] map[int] map[int] byte, x int, y int, z int) int {
    var activeCount int = 0;

    for i:=-1; i <= 1; i++ {
        for j:=-1; j <= 1; j++ {
            for k:=-1; k <= 1; k++ {
                if i == 0 && j == 0 && k == 0 {
                    continue;
                }
                if (cubesGet(cubes, x+i, y+j, z+k) == '#') {
                    activeCount++;
                }
            }
        }
    }

    return activeCount;
}

func countActiveNear4(cubes map[int] map[int] map[int] map[int] byte, x int, y int, z int, w int) int {
    var activeCount int = 0;

    for i:=-1; i <= 1; i++ {
        for j:=-1; j <= 1; j++ {
            for k:=-1; k <= 1; k++ {
                for l:=-1; l <= 1; l++ {
                    if i == 0 && j == 0 && k == 0 && l == 0 {
                        continue;
                    }
                    if (cubes4Get(cubes, x+i, y+j, z+k, w+l) == '#') {
                        activeCount++;
                    }
                }
            }
        }
    }

    return activeCount;
}


func countActiveTotal(cubes map[int] map[int] map[int] byte) (int) {

    var activeTotal = 0;

    for _, m := range cubes {
        for _, m2 := range m {
            for _, v := range m2 {
                if v == '#' {
                    activeTotal++;
                }
            }
        }
    }

    return activeTotal;
}


func countActiveTotal4(cubes map[int] map[int] map[int] map[int] byte) (int) {

    var activeTotal = 0;

    for _, m := range cubes {
        for _, m2 := range m {
            for _, m3 := range m2 {
                for _, v := range m3 {
                    if v == '#' {
                        activeTotal++;
                    }
                }
            }
        }
    }

    return activeTotal;
}

func cubesMaxDimensions(cubes map[int] map[int] map[int] byte) (int, int, int, int, int, int) {
    var minX int = 0;
    var maxX int = 0;
    var minY int = 0;
    var maxY int = 0;
    var minZ int = 0;
    var maxZ int = 0;

    //Iterate over each cube in prevCubes and set next value to cubes:
    for x, m := range cubes {
        for y, m2 := range m {
            for z, _ := range m2 {
                if x < minX { minX = x; }
                if x > maxX { maxX = x; }
                if y < minY { minY = y; }
                if y > maxY { maxY = y; }
                if z < minZ { minZ = z; }
                if z > maxZ { maxZ = z; }
            }
        }
    }

    return minX, maxX, minY, maxY, minZ, maxZ;
}

func cubesMaxDimensions4(cubes map[int] map[int] map[int] map[int] byte) (int, int, int, int, int, int, int, int) {
    var minX int = 0;
    var maxX int = 0;
    var minY int = 0;
    var maxY int = 0;
    var minZ int = 0;
    var maxZ int = 0;
    var minW int = 0;
    var maxW int = 0;

    //Iterate over each cube in prevCubes and set next value to cubes:
    for x, m := range cubes {
        for y, m2 := range m {
            for z, m3 := range m2 {
                for w, _ := range m3 {
                    if x < minX { minX = x; }
                    if x > maxX { maxX = x; }
                    if y < minY { minY = y; }
                    if y > maxY { maxY = y; }
                    if z < minZ { minZ = z; }
                    if z > maxZ { maxZ = z; }
                    if w < minW { minW = w; }
                    if w > maxW { maxW = w; }
                }
            }
        }
    }

    return minX, maxX, minY, maxY, minZ, maxZ, minW, maxW;
}


func cubesCycle(cubes map[int] map[int] map[int] byte) {
    //Copy previous instant of cubes:
    prevCubes := copyCubes(cubes);

    minX, maxX, minY, maxY, minZ, maxZ := cubesMaxDimensions(cubes);

    //Iterate over each cube in 'prevCubes' and set new value to 'cubes':
    for x:=minX-1; x <= maxX+1; x++ {
        for y:=minY-1; y <= maxY+1; y++ {
            for z:=minZ-1; z <= maxZ+1; z++ {
                ac := countActiveNear(prevCubes, x, y, z);
                v := cubesGet(prevCubes, x, y, z);
                if v == '#' {
                    if ac == 2 || ac == 3 {
                        cubesSet(cubes, x, y, z, '#');
                    } else {
                        cubesSet(cubes, x, y, z, '.');
                    }
                } else {
                    if ac == 3 {
                        cubesSet(cubes, x, y, z, '#');
                    } else {
                        cubesSet(cubes, x, y, z, '.');
                    }
                }
            }
        }
    }
}



func cubesCycle4(cubes map[int] map[int] map[int] map[int] byte) {
    //Copy previous instant of cubes:
    prevCubes := copyCubes4(cubes);

    minX, maxX, minY, maxY, minZ, maxZ, minW, maxW := cubesMaxDimensions4(cubes);

    //Iterate over each cube in 'prevCubes' and set new value to 'cubes':
    for x:=minX-1; x <= maxX+1; x++ {
        for y:=minY-1; y <= maxY+1; y++ {
            for z:=minZ-1; z <= maxZ+1; z++ {
                for w:=minW-1; w <= maxW+1; w++ {
                    ac := countActiveNear4(prevCubes, x, y, z, w);
                    v := cubes4Get(prevCubes, x, y, z, w);
                    if v == '#' {
                        if ac == 2 || ac == 3 {
                            cubes4Set(cubes, x, y, z, w, '#');
                        } else {
                            cubes4Set(cubes, x, y, z, w, '.');
                        }
                    } else {
                        if ac == 3 {
                            cubes4Set(cubes, x, y, z, w, '#');
                        } else {
                            cubes4Set(cubes, x, y, z, w, '.');
                        }
                    }
                }
            }
        }
    }
}


func main() {
    fmt.Println("Advent of Code - Day17");

    //Read input file:
    file, _ := os.Open("day17_input.txt");
    scanner := bufio.NewScanner(file);
    list := make([]string, 0);
    for scanner.Scan() {
        line := scanner.Text();
        list = append(list, line);
    }

    var cubes map[int] map[int] map[int] byte = make(map[int] map[int] map[int] byte);
    var cubes4 map[int] map[int] map[int] map[int] byte = make(map[int] map[int] map[int] map[int] byte);

    //Init from file for part1:
    for y:=0; y < len(list); y++ {
        for x:=0; x < len(list[y]); x++ {
            z := 0;
            cubesSet(cubes, x, y, z, list[y][x]);
        }
    }

    //Init from file for part2:
    for y:=0; y < len(list); y++ {
        for x:=0; x < len(list[y]); x++ {
            z := 0;
            w := 0;
            cubes4Set(cubes4, x, y, z, w, list[y][x]);
        }
    }

    //Part1:
    for i:=0; i < 6; i++ {
        cubesCycle(cubes);
        fmt.Printf("Cycle N.%d\n", i+1);
        fmt.Printf("Active total: %d\n", countActiveTotal(cubes));
    }

    fmt.Println("--------Part2 in 4 dimensions: ");

    //Part2:
    for i:=0; i < 6; i++ {
        cubesCycle4(cubes4);
        fmt.Printf("Cycle N.%d\n", i+1);
        fmt.Printf("Active total: %d\n", countActiveTotal4(cubes4));
    }
}




