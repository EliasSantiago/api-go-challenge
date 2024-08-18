package model

type Vehicle struct {
	ID           int64  `json:"id"`
	LicensePlate string `json:"licensePlate" binding:"required"`
	Model        string `json:"model" binding:"required"`
}

type VehicleUpdateRequest struct {
	ID           int64  `json:"vehicleId" binding:"required"`
	LicensePlate string `json:"licensePlate" binding:"required"`
	Model        string `json:"model" binding:"required"`
}

type VehicleAssignDriverRequest struct {
	VehicleID int64 `json:"vehicleId" binding:"required"`
	DriverID  int64 `json:"driverId" binding:"required"`
}
