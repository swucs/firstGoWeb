package main

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type SelectParam struct {
	CustomerId int `json:"customer_id"`
}

type SelectResult struct {
	CustomerName       string         `json:"customer_name"`
	RepresentativeName string         `json:"representative_name"`
	ItemName           string         `json:"item_name"`
	UnitWeight         float32        `json:"unit_weight"`
	UnitName           string         `json:"unit_name"`
	Remarks            sql.NullString `json:"remarks,string"`
}

func SelectCustomer(db *sqlx.DB) {

	var selectSQL string = `
		select
			c.name as customer_name
			, c.representative_name
			, i.name as item_name
			, i.unit_weight
			, i.unit_name 
			, i.remarks
		from customer c 
		left join customer_item ci on c.id = ci.customer_id
		left join item i on ci.item_id = i.id 
		where c.id = :customer_id
	`
	stmt, err := db.PrepareNamed(selectSQL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()

	results := []SelectResult{}
	err = stmt.Select(&results, &SelectParam{CustomerId: 1})
	if err != nil {
		log.Println(err)
		return
	}

	for _, row := range results {
		fmt.Printf("거래처명 : %s\n", row.CustomerName)
		fmt.Printf("대표자 : %s\n", row.RepresentativeName)
		fmt.Printf("품목 : %s\n", row.ItemName)
		fmt.Printf("단위무게 : %f\n", row.UnitWeight)
		fmt.Printf("단위명 : %s\n", row.UnitName)
		fmt.Printf("비고 : %s\n", row.Remarks.String)
		fmt.Printf("\n")
	}
}
