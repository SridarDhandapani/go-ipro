# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

go-ipro is a Go library for managing Panasonic iPro (WV-Series) network cameras via CGI. It provides a type-safe interface for camera configuration, user management, and stream control using Digest authentication.

## Build & Development Commands

```bash
# Build the project
go build ./...

# Run tests
go test ./...

# Run a single test
go test -run TestName ./path/to/package

# Run the example
go run examples/example.go

# Tidy dependencies
go mod tidy
```

## Architecture

### Package Structure

- **pkg/camera** - Public entry point. `Camera` struct wraps internal handlers and provides the main API
- **pkg/api** - HTTP client with Digest auth, URL scheme handling, and CGI request/response parsing
- **pkg/model** - Data structures for camera configuration (CameraInfo, CameraData, User, Stream, etc.)
- **pkg/enum** - Type-safe constants for camera settings (OSD positions, stream resolutions, toggle states)
- **pkg/error** - Custom error types
- **internal/camera** - `CameraHandler` composes all mixins with the API handler
- **internal/mixin** - Feature-specific operations (ImageMixin, StreamMixin, SystemMixin, UserMixin)
- **internal/utils** - Response parsing utilities (CGI response to map/struct conversion)

### Key Patterns

**Functional Options**: Both `Camera` and `APIHandler` use the functional options pattern:
```go
camera.NewCamera(
    camera.WithName("name"),
    camera.WithUsername("admin"),
    camera.WithAPIOpts(api.WithHost("192.168.1.100")),
)
```

**Mixin Composition**: Camera features are organized into mixins that return closures accepting `*api.APIHandler`:
```go
func (s *SystemMixin) GetInfo() func(h *api.APIHandler) (*model.CameraInfo, error)

// Usage:
info, err := c.GetInfo()(c.APIHandler)
```

**Request Validation**: Uses `gopkg.in/validator.v2` for struct validation before API calls. Model structs use `validate` tags.

**URL Encoding**: Uses `github.com/google/go-querystring` for converting structs to URL-encoded form data. Model structs use `url` tags.

### Authentication

- Digest authentication via `github.com/icholy/digest`
- Supports custom `http.Client` via `api.WithClient()` for TLS configuration
- User credentials are encrypted with XOR-based algorithm before transmission (see `model.User.Encrypt`)

### CGI Endpoints

The API communicates with cameras at `/cgi-bin/*` endpoints. Common commands:
- `getinfo` - Camera information
- `getdata` / `setdata` - Configuration read/write
- `set_h264` / `set_h264_2` / etc. - Stream configuration
- `reg_admin` / `reg_user` / `del_user` - User management
- `initial` - Camera reset (requires `Randomnum` token)
