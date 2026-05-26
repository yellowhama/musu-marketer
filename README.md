# musu-marketer

> **The High-Impact AI Speaker & Marketing Engine.**

`musu-marketer` is a professional-grade, autonomous marketing engine built in Go. It consumes verified knowledge (harvested by [musu-crawl-ai](https://github.com/yellowhama/musu-crawl-ai)) and transforms it into strategic, viral content across multiple platforms using a specialized **Marketing Strategic Brain**.

---

## 🚀 Key Features

### 🧠 The Strategic Brain (v1.2.1)
- **Strategy Architect:** Automatically performs STP (Segmentation, Targeting, Positioning) analysis before drafting any content.
- **Marketing Bible:** Injects a professional knowledge stack (AIDA, PAS, BAB, 4U frameworks) directly into the AI's cognitive core.
- **Psychological Triggers:** Intelligently selects triggers (Authority, Scarcity, Curiosity) to maximize conversion and engagement.

### 🎭 Multi-Persona Engine
- **Voice Engineering:** Switch instantly between brand voices (e.g., `hype-setter` for early adopters, `tech-analyst` for senior architects).
- **Interactive Builder:** Create custom marketing personalities via an easy-to-use CLI wizard (`persona create`).

### 🤖 Autonomous Autopilot
- **Spot-to-Draft Loop:** A zero-click pipeline that spots trends via `musu-crawl`, researches them, and drafts cross-platform campaigns into your database.
- **Project Isolation:** Maintain separate campaigns, personas, and databases for different missions or clients.

### 🌉 Product Integration
- **REST API Server:** Integrate musu-marketer as a backend service for your own SaaS or CRM.
- **Publisher Registry:** Pluggable architecture to add custom adapters for X, LinkedIn, Slack, or proprietary CMS.

---

## 🛠️ Installation & Setup

### 1. Prerequisites
- **Knowledge Base:** Requires [musu-crawl-ai](https://github.com/yellowhama/musu-crawl-ai) to provide verified data.
- **Intelligence:** Requires [Ollama](https://ollama.com) running locally.

### 2. Initialize
```bash
./musu-marketer init --project my-brand
```

---

## 📖 User Manual

### 1. Draft a Campaign
Generate strategic content based on a topic in your wiki:
```bash
./musu-marketer draft "QuantumComputing" --persona tech-analyst --project alpha
```

### 2. Manage Campaigns
```bash
./musu-marketer list --project alpha
./musu-marketer view 1 --project alpha
```

### 3. Run Autopilot
Automate trend spotting, research, and drafting in one go:
```bash
./musu-marketer autopilot "MachineLearning" --project auto-agency
```

### 4. Remote Control (API Mode)
```bash
./musu-marketer serve --port 8081
```

---

## 📂 Architecture
- `/projects/{name}/campaigns`: Drafted campaign files.
- `/projects/{name}/data/marketer.db`: SQLite lifecycle storage.
- `/projects/{name}/personas`: Project-specific brand voices.
- `INTEGRATION.md`: Developer guide for custom adapters.

---

## 📝 Roadmap
- [x] v1.0.0: Autonomous Drafting & Database
- [x] v1.1.0: Strategic Brain & Marketing Bible
- [x] v1.2.1: API Server & Connection Hardening
- [ ] v2.0.0: Visual Generation & Performance Analytics Loop
