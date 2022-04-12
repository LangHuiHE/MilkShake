package models

import (
	"encoding/base64"
	"errors"

	qrcode "github.com/skip2/go-qrcode"
)

type QrCode struct {
	Msg      string
	Png      []byte
	FileName string
	Size     int
	B64      string
}

func (qr *QrCode) checker() error {
	if qr.FileName == "" {
		return errors.New("qrcode: Missing file name for storying the image")
	}
	if qr.Msg == "" {
		return errors.New("qrcode: Missing message for the image")
	}
	if qr.Size == 0 {
		qr.Size = 256
	}
	return nil
}

// size=0 using default size 256
func (qr *QrCode) createPNG() error {
	if qr.Msg == "" {
		return errors.New("qrcode: Missing message for the image")
	}
	if qr.Size == 0 {
		qr.Size = 256
	}
	png, err := qrcode.Encode(qr.Msg, qrcode.Medium, qr.Size)
	if err == nil {
		qr.Png = png
	}
	return err
}

func (qr *QrCode) CreatePNGfile() error {
	if err := qr.checker(); err == nil {
		return qrcode.WriteFile(qr.Msg, qrcode.Medium, qr.Size, qr.FileName)
	} else {
		return err
	}
}

func (qr *QrCode) toBase64() string {
	return base64.StdEncoding.EncodeToString(qr.Png)
}

func (qr *QrCode) GeneratePNGonBase64() (string, error) {
	if err := qr.createPNG(); err == nil {
		return base64.StdEncoding.EncodeToString(qr.Png), nil
	} else {
		return "", err
	}
}
