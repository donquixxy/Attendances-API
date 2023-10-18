# Attendances (Backend) APP

Base Url = https://crimson-meadow-9742.fly.dev

Import API & Environment with Postman

## How to run locally

1. Start docker engine
2. Open terminal
3. docker compose up -d


## Features
- [x] Authentication
- [x] CRUD Employee (User)
- [x] CRUD Attendances
- [x] Cron job for send notification.

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

## Attendance API

**Endpoint : /api/attendance**   
**Method : POST**  
**Header : Authorization Bearer Token**    
**Description : Store a new attendance with given ID User (by input). Type 1 for checkin, 2 for checkout.returns created attendance**   
**Expected JSON :**  
```json
{
   "id_user" : "iduserhere",
   "type" : 1
}
```

**Endpoint : /api/absent**   
**Method : POST**  
**Header : Authorization Bearer Token**    
**Description : Store a new attendance based on Authorization token given. Type 1 for checkin, 2 for checkout.returns created attendance**   
**Expected JSON :**  
```json
{
   "type" : 2
}
```

**Endpoint : /api/attendance/:id**   
**Method : PUT**  
**Header : Authorization Bearer Token**    
**Description : Update attendance based on given id at param. Type 1 for checkin, 2 for checkout.returns updated attendance**   
**Expected JSON :**  
```json
{
   "type" : 1
}
```

**Endpoint : /api/attendance/:id**   
**Method : GET**  
**Header : Authorization Bearer Token**    
**Description : Get attendance by id based on ID given at param. return single object attendance if found, otherwise 404**

**Endpoint : /api/attendances**   
**Method : GET**  
**Header : Authorization Bearer Token**    
**Description : returns all attendances**

**Endpoint : /api/attendances/:id**   
**Method : DELETE**  
**Header : Authorization Bearer Token**    
**Description : delete attendances based on id given at param. return success msg**
