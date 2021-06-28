package raindrops

import "fmt"

func Convert(n int) string {
	raindrop := ""
	if n%3 == 0 {
		raindrop += "Pling"
	}
	if n%5 == 0 {
		raindrop += "Plang"
	}
	if n%7 == 0 {
		raindrop += "Plong"
	}
	if n%3 != 0 && n%5 != 0 && n%7 != 0 {
		raindrop = fmt.Sprintf("%d", n)
	}
	return raindrop
}
