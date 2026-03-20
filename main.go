package main

import (
	"context"
	"errors"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
)

const subject string = "demo"

func setupLogger(level string) {
	l := slog.LevelInfo

	switch strings.ToLower(level) {

	case "error":
		l = slog.LevelError

	case "warn":
		l = slog.LevelWarn

	case "debug":
		l = slog.LevelDebug

	}

	handler := slog.NewTextHandler(
		os.Stdout,
		&slog.HandlerOptions{Level: l},
	)

	slog.SetDefault(slog.New(handler))
}

// startServe will start the NATS server and will wait till it is ready to accept client connections.
func startServer(port int) (*server.Server, error) {
	if port == 0 {
		// select a random available port
		port = -1
	}

	options := &server.Options{Port: port}

	srv, err := server.NewServer(options)
	if err != nil {
		return nil, err
	}

	slog.Info("starting up NATS server")
	srv.Start()

	if !srv.ReadyForConnections(5 * time.Second) {
		return nil, errors.New("server failed to start and accept client connections")
	}

	slog.SetDefault(
		slog.With(
			slog.Group(
				"nats",
				slog.Group(
					"server",
					slog.String("address", srv.ClientURL()),
				),
			),
		),
	)

	slog.Info("NATS server running and accepting client connections")
	return srv, nil
}

// startClient will connect to the NATS address and returns a new NATS client.
func startClient(address string) (*nats.Conn, error) {
	return nats.Connect(address)
}

func publish(address string, d time.Duration, msgSize int64) {
	nc, err := startClient(address)
	if err != nil {
		panic(err)
	}
	defer nc.Close()

	msg := make([]byte, msgSize)

	logger := slog.With(
		slog.String("msgSize", FormatSize(msgSize)),
		slog.String("duration", d.String()),
	)

	logger.Info("starting to publish messages")

	var sent int64
	start := time.Now()

	done := time.NewTimer(d)
	defer done.Stop()

	defer func() {
		runtime := time.Since(start)
		totalBytes := sent * msgSize

		logger.Info(
			"publish complete",
			slog.String("runtime", runtime.String()),
			slog.Int64("published", sent),
			slog.Int64("msgPerSec", sent/int64(runtime.Seconds())),
			slog.String("totalSize", FormatSize(totalBytes)),
			slog.String("throughput", FormatSize(int64(float64(totalBytes)/runtime.Seconds()))+"/s"),
		)
	}()

	for {
		select {
		case <-done.C:
			// Make sure buffered publishes are pushed to the server.
			if err := nc.Flush(); err != nil {
				logger.Error("flush failed", slog.Any("err", err))
			}
			return

		default:
			if err := nc.Publish(subject, msg); err != nil {
				logger.Error("publish failed", slog.Any("err", err))
				return
			}
			sent++
		}
	}
}

// consumers will start one or more consumer clients.
func consumers(ctx context.Context, address string, clients int, msgSize int64) {
	var wg sync.WaitGroup

	for i := 0; i < clients; i++ {
		i := i
		wg.Add(1)

		go func() {
			defer wg.Done()

			var received int64
			start := time.Now()

			logger := slog.With(
				slog.Int("clientID", i),
				slog.String("msgSize", FormatSize(msgSize)),
				slog.String("nats.server.address", address),
			)

			nc, err := startClient(address)
			if err != nil {
				panic(err)
			}
			defer nc.Close()

			logger.Info("starting consumer")

			sub, err := nc.Subscribe(subject, func(msg *nats.Msg) {
				atomic.AddInt64(&received, 1)
			})
			if err != nil {
				panic(err)
			}

			defer func() {
				_ = sub.Unsubscribe()

				runtime := time.Since(start)
				count := atomic.LoadInt64(&received)
				totalBytes := count * msgSize

				logger.Info(
					"consumer complete",
					slog.String("runtime", runtime.String()),
					slog.Int64("received", count),
					slog.Int64("msgPerSec", count/int64(runtime.Seconds())),
					slog.String("totalSize", FormatSize(totalBytes)),
					slog.String("throughput", FormatSize(int64(float64(totalBytes)/runtime.Seconds()))+"/s"),
				)
			}()

			<-ctx.Done()
		}()
	}

	wg.Wait()
}

func main() {
	clients := flag.Int("consumers", 1, "how many consumer clients to start")
	dur := flag.Duration("duration", 5*time.Second, "time duration to run the test")
	msgBytes := flag.Int64("msgSize", 1024*1024, "message size in bytes to publish, default is 1MiB")
	port := flag.Int("port", -1, "port to start the NATS server, default is a random available port.")
	level := flag.String("level", "info", "logging level [info|warn|error|debug]")
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, *dur+time.Second)
	defer cancel()

	setupLogger(*level)

	// start a NATS server
	ns, err := startServer(*port)
	if err != nil {
		panic(err)
	}

	defer ns.Shutdown()

	// start the consumer clients
	go consumers(ctx, ns.ClientURL(), *clients, *msgBytes)

	// start publishing messages as fast as we can
	publish(ns.ClientURL(), *dur, *msgBytes)

	// wait for the context to either timeout or os signal received.
	<-ctx.Done()
}
