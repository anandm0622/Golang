package main
import "fmt"

var c = fmt.Scan
var p = fmt.Println

func f(a *[3]int){
	a[0]=3
	p(a)
}

func fr(a *[3]int){
	p(a)
}

func main() {
	var arr1= new([3]int)
	arr2:= arr1
	fr(arr1)
	f(arr2)
	p(arr1)
	p(arr2)	
}

