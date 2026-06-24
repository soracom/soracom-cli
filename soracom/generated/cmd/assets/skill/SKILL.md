---
name: soracom-cli
description: >-
  How to drive the `soracom` CLI (soracom-cli) safely and efficiently as an AI
  agent to call SORACOM APIs — managing SIMs/subscribers, groups, devices,
  Harvest data, event handlers, billing and the rest of the SORACOM platform.
  Use this skill whenever the task involves the `soracom` command, SORACOM
  resources (IMSI, SIM, subscriber, group, SIM group, VPG, Soracom Air/Beam/
  Funnel/Funk/Harvest/Napter), or constructing SORACOM API calls from the
  terminal — even if the user does not explicitly say "soracom-cli". It covers
  the agent-first features (`describe`, `--dry-run`, `--fields`, `--body`,
  structured JSON output and errors) that let you discover schemas, validate
  requests before they take effect, and keep responses small.
---

# Using soracom-cli as an agent

`soracom` is a thin, auto-generated wrapper over the SORACOM REST API. Almost
every subcommand maps 1:1 to an API operation: it builds a request from your
flags, sends it, and prints the JSON response to stdout. Because it is
generated from the API definition, the CLI is large (~500 leaf commands) but
extremely regular — once you know the patterns below, every command behaves the
same way.

This skill exists because the failure modes of an agent differ from a human's.
You will not make typos, but you may hallucinate a flag, a command name, or a
resource ID; you may flood your own context window with a huge list response;
and you may run a destructive mutation when you only meant to inspect it. The
features below exist to defend against exactly those failures. Lean on them.

## The core loop: discover → validate → execute

Do not guess command names or parameters. The CLI can tell you its own schema.

1. **Discover** with `soracom describe` (see below) instead of guessing or
   parsing `--help` prose. It returns machine-readable JSON.
2. **Validate** any mutating command (POST/PUT/DELETE) with `--dry-run` first.
   It prints the exact request that *would* be sent, without sending it.
3. **Execute** once the dry-run looks right.

## Schema introspection: `soracom describe`

`describe` is your map of the CLI. It reads the embedded API definition and
emits JSON — no network call, no credentials needed.

- `soracom describe` — list every command as `{command, method, path, summary}`.
  Pipe through `grep`/`jq` to find the right one, e.g.
  `soracom describe | jq -r '.[].command' | grep subscriber`.
- `soracom describe <command...>` — full schema for one command: its
  parameters and request-body properties with `type`, `in` (path/query/body),
  `required`, `enum`, `default` and a one-line description. Example:
  `soracom describe groups create`.

Prefer `describe <command>` over `soracom <command> --help` when you need to
know precisely what a command accepts: the JSON is unambiguous and lists enums
and required fields that help text may bury.

Reading a request-body property:

- **`option`** — present only when a dedicated `--<option>` flag exists for that
  field. If a property has **no `option`**, it has no flag and must be supplied
  through `--body` (this is always the case for object/array fields such as
  `tags` or `configuration`).
- **`type`** — `"map of <T>"` means a free-form object/dictionary: the keys are
  caller-defined (e.g. service names, tag names), not a fixed set of
  sub-fields. `"array of <T>"` is a list.
- **`schema`** — the referenced component schema name for object fields (e.g.
  `GroupConfiguration`), useful as a stable identifier when reasoning about the
  shape.
- **`example`** (on `requestBody`) — a concrete sample payload. This is the most
  reliable template for building `--body`: copy it and adjust the values.
- **`response`** — the shape of the success response: its `type` (e.g.
  `array of Sim`), `schema` name, and `fields` (recursively expanded, so deep
  paths like `sessionStatus.cell.mcc` are visible). **Use this to choose
  `--fields` paths** instead of guessing — e.g. for `sims list` the response
  shows `sessionStatus.online` exists but there is no top-level `imsi`, so the
  correct projection is `--fields simId,sessionStatus.online`, not
  `--fields imsi,...`.

## Output is JSON by default

Every successful response is printed to stdout as indented JSON. There is
nothing to parse heuristically.

- `--fields a,b,c` — **keep only these fields** in the output. Use dot notation
  for nested fields (`--fields imsi,sessionStatus.online`). When the response is
  an array, the filter applies to each element. Reach for this on list/get
  commands to avoid dumping large objects into your context window — request
  exactly the fields the task needs and nothing more.
- `--jsonl` — for commands that return an array, emit one JSON object per line
  (JSON Lines). Convenient for streaming/iterating.
- `--fetch-all` — for paginated list commands, automatically follow pagination
  and return the full result set as a single array. Combine with `--fields` so
  the aggregated result stays small.
- `--raw-output` — print the API response body verbatim (no re-indenting). Used
  for binary/raw payloads (e.g. Harvest Files downloads).

## Sending request bodies: `--body`

