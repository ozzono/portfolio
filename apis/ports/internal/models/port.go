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
	strAlias       *string       `pg:"type:varchar;default:null"` // into pg
	Regions        []interface{} `json:"regions,omitempty"`       //from json file
	strRegions     *string       `pg:"type:varchar;default:null"` // into pg
	Coordinates    []float64     `json:"coordinates,omitempty"`   //from json file
	strCoordinates *string       `pg:"type:varchar;default:null"` // into pg
	Unlocs         []string      `json:"unlocs,omitempty"`        //from json file
	strUnlocs      *string       `pg:"type:varchar;default:null"` // into pg
}

func MapToPorts(pm map[string]Port) (ports []*Port) {
	for key := range pm {
		port := pm[key]
		port.RefName = key

		if port.Alias != nil && len(port.Alias) > 0 {
			port.strAlias = ToString(port.Alias)
		}

		if port.Regions != nil && len(port.Regions) > 0 {
			port.strRegions = ToString(port.Regions)
		}

		if port.Coordinates != nil && len(port.Coordinates) > 0 {
			port.strCoordinates = ToString(port.Coordinates)
		}

		if port.Unlocs != nil && len(port.Unlocs) > 0 {
			port.strUnlocs = ToString(port.Unlocs)
		}
		ports = append(ports, &port)
	}
	return
}

func ToString(input interface{}) *string {
	b, _ := json.MarshalIndent(&input, "", "  ")
	strB := string(b)
	return &strB
}

func FromString(input *string, output interface{}) {
	json.Unmarshal([]byte(*input), &output)
	input = nil
}
