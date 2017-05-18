package main
import ("fmt" )
   
    var p = fmt.Println
    var c = fmt.Scan
    var a = [3]int{0,1,1}  
   	var flag bool 
   	var y int = 0;
    func fibo(n int) int{
	 
	if(n < len(a)){
		return a[n]
	} else{
		if(flag){
			y = n;
			flag = false
		}
		//p(len(a))
		a := make([]int, y+1)
		//p(len(a))
		a[n] = fibo(n-1) + fibo(n-2)
		return a[n]
	} 
}

func main() {
	var n,x int	
	for (true) {
		flag = true
		c(&n)
	x = fibo(n)
	p(x)
}
}
