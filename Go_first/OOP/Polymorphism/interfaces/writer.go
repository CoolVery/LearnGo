package interfaces

type Writer interface {
	Write(p []byte) int
}