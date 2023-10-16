package main

import (
	"fmt"
	"time"
)

func ping(c chan string) {
    for {
        c <- "Ping"
    }
}

func pong(c chan string) {
    for {
        c <- "Pong"
    }
}

func print(c chan string) {
    for {
        msg := <- c
        fmt.Println(msg)
        time.Sleep(time.Second)
    }
}

func main() {

    var c chan string = make(chan string)

    go ping(c)
    go pong(c)
    go print(c)

    var input string
    fmt.Scanln(&input)

}


