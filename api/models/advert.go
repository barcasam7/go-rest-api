package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Advert struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Ad_name   string `json:"name"`
	Gender    string `json:"gender"`
	Location  string `json:"location"`
	Age_start int    `json:"age_start"`
	Age_end   int    `json:"age_end"`
}

func GetAdvert(db *sql.DB, id string) (Advert, error) {
	var advert Advert
	query := "SELECT id, type, ad_name, i_age_range_start, i_age_range_end, e_gender, v_location FROM advertise_info WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&advert.ID, &advert.Type, &advert.Ad_name, &advert.Age_start, &advert.Age_end, &advert.Gender, &advert.Location)

	if err != nil {
		return advert, err
	}
	return advert, nil
}

func GetAdverts(db *sql.DB) ([]Advert, error) {
	query := "SELECT id, type, ad_name, i_age_range_start, i_age_range_end, e_gender, v_location FROM advertise_info"
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	adverts := []Advert{}
	for rows.Next() {
		var a Advert
		err := rows.Scan(&a.ID, &a.Type, &a.Ad_name, &a.Age_start, &a.Age_end, &a.Gender, &a.Location)
		if err != nil {
			return nil, err
		}
		adverts = append(adverts, a)
	}
	return adverts, nil
}
