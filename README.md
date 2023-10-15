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

## Employee API


**Endpoint : /api/user**  
**Method : POST**  
**Header : Authorization Bearer Token**  
**Description : Store a new employee. returns created user**   
**Expected JSON :**  
```json
{
    "name" : "test-user",
    "email" : "test-user@gmail.com",
    "password" : "yourpassword"
}
```

**Endpoint : /api/user/:id**  
**Method : PUT**  
**Header : Authorization Bearer Token**  
**Description : Update user by ID given at param url. Updates user data based on given json. returns updated user field**  
**Expected JSON :**    
```json
{
    "name" : "test-user (not required)", 
    "email" : "test-user@gmail.com (not required)",
    "password" : "yourpassword (not required)"
}
```

**Endpoint : /api/user/:id**  
**Method : GET**  
**Header : Authorization Bearer Token**  
**Description : Get user by ID given at param url. returns user if found, otherwise return 404**  

**Endpoint : /api/users**  
**Method : GET**  
**Header : Authorization Bearer Token**  
**Description : returns all users**

**Endpoint : /api/user/:id**  
**Method : DELETE**    
**Header : Authorization Bearer Token**    
**Description : Delete user with given ID. returns success msg**  


