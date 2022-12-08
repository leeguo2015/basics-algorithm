package squareIsWhite

func squareIsWhite(coordinates string) bool {
	if (coordinates[0]+coordinates[1])%2 == 0 {
		return false
	}
	return true
}
