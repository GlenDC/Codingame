package main

import (
    "fmt"
)

type position struct {
    x, y int   
}

func main() {
    var target, thor position
    fmt.Scanf("%d %d %d %d\n", &target.x, &target.y, &thor.x, &thor.y)

    for {
        var direction string
        var difference position

        difference.y = target.y - thor.y
        difference.x = target.x - thor.x

        switch  {
         case difference.y < 0:
            direction += "N"
            thor.y--
         case difference.y > 0:
            direction += "S"
            thor.y++
        }

        switch  {
         case difference.x > 0:
            direction += "E"
            thor.x--
         case difference.x < 0:
            direction += "W"
            thor.x++
        }
        
        fmt.Println(direction)
    }
}