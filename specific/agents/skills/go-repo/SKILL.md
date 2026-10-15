---
name: go-repo
description: Conventions for the user's Go repositories, covering code, libraries, scaffolding, lint, CI, versioning, and goreleaser releases. Use before writing Go in the user's repositories or changing their plumbing.
---

- Every exported identifier gets exactly one doc comment that starts with its name.
- In repositories that depend on `gutil`, use it for logging, JSON responses, and ETags. Use chi for routing and GORM with `AutoMigrate` for Postgres.
- Create repositories from the owner's `go-template` GitHub template rather than copying an older repository.
- Actions are disabled on the template itself. Confirm a new repository has them enabled with `gh api repos/<owner>/<repo>/actions/permissions`.
- Use the module path the owner's other repositories use. Libraries start at `v1.0.0`.
- Keep lint configuration in a committed `.golangci.yml`, never as `-E` flags in a workflow, so local runs and CI enforce the same set.
- golangci-lint caps output at 50 issues per linter. Pass `--max-issues-per-linter=0 --max-same-issues=0` for true counts.
- If golangci-lint refuses to load because it was built with an older Go than `go.mod`, rebuild it with the current toolchain at CI's version (`go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@<version>`). Never skip lint because of it.
- Measure coverage with the toolchain from `go.mod` (`GOTOOLCHAIN=go1.x.y`), since totals differ between Go versions.
- `.yamllint` disables `document-start` and `line-length` and sets `truthy: check-keys: false`.
- The yaml-json auto-format step uses `file_pattern: '. :!.github/workflows/*'`. A literal `*.json` makes `git add` exit 128 in a repository without JSON.
- Releases use goreleaser, tagged by `PaulHatch/semantic-version`. Wrap `major_pattern` and `minor_pattern` in `/ /` so scoped titles such as `feat(api):` count.
- The PR title drives both the conventional-commits check and the version bump. Fix a bad title with `gh pr edit --title`.
- GitHub releases are immutable, so a failed goreleaser run cannot be retried for the same tag. Cut a new tag.
- `GET /repos/{owner}/{repo}` reports the user's permissions, not the token's. Probe a fine-grained PAT with a `PUT /contents/<existing file>` that sends an all-zero `sha`. A 403 means no write access, and a 409 means write access with nothing changed.
- READMEs show `go get`, a usage example, and a link to godoc. No Go Report Card badge.
