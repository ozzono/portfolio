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
	r, err := NewRover("test1", "N", 1, 2)
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
	r, err := NewRover("test2", "E", 3, 3)
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
	r1, err := NewRover("r1", "N", 1, 1)
	assertNoErr(t, err)
	r2, err := NewRover("r2", "N", 1, 2)
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
	r1, err := NewRover("r1", "N", 0, 0)
	assertNoErr(t, err)
	r2, err := NewRover("r2", "E", 0, 5)
	assertNoErr(t, err)
	r3, err := NewRover("r3", "S", 5, 5)
	assertNoErr(t, err)
	r4, err := NewRover("r4", "W", 5, 0)
	assertNoErr(t, err)

	for i, r := range []*Rover{r1, r2, r3, r4} {
		t.Logf("[%d] rover test %s", i, r.Label)
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
