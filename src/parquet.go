package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"sync"
	"time"

	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/source"
	"github.com/xitongsys/parquet-go/writer"
)

type Logger[TLog any] interface {
	AddLog(TLog)
	Start() error
	Stop()
}

type logger[TLog any] struct {
	logDir           string
	buffer           chan TLog
	wg               sync.WaitGroup
	flustIntervalSec int
}

func NewLogger[TLog any](logDir string, bufferSize, flustIntervalSec int) Logger[TLog] {
	l := &logger[TLog]{
		logDir:           logDir,
		buffer:           make(chan TLog, bufferSize),
		flustIntervalSec: flustIntervalSec,
		wg:               sync.WaitGroup{},
	}
	return l
}

func (l *logger[TLog]) AddLog(r TLog) {
	select {
	case l.buffer <- r:
	default:
		// try to write to buffer and drop log instead of getting blocked in case buffer is full
		log.Println("Dropping parquet log to avoid blocking")
	}
}

func GenerateParquetFileName(ending string) string {
	// create concerted file name
	return fmt.Sprintf("%s.%d.%s",
		time.Now().Format("2006-01-02-15-04-05"),
		time.Now().Nanosecond(),
		ending,
	)

}

func (l *logger[TLog]) Start() error {
	dirExists, err := exists(l.logDir)
	if err != nil {
		log.Fatalln(err)
	}
	if !dirExists {
		log.Fatalln("Log directory does not exist")
		if e := os.MkdirAll(l.logDir, 0700); e != nil {
			return fmt.Errorf("creating log directory: %w", e)
		}
	}

	pl, err := newParquetLogger[TLog](path.Join(l.logDir, GenerateParquetFileName("parquet")))
	if err != nil {
		log.Fatalln(err)
		return err
	}

	ticker := time.NewTicker(time.Duration(l.flustIntervalSec) * time.Second)

	l.wg.Add(1)

	go func() {

		defer l.wg.Done()

		for {
			select {
			case record, ok := <-l.buffer:

				if ok {
					if e := pl.AddLog(record); e != nil {
						log.Fatalf("Failed to write parquet log:%v\n", e)
					}
				} else {
					// chan is closed, flush and return
					pl.Close()
					fmt.Println("Closing parquet logger")
					ticker.Stop()
					return
				}
			case <-ticker.C:
				// flush and rotate log
				if pl.HasLogs() {
					pl.Close()
					pl, err = newParquetLogger[TLog](path.Join(l.logDir, GenerateParquetFileName("parquet")))
					if err != nil {
						log.Fatalf("Failed to create parquet log file: %v\n", err)
					}
				}
			}
		}
	}()

	return nil
}

func (l *logger[TLog]) Stop() {
	close(l.buffer)
	l.wg.Wait()
}

type parquetLogger[TLog any] struct {
	fw       source.ParquetFile
	pw       *writer.ParquetWriter
	filename string
}

const renamePattern = "%s.tmp"

func newParquetLogger[TLog any](filename string) (parquetLogger[TLog], error) {
	fw, err := local.NewLocalFileWriter(fmt.Sprintf(renamePattern, filename))
	if err != nil {
		return parquetLogger[TLog]{}, fmt.Errorf("creating local file writer: %w", err)
	}
	//parameters: writer, type of struct, size
	pw, err := writer.NewParquetWriter(fw, new(ParquetMessage), int64(runtime.NumCPU()))
	if err != nil {
		fw.Close()
		return parquetLogger[TLog]{}, fmt.Errorf("creating parquet writer: %w", err)
	}
	pw.RowGroupSize = 32 * 1024 * 1024 // 32M
	pw.PageSize = 8 * 1024             // 8K

	pw.CompressionType = parquet.CompressionCodec_SNAPPY
	return parquetLogger[TLog]{fw: fw, pw: pw, filename: filename}, nil
}

func (p *parquetLogger[TLog]) AddLog(r TLog) error {
	return p.pw.Write(r)
}

func (p *parquetLogger[TLog]) HasLogs() bool {
	return len(p.pw.Objs) > 0
}

// nolint:ifshort
func (p *parquetLogger[TLog]) Close() {
	hadLogs := p.HasLogs()
	if err := p.pw.WriteStop(); err != nil {
		log.Printf("Error closing ParquetWriter: %v\n", err)
	}
	if err := p.fw.Close(); err != nil {
		log.Printf("Error closing ParquetFile: %v\n", err)
	}
	if hadLogs {
		// strip .tmp suffix
		if err := os.Rename(fmt.Sprintf(renamePattern, p.filename), p.filename); err != nil {
			log.Printf("Error renaming ParquetFile: %v\n", err)
		}
	} else {
		// remove empty log file
		if err := os.Remove(fmt.Sprintf(renamePattern, p.filename)); err != nil {
			log.Printf("Error removing empty ParquetFile: %v\n", err)
		}
	}
}

func exists(dir string) (bool, error) {
	_, err := os.Stat(dir)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, fmt.Errorf("getting dir stats: %w", err)
}
