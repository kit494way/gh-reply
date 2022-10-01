# gh-reply

A GitHub CLI ([gh](https://github.com/cli/cli)) extension to make use of saved replies.

## Installation

```sh
gh extension install kit494way/gh-reply
```

## Usage

### list

List saved replies.

```sh
gh reply list
```

### view

View a saved reply.

```sh
gh view <ID>
```

Without id, `gh view`, you can interactively select a saved reply.

### comment

Add a comment to an issue using a saved reply.

```sh
gh reply comment <saved_reply_id> --issue <issue_number>
```

For pull request, use `--pr` option instead of `--issue`.

```sh
gh reply comment <saved_reply_id> --pr <pull_request_number>
```

To edit a saved reply before adding a comment, use `--editor` option.

```sh
gh reply comment <saved_reply_id> --issue <issue_id> --editor
```

This launch a editor initialized by a saved reply.

Without an argument `<saved_reply_id>`, you can interactively select a saved reply to use.

```sh
gh reply comment --issue <issue_id>
```

## License

MIT
