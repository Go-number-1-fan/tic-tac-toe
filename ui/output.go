package ui

type Output interface {
	Write(str string)
	Clear()
}
