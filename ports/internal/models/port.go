package models

type Port struct {
	Id          string        `json:"id"          pg:"id"`
	Name        string        `json:"name"        pg:"name"`
	RefName     string        `json:"ref_name"    pg:"ref_name"`
	City        string        `json:"city"        pg:"city"`
	Country     string        `json:"country"     pg:"country"`
	Alias       []interface{} `json:"alias"       pg:"alias"`
	Regions     []interface{} `json:"regions"     pg:"regions"`
	Coordinates []float64     `json:"coordinates" pg:"coordinates"`
	Province    string        `json:"province"    pg:"province"`
	Timezone    string        `json:"timezone"    pg:"timezone"`
	Unlocs      []string      `json:"unlocs"      pg:"unlocs"`
	Code        string        `json:"code"        pg:"code"`
}

type PortsMap map[string]Port

func (pm PortsMap) ToPorts() (ports []Port) {
	for key := range pm {
		port := pm[key]
		port.RefName = key
		ports = append(ports, port)
	}
	return
}
