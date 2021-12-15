# COURSE API REST
Ejemplo de proyecto funcional deployado en heroku.
Podemos acceder a traves de la siguiente URL:  
```
https://courses-api-rest.herokuapp.com
```

# INTEGRANTES
1. Matias Millain (matimillain@gmail.com)
2. -
3. -
4. -
5. -
6. -


# DOCOMENTACIÓN

## REGISTER
Registro de nuevo usuario

**POST**
```
https://courses-api-rest.herokuapp.com/api/auth/register
```

**REQUEST BODY**
```
{
    "name": "Test name user",
    "email": "test-user-email@gmail.com",
    "password": "superSecretPassword"
}
```
 
**RESPONSE SUCCESS (Status: 201)**
```
{
    "status": true,
    "message": "OK!",
    "errors": null,
    "data": {
        "id": 2,
        "name": "Test user name",
        "email": "test-user-email@gmail.com",
        "token": "eyJhbGciOiJIUzI1NiI1231dapXVCJ9.eyJ1c2VyX2lkIjoiMiIsImV4cCI6MTY1MTgyMDAwMCwiaWF0IjoxNjIwMjg0MDAwLCJpc3MiOiJhZG1pbiJ9.HtnuWlBaevEO3fHAI4McH5W8axvw_3Og47RUI3m9IyI"
    }
}
```
 
**RESPONSE ERROR (Status : 400)**
```
{
    "status": false,
    "message": "Failed to process request",
    "errors": [
        "Key: 'RegisterRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag"
    ],
    "data": {}
}
```
 
## LOGIN
Autenticamos con email y password

**POST**
```
https://courses-api-rest.herokuapp.com/api/auth/login
```
 
**REQUEST JSON BODY**
```
{
    "email": "test-user-email@gmail.com",
    "password": "superSecretPassword"
}
```
 
**RESPONSE (Status: 200)**
```
{
    "status": true,
    "message": "OK!",
    "errors": null,
    "data": {
        "id": 1,
        "name": "Test user name",
        "email": "test-user-email@gmail.com",
        "token": "eyJhbGciOiJIUzI123123sInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMSIsImV4cCI6MTY1MTgyMTQ0NiwiaWF0IjoxNjIwMjg1NDQ2LCJpc3MiOiJhZG1pbiJ9.2m-r1qrCXhNkAxzK-sb4hL0Pzat3zwOmzktES_uzwts"
    }
}
```
 
**RESPONSE ERROR, CREDENCIALES ERRÓNEAS (Status: 401)**
```
{
    "status": false,
    "message": "Failed to login",
    "errors": [
        "failed to login. check your credential"
    ],
    "data": {}
}
```
 
## GET PROFILE
Obtenemos el perfil de usuario logueado

**GET**
```
https://courses-api-rest.herokuapp.com/api/user/profile
```
 
**HEADER**
```
Authorization: authTOKEN
```
 
**RESPONSE**
```
{
    "status": true,
    "message": "OK",
    "errors": null,
    "data": {
     	“id": 1,
   	"name": "Test user name",
    	"email": "test-user-email@gmail.com"
    }
}
```
 
## UPDATE PROFILE
Actualizar datos del usuario pasado por auth token

**PUT**
```
https://courses-api-rest.herokuapp.com/api/user/profile
```
 
**HEADER**
```
Authorization: authTOKEN
```
 
**REQUEST BODY**
```
{
    "name": "Test new user name",
    "email": "test-new-user-email@gmail.com"
}
```
 
**RESPONSE**
```
{
    "status": true,
    "message": "OK",
    "errors": null,
    "data": {
     	“id": 1,
   	"name": "Test new user name",
    	"email": "test-new-user-email@gmail.com"
    }
}
```
 
=============================================
## ALL COURSES
Solo se visualizan los del usuario logueado

**GET**
```
https://courses-api-rest.herokuapp.com/api/course
```
 
**HEADER**
```
Authorization: authTOKEN
```
 
**RESPONSE**
```
{
    "status": true,
    "message": "OK!",
    "errors": null,
    "data": [
        {
            "id": 2,
            "name": "Course name 2",
            "description": “Course description 2”,
            "user": {
                "id": 1,
                "name": "Test user name",
                "email": "test-user-email@gmail.com"
            }
        },
        {
            "id": 3,
            "name": "Course name 3",
            "description": “Course description 3”,
            "user": {
                "id": 1,
                "name": "Test user name",
                "email": "test-user-email@gmail.com"
            }
        }
    ]
}
```
 
## CREATE COURSE
Creamos un curso y lo asignamos al usuario conectado a través del auth token

**POST**
```
https://courses-api-rest.herokuapp.com/api/course
```
 
**HEADER**
```
Authorization: authTOKEN
```
 
**REQUEST JSON BODY**
```
{
    "name": "New course name”,
    "description": “New course description”
}
```
 
**RESPONSE**
```
{
    "status": true,
    "message": "OK!",
    "errors": null,
    "data": {
        "id": 1,
        "name": "New course name",
        "description": “New course description”,
        "user": {
            "id": 1,
            "name": "Test user name",
            "email": "test-user-email@gmail.com"
        }
    }
}
```
 
## FIND ONE COURSE BY ID
Solo podemos obtener nuestros cursos creados

**GET**
```
https://courses-api-rest.herokuapp.com/api/course/{courseID}
```
 
**HEADER**
```
Authorization: authTOKEN
```
 
**RESPONSE**
```
{
    "status": true,
    "message": "OK!",
    "errors": null,
    "data": {
        "id": 1,
        "name": "Test course name",
        "description": “Test course description”,
        "user": {
            "id": 1,
            "name": "Test user name",
            "email": "test-user-email@gmail.com"
        }
    }
}
```
 
## UPDATE COURSE
 
Solo podemos actualizar nuestros cursos, si tratamos de ingresar otro id, obtenemos un error.

**PUT**
```
https://golang-heroku.herokuapp.com/api/course/{courseID}
```

**HEADERS**
```
Authorization: authTOKEN
```
 
**REQUEST JSON BODY**
```
{
    "name": "Test modify course name",
    "description": “Test modify course description”
}
```
 
**RESPONSE**
```
{
    "status": true,
    "message": "OK!",
    "errors": null,
    "data": {
        "id": 1,
        "product_name": "Test modify course name",
        "description": “Test modify course description”,
        "user": {
            "id": 1,
            "name": "Test user name",
            "email": "test-user-email@gmail.com"
        }
    }
}
```
 
## DELETE COURSE

Solo podemos eliminar nuestros cursos creados

**DELETE**
```
https://courses-api-rest.herokuapp.com/api/product/{courseID}
```
 
**Response success (Status: 200)**
```
{
    "status": true,
    "message": "OK!",
    "errors": null,
    "data": {}
}
```
