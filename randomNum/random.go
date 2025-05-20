package randomnum

import (
	"fmt"
	"math/rand"
)

func printRandomNumber() {
	fmt.Println("My fav number is ",rand.Intn(20) )
}