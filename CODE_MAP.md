# Code And Doc Map: musu-marketer

## Runtime Entry Points
- `main.go`: CLI bootstrap
- `cmd/root.go`: global flags and environment initialization
- `cmd/init.go`: project/database/persona bootstrap
- `cmd/doctor.go`: wiki/project/AI/topic preflight
- `cmd/draft.go`: main draft command
- `cmd/autopilot.go`: higher-level orchestration flow
- `cmd/publish.go`: publish adapters

## Core Packages
- `internal/agent`: strategist, copywriter, critic, shared AI client
- `internal/bridge`: wiki integration and topic lookup
- `internal/db`: SQLite persistence for campaigns
- `internal/publisher`: local and webhook publishing
- `internal/api`: HTTP server surface

## Project Data Layout
- `projects/<project>/campaigns`
- `projects/<project>/personas`
- `projects/<project>/data/marketer.db`
- `projects/<project>/published`

## Docs
- `README.md`: operator quick start
- `SPEC.md`: product contract
- `AGENTS.md`: LLM-oriented usage guidance
- `HANDOFF.md`: implementation handoff
- `CODE_MAP.md`: code/doc index
- `QUALITATIVE_REPORT.md`: current quality verdict
- `NEXT_STEPS.md`: planned follow-up work
- `INTEGRATION.md`: API-facing integration notes
