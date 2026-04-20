# TaskBoard Backend
REST API сервис для управления задачами.
позволяет создать задачи перемещать их между статусами.
## Tech Stack
- Go
- PostgreSQL
- Docker/ docker-compose
- REST API
## Run localy
- git clone https://github.com/primo0067/taskboardbackend
- cd taskboardbackend
- docker compose up --build
## Architecture
- handlers (HTTP layer)
- db (database layer)
- models
## API
### TASKS
- GET /tasks
- POST /tasks
- PUT /tasks/{id}
- DELETE /tasks/{id}
## Tests
пока не реализованы:)
