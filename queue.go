package main

import (
	"github.com/Telefonica/nfqueue"
	"go.uber.org/zap"
)

type Queue struct {
	id    uint16
	queue *nfqueue.Queue
	log   *zap.Logger
}

func NewQueue(id uint16, log *zap.Logger) *Queue {
	q := &Queue{
		id:  id,
		log: log,
	}
	queueCfg := &nfqueue.QueueConfig{
		MaxPackets: 1000,
		BufferSize: 16 * 1024 * 1024,
		QueueFlags: []nfqueue.QueueFlag{nfqueue.FailOpen},
	}
	// Pass as packet handler the current instance because it implements nfqueue.PacketHandler interface
	q.queue = nfqueue.NewQueue(q.id, q, queueCfg)

	log.Info("queue initiated", zap.Uint16("id", id), zap.Any("config", queueCfg))

	return q
}

// Start the queue
func (q *Queue) Start() error {
	q.log.Info("starting queue")
	run := q.queue.Start()
	if run != nil {
		q.log.Fatal("error starting queue", zap.Error(run))
	}
	return nil
}

// Stop the queue.
func (q *Queue) Stop() error {
	return q.queue.Stop()
}

// Handle a nfqueue packet. It implements nfqueue.PacketHandler interface.
func (q *Queue) Handle(p *nfqueue.Packet) {
	q.log.Info("packet arrived", zap.Any("packet", p))

	// buffer := bytes.NewBuffer(p.Buffer)
	// buffer.WriteString("Inserting data")

	// Accept the packet
	p.Accept()
}
