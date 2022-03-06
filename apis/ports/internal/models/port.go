package models

type Port struct {
	Id          *int          `json:"id,omitempty"          gorm:"type:int"`
	Name        string        `json:"name,omitempty"        gorm:"type:varchar"`
	RefName     string        `json:"ref_name,omitempty"    gorm:"type:varchar"`
	City        string        `json:"city,omitempty"        gorm:"type:varchar"`
	Country     string        `json:"country,omitempty"     gorm:"type:varchar"`
	Alias       []interface{} `json:"alias,omitempty"       gorm:"type:text[];default:null"`
	Regions     []interface{} `json:"regions,omitempty"     gorm:"type:text[];default:null"`
	Coordinates []float64     `json:"coordinates,omitempty" gorm:"type:varchar"`
	Province    string        `json:"province,omitempty"    gorm:"type:varchar"`
	Timezone    string        `json:"timezone,omitempty"    gorm:"type:varchar"`
	Unlocs      []string      `json:"unlocs,omitempty"      gorm:"type:text[];default:null"`
	Code        string        `json:"code,omitempty"        gorm:"type:varchar"`
}

func MapToPorts(pm map[string]Port) (ports []Port) {
	for key := range pm {
		port := pm[key]
		port.RefName = key

		if len(port.Alias) == 0 {
			port.Alias = nil
		}

		if len(port.Regions) == 0 {
			port.Regions = nil
		}

		if len(port.Unlocs) == 0 {
			port.Unlocs = nil
		}
		ports = append(ports, port)
		return
	}
	return
}
