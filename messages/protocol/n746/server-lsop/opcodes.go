package lsop

const (
	Init int16 = iota
	LoginFail
	AccountKicked
	LoginOk
	ServerList
	PlayFail = iota + 0x01
	PlayOk
	GGAuth
)
