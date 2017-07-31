package main

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"log"
)

// Key ...
type Key struct {
	Channel   int64
	StartTime int64
	ID        []int64
}

func main() {
	var id int64 = 1
	b := bytes.NewBuffer(make([]byte, 0, 24))
	binary.Write(b, binary.LittleEndian)
	// binary.Write(b, binary.LittleEndian, key.Body)
	// binary.Write(b, binary.LittleEndian, key.StatusCode)
	str := base64.StdEncoding.EncodeToString(b.Bytes())
	//
	log.Println(str, b.Len(), len(str))
	//
	// var res Key
	// buf, _ := base64.StdEncoding.DecodeString(str)
	// out := bytes.NewBuffer(buf)
	// binary.Read(out, binary.LittleEndian, &res.Channel)
	// binary.Read(out, binary.LittleEndian, &res.StartTime)
	// for index := range res.ID {
	// 	binary.Read(out, binary.LittleEndian, &res.ID[index])
	// }
	//
	// binary.Read(out, binary.LittleEndian, &res.ID)
	//
	// log.Println(res)
}
