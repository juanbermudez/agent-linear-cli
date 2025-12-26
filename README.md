# Agent Linear CLI

A JSON-first command-line interface for Linear, designed for AI agents.

## Features

- **JSON-first output** - Structured responses for AI agents
- **--human flag** - Readable tables for terminal use
- **Error hints** - Helpful guidance in error responses
- **Secure auth** - System keychain credential storage
- **24-hour caching** - Reduced API calls for common data
- **Stdin support** - Non-interactive setup for automation

## Installation

### Using Go

```bash
go install github.com/juanbermudez/agent-linear-cli/cmd/linear@latest
```

### From Source

```bash
git clone https://github.com/juanbermudez/agent-linear-cli.git
cd agent-linear-cli
make build
make install
```

## Quick Start

```bash
# Non-interactive setup (for AI agents)
echo "lin_api_xxxxx" | linear auth login --stdin

# Or interactive setup
linear config setup --api-key lin_api_xxxxx --team ENG

# Verify configuration
linear whoami
```

## AI Agent Usage Guide

This CLI is designed for AI agents. All commands output JSON by default for easy parsing.

### Authentication

```bash
# Check auth status
linear auth status
# {"authenticated": true, "method": "api_key", "source": "keychain"}

# Login with stdin (non-interactive, best for agents)
echo "lin_api_xxxxx" | linear auth login --stdin

# Verify identity
linear whoami
# {"user": {"id": "...", "name": "...", "email": "..."}, "organization": {...}}
```

### Team & Workspace Discovery

```bash
# List all teams (get team keys for other commands)
linear team list
# {"teams": [{"id": "...", "key": "ENG", "name": "Engineering"}], "count": 1}

# List users
linear user list
# {"users": [{"id": "...", "displayName": "...", "email": "..."}], "count": N}

# List workflow states (needed for state transitions)
linear workflow list --team ENG
# {"workflowStates": [{"id": "...", "name": "In Progress", "type": "started"}]}

# List labels
linear label list --team ENG
# {"labels": [{"id": "...", "name": "Bug", "color": "#EB5757"}], "count": N}
```

### Issue Management

#### Listing Issues

```bash
# List issues (team is required)
linear issue list --team ENG

# Filter by state type: backlog, unstarted, started, completed, canceled
linear issue list --team ENG --state started
linear issue list --team ENG --state completed

# Limit results
linear issue list --team ENG --limit 10

# Human-readable output
linear issue list --team ENG --human
```

#### Viewing Issues

```bash
# View by identifier (TEAM-NUMBER format)
linear issue view ENG-123
# {"id": "...", "identifier": "ENG-123", "title": "...", "description": "...", "state": {...}}

# View with human-readable format
linear issue view ENG-123 --human
```

#### Creating Issues

```bash
# Basic issue creation
linear issue create --title "Fix login bug" --team ENG

# Full issue creation with all options
linear issue create \
  --title "Implement OAuth flow" \
  --description "Add OAuth2 support for SSO" \
  --team ENG \
  --priority 2 \
  --label "Feature"

# Priority values: 0=None, 1=Urgent, 2=High, 3=Medium, 4=Low
```

**Response:**
```json
{
  "success": true,
  "issue": {
    "id": "uuid",
    "identifier": "ENG-123",
    "url": "https://linear.app/..."
  }
}
```

#### Updating Issues

```bash
# Update title
linear issue update ENG-123 --title "New title"

# Update description
linear issue update ENG-123 --description "Updated description"

# Update priority
linear issue update ENG-123 --priority 1

# IMPORTANT: State requires workflow state ID, not name
# First get the state ID from workflow list
linear workflow list --team ENG
# Then use the ID to update
linear issue update ENG-123 --state <state-uuid>
```

#### Searching Issues

```bash
# Basic search
linear issue search "authentication"
# {"issues": [...], "totalCount": N, "hasMore": false, "query": "authentication"}

# Search with options
linear issue search "bug fix" --limit 100
linear issue search "old feature" --include-archived
linear issue search "user feedback" --include-comments
linear issue search "api error" --team ENG
```

#### Deleting Issues

```bash
linear issue delete ENG-123
# {"success": true, "operation": "delete", "issueId": "ENG-123"}
```

### Comments

