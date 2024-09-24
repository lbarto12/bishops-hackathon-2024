# Implementation of a voting software with a focus on security






## Install Dependencies:

### Front-End
```
cd fronted && bun install
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
    id               serial
            primary key,
    health_card_hash varchar,
    "name_hash"      varchar,
    candidate_1      uuid,
    candidate_2      uuid,
    canidate_3       uuid,
    has_voted        boolean
);

create table voter_reg
(
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
