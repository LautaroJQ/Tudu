# API ToDo para practicar GO

## Cual es el objetivo?

Tener una API REST funcional, que me permita crear notas, modificar su contenido, eliminarlas, verificar expiracion.

## Metas

- Emplear el modulo net/http para los endpoints
- Establecer arquitectura limpia (Routers,Models,Controllers)
- Implementar todo con JSON
- Implementar DTOs para la transferencia de datos
- Implementar Bases de datos
- Establecer los endpoints necesarios
- Definir usuarios y establecer permisos mediante middleware

## Arquitectura deseada

todo-api/
├── cmd/
│   └── api/
│       └── main.go           
├── internal/
│   ├── models/ -> Representacion del objeto
│   │   └── task.go
│   ├── dtos/ -> Representacion de las distintas transferencia de datos que hay
│   │   └── task_dto.go
│   ├── database/
│   │   └── task_store.go      
│   ├── handlers/ -> Logica de negocio
│   │   └── task_controller.go 
│   └── routers/
│       └── task_router.go -> Elije a donde va la request
├── go.mod
└── README.md

## Endpoints

- POST /tasks -> Crear tareas
- GET /tasks -> Listar tareas (Añadir query params para filtrar)
- PATCH /tasks/{id} -> Modificar una tarea parcialmente 
- 

# Anotaciones

- Un model o un DTO no deben tener logicas de negocio
- el 