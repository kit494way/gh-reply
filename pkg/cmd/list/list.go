package list

import (
	"fmt"

	"github.com/kit494way/gh-reply/pkg/savedreply"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	LimitResults int
}

func NewListCmd() *cobra.Command {
	opts := &ListOptions{}

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List saved replies",
		Long:  `List saved replies.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.LimitResults < 1 || opts.LimitResults > savedreply.MaxSavedReplies {
				return fmt.Errorf("limit must be between 1 and %d, but %d was specified.", savedreply.MaxSavedReplies, opts.LimitResults)
			}
			return listRun(opts)
		},
	}

	cmd.Flags().IntVarP(&opts.LimitResults, "limit", "L", 30, "Maximum number of saved replies to fetch, between 1 and 100")

	return cmd
}

func listRun(opts *ListOptions) error {
	res, err := savedreply.ListSavedReplies(opts.LimitResults)
	if err != nil {
		return err
	}

	return savedreply.PrintSavedReplies(res.SavedReplies, res.TotalCount)
}
