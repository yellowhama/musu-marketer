# 🤖 Guidance for AI Agents (musu-marketer)

`musu-marketer` is your "Voice and Hands." Use this tool to transform verified knowledge into high-impact marketing campaigns.

## 🎓 Master Orchestration
For advanced multi-tool orchestration (Crawl + Marketer), refer to the **[MUSU_SKILL.md](../../MUSU_SKILL.md)** in the workspace root. Activating this skill transforms you into a Lead Orchestrator of the entire ecosystem.

## 🏗️ Core Architecture for Agents
This tool provides **Marketing Primitives** grounded in professional strategy.

### 1. The Marketing Primitives
- **`draft [topic]`**: The "Thinking" step. It triggers the Strategist (STP analysis) and then the Copywriter (Content creation).
- **`autopilot [topic]`**: The "Full-Cycle" tool. Command `musu-crawl-ai` to research and then generate marketing assets automatically.
- **`persona [list/create/show]`**: Manage the brand's identity and tone of voice.
- **`publish [ID]`**: The "Execution" step. Pushes drafts to registered publisher adapters (e.g., `local`, `webhook`).

## 🏎️ How to "Drive" this tool as an Agent

### Scenario: "Identify and Launch a Viral Campaign"
1.  **Spot & Research:** Run `.\musu-marketer.exe autopilot [topic]`.
2.  **Review:** Read the generated drafts in the SQLite database or via `.\musu-marketer.exe list`.
3.  **Refine:** If the content is "3 degrees off", use the `persona create` tool to build a more accurate voice and re-draft.
4.  **Execute:** Run `.\musu-marketer.exe publish [ID] --platform local`.

## 🛑 Critical Mandates for Agents
- **Strategy First:** Always ensure a `Strategic Brief` is generated before outputting content.
- **Bible Compliance:** Strictly follow the frameworks in the `MARKETING_BIBLE.md`.
- **Self-Healing:** If `musu-crawl` is missing from the path, use the `--wiki` flag or update `config.yaml` to point to the correct data source.
