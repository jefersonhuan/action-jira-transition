package jira

import "fmt"

var params = Params{}

func findStatus(transitions []Transition) (error, Transition) {
	for _, t := range transitions {
		if t.Name == params.NewStatus {
			return nil, t
		}
	}

	return fmt.Errorf("couldn't find equivalent status of %s", params.NewStatus), Transition{}
}

func StartTransition() (err error) {
	if err := loadParams(); err != nil {
		panic(fmt.Errorf("An error ocurred while fetching variables: %+v\n", err))
	}

	err, transitions := fetchTransitions(params.IssueKey)
	if err != nil {
		return
	}

	fmt.Println("Found", len(transitions), "transitions")

	err, transition := findStatus(transitions)
	if err != nil {
		return
	}

	err = updateIssue(transition, params.IssueKey)
	if err != nil {
		return
	}

	fmt.Println("The issue", params.IssueKey, "has been successfully transitioned to", params.NewStatus)

	return
}
