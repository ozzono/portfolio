package loadbalance

import "fmt"

type uniqueType struct {
	id    int
	score int
}

type cs struct {
	uniqueType
	clients []uniqueType
}

func balance(csList, clientList []uniqueType, unnavailableCS []int) (int, error) {
	csList = sort(filterIDs(csList, unnavailableCS))
	csGroup, err := groupCS(csList)
	if err != nil {
		return 0, nil
	}

	for i := range clientList {
		for minmax := range csGroup {
			if clientList[i].score > minmax[0] && clientList[i].score <= minmax[1] {
				csGroup[minmax] = assignCS(csGroup[minmax], clientList[i])
			}
		}
	}

	luckyCS := []cs{}
	for key := range csGroup {
		for i := range csGroup[key] {
			if len(csGroup[key][i].clients) > 0 {
				if len(luckyCS) == 0 {
					luckyCS = []cs{csGroup[key][i]}
					continue
				}
				if len(csGroup[key][i].clients) > len(luckyCS[0].clients) {
					luckyCS = []cs{csGroup[key][i]}
					continue
				}
				if len(csGroup[key][i].clients) == len(luckyCS[0].clients) {
					luckyCS = append(luckyCS, csGroup[key][i])
				}
			}
		}
	}

	if len(luckyCS) != 1 {
		return 0, nil
	}

	return luckyCS[0].id, nil
}

// Sets the data ID and Score
// This is used to shape the CS and the Client
func arrayToType(input []int) (output []uniqueType) {
	for i := range input {
		output = append(output, uniqueType{i + 1, input[i]})
	}
	return output
}

// Filters out the unnavailable CS ids
func filterIDs(mainList []uniqueType, filterList []int) []uniqueType {
	if len(filterList) == 0 || len(mainList) == 0 {
		return mainList
	}

	// by using two maps I avoid nested iterations
	main := map[int]uniqueType{}
	for i := range mainList {
		main[mainList[i].id] = mainList[i]
	}

	filter := map[int]int{}
	for i := range filterList {
		filter[filterList[i]] = filterList[i]
	}

	output := []uniqueType{}
	for key := range main {
		_, ok := filter[main[key].id]
		if ok {
			continue
		}
		output = append(output, main[key])
	}
	return output
}

func groupCS(input []uniqueType) (map[[2]int][]cs, error) {
	if len(input) == 0 {
		return map[[2]int][]cs{}, fmt.Errorf("input cannot be empty")
	}
	if len(input) == 1 {
		return map[[2]int][]cs{
			{input[0].score, input[0].score - 1}: {cs{input[0], []uniqueType{}}},
		}, nil
	}

	output := map[[2]int][]cs{}
	tmp := map[int][]cs{}
	for i := range input {
		tmp[input[i].score] = append(tmp[input[i].score], cs{input[i], []uniqueType{}})
	}

	last := 0
	for i := range input {
		if last == input[i].score {
			continue
		}
		last = input[i].score

		min := 0
		if i > 0 {
			min = tmp[input[i-1].score][0].score
		}
		output[[2]int{min, tmp[input[i].score][0].score}] = append(output[[2]int{min, tmp[input[i].score][0].score}], tmp[input[i].score]...)
	}
	return output, nil
}

// Sort was needed only because of the min-max score range
func sort(pool []uniqueType) []uniqueType {
	for i := len(pool); i > 0; i-- {
		for j := 1; j < i; j++ {
			if pool[j-1].score >= pool[j].score {
				// tmp := pool[j]
				// pool[j] = pool[j-1]
				// pool[j-1] = tmp
				pool[j], pool[j-1] = pool[j-1], pool[j]
			}
		}
	}
	return pool
}

func assignCS(group []cs, client uniqueType) []cs {
	nextCs := 0
	for i := range group {
		if len(group[nextCs].clients) < len(group[i].clients) {
			nextCs = i
		}
	}
	group[nextCs].clients = append(group[nextCs].clients, client)
	return group
}
