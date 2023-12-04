package helpers

type Cordinates struct {
	X int
	Y int
}

// We assume its always left to right, not diagonal
func CalculatePerimterCordinates(startIndex, endIndex Cordinates) []Cordinates {
	combinedCordinates := []Cordinates{}
	leftMostCords := []Cordinates{}
	rightMostCords := []Cordinates{}

	// Calculate the left and right most cords.
	// Index 1 will always be the directly left or right.
	for i := -1; i < 2; i++ {
		lCoord := Cordinates{
			X: startIndex.X - 1,
			Y: startIndex.Y - i,
		}

		rCoord := Cordinates{
			X: endIndex.X + 1,
			Y: startIndex.Y - i,
		}

		leftMostCords = append(leftMostCords, lCoord)
		rightMostCords = append(rightMostCords, rCoord)

	}

	// Above the line
	for i := leftMostCords[1].X; i <= rightMostCords[1].X; i++ {
		c := Cordinates{
			X: i,
			Y: startIndex.Y - 1,
		}

		combinedCordinates = append(combinedCordinates, c)
	}

	// Below the line
	for i := leftMostCords[1].X; i <= rightMostCords[1].X; i++ {
		c := Cordinates{
			X: i,
			Y: startIndex.Y + 1,
		}

		combinedCordinates = append(combinedCordinates, c)
	}

	combinedCordinates = append(combinedCordinates, leftMostCords...)
	combinedCordinates = append(combinedCordinates, rightMostCords...)

	return combinedCordinates
}
