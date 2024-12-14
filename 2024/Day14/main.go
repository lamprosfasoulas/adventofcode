package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)
const (
    HEIGHT  =103
    WIDTH   = 101
)

func main(){
    times:= 0
    ticker := time.NewTicker(10 * time.Millisecond)
    for range ticker.C{
        //fmt.Println(times)
        run(times)
        if times == 10000{
            break
        }
        times++
    }
}
func run(sec int)[HEIGHT][WIDTH]int{
    var arr [HEIGHT][WIDTH]int
    file,_ := os.Open("day14")
    defer file.Close()

    sc := bufio.NewScanner(file)

    for sc.Scan(){
        line := strings.Fields(sc.Text())
        pos := strings.Split(line[0],"=")
        vel := strings.Split(line[1],"=")
        y := strings.Split(pos[1], ",")[0]
        x := strings.Split(pos[1], ",")[1]
        vy := strings.Split(vel[1], ",")[0]
        vx := strings.Split(vel[1], ",")[1]
        i,j := test(x,y,vx,vy,sec)
        arr[i][j] +=1
    }
    quad(&arr)
    return arr
    //fmt.Println(theone)
    //fmt.Println(result)
}
func quad(arr *[HEIGHT][WIDTH]int)int{
    var quad1,quad2,quad3,quad4 int

    for i:=0;i<HEIGHT/2;i++{
        for j:=0;j<WIDTH/2;j++{
            quad1 += (*arr)[i][j]
        }
    }
    for i:=0;i<HEIGHT/2;i++{
        for j:=WIDTH/2+1;j<WIDTH;j++{
            quad2 += (*arr)[i][j]
        }
    }
    for i:=HEIGHT/2+1;i<HEIGHT;i++{
        for j:=0;j<WIDTH/2;j++{
            quad3 += (*arr)[i][j]
        }
    }
    for i:=HEIGHT/2+1;i<HEIGHT;i++{
        for j:=WIDTH/2+1;j<WIDTH;j++{
            quad4 += (*arr)[i][j]
        }
    }
    return quad1*quad2*quad3*quad4
}
func test(xs,ys,vxs,vys string,sec int)(int,int){
    x,_ := strconv.Atoi(xs)
    y,_ := strconv.Atoi(ys)
    vx,_ := strconv.Atoi(vxs)
    vy,_ := strconv.Atoi(vys)
    for range sec{
        x += vx
        y += vy
    }
    newlocx,newlocy := x%HEIGHT,y%WIDTH
    if newlocx < 0{
        newlocx += HEIGHT
    }
    if newlocy < 0{
        newlocy += WIDTH 
    }
    return newlocx,newlocy
}
