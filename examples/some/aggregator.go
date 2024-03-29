// 用 Golang 的 channel 实现消息的批量处理
package some

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
	"time"
)

func withBatchSize(bs int) SetAggregatorOptionFunc {
	return func(option AggregatorOption) AggregatorOption {
		option.BatchSize = bs
		return option
	}
}
func withChannelBufferSize(cb int) SetAggregatorOptionFunc {
	return func(option AggregatorOption) AggregatorOption {
		option.ChannelBufferSize = cb
		return option
	}
}
func withLingerTime(t time.Duration) SetAggregatorOptionFunc {
	return func(option AggregatorOption) AggregatorOption {
		option.LingerTime = t
		return option
	}
}
func withWorkers(w int) SetAggregatorOptionFunc {
	return func(option AggregatorOption) AggregatorOption {
		option.Workers = w
		return option
	}
}

func batchProcessor(list []interface{}) error {
	fmt.Println("==batchProcessor==> ", list)
	return nil
}

func TestAggregator() {
	ag := NewAggregator(batchProcessor,
		withBatchSize(5),
		withChannelBufferSize(3),
		withLingerTime(5*time.Second),
		withWorkers(2))
	ag.Start()
	defer ag.Stop()

	t := time.NewTicker(50 * time.Millisecond)
	defer t.Stop()

	for ch := range t.C {
		ag.Enqueue(ch.UnixMilli())
	}

}

//-------------------Logger---------------------

type Logger interface {
	Debugf(str string, args ...interface{})
	Infof(str string, args ...interface{})
	Warnf(str string, args ...interface{})
	Errorf(str string, args ...interface{})
}

type consoleLogger struct {
	debug *log.Logger
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
}

func NewConsoleLogger() *consoleLogger {
	return &consoleLogger{
		debug: log.New(os.Stdout, "[debug] ", log.LstdFlags),
		info:  log.New(os.Stdout, "[info] ", log.LstdFlags),
		warn:  log.New(os.Stdout, "[warn] ", log.LstdFlags),
		error: log.New(os.Stdout, "[error] ", log.LstdFlags),
	}
}

func (cl *consoleLogger) Debugf(format string, arg ...interface{}) {
	message := fmt.Sprintf(format, arg...)
	cl.debug.Println(message)
}

func (cl *consoleLogger) Infof(format string, arg ...interface{}) {
	message := fmt.Sprintf(format, arg...)
	cl.info.Println(message)
}

func (cl *consoleLogger) Warnf(format string, arg ...interface{}) {
	message := fmt.Sprintf(format, arg...)
	cl.warn.Println(message)
}

func (cl *consoleLogger) Errorf(format string, arg ...interface{}) {
	message := fmt.Sprintf(format, arg...)
	cl.error.Println(message)
}

//-------------------Logger---------------------

// Represents the aggregator
type Aggregator struct {
	option         AggregatorOption
	wg             *sync.WaitGroup
	quit           chan struct{}
	eventQueue     chan interface{}
	batchProcessor BatchProcessFunc
}

// Represents the aggregator option
type AggregatorOption struct {
	BatchSize         int
	Workers           int
	ChannelBufferSize int
	LingerTime        time.Duration
	ErrorHandler      ErrorHandlerFunc
	Logger            Logger
}

// the func to batch process items
type BatchProcessFunc func([]interface{}) error

// the func to set option for aggregator
type SetAggregatorOptionFunc func(option AggregatorOption) AggregatorOption

// the func to handle error
type ErrorHandlerFunc func(err error, items []interface{}, batchProcessFunc BatchProcessFunc, aggregator *Aggregator)

// Creates a new aggregator
func NewAggregator(batchProcessor BatchProcessFunc, optionFuncs ...SetAggregatorOptionFunc) *Aggregator {
	option := AggregatorOption{
		BatchSize:  8,
		Workers:    runtime.NumCPU(),
		LingerTime: 1 * time.Minute,
	}

	for _, optionFunc := range optionFuncs {
		option = optionFunc(option)
	}

	if option.ChannelBufferSize <= option.Workers {
		option.ChannelBufferSize = option.Workers
	}

	return &Aggregator{
		eventQueue:     make(chan interface{}, option.ChannelBufferSize),
		option:         option,
		quit:           make(chan struct{}),
		wg:             new(sync.WaitGroup),
		batchProcessor: batchProcessor,
	}
}

