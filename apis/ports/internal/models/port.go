package models

import (
	"encoding/json"
)

type Port struct {
	Id             *int          `json:"id,omitempty"          pg:"type:int"`
	Name           string        `json:"name,omitempty"        pg:"type:varchar"`
	RefName        string        `json:"ref_name,omitempty"    pg:"type:varchar"`
	City           string        `json:"city,omitempty"        pg:"type:varchar"`
	Country        string        `json:"country,omitempty"     pg:"type:varchar"`
	Province       string        `json:"province,omitempty"    pg:"type:varchar"`
	Timezone       string        `json:"timezone,omitempty"    pg:"type:varchar"`
	Code           string        `json:"code,omitempty"        pg:"type:varchar"`
	Alias          []interface{} `json:"alias,omitempty"`         //from json file
	StrAlias       string        `pg:"type:varchar;default:null"` // into pg
	Regions        []interface{} `json:"regions,omitempty"`       //from json file
	StrRegions     string        `pg:"type:varchar;default:null"` // into pg
	Coordinates    []float64     `json:"coordinates,omitempty"`   //from json file
	StrCoordinates string        `pg:"type:varchar;default:null"` // into pg
	Unlocs         []string      `json:"unlocs,omitempty"`        //from json file
	StrUnlocs      string        `pg:"type:varchar;default:null"` // into pg
}

func MapToPorts(pm map[string]Port) (ports []*Port) {
	for key := range pm {
		port := pm[key]
		port.RefName = key

		if port.Alias != nil && len(port.Alias) > 0 {
			port.StrAlias = toString(port.Alias)
		}

		if port.Regions != nil && len(port.Regions) > 0 {
			port.StrRegions = toString(port.Regions)
		}

		if port.Coordinates != nil && len(port.Coordinates) > 0 {
			port.StrCoordinates = toString(port.Coordinates)
		}

		if port.Unlocs != nil && len(port.Unlocs) > 0 {
			port.StrUnlocs = toString(port.Unlocs)
		}
		ports = append(ports, &port)
	}
	return
}

func toString(input interface{}) string {
	b, _ := json.MarshalIndent(&input, "", "  ")
	return string(b)
}

func FromString(input string, output interface{}) {
	json.Unmarshal([]byte(input), &output)
}
