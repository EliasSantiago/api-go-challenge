package model

type Driver struct {
	ID   int64  `json:"id"`
	CPF  string `json:"cpf"`
	Name string `json:"name"`
}

type DriverUpdateRequest struct {
	ID   int64  `json:"driverId" binding:"required"`
	CPF  string `json:"cpf" binding:"required"`
	Name string `json:"name" binding:"required"`
}
