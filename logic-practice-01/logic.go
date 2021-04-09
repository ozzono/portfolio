package logicpractice01

import (
	"fmt"
)

type option struct {
	name           string
	value          bool
	dependencyList map[string]*option
	conflictList   map[string]*option
}

type ruleSet struct {
	options map[string]option
}

func newRuleSet() ruleSet {
	return ruleSet{options: map[string]option{}}
}

func newOption(name string) (out option) {
	out = option{name: name}
	out.conflictList = make(map[string]*option)
	out.dependencyList = make(map[string]*option)
	return
}

func (rs *ruleSet) addDep(a, b option) {
	rs.addOption(a, b)
	// log.Printf("adding dep between option '%s' and option '%s'", a.name, b.name)
	rs.options[a.name].dependencyList[b.name] = &b
}

func (rs *ruleSet) addConflict(a, b option) {
	rs.addOption(a, b)
	// log.Printf("adding conflict between option '%s' and option '%s'", a.name, b.name)
	rs.options[a.name].conflictList[b.name] = &b
}

func (rs *ruleSet) isCoherent() error {
	for i := range rs.options {
		allDependencies, allConflicts := dependencyNConflictList(rs.options[i])
		if len(allConflicts) == 0 {
			continue
		}
		if hasInList(rs.options[i].name, allConflicts) {
			return fmt.Errorf("an item cannot conflict to itself")
		}
		for key := range allDependencies {
			if hasInList(key, allConflicts) {
				return fmt.Errorf("'%s' is within '%s' dependency and conflict list simultaneously", key, rs.options[i].name)
			}
		}
	}
	return nil
}

func (rs *ruleSet) addOption(options ...option) {
	for i := range options {
		_, exist := rs.options[options[i].name]
		if !exist {
			rs.options[options[i].name] = options[i]
		}
	}
}

func hasInList(name string, list map[string]*option) bool {
	_, conflicted := list[name]
	return conflicted
}

func getDependencyList(o option) map[string]*option {
	output := make(map[string]*option)
	for _, item := range o.dependencyList {
		if hasInList(o.name, item.dependencyList) {
			output[o.name] = &o
			continue
		}
		output[item.name] = item
		for key, depItem := range getDependencyList(*item) {
			output[key] = depItem
		}
	}
	return output
}

func dependencyNConflictList(o option) (map[string]*option, map[string]*option) {
	output := o.conflictList
	dependencies := getDependencyList(o)
	for _, item := range dependencies {
		for key, conflict := range item.conflictList {
			output[key] = conflict
		}
	}
	return dependencies, output
}
