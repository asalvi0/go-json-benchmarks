// Code generated by msgpackgen. DO NOT EDIT.

package main

import (
	"fmt"
	msgpack "github.com/shamaton/msgpackgen/msgpack"
	dec "github.com/shamaton/msgpackgen/msgpack/dec"
	enc "github.com/shamaton/msgpackgen/msgpack/enc"
)

// RegisterGeneratedResolver registers generated resolver.
func RegisterGeneratedResolver() {
	msgpack.SetResolver(___encodeAsMap, ___encodeAsArray, ___decodeAsMap, ___decodeAsArray)
}

// encode
func ___encode(i interface{}) ([]byte, error) {
	if msgpack.StructAsArray() {
		return ___encodeAsArray(i)
	} else {
		return ___encodeAsMap(i)
	}
}

// encodeAsArray
func ___encodeAsArray(i interface{}) ([]byte, error) {
	switch v := i.(type) {
	case Person:
		encoder := enc.NewEncoder()
		size, err := ___calcArraySizePerson_9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08(v, encoder)
		if err != nil {
			return nil, err
		}
		encoder.MakeBytes(size)
		b, offset, err := ___encodeArrayPerson_9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08(v, encoder, 0)
		if err != nil {
			return nil, err
		}
		if size != offset {
			return nil, fmt.Errorf("%s size / offset different %d : %d", "Person", size, offset)
		}
		return b, err
	case *Person:
		encoder := enc.NewEncoder()
		size, err := ___calcArraySizePerson_9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08(*v, encoder)
		if err != nil {
			return nil, err
		}
		encoder.MakeBytes(size)
		b, offset, err := ___encodeArrayPerson_9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08(*v, encoder, 0)
		if err != nil {
			return nil, err
		}
		if size != offset {
			return nil, fmt.Errorf("%s size / offset different %d : %d", "Person", size, offset)
		}
		return b, err
	}
	return nil, nil
}

// encodeAsMap
func ___encodeAsMap(i interface{}) ([]byte, error) {
	switch v := i.(type) {
	case Person:
		encoder := enc.NewEncoder()
		size, err := ___calcMapSizePerson_9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08(v, encoder)
		if err != nil {
			return nil, err
		}
		encoder.MakeBytes(size)
		b, offset, err := ___encodeMapPerson_9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08(v, encoder, 0)
		if err != nil {
			return nil, err
		}
		if size != offset {
			return nil, fmt.Errorf("%s size / offset different %d : %d", "Person", size, offset)
		}
		return b, err
	case *Person:
		encoder := enc.NewEncoder()
		size, err := ___calcMapSizePerson_9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08(*v, encoder)
		if err != nil {
			return nil, err
		}
		encoder.MakeBytes(size)
		b, offset, err := ___encodeMapPerson_9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08(*v, encoder, 0)
		if err != nil {
			return nil, err
		}
		if size != offset {
			return nil, fmt.Errorf("%s size / offset different %d : %d", "Person", size, offset)
		}
		return b, err
	}
	return nil, nil
}

// decode
func ___decode(data []byte, i interface{}) (bool, error) {
	if msgpack.StructAsArray() {
		return ___decodeAsArray(data, i)
	} else {
		return ___decodeAsMap(data, i)
	}
}

// decodeAsArray
func ___decodeAsArray(data []byte, i interface{}) (bool, error) {
	switch v := i.(type) {
	case *Person:
		decoder := dec.NewDecoder(data)
		offset, err := ___decodeArrayPerson_9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08(v, decoder, 0)
		if err == nil && offset != decoder.Len() {
			return true, fmt.Errorf("read length is different [%d] [%d] ", offset, decoder.Len())
		}
		return true, err
	case **Person:
		decoder := dec.NewDecoder(data)
		offset, err := ___decodeArrayPerson_9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08(*v, decoder, 0)
		if err == nil && offset != decoder.Len() {
			return true, fmt.Errorf("read length is different [%d] [%d] ", offset, decoder.Len())
		}
		return true, err
	}
	return false, nil
}

// decodeAsMap
func ___decodeAsMap(data []byte, i interface{}) (bool, error) {
	switch v := i.(type) {
	case *Person:
		decoder := dec.NewDecoder(data)
		offset, err := ___decodeMapPerson_9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08(v, decoder, 0)
		if err == nil && offset != decoder.Len() {
			return true, fmt.Errorf("read length is different [%d] [%d] ", offset, decoder.Len())
		}
		return true, err
	case **Person:
		decoder := dec.NewDecoder(data)
		offset, err := ___decodeMapPerson_9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08(*v, decoder, 0)
		if err == nil && offset != decoder.Len() {
			return true, fmt.Errorf("read length is different [%d] [%d] ", offset, decoder.Len())
		}
		return true, err
	}
	return false, nil
}

