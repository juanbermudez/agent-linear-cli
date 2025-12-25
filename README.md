# Linear Agent CLI (Go)

A high-performance command-line interface for Linear project management, designed for AI agent consumption.

## Features

- **JSON-first output** - Designed for programmatic consumption by AI agents
- **Fast execution** - Single static binary, no runtime dependencies
- **Comprehensive Linear API coverage** - Issues, projects, documents, labels, workflows
- **VCS integration** - Git branch parsing, issue context detection
- **24-hour caching** - Workflows, statuses, users cached for performance

## Installation

### Homebrew (macOS/Linux)

```bash
brew install juanbermudez/tap/linear-cli
```

### Direct Download

```bash
# macOS (Apple Silicon)
curl -fsSL https://github.com/juanbermudez/linear-agent-cli/releases/latest/download/linear-darwin-arm64 -o /usr/local/bin/linear
chmod +x /usr/local/bin/linear

# macOS (Intel)
curl -fsSL https://github.com/juanbermudez/linear-agent-cli/releases/latest/download/linear-darwin-amd64 -o /usr/local/bin/linear
chmod +x /usr/local/bin/linear

# Linux (x86_64)
curl -fsSL https://github.com/juanbermudez/linear-agent-cli/releases/latest/download/linear-linux-amd64 -o /usr/local/bin/linear
chmod +x /usr/local/bin/linear
```

### From Source

```bash
go install github.com/juanbermudez/linear-agent-cli/cmd/linear@latest
```

## Quick Start

```bash
# Interactive setup
linear config setup

# Or set values directly
linear config set api_key "lin_api_..."
linear config set team_key "LOT"

# Verify configuration
linear whoami
```

## Usage

### Issues

```bash
# List issues
linear issue list
linear issue list --state "In Progress"

# Create issue
linear issue create --title "Fix login bug" --priority 2

# View issue
linear issue view LOT-123

# Update issue
linear issue update LOT-123 --state "Done"

# Create relationships
linear issue relate LOT-123 LOT-456 --blocks
```

### Projects

```bash
# List projects
linear project list

# Create project with document
linear project create --name "Q1 Feature" --with-doc

# View project
linear project view abc123
```

### Documents

```bash
# Create document
linear document create --title "PRD: Feature X" --project abc123

# List documents
linear document list --project abc123
```

### Workflows

```bash
# List workflow states
linear workflow list

# Force refresh cache
linear workflow cache
```

## Output Modes

**JSON (default)** - Machine-readable output:
```bash
linear issue list
# [{"id": "...", "identifier": "LOT-123", ...}]
```

**Human-readable** - Formatted for terminal:
```bash
linear issue list --human
# LOT-123  Fix login bug  In Progress  @juan
```

## Configuration

The CLI looks for `.linear.toml` in the current directory, then home directory:

```toml
api_key = "lin_api_..."
team_id = "a5880da1-..."
team_key = "LOT"
```

Environment variables override config file:
- `LINEAR_API_KEY` - API key

## Integration with Claude Code

This CLI is designed to work with the [Hyper-Engineering](https://github.com/juanbermudez/hyper-eng) Claude Code plugin for spec-driven development workflows.

```bash
# In Claude Code
claude /hyper-plan "Add user authentication"  # Creates Linear project with spec
claude /hyper-implement project:abc123        # Implements with verification loop
```

## Development

```bash
# Build
make build

# Run tests
make test

# Build all platforms
make build-all

# Install locally
make install
```

## License

MIT
