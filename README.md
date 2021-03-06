[![Netlify Status](https://api.netlify.com/api/v1/badges/d686cd09-ec33-4d57-84b4-87d24fabca90/deploy-status)](https://app.netlify.com/sites/devops-dashboard/deploys)
[![Maintainability](https://api.codeclimate.com/v1/badges/7ef9e694f2ed88234f1e/maintainability)](https://codeclimate.com/github/Kalimaha/devops-dashboard/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/7ef9e694f2ed88234f1e/test_coverage)](https://codeclimate.com/github/Kalimaha/devops-dashboard/test_coverage)

# Devops Dashboard

The devops dashboard shows:

* Open PRs in relevant Github repositories
* The latest deploys, with details about PRs, of relevant projects in Heroku

## Links

* Dashboard: https://devops-dashboard.netlify.app/
* Functions
  * `github`
    * Endpoint: https://devops-dashboard.netlify.app/.netlify/functions/vm-github-pull-requests?repositoryName=vinomofo
    * Logs: https://app.netlify.com/sites/devops-dashboard/functions/vm-github-pull-requests

## Netlify setup

The following variables must be set in the Netlify environment:

* `GITHUB_TOKEN`
* `HEROKU_TOKEN`
