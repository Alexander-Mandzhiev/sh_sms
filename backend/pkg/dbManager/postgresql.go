package dbManager

import (
	config "backend/pkg/config/service"
	"context"
	"fmt"
	"log/slog"
	"regexp"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func OpenPostgresqlConnection(cfg config.PostgresConfig, logger *slog.Logger) (*pgxpool.Pool, error) {
	const op = "db_manager.OpenPostgreSQLConnection"

	if cfg.ConnectionString == "" {
		return nil, fmt.Errorf("%s: database configuration is not initialized", op)
	}
	if cfg.MaxOpenConnections < cfg.MaxIdleConnections {
		return nil, fmt.Errorf("%s: max_open_connections (%d) must be >= max_idle_connections (%d)",
			op, cfg.MaxOpenConnections, cfg.MaxIdleConnections)
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.ConnectTimeout)
	defer cancel()

	poolConfig, err := pgxpool.ParseConfig(cfg.ConnectionString)
	if err != nil {
		logger.Error(op,
			slog.String("message", "failed to parse PostgreSQL connection string"),
			slog.String("connection_string", maskPassword(cfg.ConnectionString)),
			slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	poolConfig.MaxConns = cfg.MaxOpenConnections
	poolConfig.MinConns = cfg.MaxIdleConnections
	poolConfig.MaxConnLifetime = cfg.ConnMaxLifetime
	poolConfig.MaxConnIdleTime = cfg.ConnMaxIdleTime
	poolConfig.HealthCheckPeriod = cfg.HealthCheckPeriod
	poolConfig.ConnConfig.ConnectTimeout = cfg.ConnectTimeout

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		logger.Error(op,
			slog.String("message", "failed to create PostgreSQL connection pool"),
			slog.String("connection_string", maskPassword(cfg.ConnectionString)),
			slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err = pool.Ping(ctx); err != nil {
		logger.Error(op,
			slog.String("message", "failed to ping PostgreSQL database"),
			slog.Any("error", err))
		pool.Close()
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info(op,
		slog.String("message", "successfully connected to PostgreSQL database"),
		slog.String("host", extractHost(poolConfig.ConnConfig.Config.Host)),
		slog.Int("max_connections", int(poolConfig.MaxConns)),
		slog.Int("min_connections", int(poolConfig.MinConns)),
		slog.Duration("max_connection_lifetime", poolConfig.MaxConnLifetime),
		slog.Duration("max_connection_idle_time", poolConfig.MaxConnIdleTime),
		slog.Duration("health_check_period", poolConfig.HealthCheckPeriod))

	return pool, nil
}

func ClosePostgresqlConnection(pool *pgxpool.Pool, logger *slog.Logger) error {
	const op = "db_manager.ClosePostgreSQLConnection"
	if pool == nil {
		logger.Warn(op, slog.String("message", "pool is nil, skipping close"))
		return nil
	}

	stats := pool.Stat()
	logger.Debug(op, slog.String("message", "closing PostgreSQL connection pool"), slog.Int("total_connections", int(stats.TotalConns())), slog.Int("idle_connections", int(stats.IdleConns())), slog.Int("acquired_connections", int(stats.AcquiredConns())), slog.Int("constructing_connections", int(stats.ConstructingConns())))
	pool.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			remaining := pool.Stat().TotalConns()
			logger.Warn(op, slog.String("message", "timeout while closing connection pool"), slog.Int("remaining_connections", int(remaining)), slog.Any("error", ctx.Err()))
			return fmt.Errorf("%s: timeout while closing pool (%d connections remaining)", op, remaining)
		case <-ticker.C:
			if pool.Stat().TotalConns() == 0 {
				logger.Info(op,
					slog.String("message", "successfully closed PostgreSQL connection pool"))
				return nil
			}
		}
	}
}

var passwordRegex = regexp.MustCompile(`(password=)([^& ]+)`)

func maskPassword(dsn string) string {
	if dsn == "" {
		return ""
	}
	return passwordRegex.ReplaceAllString(dsn, "${1}***")
}

func extractHost(fullHost string) string {
	if fullHost == "" {
		return "unknown"
	}
	if parts := strings.Split(fullHost, ":"); len(parts) > 0 {
		return parts[0]
	}
	return fullHost
}
