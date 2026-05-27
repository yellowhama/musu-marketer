# musu-marketer

> **The High-Impact AI Strategic Agency.**

`musu-marketer` is a professional-grade, autonomous marketing agency built in Go. It achieves **Agentic Excellence** by orchestrating a collaborative "Crew" of agents (Strategist, Copywriter, Critic) to transform verified knowledge into high-performance marketing assets.

---

## 🚀 Key Features

### 🎭 The Strategic Crew (v2.0.0)
- **Collaborative Loop:** Every draft undergoes a rigorous audit by a **Zero-Tolerance 'Critic' Agent** to eliminate AI fluff and maximize viral hooks.
- **Social Memory:** The system learns from your past successes. The **Strategist** automatically references previous published history to maintain narrative consistency.
- **Expert Bible:** Injects a professional knowledge stack (AIDA, PAS, BAB) directly into the cognitive core.

### 🤖 Native Agent Integration
- **Model Context Protocol (MCP):** Stdio-based server that makes `musu-marketer` a native tool for Claude and Gemini.
- **REST API:** Lightweight HTTP interface for integration into custom SaaS or CRM platforms.

### ⚡ Automation & Isolation
- **Autonomous Autopilot:** Full "Spot -> Research -> Draft" pipeline using `musu-crawl-ai`.
- **Project Siloing:** Strictly isolated assets, identities, and databases for different missions or clients.

---

## 🛠️ Installation & Setup

### 1. Prerequisites
- **Knowledge:** [musu-crawl-ai](https://github.com/yellowhama/musu-crawl-ai)
- **Intelligence:** [Ollama](https://ollama.com)

### 2. Initialize
```bash
./musu-marketer init --project master-brand
```

---

## 📖 User Manual

### 1. Multi-Agent Drafting
Generate strategic content with a self-correcting review loop:
```bash
./musu-marketer draft "DeepTech" --persona tech-analyst --project alpha
```

### 2. Native Orchestration (MCP)
Add `musu-marketer mcp` to your Claude Desktop config to use it as a native tool.

### 3. Autopilot (Zero-Click)
```bash
./musu-marketer autopilot "MachineLearning" --project auto-agency
```

---

## 📝 Roadmap
- [x] v1.0.0: Autonomous Drafting & Database
- [x] v1.2.1: API Server & Connection Hardening
- [x] v2.0.0: The Strategic Crew (Critic, Memory, MCP)
- [ ] v3.0.0: Visual Generation & Analytics Feedback Loop
