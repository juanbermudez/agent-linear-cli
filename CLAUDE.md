# Linear Agent CLI (Go) - Development Guide

## Project Overview

Go rewrite of the Linear Agent CLI, designed for AI agent consumption with JSON-first output.

## Directory Structure

```
linear-agent-cli-go/
├── cmd/linear/main.go      # Entry point
├── internal/
│   ├── cmd/                # Cobra command implementations
│   ├── api/                # Linear GraphQL client
│   ├── config/             # Configuration management
│   ├── cache/              # 24-hour caching layer
│   └── vcs/                # Git/VCS integration
├── go.mod
├── Makefile
└── .github/workflows/      # CI/CD
```

## Development Commands

```bash
make build      # Build for current platform
make test       # Run tests
make lint       # Run linter
make install    # Install locally
```

## Implementation Guidelines

### Command Structure
- Use Cobra for CLI framework
- Each command group in its own file
- JSON output by default, --human for readable
- All flags documented with examples

### API Client
- Use hasura/go-graphql-client for type-safe queries
- Handle pagination automatically
- Implement retry logic with exponential backoff

### Caching
- 24-hour cache for: workflows, statuses, users, labels
- Store in ~/.linear-cache/ as JSON files
- Force refresh with `linear <resource> cache`

### Error Handling
- Return structured JSON errors
- Include helpful messages for common issues
- Non-zero exit codes for failures

## Feature Parity Checklist

Must match Deno version exactly:
- [ ] Issue: create, view, list, update, delete, search, relate, start, comment, attachment
- [ ] Project: create, view, list, update, delete, restore, milestone, update-status
- [ ] Document: create, view, list, update, delete, restore, search
- [ ] Label: create, list, update, delete
- [ ] Workflow: list, cache
- [ ] Status: list, cache
- [ ] User: list, search
- [ ] Team: list
- [ ] Initiative: create, view, list, update, archive, restore, project-add, project-remove
- [ ] Config: setup, set, get, list
- [ ] Whoami

## Testing

- Unit tests for each command
- Integration tests against Linear API (with test account)
- JSON output validation tests
