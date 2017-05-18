package main
import ("fmt"
        s "strings"
        "strconv"
       )
    var p = fmt.Println
    var c = fmt.Scan
func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var a,b string

    var x int
    c(&a)
    z := a[2:7]
    if(s.Contains(a,"P")){
        b = a[0:2]
    x,y := strconv.Atoi(b)
    p(x,y,b)
    x+=12
    
}
   k := strconv.Itoa(x)
    k = append(z)
    p(k)
}
