package listener

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"os"
	"unsafe"

	"github.com/kwarunek/keyboardstats/internal/codes"
	"github.com/kwarunek/keyboardstats/internal/store"
)

type Timeval struct {
	Sec  int64
	Usec int64
}

type InputEvent struct {
	Time  Timeval
	Type  uint16
	Code  uint16
	Value int32
}

type Listener struct {
	dev     string
	store   *store.Store
	verbose bool
}

func NewListener(dev string, store *store.Store, verbose bool) *Listener {
	return &Listener{dev: dev, store: store, verbose: verbose}
}

func (l *Listener) Listen() error {
	fd, err := os.Open(l.dev)
	if err != nil {
		log.Fatalf("Cannot open %s: %v", l.dev, err)
	}
	defer fd.Close()

	ie_size := int(unsafe.Sizeof(InputEvent{Time: Timeval{}}))

	for {
		buf := make([]byte, ie_size*2)
		_, err := io.ReadFull(fd, buf)
		if err != nil {
			log.Fatalf("Cannot read from %s: %v", l.dev, err)
		}
		for i := 0; i < len(buf); i += ie_size {
			ie := InputEvent{}
			err := binary.Read(bytes.NewReader(buf[i:i+ie_size]), binary.LittleEndian, &ie)
			if err != nil {
				log.Fatalf("Cannot read from %s: %v", l.dev, err)
			}
			if ie.Type == codes.EV_KEY && ie.Value == codes.PRESSED {
				if l.verbose {
					log.Printf("Key %s pressed\n", codes.KeyCodeToName[int(ie.Code)])
				}
				l.store.Record(codes.KeyCodeToName[int(ie.Code)])
			}
		}
	}
}
