fetch('http://localhost:8080/graphql', {
  method: 'POST',
  headers: {"Content-Type":"application/json" },
  body: JSON.stringify({
  "query":`
  query data{
    queryBuyers{
      id
      name
      age
    }
  }
  `
  })
}).then(res => res.json())
  .then(data => {
    console.log(data)
  })

