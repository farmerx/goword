package goword

import "errors"

var (
	// ErrEncrypt defined the error message on encryption spreadsheet.
	ErrEncrypt = errors.New("not support encryption currently")
	// ErrUnknownEncryptMechanism defined the error message on unsupport
	// encryption mechanism.
	ErrUnknownEncryptMechanism = errors.New("unknown encryption mechanism")
	// ErrUnsupportEncryptMechanism defined the error message on unsupport
	// encryption mechanism.
	ErrUnsupportEncryptMechanism = errors.New("unsupport encryption mechanism")
)
