package itinerary

import "testing"

func Test(t *testing.T) {
	airports, leaps, err := allAirports("input-file.txt")
	if err != nil {
		t.Fatal(err)
	}
	for i := range airports {
		t.Logf("name: %s.%p", airports[i].Name, &airports[i])
	}
	for i := range leaps {
		t.Logf("from %s.%p to %s.%p", leaps[i].origin.Name, leaps[i].origin, leaps[i].destiny.Name, leaps[i].destiny)
	}
}