```bash
# Add comment to issue
linear issue comment create ENG-123 --body "This needs review"
# {"success": true, "comment": {"id": "...", "body": "..."}}

# List comments
linear issue comment list ENG-123
# {"comments": [...], "count": N}
```

### Issue Relationships

```bash
# Create relationship
linear issue relate ENG-123 ENG-456 --blocks
linear issue relate ENG-123 ENG-456 --related-to
linear issue relate ENG-123 ENG-456 --duplicate-of

# List relationships
linear issue relations ENG-123

# Remove relationship
linear issue unrelate ENG-123 ENG-456
```

### Projects

```bash
# List projects
linear project list
linear project list --team ENG

# Search projects
linear project search "Q1 roadmap"
linear project search "feature" --include-archived

# Create project
linear project create --name "Q1 Feature" --team ENG

# Create with document
linear project create --name "Q1 Feature" --team ENG --with-doc --doc-title "Project Spec"

# View project
linear project view <project-id>

# Add milestone
linear project milestone create <project-id> --name "Phase 1" --target-date 2025-02-15
```

### Documents

```bash
# List documents
linear document list
linear document list --project <project-id>

# Create document
linear document create --title "PRD: Feature X" --content "# Overview\n\n..."

# Search documents
linear document search "authentication"

# View document
linear document view <doc-id>
```

### Initiatives

```bash
# List initiatives
linear initiative list

# Create initiative
linear initiative create --name "Q1 Platform Improvements" --status Active

# Add project to initiative
linear initiative project-add <init-id> <project-id>
```

## Output Formats

### JSON Output (Default)

All commands output JSON for machine parsing:

```bash
linear issue list --team ENG
```
```json
{
  "issues": [
    {
      "id": "uuid",
      "identifier": "ENG-123",
      "title": "Fix login bug",
      "priority": 2,
      "state": {"name": "In Progress", "type": "started"},
      "updatedAt": "2025-01-15T10:30:00Z"
    }
  ],
  "count": 1
}
```

### Human-Readable Output

Add `--human` for formatted terminal output:

```bash
linear issue list --team ENG --human
```
```
Issues for team ENG:

     ID       TITLE                 LABELS  E  A   STATE        UPDATED
▄▆█  ENG-123  Fix login bug         bug        JD  In Progress  2 hours ago

1 issues
```

### Error Responses

Errors include helpful hints for recovery:

```json
{
  "success": false,
  "error": {
    "code": "MISSING_TEAM",
    "message": "Team is required",
    "hint": "Specify a team using --team flag or set a default team",
    "usage": [
      "linear issue list --team ENG",
      "linear config set team_key ENG"
    ]
  }
}
```

Common error codes:
- `NOT_AUTHENTICATED` - Need to run `linear auth login`
- `MISSING_TEAM` - Add `--team` flag or set default
- `API_ERROR` - Linear API error (check message for details)
- `NOT_FOUND` - Issue/project/document doesn't exist

## Configuration

### Config File

Configuration is stored in `.linear.toml`:

```toml
team_key = "ENG"
```

### Environment Variables

Environment variables override config file:

```bash
export LINEAR_API_KEY=lin_api_xxxxx
export LINEAR_TEAM_KEY=ENG
```

### Config Commands

```bash
# Set default team
linear config set team_key ENG

# Get config value
linear config get team_key

# List all config
linear config list

# Show config file path
linear config path
```

## Caching

The CLI caches frequently-accessed data for 24 hours:

- Workflow states
- Project statuses
- Users
- Labels

Force cache refresh:
```bash
linear workflow cache --team ENG
linear status cache
linear user list --refresh
linear label list --refresh
```

Cache location: `~/.cache/agent-linear-cli/`

## Best Practices for AI Agents

1. **Always check auth first**: Run `linear whoami` to verify authentication
2. **Discover workspace context**: Use `team list`, `workflow list`, `label list` before operating
3. **Use state IDs for updates**: Get workflow state UUIDs from `workflow list`, not state names
4. **Parse JSON responses**: All commands return structured JSON by default
5. **Check success field**: Responses include `"success": true/false`
6. **Handle errors gracefully**: Error responses include hints and usage examples
7. **Use search for discovery**: `issue search`, `project search`, `document search`

## Documentation

Full documentation: https://juanbermudez.github.io/agent-linear-cli

## Development

```bash
# Build
make build

# Run tests
make test

# Install locally
make install
```

## License

MIT
