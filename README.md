# musu-marketer

> **The High-Impact AI Strategic Marketing Agency.**

`musu-marketer` is a professional-grade, autonomous marketing engine built in Go. It is the "Voice" of the Musu ecosystem, consuming verified knowledge from `musu-crawl-ai` and transforming it into strategic, viral content using a collaborative **Strategic Crew**.

---

## 🚀 Key Features

### 🎭 The Strategic Crew (v2.0.0)
- **Copywriter-Critic Loop:** Every draft is audited by a Zero-Tolerance Critic agent to eliminate AI fluff and maximize impact.
- **Social Memory:** Learns from past published successes to maintain narrative consistency.
- **Marketing Bible:** Injects 10+ years of professional marketing frameworks (AIDA, PAS, BAB) into the core logic.

### 🤖 Native Integration
- **MCP Server:** Native tool discovery for Claude and Gemini.
- **REST API Server:** Lightweight HTTP interface for backend product integration.
- **Interactive Builder:** Wizard-based persona creation (`persona create`).

### 🦾 Production Hardening
- **Resource Efficient:** Optimized HTTP connection pooling and singleton clients.
- **Project Siloing:** Strictly isolated assets, databases, and identities per mission.

---

## 🛠️ Installation & Setup

### 1. Prerequisites
- **Data Source:** Requires [musu-crawl-ai](https://github.com/yellowhama/musu-crawl-ai) wiki.
- **Intelligence:** Requires [Ollama](https://ollama.com).

### 2. Quick Start
```bash
./musu-marketer init --project master-brand
./musu-marketer doctor --project master-brand
./musu-marketer draft "QuantumComputing" --persona tech-analyst
```

---

## 📖 User Manual

### 1. Strategic Drafting
Generate content using a self-correcting loop:
```bash
./musu-marketer draft [topic] --persona [name] --project [project]
```

### 0. Preflight
Verify that your crawl wiki and AI endpoint are reachable before drafting:
```bash
./musu-marketer doctor --project [project]
./musu-marketer doctor --project [project] --json
./musu-marketer doctor --project [project] --fix
./musu-marketer doctor --project [project] --topic "your-topic"
```

`musu-marketer` now auto-discovers a sibling `../musu-crawl-ai/wiki` when `--wiki` is omitted. Override it with `--wiki`, `MUSU_MARKETER_WIKI`, or `MUSU_WIKI` when needed.
When safe, `doctor --fix` will scaffold the missing local project directory/database/config for you.
If you pass `--topic`, `doctor` checks whether the current wiki contains enough matching source material to draft that topic by reading `index.json` metadata and Markdown body content, not just filenames.

### 2. Autopilot (Zero-Click)
Spot trends via `crawl-ai`, research them, and draft campaigns automatically:
```bash
./musu-marketer autopilot [subreddit]
```

### 3. Campaign Management
```bash
./musu-marketer list
./musu-marketer view [ID]
./musu-marketer publish [ID] --platform local
```

---

## 📂 Architecture
All campaign data and project-specific personas are stored in the `projects/` directory, which is excluded from Git to protect your marketing secrets.

---

## 🔗 The Ecosystem
- **musu-crawl-ai:** The "Brain" providing verified knowledge.
- **musu-nurikun:** The "Hand" handling inbox triage and compliant opt-in email operations.
