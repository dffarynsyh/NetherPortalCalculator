package main
import ("fmt")

func nether() {
	var x, z int
	var x2, z2 int
	fmt.Println("Masukkan koordinat X:")
	fmt.Scanln(&x)
	fmt.Println("Masukkan koordinat Z:")
	fmt.Scanln(&z)
	x2 = (x / 8)
	z2 = (z / 8)
	fmt.Println("Koordinat Nether: ", "X=", x2, "Z=", z2)
}
func overworld() {
	var x, z int
	var x2, z2 int
	fmt.Println("Masukkan koordinat X:")
	fmt.Scanln(&x)
	fmt.Println("Masukkan koordinat Z:")
	fmt.Scanln(&z)
	x2 = x * 8
	z2 = z * 8
	fmt.Println("Koordinat Overworld: ", "X=", x2, "Z=", z2)
}

func main() {
	var input int
	fmt.Println("Penghitung Nether portal")
	fmt.Println("ketik 1 untuk menghitung koordinat nether", "ketik 2 untuk menghitung koordinat overworld")
	fmt.Scanln(&input)
	if input == 1 {
		nether()
	} else if input == 2 {
		overworld()
	} else {
		fmt.Println("Input tidak valid")
		main()
	}
}