package smile

import (
	"encoding/json"

	"github.com/brightcove/playback_go-smile/decode"
	"github.com/brightcove/playback_go-smile/domain"
)

func DecodeToJSON(smile []byte) (string, error) {
	obj, err := DecodeToObject(smile)
	if err != nil {
		return "", err
	}

	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

func DecodeToObject(smile []byte) (interface{}, error) {
	header, err := domain.DecodeHeader(smile)
	if err != nil {
		return "", err
	}

	var d decode.Decoder
	_, b, err := d.DecodeBytes(smile[header.SizeBytes:])
	return b, err
}
