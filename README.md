# A basic website with Google login

## UI

The UI requires [npm](https://www.npmjs.com/) to run.

```bash
npm install
npm start
```

Put your Google API Client ID in `ui/.env`.

```
REACT_APP_GOOGLE_CLIENT_ID="<CLIENT_ID>"
```

## Server
The server requires [Redis](https://redis.io/docs/install/install-redis/), [PostgreSQL](https://www.postgresql.org/) and [Go](https://go.dev/) to run.

### Run instruction
- Start Redis with default port (6379)
- Start PostgreSQL server with default port (5432)
- Build server and run
    ```bash
    cd server
    make all
    ./builds/server
    ```

#### First time usage
- Run `.builds/util` to instantiate the SQL DB.


## APIs

APIs between the server and UI can be found in the [here](./postman_collection.json) file.


