fetch('http://localhost:8080/graphql', {
  method: 'POST',
  headers: {"Content-Type":"application/json" },
  body: JSON.stringify({
  "query":`
    query data{
      queryTransactions{
        buyerid{
          name
          age
        }
        ip
        productids{
          name
          price
        }
      }
    }
  `
  })
}).then(res => res.json())
  .then(data => {
    console.log(data)
  })

