package model

import (
	"fmt"
	"net/http"
	"time"
)

func (m *Min) Cookies() []*http.Cookie {
	return []*http.Cookie{{
		Name:       "token",
		Value:      m.Token,
		Path:       "/",
		Domain:     ".cdcas.com",
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
		Domain:     ".cdcas.com",
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

// GetBase
/**
 * @Description: state为0代表慕课平台，state为1代表实训平台
 * @receiver m
 * @return string
 */
func (m *Min) GetBase() string {
	if m.State == 0 {
		return "https://mooc.cdcas.com"
	} else if m.State == 1 {
		return "https://shixun.cdcas.com"
	} else if m.State == 2 {
		return "https://cdcas.chuangnengkeji.com"
	} else {
		return "https://mooc.cdcas.com"
	}
}
