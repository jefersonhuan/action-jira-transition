# Jira - Issue Transition action

This action will move a Jira issue to a given transition.

## Example usage

```yaml
uses: jefersonhuan/action-jira-transition@v1
with:
  issueKey: ${{ github.ref }} # such as ...ref/bug/WWW-1987
  transition: In Progress
  jiraApiKey: ${{ secrets.JIRA_API_KEY }}
  jiraBaseUrl: https://something.atlassian.net
  jiraUserEmail: youremail@something.com
```         
