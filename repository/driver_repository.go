package repository

import (
	"database/sql"
	"fmt"

	"github.com/EliasSantiago/api-go-challenge/model"
)

type DriverRepository struct {
	connection *sql.DB
}

func NewDriverRepository(connection *sql.DB) DriverRepository {
	return DriverRepository{
		connection: connection,
	}
}

func (dr *DriverRepository) GetDrivers() ([]model.Driver, error) {
	query := "SELECT id, cpf, name FROM drivers"
	rows, err := dr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Driver{}, err
	}

	var driverList []model.Driver
	var driverObj model.Driver

	for rows.Next() {
		err = rows.Scan(
			&driverObj.ID,
			&driverObj.CPF,
			&driverObj.Name,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Driver{}, err
		}

		driverList = append(driverList, driverObj)
	}

	return driverList, nil
}

func (dr *DriverRepository) CreateDriver(driver model.Driver) (int, error) {
	var id int
	query, err := dr.connection.Prepare("INSERT INTO drivers" +
		"(cpf, name)" +
		" VALUES ($1, $2) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(driver.CPF, driver.Name).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}

func (dr *DriverRepository) GetDriverByID(id int64) (*model.Driver, error) {
	var driver model.Driver
	query, err := dr.connection.Prepare("SELECT id, cpf, name FROM drivers WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = query.QueryRow(id).Scan(&driver.ID, &driver.CPF, &driver.Name)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	query.Close()
	return &driver, nil
}

func (dr *DriverRepository) UpdateDriver(driver model.Driver) error {
	query, err := dr.connection.Prepare("UPDATE drivers SET cpf = $1, name =$2 WHERE id = $3")
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = query.Exec(driver.CPF, driver.Name, driver.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	query.Close()
	return nil
}
