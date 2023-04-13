package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	u "randomuser/user"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "test"
	db_name  = "postgres"
)

var db *sql.DB
var err error

func addId(resultId u.Id) int {
	sqlStatement := `
			INSERT INTO Ids ("Name", "Value")
			VALUES ($1, $2)
			RETURNING id`
	id := 0

	err = db.QueryRow(sqlStatement, resultId.Name, resultId.Value).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
	return id
}

func AddLocation(resultId int, location u.Location, street u.Street, coordinates u.Coordinates, timezone u.Timezone) {
	sqlStatementLocation := `
			INSERT INTO Locations 
			(StreetId, City, "State", Country, Postcode, CoordinatesId, TimezoneId, ResultId)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
			RETURNING id`
	sqlStatementTimezone := `
			INSERT INTO Timezones 
			("Offset", "Description")
			VALUES ($1, $2)
			RETURNING id`
	sqlSelectTimezone := `
			SELECT id
			FROM timezones
			WHERE "Offset" = $1`
	sqlStatementStreet := `
			INSERT INTO Streets 
			("Number", "Name")
			VALUES ($1, $2)
			RETURNING id`
	sqlStatementCoordinate := `
			INSERT INTO Coordinates 
			(Latitude, Longitude)
			VALUES ($1, $2)
			RETURNING id`
	idTimezone := -1
	idStreet := 0
	idCoordinates := 0
	idLocation := 0

	err = db.QueryRow(sqlSelectTimezone, timezone.Offset).Scan(&idTimezone)
	if err != nil {
		if idTimezone == -1 {
			err = db.QueryRow(sqlStatementTimezone, timezone.Offset, timezone.Description).Scan(&idTimezone)
			if err != nil {
				panic(err)
			}
		}
	}

	err = db.QueryRow(sqlStatementStreet, street.Number, street.Name).Scan(&idStreet)
	if err != nil {
		panic(err)
	}

	err = db.QueryRow(sqlStatementCoordinate, coordinates.Latitude, coordinates.Longitude).Scan(&idCoordinates)
	if err != nil {
		panic(err)
	}

	err = db.QueryRow(sqlStatementLocation, idStreet, location.City, location.State,
		location.Country, location.Postcode, idCoordinates, idTimezone, resultId).Scan(&idLocation)

	if err != nil {
		panic(err)
	}

	fmt.Println("Insert Location Success")
}

func AddLogin(resultId int, login u.Login) {
	sqlStatement := `
			INSERT INTO Logins 
			(Uuid, Username, "Password", Salt, md5, sha1, sha256, resultId)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
			RETURNING id`
	idLogin := 0

	err = db.QueryRow(sqlStatement, login.Uuid, login.Username, login.Password, login.Salt,
		login.Md5, login.Sha1, login.Sha256, resultId).Scan(&idLogin)

	if err != nil {
		panic(err)
	}

	fmt.Println("Insert Login Success")
}

func AddDob(resultId int, dob u.Dob) {
	sqlStatement := `
			INSERT INTO Dobs 
			("Date", Age, resultId)
			VALUES ($1, $2, $3)
			RETURNING id`
	idDob := 0

	err = db.QueryRow(sqlStatement, dob.Date, dob.Age, resultId).Scan(&idDob)

	if err != nil {
		panic(err)
	}

	fmt.Println("Insert Dob Success")
}

func AddRegistered(resultId int, registered u.Registered) {
	sqlStatement := `
			INSERT INTO Registereds
			("Date", Age, resultId)
			VALUES ($1, $2, $3)
			RETURNING id`
	idRegistered := 0

	err = db.QueryRow(sqlStatement, registered.Date, registered.Age, resultId).Scan(&idRegistered)

	if err != nil {
		panic(err)
	}

	fmt.Println("Insert Registered Success")
}

func AddPicture(resultId int, picture u.Picture) {
	sqlStatement := `
			INSERT INTO Pictures
			(Large, Medium, Thumbnail, resultId)
			VALUES ($1, $2, $3, $4)
			RETURNING id`
	idPicture := 0

	err = db.QueryRow(sqlStatement, picture.Large, picture.Medium, picture.Thumbnail, resultId).Scan(&idPicture)

	if err != nil {
		panic(err)
	}

	fmt.Println("Insert Picture Success")
}

func AddResult(resultId int, result u.Result) {
	sqlStatement := `
			INSERT INTO Results
			(Gender, Email, Phone, Cell, Nat, resultId)
			VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING id`
	id := 0

	err = db.QueryRow(sqlStatement, result.Gender, result.Email,
		result.Phone, result.Cell, result.Nat, resultId).Scan(&id)

	if err != nil {
		panic(err)
	}

	fmt.Println("Insert Result Success")
}

func InsertAll(result u.Result) {
	id := addId(result.Id)
	AddLocation(id, result.Location, result.Location.Street, result.Location.Coordinates, result.Location.Timezone)
	AddLogin(id, result.Login)
	AddDob(id, result.Dob)
	AddRegistered(id, result.Registered)
	AddPicture(id, result.Picture)
	AddResult(id, result)
}

func CloseConnect() {
	defer db.Close()
}

func ConnectToDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open(db_name, psqlInfo)

	if err != nil {
		panic(err)
	}
}
