package main
import (
        "fmt"
        "time"
)

func show(n int, data string){
  for i :=1; i<=n; i++{ 
 fmt.Println(i ,data)
time. Sleep(1 * time.Second)
}
}   
func main(){
 go show(10, "hoge")
 show(10, "hogehoge")
}
