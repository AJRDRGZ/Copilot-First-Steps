# Copilot Kata — Sample Go Repo

This repo is intentionally small and a bit **imperfect** so you can practice the Kata:
prompting, refactoring, testing, commits, and PRs with GitHub Copilot.

## Quick start
```bash
# 1) Run
go run ./cmd/katanasvc

# 2) Test
go test ./...

# 3) Benchmarks
go test -bench=. -run=^$
```

> Tip: Keep the terminal and Copilot Chat open side by side.

---

## Suggested Exercises (map to your Kata)
**Level 1 — Basics**
- Ask Copilot to generate **docstrings** for public functions.
- Request a **naming pass** on confusing variables/params.
- Small refactor: remove dead code and magic numbers in `internal/strings/slugify.go` and `internal/mathutil/series.go`.

**Level 2 — Skilled Practice**
- Create/iterate tests using the template in `.github/prompts/tests.md`.
- Add **custom instructions** (house rules) so Copilot follows your conventions.
- Ask Copilot to draft a **Conventional Commit** message from your staged diff.
- Ask Copilot to draft a **PR description** using `.github/PULL_REQUEST_TEMPLATE.md`.

**Level 3 — Mastery**
- Compare **recursion vs iterative** Fibonacci for performance (benchmarks included).
- Identify and fix **code smells** in `internal/user/repo.go` (e.g., concurrency safety).
- Add a **pre-commit checklist** to avoid secrets; propose a simple script or a doc.

## Repo structure
```
.
├── cmd/katanasvc/main.go
├── internal/
│   ├── mathutil/series.go
│   ├── mathutil/series_bench_test.go
│   ├── parser/parser.go
│   ├── parser/parser_test.go
│   ├── strings/slugify.go
│   └── user/{user.go,repo.go}
├── testdata/sample_config.json
├── .github/
│   ├── prompts/{tests.md,refactor.md}
│   └── PULL_REQUEST_TEMPLATE.md
├── go.mod
├── LICENSE
├── Makefile
└── README.md
```

## Push to your GitHub
```bash
git init
git add .
git commit -m "feat: seed Copilot Kata repo"
# Replace the URL with your repo
git branch -M main
git remote add origin https://github.com/<you>/copilot-kata-go.git
git push -u origin main
```
