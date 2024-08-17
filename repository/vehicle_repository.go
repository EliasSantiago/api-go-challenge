package repository

import (
	"database/sql"
	"fmt"

	"github.com/EliasSantiago/api-go-challenge/model"
)

type VehicleRepository struct {
	connection *sql.DB
}

func NewVehicleRepository(connection *sql.DB) VehicleRepository {
	return VehicleRepository{
		connection: connection,
	}
}

func (vr *VehicleRepository) GetVehicles() ([]model.Vehicle, error) {
	query := "SELECT license_plate, model FROM vehicles"
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
			&vehicleObj.LicenseVehicle,
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

func (vr *VehicleRepository) CreateVehicle(vehicle model.Vehicle) (int, error) {
	var id int
	query, err := vr.connection.Prepare("INSERT INTO vehicles" +
		"(license_plate, model)" +
		" VALUES ($1, $2) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(vehicle.LicenseVehicle, vehicle.Model).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}

func (dr *VehicleRepository) GetVehicleByID(id int64) (*model.Vehicle, error) {
	var vehicle model.Vehicle
	err := dr.connection.QueryRow("SELECT id, license_plate, model FROM vehicles WHERE id = $1", id).Scan(&vehicle.ID, &vehicle.LicenseVehicle, &vehicle.Model)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &vehicle, nil
}

func (dr *VehicleRepository) UpdateVehicle(vehicle model.Vehicle) error {
	query, err := dr.connection.Prepare("UPDATE vehicles SET license_plate = $1, model =$2 WHERE id = $3")
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = query.Exec(vehicle.LicenseVehicle, vehicle.Model, vehicle.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	query.Close()
	return nil
}
