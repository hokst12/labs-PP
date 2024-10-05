package mathutils

func Fact(a int) int {
	var b int = 1
	for i := 1; i <= a; i++ {
		b *= i
	}
	return b

}
