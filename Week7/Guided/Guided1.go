package main
import "fmt"
type bilangan int
type pecahan float64
func main(){
var a,b bilangan
var hasil pecahan
a = 1
b = 9
hasil = pecahan(a) / pecahan(b)
fmt.Println(hasil)
}