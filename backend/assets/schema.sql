create table users (
    uuid varchar(64) primary key,
    username varchar(256) unique,
    password text
);
create unique index users__uuid on users(uuid);
create unique index users__username on users(username);

create table services (
    uuid varchar(64) primary key,
    name varchar(256) unique,
    url text
);

create unique index services__name on services(name);
create unique index services__uuid on services(uuid);

create table user_tokens (
    user_uuid varchar(64),
    token text unique,
    created_date int,
    duration int
);

create unique index user_tokens__uuid on user_tokens(user_uuid);
create unique index user_tokens__token on user_tokens(token);

create table service_tokens (
    service_uuid varchar(64) unique,
    token text unique
);

create unique index service_tokens__service_uuid on service_tokens(service_uuid);
create unique index service_tokens__token on service_tokens(token);

create table user_permissions (
    user_uuid varchar(64),
    permission text
);

create table service_permissions (
    service_uuid varchar(64),
    permission text primary key
);

create unique index service_permissions__permission on service_permissions(permission);