package ip2country_test

import (
	"fmt"
	"net"

	"github.com/admpub/ip2country"
)

func ExampleCountry() {
	country, _ := ip2country.Country(net.IPv4(8, 8, 8, 8))
	fmt.Println(country)
	// Output: US
}
