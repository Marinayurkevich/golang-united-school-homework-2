package square

import (
	"math"
)

type chislo int

const SidesTriangle = 3
const SidesSquare = 4
const SidesCircle = 0

func CalcSquare(SideLen float64, SidesNum chislo) float64 {
	var S float64

	if SidesNum == SidesTriangle {
		S = ((SideLen * SideLen) * math.Sqrt(3)) / 4
	} else if SidesNum == SidesSquare {
		S = math.Pow(SideLen, 2)
	} else if SidesNum == SidesCircle {
		S = math.Pi * math.Pow(SideLen, 2)
	}
	return S
}
