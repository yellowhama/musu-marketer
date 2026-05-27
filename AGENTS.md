# 🤖 Guidance for AI Agents (musu-marketer)

`musu-marketer` is your "Voice and Hands." Use this tool to transform verified knowledge into high-impact marketing campaigns.

## 🎓 Master Orchestration
For advanced multi-tool orchestration (Crawl + Marketer + Nurikun), refer to the **[MUSU_SKILL.md](../../MUSU_SKILL.md)** in the workspace root. Activating this skill transforms you into a Lead Orchestrator of the entire ecosystem.

## 🏗️ Core Architecture for Agents
This tool provides **Marketing Primitives** grounded in professional strategy.

### 1. The Marketing Primitives
- **`draft [topic]`**: The "Thinking" step. It triggers the Strategist (STP analysis) and then the Copywriter (Content creation).
- **`autopilot [topic]`**: The "Full-Cycle" tool. Command `musu-crawl-ai` to research and then generate marketing assets automatically.
- **`persona [list/create/show]`**: Manage the brand's identity and tone of voice.
- **`publish [ID]`**: The "Execution" step. Pushes drafts to registered publisher adapters (e.g., `local`, `webhook`).

## 🤝 Ecosystem Collaboration: The "Influence Pipeline"
*Scenario: An agent acts as a full-service marketing agency using the Musu team.*

1.  **SPOT TRENDS:** Use `musu-crawl spot --json` to find what's hot in a specific niche.
2.  **DEEP RESEARCH:** Use `musu-crawl research --json` to gather verified facts about that trend into a Wiki.
3.  **STRATEGIZE:** Use `musu-marketer draft` to generate a high-impact campaign based on that Wiki data.
4.  **DEBUT IDENTITY:** Use `musu-nurikun forge` to create an AI persona that fits the campaign's target audience.
5.  **ESTABLISH PRESENCE:** Use `musu-nurikun signup` to register the new persona on the target platform (e.g., Reddit, X).
6.  **DEPLOY CONTENT:** Use `musu-marketer publish` to output the final copy, then use `musu-nurikun` to post it to the community.

## 🛑 Critical Mandates for Agents
- **Strategy First:** Always ensure a `Strategic Brief` is generated before outputting content.
- **Bible Compliance:** Strictly follow the frameworks in the `MARKETING_BIBLE.md`.
- **Self-Healing:** If `musu-crawl` is missing from the path, update `config.yaml` or use the `--wiki` flag.
