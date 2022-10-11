package address

type Address struct {
	street   string
	number   int
	district string
	city     string
	state    string
	country  string
}

func setAddress(street string, number int, district string, city string, state string, country string) *Address {
	newAddress := Address{street, number, district, city, state, country}
	return &newAddress
}
