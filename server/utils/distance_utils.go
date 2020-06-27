package utils

const metersPerMile = 1609.34
const metersPerYard = 0.9144

// DistanceToMeters converts any native unit (e.g. miles, yards) to meters
func DistanceToMeters(unitDistance float64, unit string) float64 {
	if unitDistance == 0 {
		return 0
	}
	// Switch on the unit to determine how to conver to meters
	switch unit {
	case "miles":
		return unitDistance * metersPerMile
	case "yards":
		return unitDistance * metersPerYard
	default:
		// TODO LOG Fatal?
		return 0
	}
}
