package main
import "fmt"

func main(){
var animals = map[string] int{
"dog" : 1,
"cat" : 2,
}
var foo = animals["dog"]
var bar = animals["cat"]

fmt.Printf("%d\n",foo)
fmt.Printf("%d\n",bar)

}


