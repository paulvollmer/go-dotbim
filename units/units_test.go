package units_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/paulvollmer/go-dotbim/units"
)

func TestConversions(t *testing.T) {
	tests := []struct {
		name string
		got  float64
		want float64
	}{
		{name: "millimeter", got: units.Millimeter(200), want: 0.2},
		{name: "millimeter to one meter", got: units.Millimeter(1000), want: 1},
		{name: "millimeter zero", got: units.Millimeter(0), want: 0},
		{name: "centimeter", got: units.Centimeter(20), want: 0.2},
		{name: "centimeter to one meter", got: units.Centimeter(100), want: 1},
		{name: "decimeter", got: units.Decimeter(5), want: 0.5},
		{name: "decimeter to one meter", got: units.Decimeter(10), want: 1},
		{name: "meter", got: units.Meter(4), want: 4},
		{name: "kilometer", got: units.Kilometer(2), want: 2000},
		{name: "inch", got: units.Inch(1), want: 0.0254},
		{name: "negative inch", got: units.Inch(-1), want: -0.0254},
		{name: "foot", got: units.Foot(1), want: 0.3048},
		{name: "yard", got: units.Yard(1), want: 0.9144},
		{name: "mile", got: units.Mile(1), want: 1609.344},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.got)
		})
	}
}

func TestShortNamesMatchLongNames(t *testing.T) {
	tests := []struct {
		name  string
		short float64
		long  float64
	}{
		{name: "mm", short: units.MM(200), long: units.Millimeter(200)},
		{name: "cm", short: units.CM(20), long: units.Centimeter(20)},
		{name: "dm", short: units.DM(5), long: units.Decimeter(5)},
		{name: "m", short: units.M(4), long: units.Meter(4)},
		{name: "km", short: units.KM(2), long: units.Kilometer(2)},
		{name: "ft", short: units.FT(3), long: units.Foot(3)},
		{name: "yd", short: units.YD(2), long: units.Yard(2)},
		{name: "mi", short: units.MI(1), long: units.Mile(1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.long, tt.short)
		})
	}
}
