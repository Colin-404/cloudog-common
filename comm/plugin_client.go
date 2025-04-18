package comm

import (
	"bufio"
	"encoding/binary"
	"io"
	"os"
	"sync"
	"time"

	"github.com/colin-404/cloudog-common/mproto"
	"google.golang.org/protobuf/proto"
)

type PluginCli struct {
	rx     io.ReadCloser
	tx     io.WriteCloser
	reader *bufio.Reader
	writer *bufio.Writer
	rmu    *sync.Mutex
	wmu    *sync.Mutex
}

func NewPluginCli() (c *PluginCli) {
	c = &PluginCli{
		rx: os.Stdin,
		tx: os.Stdout,
		// MAX_SIZE = 1 MB
		reader: bufio.NewReaderSize(os.NewFile(3, "pipe"), 1024*1024),
		writer: bufio.NewWriterSize(os.NewFile(4, "pipe"), 512*1024),
		rmu:    &sync.Mutex{},
		wmu:    &sync.Mutex{},
	}
	go func() {
		ticker := time.NewTicker(time.Millisecond * 200)
		defer ticker.Stop()
		for {
			<-ticker.C
			if err := c.Flush(); err != nil {
				break
			}
		}
	}()
	return
}

func (c *PluginCli) SendRawData(rec *mproto.RawData) (err error) {
	c.wmu.Lock()
	defer c.wmu.Unlock()
	size := proto.Size(rec)
	err = binary.Write(c.writer, binary.LittleEndian, uint32(size))
	if err != nil {
		return
	}
	var buf []byte
	buf, err = proto.Marshal(rec)
	if err != nil {
		return
	}
	_, err = c.writer.Write(buf)
	return
}
func (c *PluginCli) ReceiveTask() (task mproto.Task, err error) {
	c.rmu.Lock()
	defer c.rmu.Unlock()
	var len uint32
	err = binary.Read(c.reader, binary.LittleEndian, &len)
	if err != nil {
		return
	}
	var buf []byte
	buf, err = c.reader.Peek(int(len))
	if err != nil {
		return
	}
	_, err = c.reader.Discard(int(len))
	if err != nil {
		return
	}

	err = proto.Unmarshal(buf, &task)
	return
}
func (c *PluginCli) Flush() (err error) {
	c.wmu.Lock()
	defer c.wmu.Unlock()
	if c.writer.Buffered() != 0 {
		err = c.writer.Flush()
	}
	return
}

func (c *PluginCli) Close() {
	c.writer.Flush()
	c.rx.Close()
	c.tx.Close()
}
