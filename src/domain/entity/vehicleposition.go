package entity

type PositionAttributes struct {
	VehicleId         string  `json:"vehicle_id" db:"vehicleId"`
	StartDate         string  `json:"trip_start_date" db:"startDate"`
	LastUpdate        string  `json:"date_updated" db:"lastUpdate"`
	Longitude         float64 `json:"position_longitude" db:"longitude"`
	ParentTripId      int     `json:"trip_schedule_relationship" db:"parentTripId"`
	PositionSpeed     int     `json:"position_speed" db:"positionSpeed"`
	Latitude          float64 `json:"position_latitude" db:"latitude"`
	RouteId           string  `json:"trip_route_id" db:"routeId"`
	Label             string  `json:"vehicle_label" db:"label"`
	PostitionOdometer int     `json:"position_odometer" db:"positionOdometer"`
	TripId            string  `json:"trip_id" db:"tripId"`
	VehicleStatus     int     `json:"vehicle_current_status" db:"vehicleStatus"`
	Alcaldia          string  `db:"alcaldia"`
}

type VehiclePosition struct {
	Id                 string `json:"recordid" db:"id"`
	PositionAttributes `json:"fields"`
}

type Positions struct {
	VehiclePositions []*VehiclePosition `json:"records"`
}
