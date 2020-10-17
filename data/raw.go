package data

import (
	"bufio"
	"os"
)

//RawData ...
//	The raw data provider
type RawData struct {
	FilePath string
	r *bufio.Reader
}

//LoadData ...
//	init the reader for the single file
func (r *RawData) LoadData() error {
	f, err := os.Open(r.FilePath)
	if err != nil {
		return err
	}
	r.r = bufio.NewReader(f)
	return nil
}

//ReadData ...
//	implements the interface, to read a line from the file
func (r *RawData) ReadData() interface{} {
	v, err := r.ReadLine()
	if err != nil {
		return nil
	}
	return v
}

//ReadLine ...
//	read a single line from the file
func (r *RawData) ReadLine() (string, error) {

	var line string
	line, err := r.r.ReadString('\n')
	if err != nil {
		return "", err
	}

	return line, nil;
}

func (r *RawData) SetPath(p string) {
	r.FilePath = p
}