// calculate size from test.Person
func ___calcArraySizePerson_9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08(v Person, encoder *enc.Encoder) (int, error) {
	size := 0
	size += encoder.CalcStructHeaderFix(4)
	size += encoder.CalcString(v.Name)
	size += encoder.CalcString(v.Address)
	size += encoder.CalcInt(v.Age)
	return size, nil
}

// calculate size from test.Person
func ___calcMapSizePerson_9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08(v Person, encoder *enc.Encoder) (int, error) {
	size := 0
	size += encoder.CalcStructHeaderFix(4)
	size += encoder.CalcStringFix(4)
	size += encoder.CalcString(v.Name)
	size += encoder.CalcStringFix(7)
	size += encoder.CalcString(v.Address)
	size += encoder.CalcStringFix(3)
	size += encoder.CalcInt(v.Age)
	size += encoder.CalcStringFix(6)
	return size, nil
}

// encode from test.Person
func ___encodeArrayPerson_9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08(v Person, encoder *enc.Encoder, offset int) ([]byte, int, error) {
	var err error
	offset = encoder.WriteStructHeaderFixAsArray(4, offset)
	offset = encoder.WriteString(v.Name, offset)
	offset = encoder.WriteString(v.Address, offset)
	offset = encoder.WriteInt(v.Age, offset)
	return encoder.EncodedBytes(), offset, err
}

// encode from test.Person
func ___encodeMapPerson_9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08(v Person, encoder *enc.Encoder, offset int) ([]byte, int, error) {
	var err error
	offset = encoder.WriteStructHeaderFixAsMap(4, offset)
	offset = encoder.WriteStringFix("Name", 4, offset)
	offset = encoder.WriteString(v.Name, offset)
	offset = encoder.WriteStringFix("Address", 7, offset)
	offset = encoder.WriteString(v.Address, offset)
	offset = encoder.WriteStringFix("Age", 3, offset)
	offset = encoder.WriteInt(v.Age, offset)
	return encoder.EncodedBytes(), offset, err
}

// decode to test.Person
func ___decodeArrayPerson_9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08(v *Person, decoder *dec.Decoder, offset int) (int, error) {
	offset, err := decoder.CheckStructHeader(4, offset)
	if err != nil {
		return 0, err
	}
	{
		var vv string
		vv, offset, err = decoder.AsString(offset)
		if err != nil {
			return 0, err
		}
		v.Name = vv
	}
	{
		var vv string
		vv, offset, err = decoder.AsString(offset)
		if err != nil {
			return 0, err
		}
		v.Address = vv
	}
	{
		var vv int
		vv, offset, err = decoder.AsInt(offset)
		if err != nil {
			return 0, err
		}
		v.Age = vv
	}
	return offset, err
}

// decode to test.Person
func ___decodeMapPerson_9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08(v *Person, decoder *dec.Decoder, offset int) (int, error) {
	keys := [][]byte{
		{uint8(0x4e), uint8(0x61), uint8(0x6d), uint8(0x65)},                                        // Name
		{uint8(0x41), uint8(0x64), uint8(0x64), uint8(0x72), uint8(0x65), uint8(0x73), uint8(0x73)}, // Address
		{uint8(0x41), uint8(0x67), uint8(0x65)},                                                     // Age
		{uint8(0x48), uint8(0x69), uint8(0x64), uint8(0x64), uint8(0x65), uint8(0x6e)},              // Hidden
	}
	offset, err := decoder.CheckStructHeader(4, offset)
	if err != nil {
		return 0, err
	}
	count := 0
	for count < 4 {
		var dataKey []byte
		dataKey, offset, err = decoder.AsStringBytes(offset)
		if err != nil {
			return 0, err
		}
		fieldIndex := -1
		for i, key := range keys {
			if len(dataKey) != len(key) {
				continue
			}
			fieldIndex = i
			for dataKeyIndex := range dataKey {
				if dataKey[dataKeyIndex] != key[dataKeyIndex] {
					fieldIndex = -1
					break
				}
			}
			if fieldIndex >= 0 {
				break
			}
		}
		switch fieldIndex {
		case 0:
			{
				var vv string
				vv, offset, err = decoder.AsString(offset)
				if err != nil {
					return 0, err
				}
				v.Name = vv
			}
			count++
		case 1:
			{
				var vv string
				vv, offset, err = decoder.AsString(offset)
				if err != nil {
					return 0, err
				}
				v.Address = vv
			}
			count++
		case 2:
			{
				var vv int
				vv, offset, err = decoder.AsInt(offset)
				if err != nil {
					return 0, err
				}
				v.Age = vv
			}
			count++
		default:
			return 0, fmt.Errorf("unknown key[%s] found", string(dataKey))
		}
	}
	return offset, err
}