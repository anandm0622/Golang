package main
import ("fmt"
"net/http")

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    http.HandleFunc("/", handler);
    http.HandleFunc("/earth", handler2);
    http.ListenAndServe(":8080", nil);
    
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World")
}
func handler2(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello Earth")
}
