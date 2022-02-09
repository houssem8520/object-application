# Objects Application
Rest API to store and retrieve objects from a bucket. The object represents the text of a http request.

## Swagger Specification
### Security
Given the following security definitions (in `swagger.yml` specification document):
the endpoint starts with `/secure/` is secured with access token.

The user logs in on its Google account, which returns an access token that we will use with our API. This mechanism follows the 'accessCode' OAuth2 workflow.

```
securityDefinitions:
  OauthSecurity:
    type: oauth2
    flow: accessCode
    authorizationUrl: 'https://accounts.google.com/o/oauth2/v2/auth'
    tokenUrl: 'https://www.googleapis.com/oauth2/v4/token'
    scopes:
      admin: Admin scope
      user: User scope
```

---
**NOTE**
The oauth2 client is configured only to accept this callback `http://127.0.0.1:8080/auth/callback`. So if we change the port, the secured endpoints (start with `/secure`)will not work. 

---

### Objects API
To PUT, GET or DELETE OBJECTS use these endpoints :
```
paths:

  /objects/{bucket}/{objectID}:
    get:
      summary: Get object by bucket name and object id.
      operationId: 'getObject'
     .
     .
     .
    delete:
      summary: delete an object.
      operationId: 'deleteObject'
      tags:
        - Object Management
     ..
     .
     .
     .  
  /objects/{bucket}:
     put:
      summary: Add an object in a bucket.
      operationId: 'putObject'
      security: []
      tags:
        - Object Management
      .
      .
      .

  /secure/objects/{bucket}/{objectID}:
    get:
      summary: Get object by bucket name and object id.
      operationId: 'getSecureObject'
      tags:
        - Secure Object Management

.
.
.
          

    delete:
      summary: delete an object.
      operationId: 'deleteSecureObject'
      tags:
        - Secure Object Management
      parameters:
   .
   .
   .    
  /secure/objects/{bucket}:
     put:
      summary: Add an object in a bucket.
      operationId: 'putSecureObject'
      tags:
        - Secure Object Management
      parameters:
        - name: 'bucket'
      .
      .
      .
```
## Makefile

To build and run the project use the different targets in the `Makefile`

### BUILD
```
make build.vendor.full make build.local
```
to build the docker image

```
make build.docker
```
### RUN
the default configuration is in the `resources/.object-application.yml`
```
make run.local
```
To run the container
```
make run.docker
```

#### Login to get the access token

Get the access token through Google's oauth2 server.

Open the browser and access the API login url on: http://127.0.0.1:8080/login, which will direct you to the Google login page.

Once you login with your google ID (e.g., your gmail account), the oauth2
`access_token` is returned and displayed on the browser.

#### Exercise your authorizer

`TOKEN` is obtained from the previous step.

Now we may use this token to access to the other endpoints published by our API.

Let's try this with curl. Copy the received token and reuse it as shown below:

```
curl -X 'PUT'   'http://127.0.0.1:8080/secure/objects/BUCKET_NAME'   -H 'accept: application/json'   -H 'Content-Type: text/plain' -H 'Authorization:Bear TOKEN'  -d 'OBJECT_CONTENT'
```
---
**NOTE**
For mor infomations, Swagger docs is available in his adress 127.0.0.1:8080/docs

---
