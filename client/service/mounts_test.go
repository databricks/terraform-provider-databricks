package service

import (
	"net/url"
	"testing"
)

func TestParseURI(t *testing.T) {

	u, err := url.Parse("adl://sritikkireddyterraadls1.azuredatalakestore.net/")
	if err != nil {
		panic(err)
	}
	t.Log(u)
	t.Log(u.User.String())
	t.Log(u.Host)
	t.Log(u.Path)

}
