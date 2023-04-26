package main

import (
	"encoding/json"
	"testing"

	sonic "github.com/bytedance/sonic"
	goccy "github.com/goccy/go-json"
	sonnet "github.com/sugawarayuuta/sonnet"
)

// var in Person = Person{"NAME", "31uEbMgunupShBVTewXjtqbBv5MndwfXhb", 1000}
var in string = `{ "name": "NAME", "address": "31uEbMgunupShBVTewXjtqbBv5MndwfXhb", "age": "1000" }`

func BenchmarkJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsonB := EncodeJSON(in)
		DecodeJSON(jsonB)
	}
}

func BenchmarkGoccyJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsonB := EncodeGoccyJSON(in)
		DecodeGoccyJSON(jsonB)
	}
}

func BenchmarkSonicJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsonB := EncodeSonicJSON(in)
		DecodeSonicJSON(jsonB)
	}
}

func BenchmarkSonnetJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsonB := EncodeSonnetJSON(in)
		DecodeSonnetJSON(jsonB)
	}
}

// func BenchmarkMsgPack(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		jsonB := EncodeMsgPack(in)
// 		DecodeMsgPack(jsonB)
// 	}
// }

// func BenchmarkMsgP(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		jsonB := EncodeMsgP(in)
// 		_ = DecodeMsgP(jsonB)
// 	}
// }

// func BenchmarkMsgPackGen(b *testing.B) {
// 	RegisterGeneratedResolver()
// 	for i := 0; i < b.N; i++ {
// 		jsonB := EncodeMsgPackgen(in)
// 		_ = DecodeMsgPackgen(jsonB)
// 	}
// }

func EncodeJSON(message string) (out []byte) {
	out, _ = json.Marshal(message)
	return out
}

func DecodeJSON(b []byte) (out Person) {
	_ = json.Unmarshal(b, &out)
	return
}

func EncodeGoccyJSON(message string) (out []byte) {
	out, _ = goccy.Marshal(message)
	return out
}

func DecodeGoccyJSON(b []byte) (out Person) {
	_ = goccy.Unmarshal(b, &out)
	return
}

func EncodeSonicJSON(message string) (out []byte) {
	out, _ = sonic.Marshal(message)
	return out
}

func DecodeSonicJSON(b []byte) (out Person) {
	_ = sonic.Unmarshal(b, &out)
	return
}

func EncodeSonnetJSON(message string) (out []byte) {
	out, _ = sonnet.Marshal(message)
	return out
}

func DecodeSonnetJSON(b []byte) (out Person) {
	_ = sonnet.Unmarshal(b, &out)
	return
}

// func EncodeMsgPack(message Person) (out []byte) {
// 	out, _ = msgpack.Marshal(message)
// 	return out
// }

// func DecodeMsgPack(b []byte) (out Person) {
// 	_ = msgpack.Unmarshal(b, &out)
// 	return out
// }

// func EncodeMsgP(message Person) (out []byte) {
// 	out, _ = message.MarshalMsg(nil)
// 	return out
// }

// func DecodeMsgP(data []byte) (out Person) {
// 	_, _ = out.UnmarshalMsg(data)
// 	return
// }

// func EncodeMsgPackgen(message Person) (out []byte) {
// 	out, _ = msgpackgen.Marshal(message)
// 	return out
// }

// func DecodeMsgPackgen(data []byte) (out Person) {
// 	_ = msgpackgen.Unmarshal(data, &out)
// 	return
// }
