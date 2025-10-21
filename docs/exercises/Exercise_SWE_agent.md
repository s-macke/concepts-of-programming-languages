# Exercise - SWE-agent

## Learning Objectives

In this exercise, you will:
- Build a working application (Learning Campus Light) using SWE-agents as development assistants
- Experience firsthand how AI-powered coding agents handle real software development tasks
- Document and reflect on the strengths and limitations of different SWE-agents
- Practice planning and spec-driven development and iterative refinement with AI assistance

**Work Mode:** Individual work

# Installation

## Prerequisites
- Go installed (version 1.21 or later)
- Basic CLI tools (git, text editor)
- API keys from Learning Campus

## Available SWE-Agents

API keys for several agents have been provided in the Learning Campus.

* **Anthropic Claude Code**: https://www.claude.com/product/claude-code
  ```bash
  ANTHROPIC_API_KEY="<ANTHROPIC_API_KEY>" claude
  ```

* **OpenAI Codex CLI**: https://developers.openai.com/codex/cli/
  ```bash
  codex
  ```

* **Gemini CLI**: https://github.com/google-gemini/gemini-cli
  ```bash
  GEMINI_API_KEY=<GEMINI_API_KEY> gemini
  ```

## Additional Tools (Optional)

Worth trying:

* **GitHub Copilot Agent** as Plugin in VS Code or IntelliJ
* **Lovable.dev**: https://lovable.dev/
* **Kiro**: https://kiro.dev/

> [!CAUTION]
> Your code is sent to third-party services. Do not include sensitive data or credentials in your prompts.

# Develop a Learning Campus Light in Go

## Project Overview

You will design and implement a simplified learning management system called "Learning Campus Light" using Go. 
The key principle: **You define the scope in your spec, then implement what you specified.**

## Suggested Features (Not All Required)

These are suggestions to inspire your design.

* **Roles**: 2 roles (Student, Instructor). Instructors can write, Students can only read.
* **Authentication**: Simple JSON file-based authentication (plain text is acceptable for this exercise), or omit authentication entirely.
* **User Interface**: CLI with menu structure or simple web UI.
* **Courses**: Predefined in a JSON file or similar data store.
* **Content**: Each course has one document (plain text).
* **Operations**: View content, edit content (for instructors), list courses, etc.

**Remember**: Your spec defines what you'll build. Start simple and expand if time permits. 

## Step 1: Write anb initial Mini-Spec (Markdown Document)

Create a specification document that describes what your Learning Campus Light will do and how it will work. Your spec should include:

### Suggested Sections

1. **Product Vision**
   - What problem does this solve?
   - Who are the users?
   - What's the core value proposition?

2. **UI Design**
   - Describe the menu structure
   - Show example interactions (mockup of terminal output)
   - User workflows (e.g., "Student logs in → Views courses → Selects course → Reads document")

3. **Use Cases**
   - List the main scenarios (e.g., "Instructor edits course content", "Student views course materials")
   - Include edge cases you'll handle (e.g., "User enters invalid course ID")

4. **Technical Specification**
   - **Architecture**: High-level structure (packages, modules)
   - **Data Storage**: JSON files, in-memory, or other approach
   - **Testing Strategy**: What will you test and how?
   - **Error Handling**: How will you handle invalid inputs, missing files, etc.?

Hint: Use the SWE-agent to draft and refine this spec.

### Tips for Working with SWE-Agents

* **Use Planning Mode**: If available, use the planning mode of the SWE agent to help draft and refine the spec.
* **Iterate**: Refine the spec in several rounds with the SWE-agent before implementation.
* **Avoid Context Rot**: Start new sessions if the chat begins to hallucinate or lose track of context.
* **Keep Context Compact**: Periodically summarize progress to keep the conversation focused.
* **Be Specific**: The more detailed your spec, the better the agent can help implement it.

## Step 2: Implement with SWE-Agent Assistance

Use an SWE-agents to help you implement the application according to your spec.

**Implementation Requirements:**
- Follow your spec as closely as possible
- Write tests for core functionality
- Implement proper error handling
- Follow Go best practices and conventions
- Document your code with comments where appropriate

## Step 3: Document Your Findings on the Miro Board

Create a post on the Miro board that includes:

1. **Your Specification**
   - Post the mini-spec document you created

2. **Screenshots & Demos**
   - Screenshots of your running application
   - Show both successful operations and error handling

3. **SWE-Agent Usage**
   - Which agent(s) did you use?
   - How did you structure your prompts/conversations?
   - Did you use planning mode, iterative refinement, or direct implementation?
   - How many conversation sessions/iterations did it take?

4. **Evaluation**
   - **What worked well?** (e.g., agent understood requirements, generated clean code, good at testing)
   - **What didn't work?** (e.g., hallucinations, incorrect implementations, needed manual fixes)
   - **Spec adherence**: How closely did the final implementation follow your spec? What changed and why?
   - **Code quality**: Assess the quality of generated code (structure, readability, tests, error handling)
   - **Productivity**: Did the agent speed up development or slow it down?

5. **Lessons Learned**
   - What strategies worked best for working with the SWE-agent?
   - What would you do differently next time?
   - How could the agent be improved?

