package controller

import (
	"errors"
	"testing"
)

const (
	testWidth  = 5
	testHeight = 5
)

func TestSample1(t *testing.T) {
	// 5 5
	// 1 2 N
	// LMLMLMLMM -> LFLFLFLFF
	setTestPlateau(t)
	r, err := NewRover(1, 2, "N")
	assertNoErr(t, err)
	err = Plateau.AddRover(r)
	assertNoErr(t, err)
	for _, m := range "LFLFLFLFF" {
		err = r.Move(string(m))
		assertNoErr(t, err)
	}
	// 1 3 N
	if r.Coords.X != 1 {
		t.Logf("x: %d!=1", r.Coords.X)
		t.FailNow()
	}
	if r.Coords.Y != 3 {
		t.Logf("y: %d!=3", r.Coords.Y)
		t.FailNow()
	}
	if r.Direction != "N" {
		t.Logf("direction: %s!=N", r.Direction)
		t.FailNow()
	}

	t.Log(r.LogTxt())
}

func TestSample2(t *testing.T) {
	// 3 3 E
	// MMRMMRMRRM -> FFRFFRFRRF
	setTestPlateau(t)
	r, err := NewRover(3, 3, "E")
	assertNoErr(t, err)
	err = Plateau.AddRover(r)
	assertNoErr(t, err)
	for _, m := range "FFRFFRFRRF" {
		err = r.Move(string(m))

		assertNoErr(t, err)
	}
	// 5 1 E
	if r.Coords.X != 5 {
		t.Logf("x: %d!=5", r.Coords.X)
		t.FailNow()
	}
	if r.Coords.Y != 1 {
		t.Logf("y: %d!=1", r.Coords.Y)
		t.FailNow()
	}
	if r.Direction != "E" {
		t.Logf("direction: %s!=E", r.Direction)
		t.FailNow()
	}

	t.Log(r.LogTxt())
}

func TestSameCoordRover(t *testing.T) {
	setTestPlateau(t)
	r1, err := NewRover(1, 1, "N")
	assertNoErr(t, err)
	r2, err := NewRover(1, 2, "N")
	assertNoErr(t, err)
	err = Plateau.AddRover(r1)
	assertNoErr(t, err)
	err = Plateau.AddRover(r2)
	assertNoErr(t, err)

	err = r1.Move("F")
	assertErr(t, err)
}

func TestOutOfBounds(t *testing.T) {
	setTestPlateau(t)
	r1, err := NewRover(0, 0, "N")
	assertNoErr(t, err)
	r2, err := NewRover(0, 5, "E")
	assertNoErr(t, err)
	r3, err := NewRover(5, 5, "S")
	assertNoErr(t, err)
	r4, err := NewRover(5, 0, "W")
	assertNoErr(t, err)

	for _, r := range []*Rover{r1, r2, r3, r4} {
		t.Logf("[%d] rover test", r.ID)
		err = r.Move("B")
		if !errors.Is(err, ErrUnavailableCoord) {
			t.Logf("expecting -- %s err", ErrUnavailableCoord)
			t.Logf("outputed --- %s err", err)
			t.Fail()
		}
		assertErr(t, err)
	}
}

// expecting non-nil error
func assertErr(t *testing.T, err error, msg ...string) {
	if err == nil {
		t.Log("err should be non-nil", err)
		for i := range msg {
			t.Log(msg[i])
		}
		t.FailNow()
	}
}

// expecting nil error
func assertNoErr(t *testing.T, err error, msg ...string) {
	if err != nil {
		t.Log("err should be nil", err)
		for i := range msg {
			t.Log(msg[i])
		}
		t.FailNow()
	}
}

func setTestPlateau(t *testing.T) {
	t.Log("reseting plateau")
	SetPlateau(testWidth, testHeight)
}

func TestShowPlateau(t *testing.T) {
	setTestPlateau(t)
	emptyPlateau := `| 5|  |  |  |  |  |  |
| 4|  |  |  |  |  |  |
| 3|  |  |  |  |  |  |
| 2|  |  |  |  |  |  |
| 1|  |  |  |  |  |  |
| 0|  |  |  |  |  |  |
| X| 0| 1| 2| 3| 4| 5|`
	p, err := Plateau.Show()
	assertNoErr(t, err)
	if p != emptyPlateau {
		t.Logf("expected output: \n%s", emptyPlateau)
		t.Logf("     output: \n%s", p)
		t.Fail()
	}
	r1, err := NewRover(1, 1, "N")
	assertNoErr(t, err)
	r2, err := NewRover(1, 2, "S")
	assertNoErr(t, err)
	err = Plateau.AddRover(r1)
	assertNoErr(t, err)
	err = Plateau.AddRover(r2)
	assertNoErr(t, err)

	roveredPlateau := `| 5|  |  |  |  |  |  |
| 4|  |  |  |  |  |  |
| 3|  |  |  |  |  |  |
| 2|  |\/|  |  |  |  |
| 1|  |/\|  |  |  |  |
| 0|  |  |  |  |  |  |
| X| 0| 1| 2| 3| 4| 5|`
	p, err = Plateau.Show()
	assertNoErr(t, err)
	if p != roveredPlateau {
		t.Logf("expected output: \n%s", roveredPlateau)
		t.Logf("         output: \n%s", p)
		t.Fail()
	}

}
