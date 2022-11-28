// Golang program to store information
// about games in structs and display them
package main

import "fmt"

//Annonimaus attribute
type basedetail struct {
	name string
}

// We create a struct details to hold
// generic information about games
type details struct {
	genre       string
	genreRating string
	reviews     string
}

// We create a struct game to hold
// more specific information about
// a particular game
type game struct {
	// name  string
	basedetail
	price string
	// We use composition through
	// embedding to add the
	// fields of the details
	// struct to the game struct
	action details
}

// this is a method defined
// on the details struct
func (d details) showDetails() {
	fmt.Println("Genre:", d.genre)
	fmt.Println("Genre Rating:", d.genreRating)
	fmt.Println("Reviews:", d.reviews)
}

// this is a method defined
// on the game struct
// this method has access
// to showDetails() as well since
// the game struct is composed
// of the details struct
func (g game) show() {
	fmt.Println("Name: ", g.name) // as this is annonimaus attribute game.name can be found by compiler
	fmt.Println("Price:", g.price)
	g.action.showDetails() // as this is not annonimaus attribute we need atribute name . method name.
}

func main() {

	// defining a struct
	// object of Type details
	action := details{"Action", "18+", "mostly positive"}

	// defining a struct
	// object of Type game
	newGame := game{basedetail: basedetail{"XYZ"}, price: "$125", action: action}

	newGame.show()
}
