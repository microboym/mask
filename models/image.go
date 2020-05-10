package models

import "encoding/base64"

// Image represents an image
type Image interface {
	// String encodes the image as base64
	String() string
}

// JPGImage represents an jpeg image
type JPGImage []byte

// String encodes the jpeg image as base64
func (i JPGImage) String() string {
	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(i)
}

// PNGImage represents an jpeg image
type PNGImage []byte

// String encodes the jpeg image as base64
func (i PNGImage) String() string {
	return "data:image/jpg;base64," + base64.StdEncoding.EncodeToString(i)
}
