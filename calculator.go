package go_calc

func penambahan(bil1, bil2 int32) int32 {
	return bil1 + bil2
}

func pengurangan(bil1, bil2 int32) int32 {
	return bil1 - bil2
}

func perkalian(bil1, bil2 int32) int32 {
	return bil1 * bil2
}

func pembagian(bil1, bil2 float32) float32 {
	return bil1 / bil2
}

func modulo(bil1, bil2 int32) int32 {
	return bil1 % bil2
}

func test(name string) (string, error) {
	return "Halo " + name, nil
}
