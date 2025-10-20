# Exercise - SWE-agent

For this excercise: You have the chance to test several SWE-agent.
Protocol your experience on the Miro board.

# Installation

API keys for several agents have been provided in the Learning Campus.

* Anthropic Claude Code: https://www.claude.com/product/claude-code
  start via `ANTHROPIC_API_KEY="<ANTHROPIC_API_KEY>" claude`

* OpenAI Codex CLI: https://developers.openai.com/codex/cli/
  start via `codex`
 
* Gemini CLI: https://github.com/google-gemini/gemini-cli
  start via `GEMINI_API_KEY=<GEMINI_API_KEY> gemini`

Worth to try:

* GitHub Copilot Agent as Plugin in VS Code or IntelliJ
* Lovable.dev: https://lovable.dev/
* Kiro: https://kiro.dev/


# Develop a Learning Campus Light in Go

Make it as simple as possible. For example:

* 2 roles (Student, Instructor). Instructor can write, Student can only read.
* Authentication via a simple JSON file, or leave authentication aside for now.
* UI only as CLI with a menu structure.
* Courses predefined in a JSON file.
* Each course has exactly one document to edit. 

## Write a mini-spec markdown document of what your Learning Campus Light should be able to do and how it should look.

* Product Vision
* UI Design
* Use Cases
* Technical Spec (architecture, language, database, tests)

Tip: If available, use the planning mode of the SWE agent. Try to refine the spec in several iterations with the SWE-agent.
Tip: Avoid context rot. Start new sessions if the chat begins to hallucinate. Try to compactify the chat history from time to time.

## Describe your findings on the Miro board.

- Post your spec on the Miro board.
- Post screenshots of your application.
- How exactly did you use the SWE agent?
- What worked well?
- What didn’t work?
- How closely did it follow the spec?
