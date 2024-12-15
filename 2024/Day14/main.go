package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

    tea "github.com/charmbracelet/bubbletea"
)

//Press n to go to the next Second
//Press p to go to the previous Second
//For part 2 set time to 7000
const (
    HEIGHT  =103
    WIDTH   = 101
)
var times int = 7000
type model struct{
    cells [HEIGHT][WIDTH]int
    part1 int
    sec int
}
func (m model) Init()tea.Cmd{
    return nil
}
func (m model) Update(msg tea.Msg) (tea.Model,tea.Cmd){
    switch msg := msg.(type){
    case tea.KeyMsg:
        switch msg.String(){
        case "n":
            times++
            m.cells = run(times)
        case "p":
            times--
            m.cells = run(times)

        case "q":
            return m,tea.Quit
        }

    }
    return m,nil
}
func (m model) View()string{
    var s string
    for i:=0;i<HEIGHT;i++{
        for j:=0;j<WIDTH;j++{
            if m.cells[i][j] == 0{
                s += "."
            }else{
                s += "#"
            }
        }
        s += "\n"
    }
    if times == 100{
        s += fmt.Sprintf("First: %d\n",quad(&m.cells))
    }
    s += fmt.Sprintf("Second: %d\n",times)
    return s
}

func initialModel() model{
    return model{
        cells: [HEIGHT][WIDTH]int{}, 
        part1: 0,
        sec: 0,
    }
}

func main(){
    var com string
    fmt.Scanln(&com)
    if com == "sec"{
        start := time.Now()
        for times < 10000{
            run(times)
            times++
        }
        fmt.Println("Time: ",time.Since(start))
        fmt.Printf("The second of the tree must be: %d\n",part2)
        os.Exit(0)
    }
    p := tea.NewProgram(initialModel())
    if _,err := p.Run();err != nil{
        fmt.Println("Error: ",err)
        os.Exit(1)
    }
}
var part2,safety int
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
    if safety == 0{
        safety = quad(&arr)
        part2= sec
    }
    if quad(&arr) < safety{
        safety = quad(&arr)
        part2 = sec
    }
    if sec == 100{
        fmt.Println(" Part 1 result: ",quad(&arr))
    }
    //fmt.Print(" Part 1 result: ",quad(&arr))//the result for part 1
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
