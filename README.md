# Implementation of a voting software with a focus on security






## Install Dependencies:

### Front-End
```
cd frontend && bun install
```

### Voter API
```
cd voterapi && go mod download
```


## Setup Docker


### Build
```
cd <project root> && docker compose up
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

.env in voterapi example:
```.env
FRONTEND_URL=... # for CORS
API_HOST=0.0.0.0
API_PORT=1234

POSTGRES_HOST=0.0.0.0
POSTGRES_PORT=5050
POSTGRES_USER=voter_user
POSTGRES_PASSWORD=voter_password
POSTGRES_DBNAME=voter_db_name
```

## Build The App

Launch API
```
cd voterapi && go run .
```

Launch Frontend
```
cd frontend && bun run dev --
```
