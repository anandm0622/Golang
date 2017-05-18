package main
import "fmt"

var c = fmt.Scan
var p = fmt.Println

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var i,n,k,q,temp int
    c(&n,&k,&q)        
    a := make([]int, n)
    m := make([]int, q)
    for i=0; i<n; i++ {
        c(&a[i])
    }      
    for i=0; i<q; i++ {
        c(&m[i])
        temp = m[i] + k + 1
        if(temp > n){
            temp = temp % n
             m[i] = temp
        } else if(temp == n){
            temp = temp - k
            temp += 1
             m[i] = temp-1
        }           
   } 
    p(m)
    for i=0; i<q; i++ {
    p(a[n-1-m[i]])
    }     
}