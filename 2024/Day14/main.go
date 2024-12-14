package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	//"time"

	"github.com/gdamore/tcell/v2"
)
//Press n to go to the next Second
//Press p to go to the previous Second
//For part 2 set time to 7000
const (
    HEIGHT  =103
    WIDTH   = 101
)
func draw(s tcell.Screen, arr [HEIGHT][WIDTH]int){
    s.Clear()
    for i,row := range arr{
        for j,val := range row{
            var ch rune
            if val == 0{
                ch = '.'
            }else{
                ch = '@' 
            }
            s.SetContent(i,j,ch,nil,tcell.StyleDefault.Foreground(tcell.ColorWhite))
        }
    }
    s.Show()

}
func main(){
    times:= 7500
    //ticker := time.NewTicker(1 * time.Second/300)
    s,err := tcell.NewScreen()
    if err != nil{
        log.Fatalf("Failed to create screen: %v",err)
    }
    if err := s.Init(); err != nil{
        log.Fatalf("Failed to initialize screen: %v",err)
    }
    defer s.Fini()
    signalChan := make(chan os.Signal,1)
    signal.Notify(signalChan,syscall.SIGINT)
    /*go func(){
        for range ticker.C{
            run(times)
            draw(s,run(times))
            fmt.Println(times)
            if times == 10000{
                break
            }
            times++
        }
    }()*/
    for {
        event := s.PollEvent()
        switch ev := event.(type){
        case *tcell.EventKey:
            if ev.Key() == tcell.KeyRune && ev.Rune() == 'q'{
                log.Println("Quitting")
                s.Clear()
                s.Show()
                return
            }
            if ev.Key() == tcell.KeyRune && ev.Rune() == 'n'{
                //run(times)
                times++
                draw(s,run(times))
                fmt.Println(" Seconds: ",times)
            }
            if ev.Key() == tcell.KeyRune && ev.Rune() == 'p'{
                times--
                //run(times)
                draw(s,run(times))
                fmt.Println(" Seconds: ",times)
            }
        }
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
    fmt.Print(" Part 1 result: ",quad(&arr))//the result for part 1
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
