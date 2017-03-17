package lsop

const (
	RequestAuthLogin   int16 = iota
	RequestServerLogin       = iota + 0x01
	RequestServerList        = iota + 0x02
	RequestGGAuth            = iota + 0x01
)