// Try enqueue an item, and it is non-blocked
func (agt *Aggregator) TryEnqueue(item interface{}) bool {
	select {
	case agt.eventQueue <- item:
		return true
	default:
		if agt.option.Logger != nil {
			agt.option.Logger.Warnf("Aggregator: Event queue is full and try reschedule")
		}

		runtime.Gosched()

		select {
		case agt.eventQueue <- item:
			return true
		default:
			if agt.option.Logger != nil {
				agt.option.Logger.Warnf("Aggregator: Event queue is still full and %+v is skipped.", item)
			}
			return false
		}
	}
}

// Enqueue an item, will be blocked if the queue is full
func (agt *Aggregator) Enqueue(item interface{}) {
	agt.eventQueue <- item
}

// Start the aggregator
func (agt *Aggregator) Start() {
	for i := 0; i < agt.option.Workers; i++ {
		index := i
		go agt.work(index)
	}
}

// Stop the aggregator
func (agt *Aggregator) Stop() {
	close(agt.quit)
	agt.wg.Wait()
}

// Stop the aggregator safely, the difference with Stop is it guarantees no item is missed during stop
func (agt *Aggregator) SafeStop() {
	if len(agt.eventQueue) == 0 {
		close(agt.quit)
	} else {
		ticker := time.NewTicker(50 * time.Millisecond)
		for range ticker.C {
			if len(agt.eventQueue) == 0 {
				close(agt.quit)
				break
			}
		}
		ticker.Stop()
	}
	agt.wg.Wait()
}

func (agt *Aggregator) work(index int) {
	defer func() {
		if r := recover(); r != nil {
			if agt.option.Logger != nil {
				agt.option.Logger.Errorf("Aggregator: recover worker as bad thing happens %+v", r)
			}

			agt.work(index)
		}
	}()

	agt.wg.Add(1)
	defer agt.wg.Done()

	batch := make([]interface{}, 0, agt.option.BatchSize)
	lingerTimer := time.NewTimer(0)
	if !lingerTimer.Stop() {
		<-lingerTimer.C
	}
	defer lingerTimer.Stop()

loop:
	for {
		select {
		case req := <-agt.eventQueue:
			batch = append(batch, req)

			batchSize := len(batch)
			if batchSize < agt.option.BatchSize {
				if batchSize == 1 {
					lingerTimer.Reset(agt.option.LingerTime)
				}
				break
			}

			agt.batchProcess(batch)

			if !lingerTimer.Stop() {
				<-lingerTimer.C
			}
			batch = make([]interface{}, 0, agt.option.BatchSize)
		case <-lingerTimer.C:
			if len(batch) == 0 {
				break
			}

			agt.batchProcess(batch)
			batch = make([]interface{}, 0, agt.option.BatchSize)
		case <-agt.quit:
			if len(batch) != 0 {
				agt.batchProcess(batch)
			}

			break loop
		}
	}
}

func (agt *Aggregator) batchProcess(items []interface{}) {
	agt.wg.Add(1)
	defer agt.wg.Done()
	if err := agt.batchProcessor(items); err != nil {
		if agt.option.Logger != nil {
			agt.option.Logger.Errorf("Aggregator: error happens")
		}

		if agt.option.ErrorHandler != nil {
			go agt.option.ErrorHandler(err, items, agt.batchProcessor, agt)
		} else if agt.option.Logger != nil {
			agt.option.Logger.Errorf("Aggregator: error happens in batchProcess and is skipped")
		}
	} else if agt.option.Logger != nil {
		agt.option.Logger.Infof("Aggregator: %d items have been sent.", len(items))
	}
}
