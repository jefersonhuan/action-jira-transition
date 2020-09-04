package jira

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func doRequest(method, path string, target interface{}, body *[]byte) error {
	url := fmt.Sprintf("%s/%s", params.Domain, path)
	var bodyReader io.Reader

	if body != nil {
		bodyReader = bytes.NewReader(*body)
	}

	client := &http.Client{}
	request, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return err
	}

	request.Header.Add("Authorization", "Basic "+params.auth)
	request.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		return fmt.Errorf("jira API returned a status other than 200/204: %+v", resp.Status)
	}

	if target != nil {
		return decode(resp, target)
	}

	return nil
}

func decode(resp *http.Response, target interface{}) (err error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return
	}

	return
}

func fetchTransitions(issueKey string) (err error, transitions []Transition) {
	var transitionResponse struct {
		Transitions []Transition `json:"transitions"`
	}

	path := fmt.Sprintf("rest/api/2/issue/%s/transitions", issueKey)

	err = doRequest("GET", path, &transitionResponse, nil)

	transitions = transitionResponse.Transitions
	return
}

func updateIssue(transition Transition, issueKey string) (err error) {
	path := fmt.Sprintf("rest/api/2/issue/%s/transitions", issueKey)

	var req struct {
		Transition struct {
			ID string `json:"id"`
		} `json:"transition"`
	}
	req.Transition.ID = transition.ID

	body, err := json.Marshal(req)
	if err != nil {
		return
	}

	err = doRequest("POST", path, nil, &body)
	return
}
