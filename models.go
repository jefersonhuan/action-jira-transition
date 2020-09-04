package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"reflect"
	"regexp"
)

type Params struct {
	IssueKey  string `env:"ISSUE_KEY"`
	NewStatus string `env:"TRANSITION"`
	ApiKey    string `env:"JIRA_API_KEY"`
	Domain    string `env:"JIRA_BASE_URL"`
	UserEmail string `env:"JIRA_USER_EMAIL"`
	auth      string
}

type Transition struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func fillTag(s reflect.Type, env string) (error, int) {
	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)

		if field.Tag.Get("env") != env {
			continue
		}

		return nil, i
	}

	return fmt.Errorf("couldn't fill tag for %s", env), 0
}

func (params *Params) Load() (err error) {
	v := reflect.ValueOf(params).Elem()
	s := reflect.TypeOf(*params)

	envs := []string{
		"ISSUE_KEY",
		"TRANSITION",
		"JIRA_API_KEY",
		"JIRA_BASE_URL",
		"JIRA_USER_EMAIL",
	}

	for _, e := range envs {
		value := os.Getenv(e)
		if value == "" {
			return fmt.Errorf("variable %s is required", e)
		}

		err, index := fillTag(s, e)
		if err != nil {
			return err
		}

		v.Field(index).SetString(value)
	}

	auth := fmt.Sprintf("%s:%s", params.UserEmail, params.ApiKey)
	params.auth = base64.StdEncoding.EncodeToString([]byte(auth))

	re := regexp.MustCompile(`(\w{3}-\d+)`)
	issueKey := re.Find([]byte(params.IssueKey))

	params.IssueKey = string(issueKey)

	return
}
