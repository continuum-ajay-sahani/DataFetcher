package ifsc

import (
	"database/sql"
	"fmt"
	"log"
	//sqlite using
	_ "github.com/mattn/go-sqlite3"
)

const (
	//CBankDB db name
	CBankDB = "./bank.db"
)

//DBOperation tobe
type DBOperation struct {
	db *sql.DB
}

func (d *DBOperation) initDB() error {
	//os.Remove(CBankDB)

	db, err := sql.Open("sqlite3", CBankDB)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()
	_, err = db.Exec(getBankTable())
	if err != nil {
		fmt.Println(err)
		return err
	}
	d.db = db
	return err
}

func (d *DBOperation) openDBConn() error {
	db, err := sql.Open("sqlite3", CBankDB)
	if err != nil {
		log.Fatal(err)
	}
	d.db = db
	return err
}

func (d *DBOperation) insert(b Bank) (err error) {
	cmd := fmt.Sprintf("insert into Bank (name,state,district,branch,address,contact,ifsc,micr,latitude,longitude,details)"+
		"values ('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s')", b.bank, b.state, b.district,
		b.branch, b.address, b.contact, b.ifscCode, b.micrCode, b.latitude, b.longitude, b.details)
	_, err = d.db.Exec(cmd)
	return err
}

func (d *DBOperation) closeDB() {
	if d.db != nil {
		d.db.Close()
	}
}

func getBankTable() string {
	query := "create table Bank (RequestID integer PRIMARY KEY   AUTOINCREMENT ,name varchar" +
		" ,state varchar ,district varchar, branch varchar ,address varchar ,contact varchar" +
		" ,ifsc varchar ,micr varchar ,latitude string, longitude string, details string)"
	return query
}
