package movie

import "fmt"

func init() {
	fmt.Println("init: movie")
}

func ReviewMovie(name string, rating float64) {
	fmt.Printf("!!! I reviewed %s and it's rating is %f\n", name, rating)
}
