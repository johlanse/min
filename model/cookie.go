package model

import (
	"fmt"
	"net/http"
	"time"
)

func (m Min) Cookies() []*http.Cookie {
	return []*http.Cookie{{
		Name:       "token",
		Value:      m.Token,
		Path:       "/",
		Domain:     "yinghuaonline.com",
		Expires:    time.Time{},
		RawExpires: "",
		MaxAge:     0,
		Secure:     false,
		HttpOnly:   false,
		SameSite:   0,
		Raw:        fmt.Sprintf("token=%v; path=/; domain=yinghuaonline.com", m.Token),
		Unparsed:   nil,
	}, {
		Name:       "tgw_l7_route",
		Value:      m.Tgw,
		Path:       "/",
		Domain:     "yinghuaonline.com",
		Expires:    time.Time{},
		RawExpires: "",
		MaxAge:     0,
		Secure:     false,
		HttpOnly:   false,
		SameSite:   0,
		Raw:        fmt.Sprintf("tgw_l7_route=%v; path=/; domain=yinghuaonline.com", m.Tgw),
		Unparsed:   nil,
	}}
}
