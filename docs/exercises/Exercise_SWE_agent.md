# Exercise - SWE-agent

In this exercise, you will build a small learning campus application in Go with Claude Code. You will define its scope, guide the agent through implementation, and independently check the result. The goal is to learn how to make and justify development decisions while working with a coding agent.

## Learning Objectives

In this exercise, you will:
- Build a working application (Learning Campus Light) using SWE-agents as development assistants
- Experience firsthand how AI-powered coding agents handle real software development tasks
- Document and reflect on the strengths and limitations.
- Practice planning and iterative refinement with AI assistance
- Independently check the generated application against your own expectations

# Installation

## Prerequisites
- Go installed
- Anthropic API key from Learning Campus

## Available SWE-Agent

For this exercise, use **Claude Code with Sonnet 5.5**. An Anthropic API key is provided in the Learning Campus.

### 1. Install Claude Code

Open a terminal and run the command for your operating system.

**macOS, Linux, or WSL (Bash/Zsh):**

```bash
curl -fsSL https://claude.ai/install.sh | bash
```

**Windows (PowerShell):**

```powershell
irm https://claude.ai/install.ps1 | iex
```

On native Windows, installing [Git for Windows](https://git-scm.com/downloads/win) is recommended for Bash support.

Open a new terminal after installation and check:

```bash
claude --version
```

If the command is not found, follow the PATH instructions in the installer output or the [official setup guide](https://code.claude.com/docs/en/setup).

### 2. Create a project folder and start Claude Code

Create a separate folder for your application and change into it (these commands work in Bash, Zsh, and PowerShell):

```bash
mkdir learning-campus-light
cd learning-campus-light
```

Replace `<ANTHROPIC_API_KEY>` with the key from Learning Campus and start Claude Code with Sonnet 5.5.

**macOS, Linux, or WSL (Bash/Zsh):**

```bash
ANTHROPIC_API_KEY="<ANTHROPIC_API_KEY>" claude --model claude-sonnet-5-5
```

**Windows (PowerShell):**

```powershell
$env:ANTHROPIC_API_KEY = "<ANTHROPIC_API_KEY>"
claude --model claude-sonnet-5-5
```

Complete the first-launch prompts. If asked whether to use the API key, confirm. Review the project folder shown before accepting the folder trust prompt. See the [authentication guide](https://code.claude.com/docs/en/authentication) if authentication fails.

### 3. Verify the model

Inside Claude Code, enter `/status` to check the active model. It should be **Sonnet 5.5**. To select it explicitly in a running session, enter:

```text
/model claude-sonnet-5-5
```

Alternatively, enter `/model` to open the model picker. See the [official model configuration guide](https://support.claude.com/en/articles/11940350-claude-code-model-configuration).

### 4. Install Superpowers

Inside Claude Code, install the [Superpowers plugin](https://github.com/obra/superpowers/blob/main/README.md) from the official marketplace:

```text
/plugin install superpowers@claude-plugins-official
```

Restart Claude Code in your project folder using the API-key and model command from step 2. Enter `/plugin` and check that Superpowers is installed and enabled. Its **brainstorming** skill will be used in the first task below.

> [!CAUTION]
> Your code is sent to third-party services. Do not include sensitive data or credentials in your prompts.

# Develop a Learning Campus Light in Go

## Project Overview

You will design and implement a simplified learning management system called "Learning Campus Light" using Go.
The key principle: **You define the scope in your spec, then implement what you specified.**

## Scope of the first version

Build a small application with three core operations:

1. List predefined courses.
2. Read a course's text document.
3. Edit a course's text document.

Start with a CLI and a few sample courses. Decide during brainstorming whether to keep data in memory or save it in a JSON file; make the behavior after a restart explicit in your spec.

If time permits, add a simple web UI or a simulated role selector: students can read, instructors can also edit. Leave authentication and passwords out of this exercise. A role selector only simulates permissions; it does not authenticate users.

**Remember**: Complete and check the small version before adding features.

## Step 1: Discuss your idea with the agent

Give the agent this complete exercise sheet: copy it into your project folder as `exercise.md`, or paste its contents into the conversation. This gives it the context for the implementation, checks, and reflection as well as the first task.

Use the **Superpowers brainstorming skill** to explore what your Learning Campus Light should do. Start with a prompt such as:

```text
Read exercise.md for the full assignment. Work only on Step 1 for now.
Use the Superpowers brainstorming skill to discuss my idea for a
Learning Campus Light application.
```

Participate actively in the discussion:

- Explain who will use the application and which problem it should solve.
- Decide which features belong in the first version and which to leave out.
- Discuss alternatives for the interface, data storage, and architecture.
- Challenge the agent's suggestions and ask it to explain trade-offs.
- Agree on concrete examples of successful behavior and error cases.

Step 1 is complete when you can explain these results in your own words:

## Step 2: Write the Mini-Spec (Markdown Document)

Turn the results of your brainstorming discussion into a specification document together with the agent. It should describe what your Learning Campus Light will do and how it will work. Your spec should include:

### Suggested Sections

1. **Product Vision**
   - What problem does this solve?
   - Who are the users?
   - What's the core value proposition?

2. **UI Design**
   - Describe the menu structure
   - Show example interactions (mockup of terminal output)
   - User workflows (e.g., "User lists courses → Selects course → Reads document")

3. **Use Cases**
   - List the main scenarios (e.g., "Instructor edits course content", "Student views course materials")
   - Include edge cases you'll handle (e.g., "User enters invalid course ID")

4. **Technical Specification**
   - **Architecture**: High-level structure (packages, modules)
   - **Data Storage**: JSON files, in-memory, or other approach
   - **Testing Strategy**: What will you test and how?
   - **Error Handling**: How will you handle invalid inputs, missing files, etc.?

Check that the spec reflects your decisions from the discussion. Resolve open questions and approve the spec before asking the agent to implement it.

### Tips for Working with SWE-Agents

* **Use Planning Mode**: If available, use the planning mode of the SWE agent to help draft and refine the spec.
* **Iterate**: Refine the spec in several rounds with the SWE-agent before implementation.
* **Be Specific**: The more detailed your spec, the better the agent can help implement it.

## Step 3: Implement with SWE-Agent Assistance

Use the coding agent to help you implement the application according to your spec.

**Implementation Requirements:**
- Follow your spec as closely as possible
- Write tests for core functionality
- Implement proper error handling
- Follow Go best practices and conventions
- Document your code with comments where appropriate

## Step 4: Check the application yourself

Run the application yourself and try all three core operations.

## Step 5: Document Your Findings on the Miro Board

Create a post on the Miro board that includes:

1. **Your Specification**
   - Post the mini-spec document you created

2. **Screenshots & Demos**
   - Screenshots of your running application
   - Show both successful operations and error handling
   - Include your independent check record from Step 4 and your explanation of the selected function

3. **SWE-Agent Usage**
   - How did you use Claude Code with Sonnet 5.5?
   - How did you structure your prompts/conversations?
   - How did the Superpowers brainstorming discussion shape your requirements and design?
   - Which alternatives did you discuss, and why did you choose your approach?
   - How many conversation sessions/iterations did it take?

4. **Evaluation**
   - **A concrete decision**: Show one agent suggestion you corrected or rejected, with a short excerpt or screenshot and your reason. If you accepted all suggestions, choose one and explain which alternative you considered and why you accepted it.
   - **What worked well?** (e.g., agent understood requirements, generated clean code, good at testing)
   - **What didn't work?** (e.g., hallucinations, incorrect implementations, needed manual fixes)
   - **Spec adherence**: How closely did the final implementation follow your spec? What changed and why?
   - **Code quality**: Assess the quality of generated code (structure, readability, tests, error handling)
   - **Productivity**: Did the agent speed up development or slow it down?

5. **Lessons Learned**
   - What strategies worked best for working with the SWE-agent?
   - What would you do differently next time?
   - How could the agent be improved?
