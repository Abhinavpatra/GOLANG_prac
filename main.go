package main

import (
	"fmt"
	// "time"
)

func main(){
    var a[5] int
    fmt.Println("empty:", a)
    a[0]=1
    fmt.Println("first thing changed", a)
    fmt.Println("set:", a)
    fmt.Println("You have to do implicit type conversions in order to assign values to an array")
    fmt.Println("\nLength function like this: len(a): ",len(a))
    fmt.Println()
    b := [...]int {1, 2, 3, 4, 5}
    fmt.Println("declaration:", b)

    //You can also have the compiler count the number of elements for you with ...
    b = [...]int{1, 2, 6, 4, 5}

    fmt.Println("dcl:", b)
    fmt.Println()
    var twoD = [2][3]int{
        {1, 2, 3},
        {1, 2, 3},
    }// multi dimensional array
    fmt.Println(twoD)
}





// func main(){
//     i, j := 42, 2000
//     p := &i
//     q := &j// the type of q is Pointers, and similar for p
//     fmt.Println(p)// will print the address it is pointing to
//     fmt.Println(q)
//     fmt.Println(*p) 
//     fmt.Println(*q) 
//     p= &j
//     fmt.Println(i)
//     fmt.Println(j)
//     // * -> used to return the value
//     // & -> returns you the address 
// }



// func main(){
//     fmt.Println("starting to count")
//     for i := 0; i < 10; i++{
//         defer fmt.Print(i)
//     }
//     fmt.Println("counting done")
// }


// switch

// func main(){
//     i:= 3;
//     switch i{
//     case 1: 
//         fmt.Println("number is 1")
//     case 2: 
//         fmt.Println("number is 2")
//     case 4: 
//         fmt.Println("number is 4")
//     case 5: 
//         fmt.Println("number is 5")
//     default:
//         fmt.Println("Unknown number ")
//     };
//     t := time.Now()
//     switch {
//     case t.Hour() < 16:
//         fmt.Println("It's before noon")
//     default:
//         fmt.Println("It's after noon")
//     }
// }


// if, if-else

// func main(){
//     var count int = 0
//     if "a"=="a"{
//         count++;
//     }
//     if "b"=="b"{
//         count++;
//     }
//     if "c"=="c"{
//         count++;
//     }
//     fmt.Println(count)
// }


// for loop
// func main(){
    
//     i:= 0
//     for i < 1000 {
//         i+=i
//         i++;
//         fmt.Println(i);
//     }
// }



// func main(){
//     const anagram string = "adam smith"
//     const adam = "blah"
//     var f float32 = math.Pi;
//     var implicit int =int(f) 
//     fmt.Printf("Previously it was %f , now it is:%d:  %s %s %d", f , implicit, anagram, adam, adam)
// }





// Variable declaration with data types

// var c, py, ja, some bool
// // k := 9// error
// func main(){
// 	var i int
// 	fmt.Println(c, py, ja, some, i) 
// 	l := 9
//     fmt.Printf("%d blah blah blah", l)
// }




//

// func split(sum int)(x, y int){
// 	x = sum * 4/9;
// 	y = sum - x;
// 	return
// }

// func main(){
// 	fmt.Println(split(17))
// }









// Functions
// func swap(x, y string)(string, string){
// 	return y, x

// }
// func main(){

// 	a,b  := swap("Hello", "World")
// 	fmt.Println(a,b);
// }




// //math
// func add(x, y int ) int{
// 	return x+y;

// }
// func main(){
// 	fmt.Println(add(30,20))

// }




// VALUES

// func main() {
// 	fmt.Println("go attached with " + "lang")
// 	fmt.Println("1 + 1 =",1+1 )
// 	fmt.Println("7.0/3.0 =", 7.0/3.0) // float division
// 	fmt.Println(!false)
// 	fmt.Println(false)
// 	fmt.Println(true)
// }
