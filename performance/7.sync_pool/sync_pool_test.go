package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"sync"
	"testing"
)

func Log(w io.Writer, key, val string) {
	var b bytes.Buffer
	b.WriteByte(' ')
	b.WriteString(key)
	b.WriteByte('=')
	b.WriteString(val)
	w.Write(b.Bytes())
}

func BenchmarkLog(b *testing.B) {
	writer := ioutil.Discard
	for i := 0; i < b.N; i++ {
		Log(writer, "key", "value")
	}
}

var pool = sync.Pool{
	New: func() interface{} {
		b := make([]byte, 4096)
		return bytes.NewBuffer(b)
	},
}

func LogSyncPool(w io.Writer, key, val string) {
	b := pool.Get().(*bytes.Buffer)
	b.Reset()
	b.WriteByte(' ')
	b.WriteString(key)
	b.WriteByte('=')
	b.WriteString(val)
	w.Write(b.Bytes())
	pool.Put(b)
}

func BenchmarkLogSyncPool(b *testing.B) {
	writer := ioutil.Discard
	for i := 0; i < b.N; i++ {
		LogSyncPool(writer, "key", "value")
	}
}
