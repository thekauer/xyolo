package handler

import (
	b64 "encoding/base64"
	"encoding/json"
	"net/http"
)

type Request struct {
	img string
}

type Payload struct {
	found bool
	l     string
	x     int
	y     int
	w     int
	h     int
	c     int
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "400 - Bad request!", http.StatusBadRequest)
		return
	}
	var req Request

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	image, err := b64.StdEncoding.DecodeString(req.img)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_ = image

	w.Header().Set("Content-Type", "application/json")
	//dummy
	payload := Payload{
		found: true,
		l:     "a",
		x:     10,
		y:     10,
		w:     28,
		h:     28,
		c:     96,
	}

	json.NewEncoder(w).Encode(payload)
}
