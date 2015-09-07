package main

import (
    "log"
    "hash/fnv"
    "bytes"
    "strconv"
    "encoding/binary"
)

const (
    SALT_64  = "1V3S#F"
    SALT_32  = "AB#*FP"
    S_INT32  = 1 << iota
    S_FLOAT  = 1 << iota
    S_CHAR32 = 1 << iota
)

// Using TCP can bypass series since TCP will auto-bundle events.
type (
    TCPHeader struct {
        Length      uint8
        SeqNum      uint32
        ID          uint64
        SeriesID    uint64
        Type        uint8
        Events      []TCPEvent
    }
    TCPEvent struct {
        Time        int64
        Data        []byte
    }
    Dispatcher struct {
        ID          uint64
        Headers     map[uint64]*TCPHeader
    }
)

// Utility Functions
func Hash64(text string) uint64 {
    h := fnv.New64a()
    h.Write([]byte(SALT_64 + text))
    return h.Sum64()
}

func Hash32(text string) uint32 {
    h := fnv.New32a()
    h.Write([]byte(SALT_32 + text))
    return h.Sum32()
}

func GetType(e *SyncEvent) int {
    v := e.Value
    if s, err := strconv.ParseInt(v, 10, 32); err == nil {
        log.Printf("%T, %v\n", s, s)
        return S_INT32
    }
    if s, err := strconv.ParseFloat(v, 32); err == nil {
        log.Printf("%T, %v\n", s, s)
        return S_FLOAT
    }
    return S_CHAR32
}

// Dispatcher
func NewDispatcher(key string) *Dispatcher {
    return &Dispatcher{ID: Hash64(key), Headers: make(map[uint64]*TCPHeader)}
}

func (d *Dispatcher) GetTCPHeader(e *SyncEvent) *TCPHeader {
    if t,ok := d.Headers[e.ID]; ok {
        return t
    }
    header := &TCPHeader{ID: d.ID, SeriesID: e.ID, Type: uint8(GetType(e)) }
    d.Headers[e.ID] = header

    log.Println(header)

    return header
}

func (d *Dispatcher) HandleEvent(e *SyncEvent) {
    tcp := d.GetTCPHeader(e)
    tcp.HandleEvent(e)
}

func (d *Dispatcher) Flush() {
    for _,t := range d.Headers {
        t.Flush()
    }
}

// TCP Header
func (t *TCPHeader) HandleEvent(e *SyncEvent) {
    buf := new(bytes.Buffer)
    v := e.Value
    switch t.Type {
    case S_INT32:
        val, _ := strconv.ParseInt(v, 10, 32)
        err := binary.Write(buf, binary.LittleEndian, val)
        if err != nil {
            log.Fatal(err)
        }
        //log.Printf("%d, %x\n", val, buf.Bytes())
    case S_FLOAT:
        val, _ := strconv.ParseFloat(v, 32)
        err := binary.Write(buf, binary.LittleEndian, val)
        if err != nil {
            log.Fatal(err)
        }
        //log.Printf("%f, %x\n", val, buf.Bytes())
    case S_CHAR32:
        val := make([]byte, 32)
        copy(val, []byte(v))
        err := binary.Write(buf, binary.LittleEndian, val)
        if err != nil {
            log.Fatal(err)
        }
        //log.Printf("%s, %x\n", val, buf.Bytes())
    default:
        log.Fatal("No Type Specified for Event %+v", e)
    }
    t.Events = append(t.Events, TCPEvent{Time: e.Time, Data: buf.Bytes()})
    log.Println(t)
}

func (t *TCPHeader) Flush() {
    t.Events = nil
}

func (t *TCPHeader) Marshal() []byte {
    buf := new(bytes.Buffer)
    out := buf.Bytes()
    return out
}