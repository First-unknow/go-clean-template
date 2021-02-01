# GO-CLEAN-TEMPLATE
## Install dependencies
```
$ go mod download 
```


## Develop environment at localhost:7000
```
$ CompileDaemon -build="go build innovasive/go-clean-template" -command="./go-clean-template" 
```

# API

## Fetch All User

` https://${url}/users`

<br />

> ### Method

  GET

> ### Success Response
<br />

 * **Code:** 200 OK <br />
  * **Content:** 
```
  users : [
    {
      "id": UUID,
      "email": "",
      "first_name": "" 
      "last_name": "" 
    },
    ... 
  ]
```
<br />

> ### Error Response
  * **Code:** 500 Internal Server Error 

<br />
<br />

## Create User

` https://${url}/users`

> ### Request Body

```
  {
    "email": "",
    "first_name": "" 
    "last_name": "" 
  }
```
<br />

> ### Method

  POST

> ### Success Response
<br />

  * **Code:** 201 CREATED <br />
  * **Content:** 
```
  {
    "userID": UUID
  }
```
<br />

> ### Error Response
  * **Code:** 406 Not Acceptable 



