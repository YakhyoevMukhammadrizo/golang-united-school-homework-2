package square

import "math"

// Define custom int type to hold sides number and update CalcSquare
// signature by replacing #yourTypeNameHere#
type shape int

// Define constants to represent 0, 3 and 4 sides.
// Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
const SidesCircle shape = 0
const SidesTriangle shape = 3
const SidesSquare shape = 4

// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum shape) float64 {
	switch sidesNum {
	case SidesCircle:
		return math.Pi * math.Pi * sideLen
	case SidesTriangle:
		return (math.Sqrt(3.0) / 4) * math.Pow(sideLen, 2)
	case SidesSquare:
		return math.Pow(sideLen, 2)
	default:
		return 0.0
	}
}
