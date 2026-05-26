# 🌉 Integration Guide: Building on musu-marketer

`musu-marketer` is designed to be an extensible engine. You can integrate it into your own products via the **Publisher Registry** or the **REST API**.

---

## 1. Extending with Custom Publishers
If you need to push campaigns to your own proprietary CMS, CRM, or a custom Slack channel, you can write a simple Go adapter.

### Step-by-step:
1. Create a new file in `internal/publisher/your_adapter.go`.
2. Implement the `Publisher` interface:
```go
type MyAdapter struct {}
func (a *MyAdapter) Publish(topic, content string) (string, error) {
    // Your custom logic here (API calls, file saving, etc.)
    return "Success URL/ID", nil
}
```
3. Register it in an `init()` function:
```go
func init() {
    Register("my-custom-key", &MyAdapter{})
}
```
4. Rebuild the project. You can now use `.\musu-marketer.exe publish [ID] --platform my-custom-key`.

---

## 2. Remote Control via REST API
Start the server to allow your web app to trigger marketing tasks:
```bash
./musu-marketer serve --port 8081
```

### Endpoints (v1):
- `GET /health`: Check if the engine is running.
- `POST /api/v1/draft`: (WIP) Trigger a strategy analysis and campaign draft.
- `GET /api/v1/campaigns`: (WIP) List stored drafts.

---

## 3. Embedding the Engine
Because `musu-marketer` is written in Go, you can import its internal packages (`agent`, `bridge`, `db`) directly into your own Go projects for maximum performance.
