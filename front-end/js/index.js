const buildDevOpsDashboard = () => {
  const REPOSITORY_NAMES = [
    "vinomofo",
    "vino-delivery",
    "vino-delivery-x-ebay",
    "vino-styleguides",
    "vino-subscription",
    "vino-warehouse",
    "vino-storefront",
    "vino-toolkit",
    "smokescreen"
  ]

  const HEROKU_REPOSITORY_NAMES = [
    "vinomofo",
    "vino-delivery",
    "vino-warehouse",
    "vino-subscription"
  ]

  for (var i = 0; i < REPOSITORY_NAMES.length; i++) {
    fetchPullRequests(REPOSITORY_NAMES[i])
  }

  for (var i = 0; i < HEROKU_REPOSITORY_NAMES.length; i++) {
    fetchReleases(HEROKU_REPOSITORY_NAMES[i])
  }
}

const fetchReleases = (repositoryName) => {
  var url = `https://devops-dashboard.netlify.app/.netlify/functions/heroku?repositoryName=${repositoryName}`
  $.ajax({
    url: url
  }).then(function(data) {
    if (data != null) {
      $("#releases-loading").css("display", "none")
      var template = $("#heroku-release-template").html()
      var values = { repositoryName: repositoryName, pastCommits: data.PastCommits, futureCommits: data.FutureCommits }
      var html = Mustache.render(template, values)
      $("#releases").append(html)
    }
  })
}

const fetchPullRequests = (repositoryName) => {
  var url = `https://devops-dashboard.netlify.app/.netlify/functions/github?repositoryName=${repositoryName}`
  $.ajax({
    url: url
  }).then(function(data) {
    if (data != null) {
      $("#loading").css("display", "none")
      var template = $("#pull-request-template").html()
      for (var i = 0; i < data.length; i++) {
        var values = data2template(data[i], repositoryName)
        if (values.dateOpened < 45 && !values.draft) {
          var html = Mustache.render(template, values)
          $("#pull-requests").append(html)
        }
      }
      updateCount()
    }
  })
}

const data2template = (data, repositoryName) => ({ 
  pullRequestName: data.Title, 
  pullRequestURL: data.Url, 
  pullRequestId: data.Number,
  repositoryName: repositoryName,
  repositoryURL: `https://github.com/vinomofo/${repositoryName}/`,
  dateOpened: countDays(data.CreatedAt),
  dateUpdated: countDays(data.UpdatedAt),
  message: buildMessage(data.Reviews),
  authorName: data.AuthorName,
  authorURL: data.AuthorURL,
  avatarURL: data.AvatarURL,
  draft: data.Draft,
})

const updateCount = () => {
  const count = $(".github-pr").length

  $("#open-prs").html(`Open PRs: ${count}`)
}

const buildMessage = (reviews) => {
  var approvals           = reviews.filter(r => r.State == "APPROVED")
  var changesRequestsMsg  = changesRequestedMessage(reviews)
  var commentedMessage    = buildCommentedMessage(reviews)

  if (changesRequestsMsg === "") {
    if (approvals.length === 0) {
      return `<b>Two</b> approvals required. ${commentedMessage}`
    } else if (approvals.length === 1) {
      return `<b>One</b> more approval required. ${commentedMessage}`
    } else if (approvals.length === 2) {
      return `<span class="text-success"><b>Good to go</b></span>, thanks for your reviews <a href="${approvals[0].ReviewerURL}">${approvals[0].ReviewerName}</a> and <a href="${approvals[1].ReviewerURL}">${approvals[1].ReviewerName}</a>! ${commentedMessage}`
    } else {
      return `<span class="text-success"><b>Good to go</b></span>, thanks everyone!`
    }
  } else {
    return changesRequestsMsg
  }
}

const buildCommentedMessage = (reviews) => {
  var cr = reviews.filter(r => r.State == "COMMENTED")
  if (cr.length > 0) {
    var msgs = cr.map(x => `<a href="${x.ReviewerURL}">${x.ReviewerName}</a> left comments`)
    var b = new Set(msgs)
    return `${Array.from(b).join(", ")}.`
  }
  return ""
}

const changesRequestedMessage = (reviews) => {
  var cr = reviews.find(r => r.State == "CHANGES_REQUESTED")
  if (cr != undefined) {
    return `<a href="${cr.ReviewerURL}">${cr.ReviewerName}</a> requested <span class="text-danger">changes<span>.`
  }
  return ""
}

const countDays = (prRawDate) => {
  var prDate  = new Date(prRawDate)
  var today   = new Date()
  
  return parseInt((today.getTime() - prDate.getTime()) / (1000 * 3600 * 24))
}
