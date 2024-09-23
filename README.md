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
create table voters
(
    id          serial,
    health_card varchar(256) not null,
    name        varchar(300) not null,
    has_voted   boolean      not null
);

alter table voters
    owner to voter_user;

create table polls
(
    candidate varchar(300) not null,
    votes     bigint       not null
);

alter table polls
    owner to voter_user;
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
