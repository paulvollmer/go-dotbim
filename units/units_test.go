package units

import "testing"

func TestConversions(t *testing.T) {
	tests := []struct {
		name string
		got  float64
		want float64
	}{
		{name: "millimeter", got: Millimeter(200), want: 0.2},
		{name: "millimeter to one meter", got: Millimeter(1000), want: 1},
		{name: "centimeter", got: Centimeter(20), want: 0.2},
		{name: "centimeter to one meter", got: Centimeter(100), want: 1},
		{name: "decimeter", got: Decimeter(5), want: 0.5},
		{name: "decimeter to one meter", got: Decimeter(10), want: 1},
		{name: "meter", got: Meter(4), want: 4},
		{name: "kilometer", got: Kilometer(2), want: 2000},
		{name: "inch", got: Inch(1), want: 0.0254},
		{name: "foot", got: Foot(1), want: 0.3048},
		{name: "yard", got: Yard(1), want: 0.9144},
		{name: "mile", got: Mile(1), want: 1609.344},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Fatalf("got %v, want %v", tt.got, tt.want)
			}
		})
	}
}

func TestShortNamesMatchLongNames(t *testing.T) {
	if MM(200) != Millimeter(200) {
		t.Fatalf("MM(200) = %v, want Millimeter(200) = %v", MM(200), Millimeter(200))
	}
	if CM(20) != Centimeter(20) {
		t.Fatalf("CM(20) = %v, want Centimeter(20) = %v", CM(20), Centimeter(20))
	}
	if M(4) != Meter(4) {
		t.Fatalf("M(4) = %v, want Meter(4) = %v", M(4), Meter(4))
	}
	if DM(5) != Decimeter(5) {
		t.Fatalf("DM(5) = %v, want Decimeter(5) = %v", DM(5), Decimeter(5))
	}
	if KM(2) != Kilometer(2) {
		t.Fatalf("KM(2) = %v, want Kilometer(2) = %v", KM(2), Kilometer(2))
	}
	if FT(3) != Foot(3) {
		t.Fatalf("FT(3) = %v, want Foot(3) = %v", FT(3), Foot(3))
	}
	if YD(2) != Yard(2) {
		t.Fatalf("YD(2) = %v, want Yard(2) = %v", YD(2), Yard(2))
	}
	if MI(1) != Mile(1) {
		t.Fatalf("MI(1) = %v, want Mile(1) = %v", MI(1), Mile(1))
	}
}
