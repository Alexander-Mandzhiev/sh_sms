package grpc_client

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClientManager struct {
	connections sync.Map
	options     []grpc.DialOption
	logger      *slog.Logger
	mu          sync.Mutex
}

type connectionWrapper struct {
	conn     *grpc.ClientConn
	lastUsed time.Time
	refCount int
}

func NewGRPCClientManager(logger *slog.Logger, opts ...grpc.DialOption) *GRPCClientManager {
	defaultOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(16 << 20)),
	}

	if logger == nil {
		logger = slog.Default()
	}

	return &GRPCClientManager{
		options: append(defaultOpts, opts...),
		logger:  logger,
	}
}

func (m *GRPCClientManager) GetClientConn(ctx context.Context, addr string) (*grpc.ClientConn, error) {
	// Проверяем существующее соединение
	if wrapper, exists := m.connections.Load(addr); exists {
		w := wrapper.(*connectionWrapper)
		m.mu.Lock()
		defer m.mu.Unlock()

		if w.conn.GetState() == connectivity.Ready {
			w.refCount++
			w.lastUsed = time.Now()
			return w.conn, nil
		}

		// Соединение не в порядке, удаляем его
		m.connections.Delete(addr)
		if err := w.conn.Close(); err != nil {
			m.logger.Error("failed to close connection", "address", addr, "error", err)
		}
	}

	// Создаем новое соединение
	conn, err := grpc.DialContext(ctx, addr, m.options...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s: %w", addr, err)
	}

	wrapper := &connectionWrapper{
		conn:     conn,
		lastUsed: time.Now(),
		refCount: 1,
	}

	m.connections.Store(addr, wrapper)
	return conn, nil
}

func (m *GRPCClientManager) ReleaseConn(addr string) error {
	if wrapper, exists := m.connections.Load(addr); exists {
		m.mu.Lock()
		defer m.mu.Unlock()

		w := wrapper.(*connectionWrapper)
		w.refCount--
		w.lastUsed = time.Now()

		if w.refCount <= 0 {
			m.connections.Delete(addr)
			if err := w.conn.Close(); err != nil {
				return fmt.Errorf("failed to close connection to %s: %w", addr, err)
			}
		}
	}
	return nil
}

func (m *GRPCClientManager) CloseAll() error {
	var errs []error

	m.connections.Range(func(key, value interface{}) bool {
		addr := key.(string)
		wrapper := value.(*connectionWrapper)

		m.mu.Lock()
		defer m.mu.Unlock()

		m.connections.Delete(addr)
		if err := wrapper.conn.Close(); err != nil {
			errs = append(errs, fmt.Errorf("failed to close connection to %s: %w", addr, err))
		}
		return true
	})

	if len(errs) > 0 {
		return fmt.Errorf("errors closing connections: %v", errs)
	}
	return nil
}

func (m *GRPCClientManager) CleanupIdleConnections(idleTimeout time.Duration) {
	m.connections.Range(func(key, value interface{}) bool {
		addr := key.(string)
		wrapper := value.(*connectionWrapper)

		m.mu.Lock()
		defer m.mu.Unlock()

		if time.Since(wrapper.lastUsed) > idleTimeout && wrapper.refCount <= 0 {
			m.connections.Delete(addr)
			if err := wrapper.conn.Close(); err != nil {
				m.logger.Error("failed to close idle connection", "address", addr, "error", err)
			}
		}
		return true
	})
}
