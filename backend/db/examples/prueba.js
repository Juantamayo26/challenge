fetch('http://localhost:8080/graphql', {
  method: 'POST',
  headers: {"Content-Type":"application/json" },
  body: JSON.stringify({
  "query":`
    query {
      queryUser() {
        id
        username
        name
        avatar_url {
          url_address
        }
      }
    }
    `
  })
}).then(res => res.json())
  .then(data => {
    console.log(data)
  })

