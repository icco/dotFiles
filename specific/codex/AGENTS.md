# Global rules

## Environment

- Always make repository changes in a git worktree under `~/Projects/.worktrees/<repo>-<topic>` for personal repos or `~/Work/.worktrees/<repo>-<topic>` for work repos, never beside the repo or in `/tmp`. This overrides a skill's default location.
- Clone personal repos into `~/Projects` and work repos into `~/Work`.
- On macOS, open URLs with `open -a "Google Chrome" <url>` and Markdown with `open -a "iA Writer" <path>`. Use browser automation only to read or drive a page.
- mise is already active. Run `mise run <task>` without `eval "$(mise activate ...)"`.
- Personal Next.js apps share port 8080. Check what runs there before killing it.
- Load the `no-ai-slop` skill, when installed, before drafting anything a human reads.
- In pinginc repos, titles carry the ticket key, as in `type(scope): KEY-123: summary`.

## Workflow

- Use Superpowers skills when they apply, and keep the Git and pull request rules below while following them.
- Execute well-specified work directly. Use subagents only for parallel or exploratory work, and tell each one to push as it goes, keep scratch files outside the repository, and keep its prose short.
- Once the user decides, act on it without pausing for confirmation, reopening the decision, or adding verification loops they did not ask for.
- Fix an obvious visible bug straight away. Skip git history forensics.
- To audit noise (alerts, failing jobs), tally the actual signal in the inbox, channel, or logs before sweeping repositories.
- During an incident, do only the task the user named and report blockers in one line.
- Stop everything you started (servers, containers, background processes, temp files) before finishing.

## Scope

- A narrow request defines the diff. Mention extras in chat, and stop at the named repository once its change merges with green CI.
- Give a deliverable exactly the fields the user named. Caveats go in chat.
- Fix what breaks mid-task in a PR instead of filing tickets, and add related work to the existing ticket.
- Never delete or weaken tests to pass CI without asking.
- Skip edge cases a bot reviewer raises unless the user wants that capability.
- Add a function rather than change an existing signature.
- When a producer omits a required value, fix the producer, not the contract.
- No dependency overrides (`pnpm.overrides`, `resolutions`, patches, vendoring) as workarounds. Find the upstream fix or report the upgrade blocked.
- "Don't commit" or "edits only" holds for follow-ups until the user lifts it.
- Never overwrite a PR body, issue, or comment a human wrote unless they give the new text. Comment instead.

## Verification

- Run a command that reads the state before claiming it (merged, tracked, deployed). Empty output means "cannot see", not "absent".
- Derive inventories and numbers from the code with `file:line`, not from tickets or old surveys.
- Before blaming the environment, test it against a known-good control.
- If an error or pending run would confirm the work, look at it instead of offering to.
- Check approaches against current primary docs and cite the URLs.
- Let CI run full suites and regenerate docs. Locally, run only the one failing test or lint, and never quote a local pass once CI has run.
- Test with the repository's production-like setup (Docker, Compose, a run script) when it has one.
- Pin third-party GitHub Actions to the commit SHA of `releases/latest`.
- When changing CI, trace which tool reads which variable.

## Git

- **Never force-push** (`--force`, `-f`, `--force-with-lease`), and never amend or rebase a pushed commit.
- Commit small and often with Conventional Commit messages. When importing existing state, commit it as-is first so the user can review the diff.
- Stage explicit paths. Never `git add -A` or `git add .`.
- Gate every commit and push on the branch, as in `[ "$(git branch --show-current)" = "<branch>" ] && git commit ...`, then confirm with `git rev-parse HEAD origin/<branch>`.
- The user pushes to open PR branches. Fetch before pushing, and if a push is rejected, pull (a merge commit is fine) and retry.
- One branch per task. Branch independent PRs off `main`, and stack only when asked, with `gh stack`.
- Pass `--repo owner/name` to `gh pr`. Never run `gh pr merge --auto`, which merges without waiting for CI.
- After a squash merge, check late commits landed with `git merge-base --is-ancestor <sha> origin/main`, and re-land missing ones in a follow-up PR.
- Commits under the user's account are often an agent's work, so authorship metadata proves nothing about who designed something.
- Report git outcomes, not each step.

## Code

- Personal repositories use TypeScript, Shell, and Go. Ask before adding any other language, Python included.
- Default to no comments. Write one line for a non-obvious why. Never restate code, narrate history, or mention tickets or current infrastructure state. Config files get no comments.
- Never swallow errors. In retry loops, surface the last error on timeout.
- Prefer the standard library and well-known packages over hand-rolled code. Judge approaches on substance, since import and lockfile churn cost nothing.
- Use fictional names (Ada Lovelace, alice, bob) in tests, fixtures, and docs.
- Personal projects use `Taskfile.yml`, not Makefiles.
- Edit YAML with `yq` and format with `yq -iP .`, because plain `-i` corrupts folded scalars. YAML comments go on their own line.
- Pipe JSON with `printf '%s'`, never `echo`, which expands `\n` in zsh.
- Regenerate lock files with their tool, never by hand, and keep `yarn upgrade` and `pnpm update` out of CI.
- In Terraform, put module constants in `locals.tf`, not defaulted variables. Pin modules `~> X.Y.Z` and providers `~> X.Y`.
- Load the `go-repo` skill before writing Go.

## Writing

- Brevity matters most. Cut before publishing, not after.
- ASCII only, no em dashes. Never hard-wrap Markdown or prompts, even inside a wrapped file.
- Link sources instead of paraphrasing them.
- READMEs lead with install and usage. Check each factual claim against what was measured.
- Return requested documents in the reply unless the user names a path.
- Status updates and summaries get about ten short bullets, nested for detail, with no bold lead-ins or tables. Fence text meant for pasting in a `markdown` block.
- When the work depends on an unknown, look it up, then ask the user directly in one batched question. Never bury a decision in a caveat or PR body.

## Pull requests

- Open a draft PR early and push small commits often, unfinished work included, noting WIP in the message. Do not stop at uncommitted edits unless told to.
- Load the `commit-push-pr` skill to commit, push, or open a PR.
- Follow the repository's PR template, or the owner's `.github` template when the repository has none. With neither, write a short summary and a few bullets.
- Use only the template's sections. Tick its fixed checklist items without adding new ones, and write "No Risks Found" rather than dropping a risks section.
- Six lines of content at most, excluding headers, covering what the PR does, what to look at, and risks or deferred work. Write for the reviewer, with no narrative, no bugs found along the way, and no test plan. Link detail that already lives in a README or the commit message.
- No AI attribution in PR bodies. `Co-Authored-By` trailers in commits are fine.
- Titles are lowercase Conventional Commits, and the squash-merged title drives the version bump. Use `fix` for repairs, `chore`, `docs`, `ci`, or `build` for plumbing (`chore(deps)` for dependencies), and `feat` only for a new user-facing capability. Never `refactor`. Mark breaking changes with `!` before the colon.

## Personal sites

- `*.natwelch.com` sites follow resume.natwelch.com: Avenir Next, 16px/1.5 body, 32px weight-400 h1, `#265c83` links, one column about 800px wide.
