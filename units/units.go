// Package units converts lengths into meters so they can be used as .bim mesh
// coordinates. The meter is the default unit.
package units

// Millimeter converts millimeters to meters.
func Millimeter(millimeters float64) float64 {
	return millimeters / 1000
}

// MM converts millimeters to meters.
func MM(millimeters float64) float64 {
	return Millimeter(millimeters)
}

// Centimeter converts centimeters to meters.
func Centimeter(centimeters float64) float64 {
	return centimeters / 100
}

// CM converts centimeters to meters.
func CM(centimeters float64) float64 {
	return Centimeter(centimeters)
}

// Decimeter converts decimeters to meters.
func Decimeter(decimeters float64) float64 {
	return decimeters / 10
}

// DM converts decimeters to meters.
func DM(decimeters float64) float64 {
	return Decimeter(decimeters)
}

// Meter returns meters. Meter is the default .bim coordinate unit, so the
// value is unchanged.
func Meter(meters float64) float64 {
	return meters
}

// M returns meters.
func M(meters float64) float64 {
	return Meter(meters)
}

// Kilometer converts kilometers to meters.
func Kilometer(kilometers float64) float64 {
	return kilometers * 1000
}

// KM converts kilometers to meters.
func KM(kilometers float64) float64 {
	return Kilometer(kilometers)
}

// Inch converts inches to meters. One inch is 0.0254 meters.
func Inch(inches float64) float64 {
	return inches * 0.0254
}

// Foot converts feet to meters. One foot is 0.3048 meters.
func Foot(feet float64) float64 {
	return feet * 0.3048
}

// FT converts feet to meters.
func FT(feet float64) float64 {
	return Foot(feet)
}

// Yard converts yards to meters. One yard is 0.9144 meters.
func Yard(yards float64) float64 {
	return yards * 0.9144
}

// YD converts yards to meters.
func YD(yards float64) float64 {
	return Yard(yards)
}

// Mile converts miles to meters. One mile is 1609.344 meters.
func Mile(miles float64) float64 {
	return miles * 1609.344
}

// MI converts miles to meters.
func MI(miles float64) float64 {
	return Mile(miles)
}
