package vector

import (
	"math"
	"testing"
)

func TestNewVec3(t *testing.T) {
	t.Run("(1,2,3)", func(t *testing.T) {
		v := NewVec3(1.0, 2.0, 3.0)

		if v.X != 1.0 {
			t.Errorf("X should be %f, but was %f", 1.0, v.X)
		}
		if v.Y != 2.0 {
			t.Errorf("Y should be %f, but was %f", 2.0, v.X)
		}
		if v.Z != 3.0 {
			t.Errorf("Z should be %f, but was %f", 3.0, v.X)
		}
	})
}

func TestLength(t *testing.T) {
	t.Run("(0,0,0)", func(t *testing.T) {
		// Given
		v := NewVec3(0.0, 0.0, 0.0)
		expected := 0.0

		// When
		result := Length(v)

		if result != expected {
			t.Errorf("Length of %v should be %f, but is %f", v, expected, result)
		}
	})

	t.Run("(1,1,1)", func(t *testing.T) {
		// Given
		v := NewVec3(1.0, 1.0, 1.0)
		expected := math.Sqrt(3.0)

		// When
		result := Length(v)

		if result != expected {
			t.Errorf("Length of %v should be %f, but is %f", v, expected, result)
		}
	})

	t.Run("(5,5,5)", func(t *testing.T) {
		// Given
		v := NewVec3(5.0, 5.0, 5.0)
		expected := 5.0 * math.Sqrt(3.0)

		// When
		result := Length(v)

		if result != expected {
			t.Errorf("Length of %v should be %f, but is %f", v, expected, result)
		}
	})
}
