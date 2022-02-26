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
		S = ((SideLen * SideLen) * Sqrt(3)) / 4
	} else if SidesNum == SidesSquare {
		S = SideLen * SideLen
	} else if SidesNum == SidesCircle {
		S = math.Pi * (SideLen * SideLen)
	}
	return S
}
