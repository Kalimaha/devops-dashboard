const buildDevOpsDashboard = () => {
  var url = "https://geobricks.stoplight.io/mocks/geobricks/github-pulls/636798/vm-github-pull-requests/repositoryName"
  $.ajax({
    url: url
  }).then(function(data) {
    var template = $("#pull-request-template").html()
    for (var i = 0; i < data.length; i++) {
      var values = data2template(data[i])
      var html = Mustache.render(template, values)
      $("#pull-requests").append(html)
    }
  })
}

const data2template = (data) => ({ 
  pullRequestName: data.pullRequestName, 
  pullRequestURL: data.pullRequestURL, 
  repositoryName: data.repositoryName, 
  repositoryURL: data.repositoryURL, 
  dateOpened: countDays(data.dateOpened),
  message: buildMessage(data.reviews),
})

const buildMessage = (reviews) => {
  if (reviews.length === 0) {
    return "<b>Two</b> reviews required."
  } else if (reviews.length === 1) {
    return `<b>One</b> more review required. Thanks for your review <a href="${reviews[0].reviewerURL}">${reviews[0].reviewerName}</a>!`
  } else if (reviews.length === 2) {
    return `<span class="text-success"><b>Good to go!</b></span> Thanks for your reviews <a href="${reviews[0].reviewerURL}">${reviews[0].reviewerName}</a> and <a href="${reviews[1].reviewerURL}">${reviews[1].reviewerName}</a>!`
  } else {
    return `<span class="text-success"><b>Good to go!</b></span> Thanks everyone!`
  }
}

const countDays = (prRawDate) => {
  var prDate  = new Date(prRawDate)
  var today   = new Date()
  
  return parseInt((today.getTime() - prDate.getTime()) / (1000 * 3600 * 24))
}
