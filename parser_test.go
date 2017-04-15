package hostkeys

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	cases := []struct {
		Name          string
		Path          string
		ExpectedKeys  []string
		ExpectedError error
	}{
		{
			Name:          "All keys present 1",
			Path:          "example1",
			ExpectedKeys:  example1Keys,
			ExpectedError: nil,
		},
		{
			Name:          "All keys present 2",
			Path:          "example2",
			ExpectedKeys:  example2Keys,
			ExpectedError: nil,
		},
		{
			Name:          "Empty",
			Path:          "example3",
			ExpectedKeys:  nil,
			ExpectedError: ErrNoStartHostKeysBlock,
		},
		{
			Name:          "Invalid Key",
			Path:          "example4",
			ExpectedKeys:  nil,
			ExpectedError: errors.New("invalid host key format: ecdsa-sha2-nistp256"),
		},
	}

	for _, tc := range cases {
		source, err := os.Open(filepath.Join("test-fixtures", tc.Path))
		if err != nil {
			t.Fatalf("%s: os.Open: %s", tc.Name, err)
		}

		f, err := ioutil.ReadAll(source)
		if err != nil {
			t.Fatalf("%s: ioutil.ReadAll: %s", tc.Name, err)
		}

		keys, err := Parse(string(f))

		assert.Equal(t, tc.ExpectedError, err)
		assert.Equal(t, tc.ExpectedKeys, keys)
	}

}

var example1Keys = []string{
	"ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHlszmP1X6dG89YllW0P5giiK9Zw7kYG3ZqRl20uCN3ccd4kJtu4SPQ/4C7SF8EqFT1jV7mYgFwQBGk2CxBFMXs=",
	"ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIN21qV9XJj1ns+UcUVYBLxO+YFsbdO5o+xD3ywRiLq8I",
	"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC61LgHyaAiXuWp2ym4Vfg70DrTnb8a/mWXugk9Yug44Au+Sfa/eGblhMPMbH/VHMHu1PhHUvN8raCG4xXbdN0303S73UpL85D7RIUoCNMHf/8KuKPMYpKlH7On1fmb7zusZOIhnXr2nl1BJdbEW2hvTh09B2tJEywsBqHr6VT79UqVub1nhzw4+idpPlUTqZ1XJkaCuRa748nvVbXUt6+iqCBIP2CcQjWKsmTyRRsMxeoJD4+zIIFElcD3Yd16/uHI9ra/rtC1CfHukKo/L+rnIAuPTQQd3yTmscB3Q39QsisUq8gWxRpricNIqQXTB5oiA+1QFztkO3pT9Zrj3Lol",
}

var example2Keys = []string{
	"ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBDPjs58ymtFMCw/lITE7KVWNCvp9ogIe6IpBUnYZRSgE80T+G6x2d3HtHJUqb+mx8HKjjnvxlU98r8WNRcbKLeI=",
	"ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIIbhB2/ObNlRtBoPoncWIQRgOS05NLvTqCThnDk/509a",
	"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCVWGpsxTYa3smPdeTHw5JgMHXtwChUHGaw0ve/Wo17oWS+k7KOHKF9Wci3KQD/Kl7RdKAG86fcMfJ7Xt5pDr01wefb+0jjZjbwK5vLEJm/Y5GKiSBo1KrKb4f7dqZvaF41+mk9xoRokGb1iqWdrr50EhSQQvMS8j/wrcV8mQQwIYgQbZ+oD+j64qFMjdkRDS4Kdx4ZJmzKNrEzMmfwREyvdxDyyoio8GmZUM6UZc+8d8O9evLHWBOfAqtmgMKGqEh+KETcYLouf8c1fNmj6psfxz5oduPpi0s2bO9hZgHRMXrXszwTrvyikxaAW+GJ36W6jw4TcfCPMQAHWEVsqy7V",
}
