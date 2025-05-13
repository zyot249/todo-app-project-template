# Applied Rules for Git

## Commit Message

```format
<type>(<scope>): <subject>

<body>

<footer>
```

Where:

- `<type>` Indicates the kind of change:
  - `feat`: A new feature
  - `fix`: A bug fix
  - `test`: Adding or updating tests
  - `perf`: A performance improvement
  - `refactor`: A code change that neither fixes a bug nor adds a feature
  - `docs`: Documentation only changes
  - `chore`: Changes to build process, tools, etc. (Other changes that don't modify src or test files)

- `<scope>`(Optional) The scope of the change (e.g., component name, module).
  - `{service}`: The service that is affected by the change.
  - `{module}`: The module that is affected by the change.

- `<subject>` Short description in present tense, no period at the end.
  - Use the imperative, present tense: "change", "add", "remove" instead of "changed", "added", "removed".
  - Don't capitalize the first letter.
  - No dot (.) at the end.

- `<body>`(Optional) Detailed explanation if needed.

- `<footer>`(Optional) Breaking changes, issue references, etc.
