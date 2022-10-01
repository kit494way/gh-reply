package root

import (
	"os"

	"github.com/kit494way/gh-reply/pkg/cmd/comment"
	"github.com/kit494way/gh-reply/pkg/cmd/list"
	"github.com/kit494way/gh-reply/pkg/cmd/view"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gh-reply",
		Short: "gh extension to make use of saved replies",
		Long: `gh extension to make use of saved replies.

This command list and view saved replies.
You can also add a comment to an issue or a pull request using a saved reply.`,
	}

	cmd.AddCommand(list.NewListCmd())
	cmd.AddCommand(view.NewViewCmd())
	cmd.AddCommand(comment.NewCommentCmd())

	return cmd
}

func Execute() {
	rootCmd := NewRootCmd()

	err := rootCmd.Execute()
	if err != nil {
		// err is displayed by cobra.
		os.Exit(1)
	}
}
