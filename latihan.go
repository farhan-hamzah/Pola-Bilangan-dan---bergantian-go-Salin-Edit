package main
import "fmt"

func bil(n int) {
	if n == 0 {
		return
	}

	bil(n - 1)

	if n%2 != 0 {
		fmt.Print(n)
	} else {
		fmt.Print(-n)
	}

	if n != 0 {
		fmt.Print(", ")
	}
}

func main() {
	var bilangan int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&bilangan)
	bil(bilangan)
}