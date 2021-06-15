package itinerary

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type airport struct {
	Name      string
	Destinies []*airport
	Hopped    bool
}

type leap struct {
	origin *airport
	cost   int
}

type route struct {
	leaps []leap
	cost  int
}

func (l *leap) routes() []route {

	return nil
}

func allAirports(path string) ([]*airport, []leap, error) {
	airports := []*airport{}
	airportMap := map[string]*airport{}
	leaps := []leap{}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, nil, fmt.Errorf("ioutil.ReadFile err: %v", err)
	}
	rows := strings.Split(string(file), "\n")
	for i := range rows {
		columns := strings.Split(rows[i], ",")
		_, ok := airportMap[columns[0]]
		if !ok {
			airportMap[columns[0]] = &airport{Name: columns[0]}
		}
		_, ok = airportMap[columns[1]]
		if !ok {
			airportMap[columns[1]] = &airport{Name: columns[1]}
		}
		if len(columns) != 3 {
			return nil, nil, fmt.Errorf("invalid row format")
		}
	}
	for key := range airportMap {
		airports = append(airports, airportMap[key])
		fmt.Printf("%s.%p\n", key, airportMap[key])
	}
	for i := range rows {
		columns := strings.Split(rows[i], ",")
		cost, err := strconv.Atoi(columns[2])
		if err != nil {
			return nil, nil, fmt.Errorf("strconv.Atoi - %s", err)
		}
		leaps = append(leaps, leap{airportMap[columns[0]], airportMap[columns[1]], cost})
	}
	return airports, leaps, nil
}
