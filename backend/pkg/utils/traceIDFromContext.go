package utils

import "context"

func TraceIDFromContext(ctx context.Context) string {
	if traceID, ok := ctx.Value("trace_id").(string); ok {
		return traceID
	}
	return "unknown"
}
