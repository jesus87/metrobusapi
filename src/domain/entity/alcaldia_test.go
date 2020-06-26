package entity_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jesus87/apidf/src/domain/entity"
)

func TestGetPolygon(t *testing.T) {
	alcaldia := &entity.Alcaldia{}
	alcaldia.GeoShape.Coordinates = [][][]float64{
		[][]float64{
			[]float64{100.213212, 2000.21321321},
			[]float64{101.213212, 2001.21321321},
			[]float64{-101.213212, -2001.21321321},
		},
	}

	polygonText := alcaldia.GetPolygon()
	for _, coordinates1 := range alcaldia.GeoShape.Coordinates {
		for _, coordinates2 := range coordinates1 {
			coorinate := fmt.Sprintf("%f %f", coordinates2[0], coordinates2[1])

			if !strings.Contains(polygonText, coorinate) {
				t.Errorf("Error in TestGetPolygon, expected: %v, but %v", coorinate, polygonText)
			}
		}
	}
}
