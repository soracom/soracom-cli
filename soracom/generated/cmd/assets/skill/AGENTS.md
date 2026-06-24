# AGENTS.md

Agent-first guidance for the `soracom` command (soracom-cli). If you are an AI
agent calling SORACOM APIs from the terminal, read this first. For the full
playbook see [`.agents/skills/soracom-cli/SKILL.md`](.agents/skills/soracom-cli/SKILL.md).

## What this CLI is

`soracom` is a thin, auto-generated wrapper over the SORACOM REST API. Each leaf
subcommand maps 1:1 to one API operation: it builds a request from your flags,
sends it, and prints the JSON response to stdout. There are ~500 commands, but
they are all generated from the same templates, so they behave identically.

## The core loop: discover → validate → execute

Do not guess command names, flags, or resource IDs. Let the CLI tell you.

1. **Discover** — `soracom describe` lists every command as machine-readable
   JSON `{command, method, path, summary}`. `soracom describe <command...>`
   returns one command's full schema (parameters and request-body properties
   with `type`, `in`, `required`, `enum`, `default`). No network or credentials
   needed. Prefer this over parsing `--help` prose.
2. **Validate** — append `--dry-run` to any command to print the exact request
   (`method`, `url`, redacted `headers`, parsed `body`) **without sending it**.
   Always dry-run a mutation (POST/PUT/DELETE) you are unsure about.
3. **Execute** — run it once the dry-run looks right.

## Invariants — rely on these

- **Output is JSON by default** on stdout. Nothing to parse heuristically.
- **Errors**: the process exits non-zero on failure. By default the error is
  printed in cobra's conventional format on stderr (stdout stays clean) — branch
  on the exit code. Set `SORACOM_JSON_ERRORS=1` to opt into machine-readable
  errors: API errors as the SORACOM API's own JSON body, client-side errors as
  `{"error":{"message":"..."}}`, with the usage dump suppressed.
- **`--fields a,b.c`** keeps only the named fields (dot notation for nested,
  applied per element for arrays). Use it on list/get commands to protect your
  context window — request only what the task needs.
- **`--fetch-all`** follows pagination and returns the whole list; combine with
  `--fields` to keep it small. **`--jsonl`** emits an array as JSON Lines.
- **`--body`** sends a raw JSON payload that maps directly to the API schema:
  `--body '{...}'`, `--body @file.json`, or `--body -` (stdin). This is the
  reliable way to set nested/object fields (e.g. `tags`, `configuration`) that
  have no dedicated flag. Learn the schema with `soracom describe <command>`.

## Authentication

The normal setup is interactive and done **by the user**, once: `soracom configure`
creates `~/.soracom/default.json` (Windows: `%USERPROFILE%\.soracom\default.json`).
Don't run `soracom configure` yourself — it prompts for secrets. If no profile
exists, ask the user to run it.

Discover what's configured **safely** before authenticated commands:

- `soracom configure list` → JSON array of profile **names** (no secrets). Empty
  → ask the user to run `soracom configure`. `["default"]` → just run commands.
  Several → ask which one and pass `--profile <name>`.

**Never read the credentials.** `~/.soracom/*.json` and `soracom configure get`
contain the AuthKeyId / AuthKey / passwords in **plaintext** (the files are not
encrypted) — do not open, `cat`, or print them, and do not run `configure get`.
Select an identity with `--profile <name>` and let the CLI read the secret.

For CI / headless / one-shot runs without a profile, pass credentials sourced
from a secret store or environment (never hardcoded or echoed):

- `--api-key` + `--api-token` (a pre-obtained key/token), or
- `--auth-key-id` + `--auth-key` + `--coverage-type g|jp` (a SAM AuthKey).

`--coverage-type g` = Global, `jp` = Japan. `SORACOM_ENDPOINT` overrides the
endpoint.

## Minimal example

```
soracom describe groups create                 # learn the schema
soracom groups create --body '{"tags":{"name":"floor-1"}}' --dry-run   # preview
soracom groups create --body '{"tags":{"name":"floor-1"}}' --fields groupId
```

---

## For agents modifying this repository

The CLI is **generated** — do not edit `soracom/generated/cmd/*.go` by hand
(they carry `// Code generated ... DO NOT EDIT.`). Change behavior at the source:

- API surface / help text: `generators/assets/soracom-api.{en,ja}.yaml`,
  `generators/assets/sandbox/...`, and `generators/assets/cli/{en,ja}.yaml`.
- Per-command code: `generators/cmd/templates/{root,trunk,leaf}.gotmpl`.
- Hand-written commands/helpers: `generators/cmd/predefined/*.go` (copied as-is
  into the generated package, so they may reference generated symbols like
  `RootCmd`; `*_test.go` here are copied and run too).

Then regenerate and verify:

```
make generate        # regenerate soracom/generated/cmd (needs goimports on PATH)
make build VERSION=0.0.0-dev
make test            # generator + lib tests
make test-generated  # generated package tests
make lint            # staticcheck (needs staticcheck on PATH)
```

The agent-first features above are implemented in the generator, so they apply
uniformly to every command: `--dry-run`/`--fields` (global flags in
`root.gotmpl`), the dry-run short-circuit and field filtering (`leaf.gotmpl`),
and the `describe` command, output filter and opt-in structured-error helpers
(`generators/cmd/predefined/`).
