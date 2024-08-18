package repository

import (
	"database/sql"
	"fmt"

	"github.com/EliasSantiago/api-go-challenge/model"
)

type VehicleRepository interface {
	GetVehicles() ([]model.Vehicle, error)
	CreateVehicle(vehicle model.Vehicle) (int64, error)
	GetVehicleByID(id int64) (*model.Vehicle, error)
	UpdateVehicle(vehicle model.Vehicle) error
	DeleteVehicle(id int64) error
	AssignDriver(request model.VehicleAssignDriverRequest) error
}

type vehicleRepository struct {
	connection *sql.DB
}

func NewVehicleRepository(connection *sql.DB) VehicleRepository {
	return &vehicleRepository{
		connection: connection,
	}
}

func (vr *vehicleRepository) GetVehicles() ([]model.Vehicle, error) {
	query := "SELECT id, license_plate, model FROM vehicles"
	rows, err := vr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Vehicle{}, err
	}

	var vehicleList []model.Vehicle
	var vehicleObj model.Vehicle

	for rows.Next() {
		err = rows.Scan(
			&vehicleObj.ID,
			&vehicleObj.LicensePlate,
			&vehicleObj.Model,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Vehicle{}, err
		}

		vehicleList = append(vehicleList, vehicleObj)
	}

	return vehicleList, nil
}

func (vr *vehicleRepository) CreateVehicle(vehicle model.Vehicle) (int64, error) {
	var id int64
	query, err := vr.connection.Prepare("INSERT INTO vehicles" +
		"(license_plate, model)" +
		" VALUES ($1, $2) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(vehicle.LicensePlate, vehicle.Model).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}

func (vr *vehicleRepository) GetVehicleByID(id int64) (*model.Vehicle, error) {
	var vehicle model.Vehicle
	err := vr.connection.QueryRow("SELECT id, license_plate, model FROM vehicles WHERE id = $1", id).Scan(&vehicle.ID, &vehicle.LicensePlate, &vehicle.Model)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &vehicle, nil
}

func (vr *vehicleRepository) UpdateVehicle(vehicle model.Vehicle) error {
	query, err := vr.connection.Prepare("UPDATE vehicles SET license_plate = $1, model =$2 WHERE id = $3")
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = query.Exec(vehicle.LicensePlate, vehicle.Model, vehicle.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	query.Close()
	return nil
}

func (vr *vehicleRepository) DeleteVehicle(id int64) error {
	_, err := vr.connection.Exec("DELETE FROM vehicles WHERE id = $1", id)
	return err
}

func (vr *vehicleRepository) AssignDriver(request model.VehicleAssignDriverRequest) error {
	query := `INSERT INTO driver_vehicles (driver_id, vehicle_id) VALUES ($1, $2)`
	_, err := vr.connection.Exec(query, request.DriverID, request.VehicleID)

	if err != nil {
		return err
	}

	return nil
}
