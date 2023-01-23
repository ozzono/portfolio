# Logic Practice

This code test arises from a practical idea that can be summarized as follows.

You have an application for renting cars. One of the steps in the process is choosing the color, extras, etc., you want. Some options required some other
option to be included. For example, you can't have a GPS if you don't also have a USB plug. (But you can have the USB without having the GPS.)
So, when the user clicks on the GPS checkbox, the USB checkbox is automatically checked, and if the USB checkbox is unchecked, the GPS checkbox gets unchecked
too.

Now, say you have another system that uses the USB plug, for example, a music device or whatever. The USB plug is unique, so you can't have the GPS and the music. When the user checks the GPS checkbox, the music checkbox is disabled, and the other way around.

## Rule Sets

Let's say we have a set of options that the user can select. Options can be related between them in two ways: one can depend on another, and two options
can be mutually exclusive. That means that these equalities must always hold true (note: this is not code, those are logical equations):

* "A depends on B", or "for A to be selected, B needs to be selected"

```
ruleSet.AddDep(A, B) =>
if isSelected(A) then isSelected(B)
```

* "A and B are exclusive",  or "B and A are exclusive",  or "for A to be
selected, B needs to be unselected; and for B to be selected, A needs to be
unselected"

```
ruleSet.AddConflict(A, B) =>
ruleSet.AddConflict(B, A) =>
if isSelected(A) then !isSelected(B) && if isSelected(B) then !isSelected(A)
```

We say that a set of relations are _coherent_ if the laws above are valid for
that set. For example, this set of relations is coherent:

```
AddDep(A, B) // "A depends on B"
AddDep(B, C) // "B depends on C"
AddConflict(C, D) // "C and D are exclusive"
```

And these sets are _not_ coherent:

```
AddDep(A, B)
AddConflict(A, B)
```

A depends on B, so it's a contradiction that they are exclusive.  If B is
selected, A would need to be selected, but that's impossible because, by
the exclusion rule, A is supposed to be unselected when B is selected.

```
AddDep(A, B)
AddDep(B, C)
AddConflict(A, C)
```
- 

The dependency relation is transitive; it's easy to see, from the rules above,
that if A depends on B, and B depends on C, then A also depends on C. So this
is a contradiction for the same reason as the previous case.

## Questions

A. Write a data structure (`RuleSet`) for expressing these rules between options,
i.e., for defining a rule set.  You also need to define a constructor and 2
methods:

* `NewRuleSet()`: Returns an empty rule set.
* `RuleSet.AddDep(A, `B)` is a method for rule sets that adds a new dependency A
  between and B.
* `RuleSet.AddConflict(A, `B)`, a method for rule sets that add a new conflict
  between A and B.

B. Implement the algorithm that checks that an instance of that data structure is
coherent, that is, that no option can depend, directly or indirectly, on
another option and also be mutually exclusive with it.

* `RuleSet.IsCoherent`()` is a method for rule sets that returns true if it is a
  coherent rule set, false otherwise.

## Tests Scenarios

TestDependsAA
    - Get an empty ruleset
    - AddDep("a","a")
    - The result should be coherent

TestDependsAB_BA
    - Get an empty ruleset
    - AddDep("a", "b")
    - AddDep("b", "a")
    - The result should be coherent

TestExclusiveAB
    - Get an empty ruleset
    - AddDep("a", "b")
    - AddConflict("a", "b")
    - Be coherent should return an error

TestExclusiveAB_BC
    - Get an empty ruleset
    - AddDep("a", "b")
    - AddDep("b", "c")
    - AddConflict("a", "c")
    - Be coherent should return an error

TestDeepDeps
    - Get an empty ruleset
    - AddDep("a", "b")
 - AddDep("b", "c")
 - AddDep("c", "d")
 - AddDep("d", "e")
 - AddDep("a", "f")
 - AddConflict("e", "f")
    - Be coherent should return an error
