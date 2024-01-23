package go_calc

func Penambahan(bil1, bil2 int32) int32 {
	return bil1 + bil2
}

func Pengurangan(bil1, bil2 int32) int32 {
	return bil1 - bil2
}

func Perkalian(bil1, bil2 int32) int32 {
	return bil1 * bil2
}

func Pembagian(bil1, bil2 float32) float32 {
	return bil1 / bil2
}

func Modulo(bil1, bil2 int32) int32 {
	return bil1 % bil2
}

func Test(name string) (string, error) {
	return "Halo " + name, nil
}
