const buildDevOpsDashboard = () => {
  const REPOSITORY_NAMES = ["vinomofo", "vino-delivery", "vino-subscription", "vino-warehouse"]
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
  if (reviews.length === 0) {
    return "<b>Two</b> reviews required."
  } else if (reviews.length === 1) {
    return `<b>One</b> more review required. Thanks for your review <a href="${reviews[0].ReviewerURL}">${reviews[0].ReviewerName}</a>!`
  } else if (reviews.length === 2) {
    return `<span class="text-success"><b>Good to go!</b></span> Thanks for your reviews <a href="${reviews[0].ReviewerURL}">${reviews[0].ReviewerName}</a> and <a href="${reviews[1].ReviewerURL}">${reviews[1].ReviewerName}</a>!`
  } else {
    return `<span class="text-success"><b>Good to go!</b></span> Thanks everyone!`
  }
}

const countDays = (prRawDate) => {
  var prDate  = new Date(prRawDate)
  var today   = new Date()
  
  return parseInt((today.getTime() - prDate.getTime()) / (1000 * 3600 * 24))
}
