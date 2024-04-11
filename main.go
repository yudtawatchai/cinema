package main

import (
	"fmt"

	"github.com/yudtawatchai/cinema/movie"
	"github.com/yudtawatchai/cinema/ticket"
)

func init() {
	fmt.Println("init: main")
}

func main() {
	movieName := movie.FindMovieName("tt4154796")
	fmt.Println(movieName)
	movie.ReviewMovie(movieName, 8.4)
	ticket.BuyTicket(movieName)
}
