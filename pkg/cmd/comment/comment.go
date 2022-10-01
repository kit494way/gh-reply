package comment

import (
	"fmt"
	"os"

	"github.com/cli/go-gh"
	"github.com/kit494way/gh-reply/pkg/editor"
	"github.com/kit494way/gh-reply/pkg/prompter"
	"github.com/spf13/cobra"
)

type CommentOptions struct {
	Editor            bool
	Repo              string
	SavedReplyID      string
	Target            string
	IssueOrPRSelector string
}

func NewCommentCmd() *cobra.Command {
	opts := &CommentOptions{}

	var issueSelector, prSelector string

	cmd := &cobra.Command{
		Use:   "comment <saved_reply_id> [flags]",
		Short: "Add a comment to an issue or a pull request by using a saved reply",
		Long: `Add a comment to an issue or a pull request by using a saved reply.

With '--editor', you can edit a saved reply before adding a comment.`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if issueSelector != "" && prSelector == "" {
				opts.Target = "issue"
				opts.IssueOrPRSelector = issueSelector
			} else if issueSelector == "" && prSelector != "" {
				opts.Target = "pr"
				opts.IssueOrPRSelector = prSelector
			} else {
				return fmt.Errorf("only one of --issue or --pr can be enabled and one of each is required")
			}

			if len(args) > 0 {
				opts.SavedReplyID = args[0]
			}

			return runComment(opts)
		},
	}

	cmd.Flags().StringVarP(&issueSelector, "issue", "i", "", "Issuer number or url")
	cmd.Flags().StringVarP(&prSelector, "pr", "p", "", "Pull request number, url or branch")
	cmd.Flags().StringVarP(&opts.Repo, "repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	cmd.Flags().BoolVarP(&opts.Editor, "editor", "e", false, "edit a saved reply before adding a comment")

	return cmd
}

func runComment(opts *CommentOptions) error {
	interactive := opts.SavedReplyID == ""
	savedReply, err := prompter.SmartSelectSavedReply(opts.SavedReplyID)
	if err != nil {
		return err
	}

	body := savedReply.Body
	if opts.Editor {
		body, err = editor.Edit(savedReply.Body, os.Stdin, os.Stdout, os.Stderr)
		interactive = true
		if err != nil {
			return err
		}
	}

	ghargs := []string{opts.Target, "comment", opts.IssueOrPRSelector, "--body", body}
	if opts.Repo != "" {
		ghargs = append(ghargs, "--repo", opts.Repo)
	}

	if interactive {
		var msg string
		if opts.Repo == "" {
			msg = fmt.Sprintf("Executing: gh %s comment %s --body \\\n%s\nDo you continue?", opts.Target, opts.IssueOrPRSelector, body)
		} else {
			msg = fmt.Sprintf("Executing: gh %s comment %s --repo %s --body \\\n%s\nDo you continue?", opts.Target, opts.IssueOrPRSelector, opts.Repo, body)
		}
		answer, err := prompter.Confirm(msg, true)
		if err != nil || !answer {
			return err
		}
	}

	stdout, _, err := gh.Exec(ghargs...)
	if err != nil {
		return fmt.Errorf("Failed to execute: gh %s -R %s, err: %w", opts.Target, opts.Repo, err)
	}
	fmt.Print(stdout.String())
	return nil
}
