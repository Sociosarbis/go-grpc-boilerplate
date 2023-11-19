package gocanal

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	proto "github.com/gogo/protobuf/proto"
	"github.com/segmentio/kafka-go"
	"github.com/sociosarbis/grpc/boilerplate/internal/config"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/logger"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/slice"
	"github.com/withlin/canal-go/protocol"
	pbe "github.com/withlin/canal-go/protocol/entry"
)

type Canal struct {
	reader *kafka.Reader
}

func NewCanal(config config.AppConfig) (*Canal, error) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{fmt.Sprintf("%s:%s", config.KafkaHost, config.KafkaPort)},
		Topic:     "example",
		Partition: 0,
		GroupID:   "example-consumer-1",
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})
	canal := &Canal{
		reader: reader,
	}
	return canal, nil
}

func (c *Canal) Run() {
	endChan := make(chan os.Signal, 1)
	signal.Notify(endChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		_, ok := <-endChan
		if ok {
			close(endChan)
			endChan = nil
		}
		c.Close()
	}()
	for {
		ctx := context.Background()
		m, err := c.reader.FetchMessage(ctx)
		if err != nil {
			logger.Err(err, "reader.ReadMessage")
			break
		}
		message, err := protocol.Decode(m.Value, false)
		if err != nil {
			logger.Err(err, "decode pbe.Message")
		}
		batchId := message.Id
		if batchId == -1 || len(message.Entries) <= 0 {
			time.Sleep(200 * time.Millisecond)
			continue
		}
		printEntry(message.Entries)
		err = c.reader.CommitMessages(ctx, m)
		if err != nil {
			logger.Err(err, "reader.CommitMessages")
		}
	}
	if endChan != nil {
		close(endChan)
	}
}

func (c *Canal) Close() {
	if err := c.reader.Close(); err != nil {
		logger.Err(err, "reader.Close")
	}
}

func printEntry(entrys []pbe.Entry) {

	for _, entry := range entrys {
		if entry.GetEntryType() == pbe.EntryType_TRANSACTIONBEGIN || entry.GetEntryType() == pbe.EntryType_TRANSACTIONEND {
			continue
		}
		rowChange := new(pbe.RowChange)

		err := proto.Unmarshal(entry.GetStoreValue(), rowChange)
		checkError(err)
		if rowChange.GetEventType() != 0 {
			eventType := rowChange.GetEventType()
			header := entry.GetHeader()
			logger.Info(fmt.Sprintf("================> binlog[%s : %d],name[%s,%s], eventType: %s", header.GetLogfileName(), header.GetLogfileOffset(), header.GetSchemaName(), header.GetTableName(), header.GetEventType()))

			for _, rowData := range rowChange.GetRowDatas() {
				if eventType == pbe.EventType_DELETE {
					printColumn(rowData.GetBeforeColumns())
				} else if eventType == pbe.EventType_INSERT {
					printColumn(rowData.GetAfterColumns())
				} else {
					fmt.Println("-------> before")
					printColumn(rowData.GetBeforeColumns())
					fmt.Println("-------> after")
					printColumn(rowData.GetAfterColumns())
				}
			}
		}
	}
}

func printColumn(columns []*pbe.Column) {
	msg := "\n" + strings.Join(slice.Map(columns, func(col *pbe.Column) string {
		return fmt.Sprintf("%s : %s  update= %t", col.GetName(), col.GetValue(), col.GetUpdated())
	}), "\n")
	logger.Info(msg)
}

func checkError(err error) {
	if err != nil {
		logger.Err(err, "checkError")
	}
}
