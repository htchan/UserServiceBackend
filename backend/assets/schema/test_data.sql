insert into services (uuid, name, url)
values
("1", "user service", "http://url");

insert into user_tokens (user_uuid, service_uuid, token, created_date, duration)
values
("1", "1", "token", '2000-01-01 00:00:00.000000000+00:00', 60000000000);

insert into service_tokens (service_uuid, token)
values
("1", "token");