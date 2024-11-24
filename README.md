# First Go Backend

## Description

I'm started to learn go recently and thus I tried to code a backend a very simple backend using go.

I used `PostgreSQL` as my database whithin a [docker-compose](./docker-compose.yml) configuration.

I used [gin](https://gin-gonic.com/) as the backend framework.

## API

### Auth

- POST `/auth/register`
    - To Create a user 
    - req payload 
    ```json
    {
        "username": string,
        "email": string + unique,
        "password": string
    }
    ```
- POST `/auth/login`
    - Login the user
    - req payload
    ```json
    {
        "email": string + unique,
        "password": string
    }
    ```
    - response payload
    ```json
    {
        "access_token": string,
        "refresh_token": string
    }
    ``` 
- POST `/auth/refresh`
    - refresh the access token
    - Access token must be existing in `Authorization: Bearer <>`
    - req payload 
    ```json
    {
        "refresh_token": string
    }
    ```
    - response payload
    ```json
    {
        "access_token": string,
        "refresh_token": string
    }

## User

- GET `/user/`
    - Access token must be existing in `Authorization: Bearer <>`
    - Get user's info
    ```json
    {
        "id": int,
        "username": string,
        "email": string,
        "created_at":,
        "deleted_at":
    }
    ```
- PUT `/user/`
    - Access token must be existing in `Authorization: Bearer <>`
    - Request payload
    ```json
    {
        "username": *can be set can be not,
        "email": *can be set can be not
    }
    at least one field must be set
    ```
    - Response payload
    ```json
    {
        "id": int,
        "username": string,
        "email": string,
        "created_at":,
        "deleted_at":
    }
    ```
- DELETE `/user/`
    - Access token must be existing in `Authorization: Bearer <>`
    - Delete user account

## ToDo

- POST `/todos/`
    - Access token must be existing in `Authorization: Bearer <>`
    - Create ToDo
    ```json
    {
        "title": string,
        "description": string
    }
    ```
- PUT `/todos/`
    - Access token must be existing in `Authorization: Bearer <>`
    - Update ToDo
    ```json
    {
        "title": *string,
        "description": *string,
        "status": *{"completed"|"in_progress"}
    }
    ```
- DELETE `/todos/:id/trash`
    - Access token must be existing in `Authorization: Bearer <>`
    - Soft Delete
- DELETE `/todos/:id/permanent`
    - Access token must be existing in `Authorization: Bearer <>`
    - Hard Delete
- GET `/todos/`
    - Access token must be existing in `Authorization: Bearer <>`
    - Get all active todos
- GET `/todos/all`
    - Access token must be existing in `Authorization: Bearer <>`
    - Get all todos { trash + active }
- GET `/todos/trash`
    - Access token must be existing in `Authorization: Bearer <>`
    - Get users trash todos
- GET `/todos/:id`
    - Access token must be existing in `Authorization: Bearer <>`
    - Get todo by id { trash or active}