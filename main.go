package main

import (
	"fmt"
	// "time"
)

func main(){
    i, j := 42, 2000
    p := &i
    q := &j
    fmt.Println(p)
    fmt.Println(q)
    fmt.Println(*p) 
    fmt.Println(*q) 
    p= &j
    fmt.Println(*p)
}



// func main(){
//     fmt.Println("starting to count")
//     for i := 0; i < 10; i++{
//         defer fmt.Print(i)
//     }
//     fmt.Println("counting done")
// }


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




// var c, py, ja, some bool
// // k := 9// error
// func main(){
// 	var i int
// 	fmt.Println(c, py, ja, some, i) 
// 	l := 9
//     fmt.Printf("%d blah blah blah", l)
// }



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
