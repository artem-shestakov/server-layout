package auth

// Camera information from VDC servers
// swagger:response cameraResponse
type cameraResponse struct {
	// in: body
	Body Camera
}

// Serial number of IP camera
// swagger:parameters uidbysn
type serialNumber struct {
	// in: query
	// required: true
	SerialNumber string `json:"serial_number"`
}

type Camera struct {
	// Camera ID
	ID string `json:"id"`
	// Camera UID
	Uid string `json:"uid"`
	// State of camera
	State string `json:"state"`
	// Сerver to which the camera is connected
	Server string `json:"server"`
}
