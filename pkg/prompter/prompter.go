package prompter

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/kit494way/gh-reply/pkg/savedreply"
)

func SelectSavedReply() (*savedreply.SavedReply, error) {
	repliesAndCount, err := savedreply.ListSavedReplies(savedreply.MaxSavedReplies)
	if err != nil {
		return nil, err
	}

	options := make([]string, 0, repliesAndCount.TotalCount)
	lowerTitles := make([]string, 0, repliesAndCount.TotalCount)
	for _, reply := range repliesAndCount.SavedReplies {
		options = append(options, reply.ID)
		lowerTitles = append(lowerTitles, strings.ToLower(reply.Title))
	}

	prompt := &survey.Select{
		Message: "Select a saved reply:",
		Options: options,
		Description: func(value string, index int) string {
			return repliesAndCount.SavedReplies[index].Title
		},
		Filter: func(filter, value string, index int) bool {
			return strings.Contains(lowerTitles[index], strings.ToLower(filter))
		},
	}

	var answer string
	err = survey.AskOne(prompt, &answer, survey.WithValidator(survey.Required))
	if err != nil {
		return nil, err
	}

	for _, savedReply := range repliesAndCount.SavedReplies {
		if savedReply.ID == answer {
			return &savedReply, nil
		}
	}
	return nil, fmt.Errorf("Failed to select a saved reply")
}

func SmartSelectSavedReply(id string) (savedReply *savedreply.SavedReply, err error) {
	if id == "" {
		savedReply, err = SelectSavedReply()
	} else {
		savedReply, err = savedreply.GetSavedReply(id)
	}
	return
}

func Confirm(message string, defaultAnsewer bool) (answer bool, err error ) {
	p := &survey.Confirm{
		Message: message,
		Default: defaultAnsewer,
	}
	err = survey.AskOne(p, &answer)
	return
}
