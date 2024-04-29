package models

import "time"

type CNTime time.Time

func (ct CNTime) MarshalJSON() ([]byte, error) {
	t := time.Time(ct)
	return []byte(`"` + t.Format("2006-01-02 15:04:05") + `"`), nil
}

type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}

type BasePage struct {
	Size int `json:"size,omitempty"`
	Page int `json:"page,omitempty"`
}
