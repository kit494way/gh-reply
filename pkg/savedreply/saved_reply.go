package savedreply

import (
	"fmt"
	"strings"

	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/markdown"
	"github.com/cli/go-gh/pkg/tableprinter"
	"github.com/cli/go-gh/pkg/term"
	graphql "github.com/cli/shurcooL-graphql"
)

// see https://docs.github.com/en/get-started/writing-on-github/working-with-saved-replies/about-saved-replies
const MaxSavedReplies = 100

type SavedReply struct {
	ID    string
	Title string
	Body  string
}

type SavedRepliesAndTotalCount struct {
	SavedReplies []SavedReply
	TotalCount   int
}

func ListSavedReplies(first int) (*SavedRepliesAndTotalCount, error) {
	client, err := gh.GQLClient(nil)
	if err != nil {
		return nil, err
	}

	var query struct {
		Viewer struct {
			SavedReplies struct {
				Nodes      []SavedReply
				TotalCount int
			} `graphql:"savedReplies(first: $first)"`
		}
	}

	variables := map[string]interface{}{
		"first": graphql.Int(first),
	}

	err = client.Query("SavedReplies", &query, variables)
	if err != nil {
		return nil, err
	}

	res := SavedRepliesAndTotalCount{
		SavedReplies: query.Viewer.SavedReplies.Nodes,
		TotalCount:   query.Viewer.SavedReplies.TotalCount,
	}
	return &res, nil
}

func GetSavedReply(id string) (*SavedReply, error) {
	client, err := gh.GQLClient(nil)
	if err != nil {
		return nil, err
	}

	var query struct {
		Node struct {
			SavedReply SavedReply `graphql:"... on SavedReply"`
		} `graphql:"node(id: $id)"`
	}

	variables := map[string]interface{}{
		"id": graphql.ID(id),
	}

	err = client.Query("SavedReply", &query, variables)
	if err != nil {
		return nil, err
	}

	return &query.Node.SavedReply, nil
}

func PrintSavedReply(reply *SavedReply) error {
	terminal := term.FromEnv()

	if !terminal.IsTerminalOutput() {
		fmt.Fprintf(terminal.Out(), "%s", reply.Body)
		return nil
	}

	md, err := markdown.Render(reply.Body, markdown.WithTheme(terminal.Theme()))
	if err != nil {
		return err
	}

	fmt.Fprintf(terminal.Out(), "ID: %s\n", reply.ID)
	fmt.Fprintf(terminal.Out(), "Title: %s\n", reply.Title)
	fmt.Fprintf(terminal.Out(), "%s\n", md)
	return nil
}

func PrintSavedReplies(replies []SavedReply, totalCount int) error {
	terminal := term.FromEnv()

	if terminal.IsTerminalOutput() {
		fmt.Fprintf(terminal.Out(), "\nShowing %d of %d saved replies\n\n", len(replies), totalCount)
	}

	termwidth, _, _ := terminal.Size()
	t := tableprinter.New(terminal.Out(), terminal.IsTerminalOutput(), termwidth)

	replacer := strings.NewReplacer("\r", "\\r", "\n", "\\n")
	for _, reply := range replies {
		t.AddField(reply.ID)
		t.AddField(reply.Title)
		body := replacer.Replace(reply.Body)
		t.AddField(body)
		t.EndRow()
	}

	return t.Render()
}
