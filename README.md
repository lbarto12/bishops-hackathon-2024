# Implementation of a voting software with a focus on security






## Install Dependencies:

### Front-End
```
cd frontend && bun install
```

### Voter API
```
cd api && go mod tidy
```

### Formgen
```
cd formgen && go mod tidy
```

### In Editor 
```sql
create table voter
(
    id  serial
        primary key,
    candidate_1      varchar,
    candidate_2      varchar,
    candidate_3      varchar,
    has_voted        boolean
);

create table voter_reg
(
    candidate_1      uuid,
    candidate_2      uuid,
    candidate_3       uuid,
    can_verify_1 varchar,
    can_verify_2 varchar,
    can_verify_3 varchar,
    health_card varchar,
    name        varchar,
    id          serial
        primary key
);

create table polls
(
    candidate integer
        primary key,
    votes int
);

```

.env in api example:
```.env
FRONTEND_URL="http://localhost:5173"
TABULATION_URL="http://localhost:5174"
API_HOST=localhost
API_PORT=5005

POSTGRES_HOST=localhost
POSTGRES_PORT=5050
POSTGRES_USER=voter_user
POSTGRES_PASSWORD=voter_password
POSTGRES_DBNAME=voter_db_name

CANDIDATE_SALT='<salt>'
```

.env in tabulation AND frontend example:
```.env
PUBLIC_API_HOST="http://localhost:5005/api"
```

.env in formgen example:
```.env
FRONTEND_URL="http://localhost:5173"

POSTGRES_HOST=localhost
POSTGRES_PORT=5050
POSTGRES_USER=voter_user
POSTGRES_PASSWORD=voter_password
POSTGRES_DBNAME=voter_db_name

CANDIDATE_SALT='<salt>'
```

## Build The App

Launch Postgress with Docker
```
cd <project root> && docker compose up
```

Launch API
```
cd api && go run .
```
Run Formgen:
```
cd formgen && go run .
```

Launch Frontend
```
cd frontend && bun run dev --host
```

Launch Tabulation
```
cd tabulation && bun run dev --host
```