For POST/PUT commands you can set individual body fields via generated flags,
but the most reliable path — especially for nested objects, arrays, or fields
that have no dedicated flag (object-typed fields like `tags` or `configuration`
never get one) — is the raw `--body`, which maps directly to the API schema:

- `--body '{"tags":{"name":"sensor-01"}}'` — inline JSON.
- `--body @payload.json` — read JSON from a file.
- `--body -` — read JSON from stdin.

Use `soracom describe <command>` to learn the exact body schema, then pass it
through `--body`. This is zero-translation-loss: the API reference *is* your
documentation.

## Validate before acting: `--dry-run`

`--dry-run` is available on every command. It builds the full request and
prints it as JSON — `method`, `url`, `headers` (secret and profile-supplied
headers redacted), and the parsed `body` — **without performing the call and
without any network-backed authentication**. It consumes a locally supplied
`--api-key`/`--api-token` (so the preview is faithful) but never exchanges a
profile or AuthKey for a token over the network, so it has zero side effects.

Always dry-run a mutation you are unsure about. It confirms the resource ID
landed in the right place in the path, the body parsed as you intended, and the
endpoint/coverage are correct, before anything changes on the account.

```
soracom groups create --body '{"tags":{"name":"sensor-01"}}' --dry-run
```

## Errors and exit codes

On failure the CLI exits non-zero. By default it prints the error in cobra's
conventional format on stderr while stdout stays clean, so you can branch on the
exit code alone.

If you want failures to be machine-readable, set the environment variable
`SORACOM_JSON_ERRORS=1`. Then errors are written to **stderr** as JSON: API
errors as the SORACOM API's own JSON error body, client-side errors (e.g. a
missing required parameter) wrapped as `{"error":{"message":"..."}}`, and the
usage dump is suppressed. Prefer this in agent runs so you never have to scrape
human text.

## Authentication

A SORACOM profile is normally created **by the user**, interactively and once:
`soracom configure` walks through coverage type and credentials and writes
`~/.soracom/default.json` (Windows: `%USERPROFILE%\.soracom\default.json`). Do
not run `soracom configure` yourself — it prompts for secrets. If the user is not
configured yet, ask them to run it.

**Discover what's configured, safely.** Before running authenticated commands,
run `soracom configure list` — it prints the profile **names** as a JSON array
and never reveals credentials:

- empty array → no profile yet; ask the user to run `soracom configure`;
- `["default"]` → just run commands normally (no `--profile` needed);
- several names → ask the user which one, then pass `--profile <name>` on every
  command (e.g. `soracom sims list --profile user1`).

**Never read the credentials.** The profile files `~/.soracom/*.json` store the
AuthKeyId, AuthKey and passwords in plaintext (they are not encrypted), and
`soracom configure get` prints that same content. Do not open, read, `cat`, or
print these files, and do not run `soracom configure get` — no agent task needs
the raw secret. Select an identity with `--profile <name>` and let the CLI read
the secret for you.

**Headless / CI without a profile.** Credentials may be passed on the command
line, but only sourced from a secret store or environment variable — never
hardcoded into a command you save, and never echoed back to the user:

- `--api-key` + `--api-token` — a pre-obtained API key/token, or
- `--auth-key-id` + `--auth-key` + `--coverage-type g|jp` — a SAM AuthKey.

`--coverage-type g` = Global, `jp` = Japan. `SORACOM_ENDPOINT` overrides the
endpoint.

**Confirm you're authenticated** with a harmless read, e.g.
`soracom operator get` or `soracom sims list --limit 1 --fields simId` — a zero
exit code means the credentials work. (`soracom configure list` only tells you a
profile file exists, not that its credentials are still valid.)

## Quick reference

| Goal | Flag / command |
| --- | --- |
| Find the right command | `soracom describe \| jq ...` |
| Learn a command's schema | `soracom describe <command>` |
| Preview a request, no side effect | `--dry-run` |
| Shrink the response | `--fields a,b.c` |
| Full paginated result | `--fetch-all` (+ `--fields`) |
| Array as JSON Lines | `--jsonl` |
| Send a raw/nested payload | `--body '{...}'` / `--body @file` / `--body -` |
| See which profiles exist (no secrets) | `soracom configure list` |
| Use a specific identity | `--profile <name>` |
| Headless auth (no profile) | `--api-key`+`--api-token` or `--auth-key-id`+`--auth-key --coverage-type` |

## Worked example

Task: enable a new group with a name tag, safely.

```
# 1. Discover the command and its body schema
soracom describe groups create

# 2. Validate the request without creating anything
soracom groups create --body '{"tags":{"name":"factory-floor-1"}}' --dry-run

# 3. Create it, keeping only the fields you need
soracom groups create --body '{"tags":{"name":"factory-floor-1"}}' \
  --fields groupId,createdAt
```
