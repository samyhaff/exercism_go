package space

type Planet string

func Age(s float64, planet Planet) float64 {
	switch planet {
	case "Earth":
		return s / 31557600
	case "Mercury":
		return Age(s, "Earth") / 0.2408467
	case "Venus":
		return Age(s, "Earth") / 0.61519726
	case "Mars":
		return Age(s, "Earth") / 1.8808158
	case "Jupiter":
		return Age(s, "Earth") / 11.862615
	case "Saturn":
		return Age(s, "Earth") / 29.447498
	case "Uranus":
		return Age(s, "Earth") / 84.016846
	case "Neptune":
		return Age(s, "Earth") / 164.79132
	default:
		return -1
	}
}
