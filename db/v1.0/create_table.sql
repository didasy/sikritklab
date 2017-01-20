use sikritklab;

create table posts(
    id bigint primary key not null auto_increment, 
    created_at datetime not null,
    updated_at datetime not null, 
    deleted_at datetime,
    thread_id bigint not null,
    title varchar(128) not null,
    content text(2000) not null
);

create index created_at_idx on posts(created_at);
create index updated_at_idx on posts(updated_at);
create index deleted_at_idx on posts(deleted_at);
create index thread_id_idx on posts(thread_id);
create index title_idx on posts(title);

create table threads(
    id bigint primary key not null auto_increment, 
    created_at datetime not null,
    updated_at datetime not null, 
    deleted_at datetime,
    title varchar(128) not null,
    tags text not null,
    fulltext tags_idx (tags)
);

create index created_at_idx on threads(created_at);
create index updated_at_idx on threads(updated_at);
create index deleted_at_idx on threads(deleted_at);
create index title_idx on posts(title);
