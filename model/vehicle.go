package model

type Vehicle struct {
	ID             int64  `json:"id"`
	LicenseVehicle string `json:"licensePlate"`
	Model          string `json:"model"`
}

type VehicleUpdateRequest struct {
	ID             int64  `json:"vehicleId" binding:"required"`
	LicenseVehicle string `json:"licensePlate" binding:"required"`
	Model          string `json:"model" binding:"required"`
}
