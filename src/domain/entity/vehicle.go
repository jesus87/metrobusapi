package entity

type Vehicle struct {
	VehicleID int    `db:"vehicleId"`
	Label     string `db:"label"`
}
