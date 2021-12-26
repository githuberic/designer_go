package e3

type AbstractMessage interface {
	SendMessage(text, to string)
}
