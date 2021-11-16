create table users (
    uuid varchar(64) primary key,
    username varchar(256) unique,
    password text
);
create unique index users__uuid on users(uuid);

create table services (
    uuid varchar(64),
    name varchar(256)
);

create unique index services__name on services(name);

create table user_tokens (
    username varchar(256) unique,
    token text unique,
    created_date int,
    duration int
);

create unique index user_tokens__username on user_tokens(username);
create unique index user_tokens__token on user_tokens(token);

create table service_tokens (
    service_name varchar(256),
    token text unique
);

create unique index service_tokens__service_name on service_tokens(service_name);
create unique index service_tokens__token on service_tokens(token);

create table user_permissions (
    username varchar(256),
    permission text
);

create table service_permissions (
    service_name varchar(256),
    permission text
);

create unique index service_permissions__permission on service_permissions(permission);