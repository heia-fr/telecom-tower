package ledmatrix

import (
	"testing"
)

func TestBasic(t *testing.T) {
	m := NewMatrix(8, 0)
	m.SetPixel(4, 4, White)
	if m.columns != 5 {
		t.Errorf("Columns shoud now be 5 instead of %d", m.rows)
	}
	m.SetPixel(2, 2, Red)
	if m.columns != 5 {
		t.Errorf("Columns shoud still be 5 instead of %d", m.rows)
	}

	if m.GetPixel(2, 2) != Red {
		t.Error("Pixel shoud be red")
	}
	if m.GetPixel(3, 3) != Black {
		t.Error("Pixel shoud be black")
	}
	if m.GetPixel(4, 4) != White {
		t.Error("Pixel shoud be white")
	}
}

func TestSlice(t *testing.T) {
	m := NewMatrix(8, 10)
	m.SetPixel(4, 4, White)
	m.SetPixel(5, 4, Red)
	m.SetPixel(6, 4, Red)
	m.SetPixel(7, 4, White)

	n := m.Slice(4, 8)
	if n.rows != m.rows {
		t.Error("Number of rows is not the same")
	}
	if n.columns != 4 {
		t.Errorf("Number of columns is %d instead of 4", n.columns)
	}

	if n.GetPixel(0, 4) != White {
		t.Error("Pixel shoud be white")
	}
	if n.GetPixel(1, 4) != Red {
		t.Error("Pixel shoud be red")
	}
	if n.GetPixel(2, 4) != Red {
		t.Error("Pixel shoud be red")
	}
	if n.GetPixel(3, 4) != White {
		t.Error("Pixel shoud be white")
	}

	if n.GetPixel(0, 0) != Black {
		t.Error("Pixel shoud be black")
	}
	if n.GetPixel(3, 5) != Black {
		t.Error("Pixel shoud be black")
	}

}
