package main 
import "fmt"

type Friend struct{
Yusuke
Haretaro
}

type Yusuke struct{
shincyo  int  
taijyu  int 
}

type Haretaro struct{
shincyo int  
taijyu int 
}

func (y *Yusuke) printYusukeShincyo() {
fmt.Print("　ゆうすけ君の身長は",y.shincyo,"cm\n")
}

func (h *Haretaro) printHaretaroShincyo() {
fmt.Print("　はれたろう君の身長は",h.shincyo,"cm\n")
}

func (y *Yusuke) printYusukeTaijyu() {
fmt.Print("　ゆうすけ君の体重は",y.shincyo,"cm\n")
}

func (h *Haretaro) printHaretaroTaijyu() {
fmt.Print("　はれたろう君の体重は",h.taijyu,"cm\n")
}

func main(){
f := new(Friend)
f.Yusuke.shincyo = 160
f.Haretaro.shincyo = 170
f.Yusuke.taijyu = 55
f.Haretaro.taijyu = 60 
f.Yusuke.printYusukeShincyo()
f.Haretaro.printHaretaroShincyo()
f.Yusuke.printYusukeTaijyu()
f.Haretaro.printHaretaroTaijyu()
}

