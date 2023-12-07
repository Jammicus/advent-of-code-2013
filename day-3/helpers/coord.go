package helpers

type Coordinates struct {
	X     int
	Y     int
	Value rune
}

type NumberCoordinates struct {
	StartIndexXAxis int
	EndIndexXAxis   int
	YAxis           int
	Value           int
}

// We assume its always left to right, not diagonal
func CalculatePerimterCoordinates(coord Coordinates) []Coordinates {
	combinedCoordinates := []Coordinates{}
	leftMostCoords := []Coordinates{}
	rightMostCoords := []Coordinates{}

	// Calculate the left and right most cords.
	// Index 1 will always be the directly left or right.
	for i := -1; i < 2; i++ {
		lCoord := Coordinates{
			X: coord.X - 1,
			Y: coord.Y - i,
		}

		rCoord := Coordinates{
			X: coord.X + 1,
			Y: coord.Y - i,
		}

		combinedCoordinates = append(combinedCoordinates, lCoord)
		combinedCoordinates = append(combinedCoordinates, rCoord)

	}

	// Above the line
	c := Coordinates{
		X: coord.X,
		Y: coord.Y - 1,
	}

	combinedCoordinates = append(combinedCoordinates, c)

	// Below the line
	c = Coordinates{
		X: coord.X,
		Y: coord.Y + 1,
	}

	combinedCoordinates = append(combinedCoordinates, c)

	combinedCoordinates = append(combinedCoordinates, leftMostCoords...)
	combinedCoordinates = append(combinedCoordinates, rightMostCoords...)

	return combinedCoordinates
}
