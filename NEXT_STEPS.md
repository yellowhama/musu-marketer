# Next Steps: musu-marketer

> See `C:\Users\empty\MUSU_MCP_AUDIT_2026-05-28.md` for the 2026-05-28 real-usage audit that produced the MCP-related items below.

## P1
- decide whether ranked lexical topic retrieval should stay simple or move toward semantic/vector retrieval
- extract doctor reporting/fixing helpers from `cmd/doctor.go`
- declare MCP tool parameter schemas in `cmd/mcp.go` for `draft_campaign` / `list_campaigns` (currently empty JSON schema → MCP clients cannot pass `topic`/`project`/`persona`); use `WithString`/`Required` from `mcp-go`
- guard empty input in `handleDraft` (currently no validation; an empty `topic` would silently proceed)

## P2
- extend the sample wiki into a richer multi-topic fixture
- expose topic-readiness explanations more explicitly in JSON output
- add `json:"…"` tags to `preflight.DoctorResult` so the MCP envelope is snake_case-consistent with the inner Report

## P3
- improve publish adapters beyond local/webhook
- extract shared module(s) for `AgentClient` + `preflight/doctor` + env-loader to remove triple-duplicated logic across the three repos

## Verified Integration Harness
- set `MUSU_MARKETER_INTEGRATION_AI_URL`
- optionally set `MUSU_MARKETER_INTEGRATION_MODEL` (verified locally with `llama3.2:1b`)
- run `go test -tags integration ./cmd`
- or run `powershell -ExecutionPolicy Bypass -File .\scripts\run-real-integration.ps1`
