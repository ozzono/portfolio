package logicpractice01

import (
	"fmt"
	"testing"
)

//  - Get an empty ruleset - AddDep("a","a") - The result should be coherent
func TestDependsAA(t *testing.T) {
	rs := newRuleSet()
	a := newOption("a")
	rs.addDep(a, a)
	if err := rs.isCoherent(); err != nil {
		t.Fatalf("rule set is not coherent: %v", err)
	}
	t.Log("rule set is coherent")
}

//  - Get an empty ruleset - AddDep("a", "b") - AddDep("b", "a") - The result should be coherent
func TestDependsAB_BA(t *testing.T) {
	rs := newRuleSet()
	a, b := newOption("a"), newOption("b")
	rs.addDep(a, b)
	rs.addDep(b, a)
	if err := rs.isCoherent(); err != nil {
		t.Fatalf("rule set is not coherent: %v", err)
	}
	t.Log("rule set is coherent")
}

//  - Get an empty ruleset - AddDep("a", "b") - AddConflict("a", "b") - Be coherent should return an error
func TestExclusiveAB(t *testing.T) {
	rs := newRuleSet()
	a, b := newOption("a"), newOption("b")
	rs.addDep(a, b)
	rs.addConflict(a, b)
	err := fmt.Errorf("")
	if err = rs.isCoherent(); err == nil {
		t.Fatalf("rule set is coherent but should not")
	}
	t.Logf("rule set is non coherent as intended: %v", err)
}

//  - Get an empty ruleset - AddDep("a", "b") - AddDep("b", "c") - AddConflict("a", "c") - Be coherent should return an error
func TestExclusiveAB_BC(t *testing.T) {
	rs := newRuleSet()
	a, b, c := newOption("a"), newOption("b"), newOption("c")
	rs.addDep(a, b)
	rs.addDep(b, c)
	rs.addConflict(a, c)
	err := fmt.Errorf("")
	if err = rs.isCoherent(); err == nil {
		t.Fatalf("rule set is coherent but should not")
	}
	t.Logf("rule set is non coherent as intended: %v", err)
}

//  - Get an empty ruleset - AddDep("a", "b") - AddDep("b", "c") - AddDep("c", "d") - AddDep("d", "e") - AddDep("a", "f") - AddConflict("e", "f") - Be coherent should return an error
func TestDeepDeps(t *testing.T) {
	rs := newRuleSet()
	a, b, c, d, e, f := newOption("a"), newOption("b"), newOption("c"), newOption("d"), newOption("e"), newOption("f")
	rs.addDep(a, b)
	rs.addDep(b, c)
	rs.addDep(c, d)
	rs.addDep(d, e)
	rs.addDep(a, f)
	rs.addConflict(e, f)
	err := fmt.Errorf("")
	if err = rs.isCoherent(); err == nil {
		t.Fatalf("rule set is coherent but should not")
	}
	t.Logf("rule set is non coherent as intended: %v", err)
}
