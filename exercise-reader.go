// https://gist.github.com/zyxar/2317744#comment-1363886

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (read MyReader) Read(b []byte) (n int, err error) {
	b[0] = 'A'
	return 1, nil
}


func main() {
	reader.Validate(MyReader{})
}
