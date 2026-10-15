---
name: commit-push-pr
description: Commit, push, and open or update a pull request following the user's PR rules. Use when asked to commit, push, ship, or open a PR.
---

1. Run `git status -sb`, `git diff HEAD`, and `git branch --show-current`. On `main`, create a branch first.
2. Read the PR template from `.github/pull_request_template.md`, or from the owner's `.github` repository when there is none (`gh api repos/<owner>/.github/contents/pull_request_template.md --jq .content | base64 -d`).
3. Stage explicit paths, never `git add -A` or `git add .`. Commit with a Conventional Commits message, lowercase after the prefix, gated on the branch name: `[ "$(git branch --show-current)" = "<branch>" ] && git commit ...`.
4. Gate the push the same way, push, and confirm `git rev-parse HEAD origin/<branch>` match.
5. Run `gh pr view --repo owner/name`. If a PR exists, leave its body alone and comment instead. Fix a title that fails the Conventional Commits check with `gh pr edit --title`.
6. Otherwise run `gh pr create --draft --repo owner/name` with a body that follows the rules below.

## Body rules

- Six lines of content at most, excluding template headers. One sentence on what the PR does, then what to look at, then risks and deferred work.
- Use the template's sections and nothing else.
- Write for the reviewer. No story of the work, no bugs found along the way, no test-plan checklist unless asked.
- No AI attribution line. No em dashes. One paragraph per line, no hard wrapping.
- Draft the body, cut it, then open the PR. Never publish a long first draft and trim afterwards.
