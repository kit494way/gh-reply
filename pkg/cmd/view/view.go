package view

import (
	"github.com/kit494way/gh-reply/pkg/prompter"
	"github.com/kit494way/gh-reply/pkg/savedreply"
	"github.com/spf13/cobra"
)

type ViewOptions struct {
	ID string
}

func NewViewCmd() *cobra.Command {
	var opts = &ViewOptions{}

	cmd := &cobra.Command{
		Use:   "view <id>",
		Short: "Display the title, body and noed ID of a saved reply",
		Long: `Display the title, body and noed ID of a saved reply

Without an argument, you can interactively select a saved reply.`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				opts.ID = args[0]
			}

			return viewRun(opts)
		},
	}

	return cmd
}

func viewRun(opts *ViewOptions) error {
	savedReply, err := prompter.SmartSelectSavedReply(opts.ID)
	if err != nil {
		return err
	}

	return savedreply.PrintSavedReply(savedReply)
}
