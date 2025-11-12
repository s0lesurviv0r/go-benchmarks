package models

import (
	"encoding/binary"
	"io"
)

type Nested struct {
	Id   uint64 `json:"id";cbor:"1,keyasint"`
	Name string `json:"name":cbor:"2,keyasint"`
}

// WriteTo writes the Nested struct to the writer in binary format
func (n *Nested) WriteTo(w io.Writer) (int64, error) {
	var total int64

	// Write Id (8 bytes)
	if err := binary.Write(w, binary.LittleEndian, n.Id); err != nil {
		return total, err
	}
	total += 8

	// Write Name length (4 bytes) and Name bytes
	nameLen := uint32(len(n.Name))
	if err := binary.Write(w, binary.LittleEndian, nameLen); err != nil {
		return total, err
	}
	total += 4

	written, err := w.Write([]byte(n.Name))
	if err != nil {
		return total, err
	}
	total += int64(written)

	return total, nil
}

// ReadFrom reads the Nested struct from the reader in binary format
func (n *Nested) ReadFrom(r io.Reader) (int64, error) {
	var total int64

	// Read Id (8 bytes)
	if err := binary.Read(r, binary.LittleEndian, &n.Id); err != nil {
		return total, err
	}
	total += 8

	// Read Name length (4 bytes)
	var nameLen uint32
	if err := binary.Read(r, binary.LittleEndian, &nameLen); err != nil {
		return total, err
	}
	total += 4

	// Read Name bytes
	nameBuf := make([]byte, nameLen)
	read, err := io.ReadFull(r, nameBuf)
	if err != nil {
		return total, err
	}
	total += int64(read)
	n.Name = string(nameBuf)

	return total, nil
}

type Object struct {
	Id     uint64 `json:"id";cbor:"1,keyasint"`
	Name   string `json:"name";cbor:"2,keyasint"`
	Nested Nested `json:"nested":cbor:"3,keyasint"`
}

// WriteTo writes the Object struct to the writer in binary format
func (o *Object) WriteTo(w io.Writer) (int64, error) {
	var total int64

	// Write Id (8 bytes)
	if err := binary.Write(w, binary.LittleEndian, o.Id); err != nil {
		return total, err
	}
	total += 8

	// Write Name length (4 bytes) and Name bytes
	nameLen := uint32(len(o.Name))
	if err := binary.Write(w, binary.LittleEndian, nameLen); err != nil {
		return total, err
	}
	total += 4

	written, err := w.Write([]byte(o.Name))
	if err != nil {
		return total, err
	}
	total += int64(written)

	// Write Nested
	nestedWritten, err := o.Nested.WriteTo(w)
	if err != nil {
		return total, err
	}
	total += nestedWritten

	return total, nil
}

// ReadFrom reads the Object struct from the reader in binary format
func (o *Object) ReadFrom(r io.Reader) (int64, error) {
	var total int64

	// Read Id (8 bytes)
	if err := binary.Read(r, binary.LittleEndian, &o.Id); err != nil {
		return total, err
	}
	total += 8

	// Read Name length (4 bytes)
	var nameLen uint32
	if err := binary.Read(r, binary.LittleEndian, &nameLen); err != nil {
		return total, err
	}
	total += 4

	// Read Name bytes
	nameBuf := make([]byte, nameLen)
	read, err := io.ReadFull(r, nameBuf)
	if err != nil {
		return total, err
	}
	total += int64(read)
	o.Name = string(nameBuf)

	// Read Nested
	nestedRead, err := o.Nested.ReadFrom(r)
	if err != nil {
		return total, err
	}
	total += nestedRead

	return total, nil
}
