# digitel-test

Base Url = https://crimson-meadow-9742.fly.dev

Pre - test for golang developer @ digitels

Import API & Environment with Postman

## Auth API

**Endpoint : /api/auth**  
**Method : POST**  
**Description : Authenticate user with given email and password. returns access token & refesh token**    
**Expected JSON :** 
```json
{
    "email" : "youremail@gmail.com",
    "password" : "yourpassword"
}
```

**Endpoint : /api/auth**  
**Method : POST**  
**Description : Generate a new access token & refresh token using refresh token. returns access token & refesh token**    
**Expected JSON :** 
```json
{
    "refresh_token" : "refreshtokenhere"
}
```