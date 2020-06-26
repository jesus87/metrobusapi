package entity

import (
	"fmt"
	"strings"
)

type GeoShape struct {
	TypeId      string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}

type Attributes struct {
	Name     string   `json:"nomgeo" db:"name"`
	GeoShape GeoShape `json:"geo_shape"`
}

type Alcaldia struct {
	Id         string `json:"recordid" db:"id"`
	Attributes `json:"fields"`
}

type AlcaldiaCatalog struct {
	Alcaldias []*Alcaldia `json:"records"`
}

func (a *Alcaldia) GetPolygon() string {
	//'POLYGON((-74.13591384887695 40.93750722242824,-74.13522720336914 40.929726129575016))'

	var sbCoordinates strings.Builder
	if len(a.GeoShape.Coordinates) > 0 {
		for _, coordinate := range a.GeoShape.Coordinates[0] {
			if len(coordinate) == 2 {
				if sbCoordinates.Len() > 0 {
					sbCoordinates.WriteString(",")
				}

				sbCoordinates.WriteString(fmt.Sprintf("%f %f", coordinate[0], coordinate[1]))
			}
		}
	}

	return "POLYGON((" + sbCoordinates.String() + "))"
}
