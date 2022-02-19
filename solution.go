package donich

import (
	"math"
)

type myType int

func CalcSquare(sideLen float64, sidesNum myType) (res float64) {

	const SidesTriangle myType = 3
	const SidesSquare myType = 4
	const SidesCircle myType = 0

	if SidesTriangle == sidesNum {
<<<<<<< HEAD
		res = (math.Pow(sideLen, 2.0) * math.Sqrt(3.0)) / 4
=======
		res = (math.Pow(sideLen, 2.0) * math.Sqrt(3.0))/4
>>>>>>> 32b5f5b403bf722d5356d1872dfbea3f351400b3
	} else if SidesSquare == sidesNum {
		res = math.Pow(sideLen, 2.0)
	} else if SidesCircle == sidesNum {
		res = math.Pi * math.Pow(sideLen, 2.0)
	}
	return
}
