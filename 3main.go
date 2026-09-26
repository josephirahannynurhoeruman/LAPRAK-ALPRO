package main
import "fmt"
func main() {
	var r, luas float64
	fmt.Print("Masukkan jari-jari: ")
	fmt.Scan(&r)
	luas = 3.14 * r * r
	fmt.Printf("Luas lingkaran = %. 1f\n", luas)
}