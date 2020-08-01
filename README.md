# DevOps Dashboard

About Netlify CLI
Install CLI: npm install netlify-cli -g
Init CLI: netlify init
List functions: netlify functions:list
Invoke function
Run netlify dev in a separate tab, then
netlify functions:invoke stock-great-wines --identity, or
curl -X POST "http://localhost:8888/.netlify/functions/stock-great-wines" -d "{'spam': 'eggs'}"
