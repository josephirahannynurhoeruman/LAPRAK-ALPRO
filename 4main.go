package main
import "fmt"
func main() {
	var fahrenheit, celcius float64
	fmt.Print("Masukkan suhu Fahrenheit: ")
	fmt.Scan(&fahrenheit)
	celcius = (fahrenheit - 32) * 5 / 9
	fmt.Println("Suhu Celcius =", celcius)
}