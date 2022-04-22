package protocol

type LoginState uint8

const (
	WAITING_FOR_START LoginState = iota
	WAITING_FOR_KEY   LoginState = iota
	AUTHENTICATING    LoginState = iota
	NEGOTIATING       LoginState = iota
	READY_TO_ACCEPT   LoginState = iota
	DELAY_ACCEPT      LoginState = iota
	ACCEPTED          LoginState = iota
)
