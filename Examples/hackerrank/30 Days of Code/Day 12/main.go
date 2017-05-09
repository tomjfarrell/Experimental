package main
import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0 ; i < n ; i++ {
		res := "NO"
		m := make(map[rune]bool)
		var short,long string
		fmt.Scan(&short,&long)
		if len(long) < len(short) {
			short,long = long,short
		}
		for i,c := range short {
			m[c] = true
			if m[rune(long[i])] {
				res = "YES"
			}
		}
		fmt.Println(res)
	}
}