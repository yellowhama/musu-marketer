# Project Handoff: musu-marketer

## What This Repo Is
`musu-marketer` is the Musu ecosystem's campaign drafting engine. It consumes a grounded wiki, keeps project-local personas and campaign state, and outputs publishable copy through local or webhook publishers.

## Current Truth
- binary: `musu-marketer.exe`
- version constant: `v2.0.1`
- default AI contract: OpenAI-compatible endpoint at `--ai-url`
- default wiki contract: sibling `../musu-crawl-ai/wiki` auto-discovery when `--wiki` is omitted
- key recovery path: `doctor --fix`

## What Changed In This Round
- added `doctor`
- added `--json`
- added `doctor --fix`
- changed topic readiness from filename-only to `index.json` + body-content matching
- added test coverage for indexed/content topic lookup
- `init` now writes a project-local `NEXT_STEPS.md` and returns structured bootstrap metadata in `--json`
- embedded the Marketing Bible so draft execution stays stable outside the repo root

## Operator Flow
1. `musu-marketer init --project <name>`
2. `musu-marketer doctor --project <name> --topic "<topic>"`
3. `musu-marketer draft "<topic>" --persona <persona> --project <name>`
4. `musu-marketer publish <id> --platform local|webhook`

## Known Constraints
- topic readiness is heuristic substring matching, not ranked search
- `doctor` still mixes reporting and fixing in one command file
- local smoke coverage exists, but topic retrieval is still heuristic rather than ranked

## Key Files
- `cmd/root.go`: global flags, wiki auto-discovery
- `cmd/init.go`: project bootstrap
- `cmd/doctor.go`: preflight and scaffold repair
- `cmd/output.go`: JSON success/error envelope
- `internal/bridge/wiki.go`: wiki lookup and topic matching
- `internal/agent/*`: strategist/copywriter/critic logic
- `internal/agent/skills.go`: embedded Marketing Bible loader
- `projects/<project>/NEXT_STEPS.md`: generated project-local bootstrap guide
