CREATE TABLE cx_applications (
    id          serial primary key,
    hash        varchar not null UNIQUE,
    config      json not null,
    chain_type  varchar not null,
    created_at  timestamp not null,
    updated_at  timestamp not null,
    deleted_at  timestamp null
);

CREATE TABLE servers (
    id                  serial primary key,
    address             varchar not null,
    created_at          timestamp not null,
    updated_at          timestamp not null,
    deleted_at          timestamp null,
    cx_application_id   integer not null,
    foreign key (cx_application_id) references cx_applications (id)
);