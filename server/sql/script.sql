create table department
(
    id                varchar(36)                         not null
        primary key,
    department_name   varchar(36)                         not null,
    parent_department varchar(36)                         null,
    create_time       timestamp default (now())           null,
    update_time       timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    delete_time       timestamp                           null,
    department_order  int                                 null,
    constraint department_pk_2
        unique (id),
    constraint department_pk_3
        unique (department_name),
    constraint department_department_id_fk
        foreign key (parent_department) references department (id)
            on delete cascade
);

create table pages
(
    page_id        varchar(36)                         not null
        primary key,
    page_name      varchar(255)                        not null,
    page_path      varchar(255)                        not null,
    page_icon      varchar(255)                        null,
    page_component varchar(255) null,
    parent_page    varchar(36)                         null,
    page_order     int          not null,
    create_time    timestamp default CURRENT_TIMESTAMP null,
    update_time    timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    delete_time    timestamp                           null on update CURRENT_TIMESTAMP,
    can_edit       int       default 1                 null,
    constraint pages_pk_2
        unique (page_id),
    constraint pages_parent_page_fk
        foreign key (parent_page) references pages (page_id)
            on delete cascade
);

create table roles
(
    role_id     varchar(36)                         not null
        primary key,
    role_name   varchar(255)                        not null,
    description text                                null,
    create_time timestamp default CURRENT_TIMESTAMP null,
    update_time timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    delete_time timestamp                           null,
    constraint roles_pk
        unique (role_name),
    constraint roles_pk_2
        unique (role_name)
);

create table roles_pages
(
    role_id varchar(36) null,
    page_id varchar(36) null,
    constraint roles_pages_pages_page_id_fk
        foreign key (page_id) references pages (page_id)
            on delete cascade,
    constraint roles_pages_roles_role_id_fk
        foreign key (role_id) references roles (role_id)
            on delete cascade
);

create table users
(
    id            varchar(36)               not null
        primary key,
    account       varchar(255)              not null,
    password      varchar(255)              not null,
    nickname      varchar(255)              not null,
    avatar        varchar(255)              null,
    create_time   timestamp default (now()) null on update CURRENT_TIMESTAMP,
    update_time   timestamp default (now()) null on update CURRENT_TIMESTAMP,
    delete_time   timestamp                 null,
    status        int       default 1       null,
    department_id varchar(36)               null,
    constraint account
        unique (account),
    constraint users_department_id_fk
        foreign key (department_id) references department (id)
            on delete set null
);

create table users_roles
(
    user_id varchar(36) not null,
    role_id varchar(36) not null,
    primary key (user_id, role_id),
    constraint users_roles_ibfk_1
        foreign key (user_id) references users (id),
    constraint users_roles_ibfk_2
        foreign key (role_id) references roles (role_id)
);

create index role_id
    on users_roles (role_id);


