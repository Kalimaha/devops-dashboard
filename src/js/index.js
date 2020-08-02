const buildDevOpsDashboard = () => {
  const REPOSITORY_NAMES = [
    "vinomofo",
    "vino-delivery",
    "vino-subscription",
    "vino-warehouse",
    "vino-delivery-x-ebay",
    "smokescreen"
  ]

  for (var i = 0; i < REPOSITORY_NAMES.length; i++) {
    fetchPullRequests(REPOSITORY_NAMES[i])
  }
}

const fetchPullRequests = (repositoryName) => {
  var url = `https://devops-dashboard.netlify.app/.netlify/functions/vm-github-pull-requests?repositoryName=${repositoryName}`
  $.ajax({
    url: url
  }).then(function(data) {
    if (data != null) {
      $("#loading").css("display", "none")
      var template = $("#pull-request-template").html()
      for (var i = 0; i < data.length; i++) {
        var values = data2template(data[i], repositoryName)
        if (values.dateOpened < 30) {
          var html = Mustache.render(template, values)
          $("#pull-requests").append(html)
        }
      }
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
  message: buildMessage(data.Reviews),
})

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
