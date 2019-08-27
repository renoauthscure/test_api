# How to run

change paramter in const variable :

```

// how many client request
const client = 24

// url api or end point
const urlTarget = "http://localhost:8000/api/v1/vault/add/voucher"

// json payload file path
const pathPayload = "payload/payload.json"

```

then run by type :

```

// will build bynary
// and run it
go build && ./test_api 


```