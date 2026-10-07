# Quiz Engine

A small quiz engine for the lecture: a **Go** HTTP server with an embedded
**TypeScript** web UI. Quizzes are plain YAML files. There is no login and no
database — everything a run needs lives in memory and is gone when the server
stops.

```
go run ./src/quiz/cmd/quizserver -dir src/quiz/quizzes
```

then open <http://localhost:8080>.

## Contents

| Path | What it is |
|---|---|
| `quiz/` | Domain model: parsing, validation, grading |
| `session/` | Session state and the in-memory session store |
| `server/` | HTTP API and static file serving |
| `web/` | TypeScript frontend (Vite, no framework) |
| `quizzes/` | The quiz files served by default |
| `cmd/quizserver/` | The server binary |
| `cmd/quizpack/` | Tool that inlines images into a quiz file |

## The quiz file format

A quiz is one YAML file. It is **self-contained**: images are embedded as
base64, so any single file can be uploaded and used without further assets.

```yaml
title: Concurrency in Go            # required
description: Lecture 7 recap        # optional

questions:                          # required, at least one
  - type: single                    # single | multiple | truefalse | text
    text: What does this print?     # required
    points: 2                       # optional, default 1
    code:                           # optional source snippet
      lang: go
      source: |
        ch := make(chan int)
        close(ch)
        fmt.Println(<-ch)
    image:                          # optional picture
      data: iVBORw0KGgoAAAANS…      # base64, no data: prefix, no line breaks
      mime: image/png
      alt: A channel diagram
    options:
      - {id: a, text: "0"}
      - {id: b, text: "deadlock"}
    answer: a                       # required
    explanation: >-                 # optional, shown after answering
      Receiving from a closed channel returns the zero value.
```

### Question types

| `type` | Control | `answer` | `options` |
|---|---|---|---|
| `single` (default) | radio buttons | one option id | required, at least 2 |
| `multiple` | checkboxes | list of option ids | required, at least 2 |
| `truefalse` | radio buttons | `"true"` or `"false"` | optional, generated if omitted |
| `text` | text field | list of accepted strings | must be absent |

`answer` accepts a scalar (`answer: a`) or a list (`answer: [a, c]`).

* **`multiple`** is graded all or nothing: the selected set must match exactly.
* **`text`** is compared case-insensitively, with leading, trailing and repeated
  whitespace collapsed. List every phrasing you want to accept:
  `answer: ["go", "the go keyword"]`.
* **`truefalse`** without `options` gets `True` / `False` automatically. Quote
  the answer (`answer: "true"`), otherwise YAML turns it into a boolean.

### Fields in detail

| Field | Where | Meaning |
|---|---|---|
| `title` | quiz | Shown in the list and above every question. Required. |
| `description` | quiz | One line under the title in the quiz list. |
| `text` | question | The question itself. Required. |
| `points` | question | Score for a correct answer. Default `1`, must not be negative. |
| `code.source` | question | Rendered monospace, whitespace preserved. Use a YAML block scalar (`|`). |
| `code.lang` | question | Language name, e.g. `go`. Emitted as a `language-go` class. |
| `image.data` | question, option | Base64 of the image file. No `data:` prefix. |
| `image.mime` | question, option | e.g. `image/png`. Defaults to `image/png` in the UI. |
| `image.alt` | question, option | Alternative text. |
| `options[].id` | question | Short, unique within the question. Referenced by `answer`. |
| `options[].text` | question | The option label. |
| `explanation` | question | Shown after answering, for both right and wrong answers. |

Questions are addressed by their position in the file; they have no ids.

### Validation

A quiz file is validated when it is loaded, and a broken file is refused with a
message naming the question. Rejected are: a missing title, a quiz without
questions, an unknown `type`, a missing `answer`, an `answer` that is not one of
the `options`, fewer than two options, duplicate option ids, options on a `text`
question, and negative points. A broken file in `quizzes/` is skipped with a log
message; the server still starts with the remaining quizzes.

Files may be at most 8 MiB, which base64 images make relevant.

### Images without the base64 pain

Writing base64 by hand and reviewing it in a diff is unpleasant. `quizpack`
lets you keep a readable source file that references images by path:

