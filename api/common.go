package api // import "github.com/docker/docker/api"

// Common constants for daemon and client.
const (
	// DefaultVersion of Current REST API
	DefaultVersion = "1.42"

	// NoBaseImageSpecifier is the symbol used by the FROM
	// command to specify that no base image is to be used.
	NoBaseImageSpecifier = "scratch"

	// MediaTypeRawStream is vendor specific MIME-Type set for raw TTY streams
	MediaTypeRawStream = "application/vnd.docker.raw-stream"

	// MediaTypeMultiplexedStream is vendor specific MIME-Type set for stdin/stdout/stderr multiplexed streams
	MediaTypeMultiplexedStream = "application/vnd.docker.multiplexed-stream"
)
