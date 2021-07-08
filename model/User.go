package model

type User struct {
	Name    string `json:"name"`
	PingCnt int    `json:"pingCnt"`
}
