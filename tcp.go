package main

import (
    "log"
    "hash/fnv"
    "bytes"
    "strconv"
    "time"
    "fmt"
    "encoding/binary"
    "github.com/gdamore/mangos"
    "github.com/gdamore/mangos/protocol/push"
    "github.com/gdamore/mangos/transport/ipc"
    "github.com/gdamore/mangos/transport/tcp"
)

const (
    SALT_64  = "1V3S#F"
    SALT_32  = "AB#*FP"
    //TCP_URL = "tcp://localhost:40899"
    TCP_URL = "tcp://52.8.222.214:40899"
    TCP_INTERVAL = time.Second * 1
    TCP_TIMEOUT  = time.Second * 5
    TCP_RECON    = time.Second * 10
)

// TCP Data represents a single series (ID=seriesID)
// Relying on TCP protocol to auto-bundle events.
type (
    TCPData struct {
        Length      uint8
        Type        uint8
        Token       uint32
        ID          uint64
        Events      []TCPEvent
    }
    TCPEvent struct {
        Time        int64
        Data        []byte
    }
    // Dial TCP with cluster ID
    TCPDispatcher struct {
        ID          uint32
        Socket      mangos.Socket
        Timer       chan bool
        Recon       chan bool
        Events      chan UploadEvent
        Headers     map[uint64]*TCPData
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

func HashSeriesID(token uint32, group string, desc string) uint64 {
    strID := fmt.Sprintf("%d.%s.%s", token, group, desc)
    hid := Hash64(strID)
    return hid
}

// TCPDispatcher
func NewTCPDispatcher(key string, events chan UploadEvent) *TCPDispatcher {
    return &TCPDispatcher{ID: Hash32(key), Headers: make(map[uint64]*TCPData), 
        Recon: make(chan bool, 1), Timer: make(chan bool, 1), Events: events}
}

func (d *TCPDispatcher) Connect() {
    var err error

    if d.Socket, err = push.NewSocket(); err != nil {
        log.Fatalf("Push Socket Failure: %s", err.Error())
    }

    d.Socket.AddTransport(ipc.NewTransport())
    d.Socket.AddTransport(tcp.NewTransport())

    if err = d.Socket.Dial(TCP_URL); err != nil {
        log.Fatalf("TCP dial failed: %s", err.Error())
    }

    d.Socket.SetOption(mangos.OptionSendDeadline, TCP_TIMEOUT)

    log.Printf("TCPDispatcher connected to TCP: %s", TCP_URL)
}

func (d *TCPDispatcher) StartTimer() {
    for {
        time.Sleep(TCP_INTERVAL)
        d.Timer <- true
    }
}

func (d *TCPDispatcher) Run() {
    d.Connect()
    defer d.Disconnect()
    go d.StartTimer()
    go d.AutoReconnect()

    changed := false
    defer d.Socket.Close()

    for {
        select {
        case sEvent := <-d.Events:
            d.HandleEvent(&sEvent)
            changed = true
        case <-d.Timer:
            if changed {
                d.Flush()
                changed = false
            }
        case <-d.Recon:
            d.Connect()
        }
    }
}

func (d *TCPDispatcher) GetTCPData(e *UploadEvent) *TCPData {
    if t,ok := d.Headers[e.ID]; ok {
        return t
    }
    header := &TCPData{ID: e.ID, Type: e.Type, Token: d.ID}
    d.Headers[e.ID] = header

    return header
}

func (d *TCPDispatcher) HandleEvent(e *UploadEvent) {
    tcp := d.GetTCPData(e)
    tcp.HandleEvent(e)
}

func (d *TCPDispatcher) Disconnect() {
    if d.Socket != nil {
        log.Println("TCPDispatcher disconnecting...")
        d.Socket.Close()
        d.Socket = nil
    }   
}

func (d *TCPDispatcher) AutoReconnect() {
    for {
        time.Sleep(TCP_RECON)
        if d.Socket == nil {
            d.Recon <- true
        }
    }
}

func (d *TCPDispatcher) Flush() {
    for _,t := range d.Headers {
        if !t.Flush(d.Socket) {
            d.Disconnect()
        }
    }
}

func (d *TCPDispatcher) GetNumBytes() int {
    n := 0
    for _,t := range d.Headers {
        n += len(t.Marshal())
    }
    return n
}

// TCP Header
func (t *TCPData) HandleEvent(e *UploadEvent) {
    var buf bytes.Buffer
    v := e.Value
    t.Type = e.Type
    switch t.Type {
    case S_INT:
        val, _ := strconv.ParseInt(v, 10, 32)
        err := binary.Write(&buf, binary.LittleEndian, int32(val))
        if err != nil {
            log.Fatal(err)
        }
    case S_FLOAT:
        val, _ := strconv.ParseFloat(v, 32)
        err := binary.Write(&buf, binary.LittleEndian, float32(val))
        if err != nil {
            log.Fatal(err)
        }
    case S_CHAR:
        val := make([]byte, 16)
        copy(val, []byte(v))
        // Force string truncation to 15 chars
        val[14] = 0
        err := binary.Write(&buf, binary.LittleEndian, val)
        if err != nil {
            log.Fatal(err)
        }
    default:
        log.Fatal("No Type Specified for Event %+v", e)
    }
    t.Events = append(t.Events, TCPEvent{Time: e.Time, Data: buf.Bytes()})
    t.Length = uint8(len(t.Events))
}

// Send data to TCP server
func (t *TCPData) Flush(sock mangos.Socket) bool {
    msg := t.Marshal()
    t.Events = nil

    if sock == nil {
        return false
    }
    if err := sock.Send(msg); err != nil {
        log.Printf("Cannot push message on socket: %s", err.Error())
        return false
    }
    return true
}

func (t *TCPData) Marshal() []byte {
    buf := new(bytes.Buffer)
    binary.Write(buf, binary.LittleEndian, t.Length)
    binary.Write(buf, binary.LittleEndian, t.Type)
    binary.Write(buf, binary.LittleEndian, t.Token)
    binary.Write(buf, binary.LittleEndian, t.ID)

    for _,e := range t.Events {
        binary.Write(buf, binary.LittleEndian, e.Time)
        binary.Write(buf, binary.LittleEndian, e.Data)
    }

    out := buf.Bytes()
    return out
}

