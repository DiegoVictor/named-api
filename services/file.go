package services

import "os"

type File struct{}

func (f File) Append(file *os.File, text string) (int, error) {
	return file.WriteString("\n" + text)
}

func (f File) Open(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
}
