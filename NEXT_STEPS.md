# Next Steps: musu-marketer

> 2026-05-28 audit items (MCP schemas, draft guard, db MkdirAll, snake_case envelope, shared module extraction) are now CLOSED. See `MUSU_MCP_AUDIT_2026-05-28.md` and `MUSU_THERMONUCLEAR_REVIEW_2026-05-28.md`.

## P1
- decide whether ranked lexical topic retrieval should stay simple or move toward semantic/vector retrieval
- extract doctor reporting/fixing helpers from `cmd/doctor.go`

## P2
- extend the sample wiki into a richer multi-topic fixture
- expose topic-readiness explanations more explicitly in JSON output

## P3
- improve publish adapters beyond local/webhook
- production hardening on the docker-compose bundle: TLS termination, log rotation, image registry push

## Verified Integration Harness
- set `MUSU_MARKETER_INTEGRATION_AI_URL`
- optionally set `MUSU_MARKETER_INTEGRATION_MODEL` (verified locally with `llama3.2:1b`)
- run `go test -tags integration ./cmd`
- or run `powershell -ExecutionPolicy Bypass -File .\scripts\run-real-integration.ps1`
