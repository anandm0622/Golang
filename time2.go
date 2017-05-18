package main
import "fmt"

var c = fmt.Scan
var p = fmt.Println

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var i,n,k,q,temp int
    c(&n,&k,&q)       
    a := make([]int, n)
    b := make([]int, n)
    m := make([]int, q)
    for i=0; i<n; i++ {
        c(&a[i])
    }    
    for i=0; i<q; i++ {
        c(&m[i])       
    }
    for i=0; i<n; i++ {
       temp = i+k
        if(temp >= n){
             temp = temp % n
        } 
       b[temp] = a[i]
    } 
    for i=0; i<q; i++ {
        p(b[m[i]])       
    }          
}