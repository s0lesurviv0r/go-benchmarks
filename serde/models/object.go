package models

type Nested struct {
	Id   uint64 `json:"id";cbor:"1,keyasint"`
	Name string `json:"name":cbor:"2,keyasint"`
}

type Object struct {
	Id     uint64 `json:"id";cbor:"1,keyasint"`
	Name   string `json:"name";cbor:"2,keyasint"`
	Nested Nested `json:"nested":cbor:"3,keyasint"`
}
