package lsop

const (
	Init uint16 = iota
	LoginFail
	AccountKicked
	LoginOk
	ServerList
	PlayFail = iota + 0x01
	PlayOk
	GGAuth
)
