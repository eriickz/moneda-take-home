

# moneda-take-home

Moneda take home for candidates

## Project Requirements:

Environment Setup:

- Set up a go environment: https://go.dev/doc/install.
- Provide the rights enviroment variables as stated in .env.example

API Documentation:

- All endpoints are located under "/api" path. 
- Flights endpoints are secured by a token, must be provided in an "Authentication" header (not Bearer). 
- You can generate an auth token by making a request to "/auth/sign-in" with "admin" as a username (required) and any password.
- You can import ```api-docs.json``` to your Postman or any other API testing tool. 

### Security:

All flights endpoints are secured by a JWT token, your must provide a secret in .env to make it work.

### Code Organization:

- Global Folder
All code relating to configurations that make the app works and startups scripts are here. I used to have a ```config.go``` file to initalize AWS services or a ```models.go``` to quick grab a global type that are share among modules if I need to.

- Modules Folder
I usually divide the apps into modules because it provide more organization, reusability and readability, and all code relating to these modules are in their proper folder.

>[module].controller.go  - Handlers that are triggered when an endpoint is reached.
>
>[module].routes.go      - Module's routes paths definition.
>
>[module].models.go      - Module's types (request/response and other business logic structs)
>
>[module].lib.go         - More used functions in this module, usually parsers or custom validations handlers
>
>[module].utils.go       - Code that serve as configurations in this module or for the functions located in .lib.go
>
>/test                   - Well, tests... In files like [module].controller.test.go
