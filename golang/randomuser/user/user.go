package user

type Response struct {
	Results []Result
	Info    Info
}

type Info struct {
	Seed    string
	Results int
	Page    int
	Version string
}

type Name struct {
	Title string
	First string
	Last  string
}

type Street struct {
	Number int
	Name   string
}

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

type Timezone struct {
	Offset      string
	Description string
}

type Location struct {
	Street      Street
	City        string
	State       string
	Country     string
	Postcode    int
	Coordinates Coordinates
	Timezone    Timezone
}

type Login struct {
	Uuid     string
	Username string
	Password string
	Salt     string
	Md5      string
	Sha1     string
	Sha256   string
}

type Dob struct {
	Date string
	Age  int
}

type Registered struct {
	Date string
	Age  int
}

type Id struct {
	Name  string
	Value string
}

type Picture struct {
	Large     string
	Medium    string
	Thumbnail string
}

type Result struct {
	Gender     string
	Name       Name
	Location   Location
	Email      string
	Login      Login
	Dob        Dob
	Registered Registered
	Phone      string
	Cell       string
	Id         Id
	Picture    Picture
	Nat        string
}
