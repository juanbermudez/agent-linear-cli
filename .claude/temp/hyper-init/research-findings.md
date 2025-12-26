# Research Agent Initialization Findings

## Documentation Structure

### Documentation Locations
- **README.md** - User-facing guide with features, installation, usage
- **CLAUDE.md** - Development guide with implementation guidelines
- **docs/** - HTML documentation site (landing, getting-started, commands, changelog)
- **.claude/agents/** - Agent workflow documentation
- **.claude/skills/** - Skill definitions for Linear CLI expert

### Key Reference Files for Research
1. `internal/api/client.go` - Main GraphQL client (3,769 LOC)
2. `internal/cmd/` - All CLI command implementations
3. `internal/auth/auth.go` - Authentication flow
4. `internal/cache/cache.go` - Caching layer
5. `CLAUDE.md` - Project conventions

## Codebase Organization

### Package Hierarchy
```
cmd/linear/main.go          # Entry point
internal/
├── cmd/                    # Cobra commands (7,003 LOC)
├── api/                    # GraphQL client (3,769 LOC)
├── auth/                   # Authentication (keyring, OAuth)
├── config/                 # TOML configuration
├── cache/                  # 24-hour TTL caching
├── output/                 # JSON/human output
└── display/                # Tables, formatting
```

### Key Dependencies
- `spf13/cobra` - CLI framework
- `hasura/go-graphql-client` - GraphQL
- `zalando/go-keyring` - Credential storage
- `pelletier/go-toml` - Configuration

## Linear API Integration

### Entity Types
- Issues (full CRUD, search, relationships, comments, attachments)
- Projects (CRUD, milestones, status updates)
- Documents (CRUD, search)
- Initiatives (CRUD, project linking)
- Teams, Users, Workflows, Labels

### Authentication Priority
1. LINEAR_API_KEY environment variable
2. LINEAR_CLIENT_ID + LINEAR_CLIENT_SECRET (OAuth)
3. System keychain storage
4. Return error with guidance

### Error Response Pattern
```json
{
  "success": false,
  "error": {
    "code": "ERROR_CODE",
    "message": "Description",
    "hint": "AI agent guidance",
    "usage": ["example commands"]
  }
}
```

## Research Recommendations

1. **Start with CLAUDE.md** for project conventions
2. **Check internal/api/client.go** for API patterns
3. **Review internal/cmd/** for command structure examples
4. **Use .linear.toml** for workspace context
5. **JSON output** is default (designed for AI consumption)
