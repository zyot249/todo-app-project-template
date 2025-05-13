# Applied Rules for Git

## Branching

### Branch Name

```format
<type>/<ticket-number>-<short-description>
```

Where:

- `<type>` Indicates the kind of change:
  - `feature`: For new features
  - `bugfix` or `fix`: For bug fixes
  - `hotfix`: For urgent production fixes
  - `test`: For adding or updating tests
  - `perf`: For performance improvements
  - `refactor`: For code refactoring
  - `chore`: For other changes that don't modify src or test files
  - `docs`: For updating documentation
  - `release`: For release preparation

- `<ticket-number>`(Optional) The number of the ticket in the project management system.
- `<short-description>` Short description of the change (use lower case and separate words with hyphens).

## Commit Message

```format
<type>(<scope>): <subject>

<body>

<footer>
```

Where:

- `<type>` Indicates the kind of change:
  - `feat`: For new features
  - `fix`: A bug fix
  - `test`: Adding or updating tests
  - `perf`: A performance improvement
  - `refactor`: A code change that neither fixes a bug nor adds a feature
  - `docs`: Documentation only changes
  - `chore`: Changes to build process, tools, etc. (Other changes that don't modify src or test files)

- `<scope>`(Optional) The scope of the change (e.g., component name, module).
  - `{service}`: The service that is affected by the change.
  - `{module}`: The module that is affected by the change.

- `<subject>` Short description of the change.
  - Use the imperative, present tense: "change", "add", "remove" instead of "changed", "added", "removed".
  - No dot (.) at the end.

- `<body>`(Optional) Detailed explanation if needed.

- `<footer>`(Optional) Breaking changes, issue references, etc.

## Pull Request

### Title

```format
[<scope>]: <subject>
```

Where:

- `<scope>` The scope of the change (e.g., component name, module).
  - `{service}`: The service that is affected by the change.
  - `{module}`: The module that is affected by the change.
- `<subject>` Short description of the change.

### Description

Follow the [template](./PR_TEMPLATE.md)
