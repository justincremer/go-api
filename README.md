# A simple rest api written in go

## Cassandra DB
### Create Volume
``` bash
docker volume create --name=cassandra-data
```

### Create DB
``` bash
docker-compose up -d
```

### Create keyspace
``` bash
CREATE KEYSPACE go_db WITH replication = {'class': 'SimpleStrategy',
'replication_factor': 1};
```

### Create user table
``` bash
CREATE TABLE IF NOT EXISTS users (
    user_id uuid,
    username text,
    name text,
    email text,
    phone text,
    PRIMARY KEY (user_id, username)
);
```

### Add a user 
``` bash
INSERT INTO users (user_id, username, name, email, phone) VALUES (now(),
'username', 'name', 'email', 'phone');
```

## Api
### Start
``` bash
go run src/main.go
```

### Build
``` bash
go build -o bin/start src/main.go
```

