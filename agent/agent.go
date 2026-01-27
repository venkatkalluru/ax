// Package agent provides interfaces and implementations for local and remote agents.
// Agents process content through callback handlers and can be registered with the dispatcher.
package agent

import (
	"context"

	"github.com/google/gar/proto"
)

// OutputHandler is a callback function that handles output content from an agent.
// It is called for each piece of content the agent generates.
type OutputHandler func(content *proto.Content) error

// Agent defines the common interface for both local and remote agents.
// Agents process content using callback handlers.
type Agent interface {
	// Process handles processing of input content.
	// It calls the output handler for each piece of content generated.
	// The handler may be called multiple times during processing.
	Process(ctx context.Context, sessionID string, inputs []*proto.Content, handler OutputHandler) error

	// HealthCheck checks if the agent is healthy and responsive.
	// Returns an error if the agent is unhealthy or unreachable.
	HealthCheck(ctx context.Context) error

	// ID returns the unique identifier for this agent.
	ID() string

	// Close gracefully shuts down the agent and releases resources.
	Close() error
}
