[![Netlify Status](https://api.netlify.com/api/v1/badges/d686cd09-ec33-4d57-84b4-87d24fabca90/deploy-status)](https://app.netlify.com/sites/devops-dashboard/deploys)

# Devops Dashboard

The devops dashboard shows:

* Open PRs in relevant Github repositories
* The latest deploys, with details about PRs, of relevant projects in Heroku

## Links

* Dashboard: https://devops-dashboard.netlify.app/?repositoryName=vino-delivery
* Function `vm-github-pull-requests`
  * Endpoint: https://devops-dashboard.netlify.app/.netlify/functions/vm-github-pull-requests?repositoryName=vinomofo
  * Logs: https://app.netlify.com/sites/devops-dashboard/functions/vm-github-pull-requests