```yaml
image: {path: diagrams/channels.png, alt: A channel}
```

and turns it into the base64 form on the way out:

```
go run ./src/quiz/cmd/quizpack -in lecture7.src.yaml -out src/quiz/quizzes/lecture7.yaml
```

Paths are resolved relative to the input file, the MIME type is taken from the
file extension, and the packed result is validated before it is written.

## Running a quiz

Pick a quiz from the list, or upload a YAML file to try it out. An uploaded
quiz is parsed into memory and attached to that one session: it is never
written to disk and never appears in the list for anyone else. Sessions are
dropped after an hour of inactivity (`-ttl`).

Each question is answered once. A reload does not give a second attempt.

## Anonymous group quizzes

Choose **Host group quiz** beside a quiz, or check **Host uploaded quiz as a
group** before uploading YAML. The host lobby displays a locally generated QR
code and a copyable join link. Participants scan it and tap **Join**, without
entering a name or creating an account.

The host starts the quiz, reveals results for each question, and advances to
the next question. Participants answer independently on their phones, once per
question. Before reveal, only submission progress is visible. After reveal,
everyone sees aggregate results and the authored explanation; each participant
also sees their own feedback. Late arrivals can join the current phase, with
earlier questions left unanswered. There is no timer or leaderboard.

Choice results show option counts and percentages of submissions. Multiple-choice
percentages can sum above 100%. Text results show only correct/incorrect counts
and the authored accepted answer; submitted text is discarded after grading and
never displayed. The final summary shows participation, completion, per-question
accuracy and average percentage score across people who submitted at least once.
All quiz questions count toward that score, even when the host finishes early;
unanswered questions earn zero. The per-question summary covers opened questions.
With no submissions or a zero-point quiz, the average is shown as zero.

**Finish early** closes submissions and joining. **Delete room** immediately
removes the room and all results. Otherwise, rooms expire after `-ttl` without
successful authenticated activity, or disappear on server restart. Open host or
participant tabs poll every two seconds and keep the room alive. Public join-page
views do not extend expiry.

### Phone access and QR links

Phones need access to the running server. `localhost` on a phone refers to that
phone, not the presenter's computer. Open the host page using a reachable LAN or
public address, or explicitly configure the QR link base URL:

```sh
go run ./src/quiz/cmd/quizserver -dir src/quiz/quizzes \
  -public-url http://192.168.1.42:8080
```

Replace the example address with your server's reachable address. The default is
the host browser's origin and page path. Production deployments should use HTTPS.
The QR is generated in the browser; no external QR service receives the link.

### Anonymity and recovery

The application collects no names, emails, IP-based identities, fingerprints, or
analytics. Random room-specific bearer tokens distinguish participants and give
the host control. Tokens stay in tab-scoped session storage, so refreshing or
reconnecting in the same tab restores access. Clearing storage or starting in a
fresh browser creates a different participation; the application cannot enforce
one human per entry. There is no participant list or individual-score API for
the host. Only group totals are shared, although small groups can still make
individual outcomes inferable.

Application request logs contain route templates, not actual session/room IDs,
credentials, answer bodies, or client IPs. Configure hosting and reverse-proxy
logs accordingly: this application does not guarantee network-level anonymity.

### Group HTTP API

Host creation returns `{id, token, publicURL}`. Joining returns `{token}`. The join
link contains only the room ID. Authenticated requests use
`Authorization: Bearer <token>`; never put a credential in a URL.

| Method | Path | Purpose |
|---|---|---|
| `POST` | `/api/groups` | Create a room, body `{"quizId":"…"}` |
| `POST` | `/api/groups/upload` | Create from multipart field `file` (8 MiB request limit) |
| `POST` | `/api/groups/{id}/join` | Join without authentication while the room is open |
| `GET` | `/api/groups/{id}` | Authenticated state; only the caller's personal feedback |
| `POST` | `/api/groups/{id}/answers/{n}` | Participant submission, body `{"answer":["a"]}` |
| `POST` | `/api/groups/{id}/actions` | Host action, body `{"action":"reveal","revision":1}` |

