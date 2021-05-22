function queryBuyers(offset, first){
    return JSON.stringify({
        query:`query data{
            queryBuyers(offset:` + offset + `, first: `+ first +`){ 
                id
                name
                age
            }
        }`,
        variables: null
    })
}

function queryTransactions(){
    return JSON.stringify({
        query:`
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
        }`,
        variables: null
    })
}

function queryOneBuyer(id){
    return JSON.stringify({
        query: `
        query data{
            queryBuyers(filter: { id: {eq: "` + id + `"}}){
                id
                name
                age
                transaction{
                    ip
                    productids{
                        name
                        price
                    }
                }
            }
        }
        `
    })
}

function queryIp(ip){
    return JSON.stringify({
        query:`
        query data{
            queryTransactions(filter: { ip : {eq: "` + ip + `"}}) {
                id
                ip
                buyerid{
                    name
                    age
                }
            }
        }`,
        variables: null
    })
}

function queryProducts(){
    return JSON.stringify({
        query:`
        query data{
            queryProducts{
                id
                name
                price
            }
        }`,
        variables: null
    })
}

function querySchema() {
    return JSON.stringify({
        query: `query { getGQLSchema { schema }}`,
        variables: null
    });
}

module.exports = { queryBuyers, queryIp, queryProducts , queryOneBuyer};