package main

import (
  "bytes"
  "encoding/binary"
  "log"
)

func main() {
  str := "63520-"
  buf := new(bytes.Buffer)
  val := make([]byte, 16)
  copy(val, []byte(str))
  binary.Write(buf, binary.LittleEndian, val)
  log.Println(buf)
}