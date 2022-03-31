package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Customer struct {
	Id                 int
	Name               string
	BusinessNumber     string `json:"business_number"`
	RepresentativeName string `json:"representative_name"`
	BusinessConditions string `json:"business_conditions"`
	TypeOfBusiness     string `json:"type_of_business"`
	Address            string
	PhoneNumber        string `json:"phone_number"`
	FaxNumber          string `json:"fax_number"`
	Use                bool
	Deleted            bool
}

func CreateCustomer(tx *sqlx.Tx) {
	var insertSQLNoNamed string = `
		INSERT INTO public.customer(
		"name", business_number, representative_name, business_conditions, type_of_business, address, phone_number, fax_number, use, deleted
		) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10
		)
	`
	exec := tx.MustExec(insertSQLNoNamed,
		"거래처명",
		"221-82-01179",
		"홍길동",
		"제조,서비스",
		"축산물",
		"경기도 포천시 설운동 550-1 2/2",
		"031-541-4052",
		"031-541-4050",
		true,
		false,
	)
	fmt.Println(exec)

	var insertSQL string = `
		INSERT INTO public.customer(
		name, business_number, representative_name, business_conditions, type_of_business, address, phone_number, fax_number, use, deleted
		) VALUES (
		:name, :business_number, :representative_name, :business_conditions, :type_of_business, :address, :phone_number, :fax_number, :use, :deleted
		)
		RETURNING id
	`

	insertParam := Customer{
		Name:               "거래처명",
		BusinessNumber:     "221-82-01179",
		RepresentativeName: "홍길동",
		BusinessConditions: "제조,서비스",
		TypeOfBusiness:     "축산물",
		Address:            "경기도 포천시 설운동 550-1 2/2",
		PhoneNumber:        "031-541-4052",
		FaxNumber:          "031-541-4050",
		Use:                true,
		Deleted:            false,
	}

	//fmt.Println(insertParam)
	stmt, err := tx.PrepareNamed(insertSQL)
	var id int
	err = stmt.Get(&id, &insertParam)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("성공ID : ", id)

	if 1 == 1 {
		tx.Commit()
		return
	}

	result, err := tx.NamedExec(insertSQL, &insertParam)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	//lastinsertid는 postgres driver에서 지원하지 않는다고 함.
	//id, err := result.LastInsertId()
	affected, _ := result.RowsAffected()
	fmt.Println("성공 : ", id)
	fmt.Println("성공 : ", err)
	fmt.Println("성공 : ", affected)

	tx.Commit()
}
