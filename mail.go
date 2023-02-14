package maildis

import (
	"github.com/LorenzoSrtPink/acxes"
)

func ListFor(report string) []string {
	setup()
	db, err := acxes.For(Host)
	checkFor(err, "Connecting to database")
	defer db.Close()
	query := `call mailing_list_for(?);`
	rows, err := db.Query(query, report)
	defer rows.Close()
	checkFor(err, "quering for results")
	var list []string
	for rows.Next() {
		var email string
		err = rows.Scan(&email)
		checkFor(err, "Scanning email")
		list = append(list, email)
	}
	return list
}
