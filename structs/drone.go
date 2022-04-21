package structs

type Drone struct {
	Id        int64  `json:"id"`
	SerialNumber string `json:"serial_number"`
	WeightLimit  string `json:"weight_limit"`
	BatteryCapacity     string `json:"battery_capacity"`
	State     string `json:"state"`
	CreatedAt string `json:"created_at"`
}