State includes `phase` (`lobby`, `question`, `reveal`, `finished`), `revision`,
question number, participant/submission counts and the sanitized current question.
`correctAnswer`, `explanation`, `stats` and `personal` appear only after reveal or
finish; `summary` appears on finish. Host actions are `start`, `reveal`, `next`,
`finish`, and `delete`. Transitions require the current revision; stale requests
return 409 and cannot advance twice. Deletion requires only the host credential.
Repeated answers return the original acceptance without changing any count.
Invalid credentials return 403; unknown, deleted or expired rooms return 404.

## Solo HTTP API

The correct answer is **never** sent to the browser before the question has
been answered. `GET .../questions/{n}` returns a sanitized question without
`answer` and `explanation`; grading happens on the server, and the answer comes
back only in the response to `POST .../answers/{n}`.

| Method | Path | Purpose |
|---|---|---|
| `GET` | `/api/quizzes` | The quizzes on the server |
| `POST` | `/api/sessions` | Start a session, body `{"quizId": "…"}` |
| `POST` | `/api/upload` | Start a session from an uploaded file (multipart field `file`) |
| `GET` | `/api/sessions/{id}` | Session info |
| `GET` | `/api/sessions/{id}/questions/{n}` | Sanitized question, `n` is 1 based |
| `POST` | `/api/sessions/{id}/answers/{n}` | Grade an answer, body `{"answer": ["a"]}` |
| `GET` | `/api/sessions/{id}/result` | Score so far |

The session id is the only thing protecting a session, so it is 128 random bits.

## Flags

| Flag | Default | Meaning |
|---|---|---|
| `-addr` | `:8080` | Listen address |
| `-dir` | `src/quiz/quizzes` | Directory with quiz YAML files |
| `-ttl` | `1h` | How long an idle session or group room is kept |
| `-public-url` | browser origin/path | Reachable base URL used for group QR join links |

## Docker

```
docker compose -f src/quiz/compose.yaml up --build
```

The image is built in three stages — Node builds the frontend, Go builds a
static binary, and the result is copied into a distroless base — so the final
image contains the binary and the quiz files and nothing else. `quizzes/` is
mounted read-only, so editing a quiz and restarting the container is enough.

## Frontend development

`web/dist` is checked in and embedded with `go:embed`, so the server builds and
runs without any Node.js toolchain. Node is only needed when changing the UI:

```
cd src/quiz/web
npm install
npm run dev      # Vite dev server, proxies /api to localhost:8080
npm run build    # type check and rebuild web/dist — commit the result
```

## Design notes for the lecture

The interesting parts, as Go concepts:

* **`go:embed`** puts the whole frontend into one binary (`web/embed.go`).
* **Custom YAML unmarshalling** lets `answer` be either a scalar or a list
  (`Answer.UnmarshalYAML` in `quiz/model.go`).
* **Separating `Question` from `PublicQuestion`** makes "the client must not see
  the answer" a property of the type system rather than of a code review.
* **`sync.RWMutex` and a janitor goroutine** in `session/store.go` show shared
  mutable state done deliberately, next to all the channel examples.
* **Method and wildcard routing** (`GET /api/sessions/{id}`) from Go 1.22 means
  no third party router.
* **The `Session` interface** supports solo sessions. Group rooms have their own
  store and API because host controls and multiple participants need different
  authorization rules. Both modes reuse the same quiz domain model and grading.
* **A mutex around group transitions** makes answering versus revealing atomic;
  a revision number prevents duplicate host requests from advancing twice.

## Tests

```
go test ./src/quiz/...
```

Group concurrency and privacy checks:

```sh
go test -race ./src/quiz/...
cd src/quiz/web
npm ci
npm run build
npx playwright install chromium
npm run test:group
```

Alternatively set `QUIZ_BROWSER_BIN=/path/to/chrome` to use an installed browser.
The browser test builds and starts its own temporary Go server and checks QR
decoding, all question types, independent participants, uploads, late joining,
refresh, offline recovery, input preservation, mobile layout, statistics, and
room deletion. Optional `QUIZ_SCREENSHOT_DIR` retains summary screenshots.
