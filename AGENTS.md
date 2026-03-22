# Project Conventions

## Tooling
- `mise` is the primary task runner.
- Tools must be pinned to specific versions (never 'latest' or 'lts').
- `workspaced` via `mise` handles all linting and formatting.

## Workflows
- GitHub Actions must be consolidated into a single workflow file at `.github/workflows/autorelease.yml`.

## Error Handling
- Never fail silently or ignore unexpected errors.
- Do not leave empty catch blocks.
- The project has a centralized error reporting function `reportError(err error, context map[string]interface{})` located in `error.go`. All unexpected or unrecoverable errors must be funneled through this function instead of failing silently.

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
