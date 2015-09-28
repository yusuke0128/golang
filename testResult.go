package main
import "fmt"
func main(){
type result struct{
math int
english int
}
var takeshi result
var yusuke result
fmt.Printf("たけしくんの数学の点数は何点ですか？\n")
fmt.Scanf("%v",&takeshi.math)
fmt.Printf("たけしくんの英語の点数は何点ですか？\n")
fmt.Scanf("%v",&takeshi.english)
fmt.Printf("ゆうすけくんの数学の点数は何点ですか？\n")
fmt.Scanf("%v",&yusuke.math)
fmt.Printf("ゆうすけくんの英語の点数は何点ですか？\n")
fmt.Scanf("%v",&yusuke.english)
fmt.Printf("たけしくんの数学の点数は%d点\n",takeshi.math)
fmt.Printf("たけしくんの英語の点数は%d点\n",takeshi.english)
fmt.Printf("ゆうすけくんの数学の点数は%d点\n",takeshi.math)
fmt.Printf("ゆうすけくんの英語の点数は%d点\n",takeshi.english)
return 
}
