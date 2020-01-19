/*
Package ip2country uses github.com/oschwald/maxminddb-golang with an embedded GeoLite2 db
to map net.IPs to country ISO codes
*/
package ip2country

import (
	"encoding/base64"
	"net"

	"github.com/oschwald/maxminddb-golang"
)

type record struct {
	Country struct {
		ISOCode string `maxminddb:"iso_code"`
	} `maxminddb:"country"`
}

var reader *maxminddb.Reader

func Init() {
	reader, _ = maxminddb.FromBytes(Bytes())
}

func Bytes() []byte {
	d2, _ := base64.StdEncoding.DecodeString(data)
	return d2
}

func Reader() *maxminddb.Reader {
	if reader == nil {
		Init()
	}
	return reader
}

// Country returns ISO code of the country that given IP belongs to.
func Country(ip net.IP) (string, error) {
	r := &record{}
	err := Reader().Lookup(ip, &r)
	if err != nil {
		return "", err
	}
	return r.Country.ISOCode, nil
}
