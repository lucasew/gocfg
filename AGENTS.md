# Project Conventions and Guidelines

## General Rules
- Error handling must never be silent; all unexpected errors must be routed to a centralized, Sentry-aware error-reporting function. Empty catch blocks are prohibited.
- The project has a centralized error reporting function `reportError(err error, context map[string]interface{})` located in `error.go`. All unexpected or unrecoverable errors must be funneled through this function instead of failing silently.
- Tooling relies on `mise` as the primary task runner, with tools strictly pinned to specific versions (never 'latest' or 'lts'). In `mise.toml`, tasks like `install`, `test`, and `codegen` must only depend on wildcards (e.g., `[task]:*`), while the `ci` task must explicitly depend on `['lint', 'test']`.
- GitHub Actions must be consolidated into a single workflow file at `.github/workflows/autorelease.yml` with a standard flow: Install, Codegen, PR if Diff, CI, and conditionally Release/Artifacts (omitted for Vercel/Cloudflare).
- Pull Requests must include specific body sections: `Assumptions`, `Alternatives Not Chosen`, `How To Pivot`, and `Next Knobs`. Titles must adhere to the active agent's convention (e.g., `🛡️ Sentinel: [Severity] [Description]`, `🛟 Arrumador: [Description]`, or `🛠️ Refactor: [Description]`).
- Git staging must be explicit (`git add <path>`); using `git add -A` or `git add .` is prohibited. Tooling download artifacts (e.g., install scripts, binaries) must never be committed.
- All linting and formatting must be handled by `workspaced` via `mise`; manual installation of individual linters is strictly forbidden.
- Before planning work, read `.jules/CONSISTENTLY_IGNORED.md` to avoid proposing changes that match patterns consistently rejected by reviewers.

## Refactoring Guidelines
- Small, atomic changes: Prefer incremental refactorings.
- Enforce naming conventions.
- Enforce directory structure:
  - Things that change together should live together (colocation).
  - A directory should have a single, clear purpose.
  - Sparse directories should be merged.
  - Consistency: similar modules should follow the same internal layout pattern across the project.
- Apply design patterns where appropriate. Ensure cohesion (SRP) and reduce coupling.
- Reduce Complexity: Detect and fix code smells, replace magic numbers.
- Code Reuse: Respect the Rule of Three.
- Justify Extraction Changes (Extract Method, Extract Class) citing Fowler, Martin, GoF. Positional changes (file moves) need clear rationale.

## Testing Philosophy
- Passing automated checks is the baseline. Manual/interactive verification or explicit runtime reasoning is required.
- New tests must exercise actual implemented logic/edge cases and catch real bugs, not just test external libraries.

## Operational Memory
- `gocfg.go` -> INI-like configuration parser logic based on maps
- `sectionprovider.go` -> Interfaces and implementations for storing configuration sections (maps, environment)
- `error.go` -> Centralized error reporting
- `.github/workflows/autorelease.yml` -> Unified CI/CD workflow
