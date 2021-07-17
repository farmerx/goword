package goword

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"

	"golang.org/x/net/html/charset"
)

// charsetTranscoderFn define for charset transfer code function
type charsetTranscoderFn func(charset string, input io.Reader) (reader io.Reader, err error)

// Options define the options for open wordprocessing.
type Options struct {
	Password string
}

// File define a populated wordprocessing file struct.
type File struct {
	sync.Mutex
	options       *Options
	ContentTypes  *contentTypes
	Relationships sync.Map
	CharsetReader charsetTranscoderFn
}

// OpenFile take the name of an wordprocessing file and returns a populated spreadsheet file struct
// for it. For example, open wordprocessing with password protection:
//
//    f, err := goword.OpenFile("demo.docx", goword.Options{Password: "password"})
//    if err != nil {
//        return
//    }
//
// Note that the goword just support decrypt and not support encrypt currently, the wordprocessing
// saved by Save and SaveAs will be without password unprotected.
func OpenFile(filepath string, opt ...Options) (*File, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return OpenReader(file, opt...)
}

func newFile() *File {
	return &File{
		Relationships: sync.Map{},
		CharsetReader: charset.NewReaderLabel,
	}

}

func OpenReader(r io.Reader, opt ...Options) (*File, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	f := newFile()
	if bytes.Contains(b, oleIdentifier) && len(opt) > 0 {
		for _, o := range opt {
			f.options = &o
		}
		b, err = Decrypt(b, f.options)
		if err != nil {
			return nil, fmt.Errorf("decrypted file failed")
		}
	}

	_, err = zip.NewReader(bytes.NewReader(b), int64(len(b)))
	if err != nil {
		return nil, err
	}
	return f, nil
}
