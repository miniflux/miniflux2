// Code generated by go generate; DO NOT EDIT.
// 2017-12-10 18:56:24.36359961 -0800 PST m=+0.005578697

package sql

var SqlMap = map[string]string{
	"schema_version_1": `create table schema_version (
    version text not null
);

create table users (
    id serial not null,
    username text not null unique,
    password text,
    is_admin bool default 'f',
    language text default 'en_US',
    timezone text default 'UTC',
    theme text default 'default',
    last_login_at timestamp with time zone,
    primary key (id)
);

create table sessions (
    id serial not null,
    user_id int not null,
    token text not null unique,
    created_at timestamp with time zone default now(),
    user_agent text,
    ip text,
    primary key (id),
    unique (user_id, token),
    foreign key (user_id) references users(id) on delete cascade
);

create table categories (
    id serial not null,
    user_id int not null,
    title text not null,
    primary key (id),
    unique (user_id, title),
    foreign key (user_id) references users(id) on delete cascade
);

create table feeds (
    id bigserial not null,
    user_id int not null,
    category_id int not null,
    title text not null,
    feed_url text not null,
    site_url text not null,
    checked_at timestamp with time zone default now(),
    etag_header text,
    last_modified_header text,
    parsing_error_msg text default '',
    parsing_error_count int default 0,
    primary key (id),
    unique (user_id, feed_url),
    foreign key (user_id) references users(id) on delete cascade,
    foreign key (category_id) references categories(id) on delete cascade
);

create type entry_status as enum('unread', 'read', 'removed');

create table entries (
    id bigserial not null,
    user_id int not null,
    feed_id bigint not null,
    hash text not null,
    published_at timestamp with time zone not null,
    title text not null,
    url text not null,
    author text,
    content text,
    status entry_status default 'unread',
    primary key (id),
    unique (feed_id, hash),
    foreign key (user_id) references users(id) on delete cascade,
    foreign key (feed_id) references feeds(id) on delete cascade
);

create index entries_feed_idx on entries using btree(feed_id);

create table enclosures (
    id bigserial not null,
    user_id int not null,
    entry_id bigint not null,
    url text not null,
    size int default 0,
    mime_type text default '',
    primary key (id),
    foreign key (user_id) references users(id) on delete cascade,
    foreign key (entry_id) references entries(id) on delete cascade
);

create table icons (
    id bigserial not null,
    hash text not null unique,
    mime_type text not null,
    content bytea not null,
    primary key (id)
);

create table feed_icons (
    feed_id bigint not null,
    icon_id bigint not null,
    primary key(feed_id, icon_id),
    foreign key (feed_id) references feeds(id) on delete cascade,
    foreign key (icon_id) references icons(id) on delete cascade
);
`,
	"schema_version_2": `create extension if not exists hstore;
alter table users add column extra hstore;
create index users_extra_idx on users using gin(extra);
`,
	"schema_version_3": `create table tokens (
    id text not null,
    value text not null,
    created_at timestamp with time zone not null default now(),
    primary key(id, value)
);`,
	"schema_version_4": `create type entry_sorting_direction as enum('asc', 'desc');
alter table users add column entry_direction entry_sorting_direction default 'asc';
`,
	"schema_version_5": `create table integrations (
    user_id int not null,
    pinboard_enabled bool default 'f',
    pinboard_token text default '',
    pinboard_tags text default 'miniflux',
    pinboard_mark_as_unread bool default 'f',
    instapaper_enabled bool default 'f',
    instapaper_username text default '',
    instapaper_password text default '',
    fever_enabled bool default 'f',
    fever_username text default '',
    fever_password text default '',
    fever_token text default '',
    primary key(user_id)
)
`,
}

var SqlMapChecksums = map[string]string{
	"schema_version_1": "7be580fc8a93db5da54b2f6e87019803c33b0b0c28482c7af80cef873bdac4e2",
	"schema_version_2": "e8e9ff32478df04fcddad10a34cba2e8bb1e67e7977b5bd6cdc4c31ec94282b4",
	"schema_version_3": "a54745dbc1c51c000f74d4e5068f1e2f43e83309f023415b1749a47d5c1e0f12",
	"schema_version_4": "216ea3a7d3e1704e40c797b5dc47456517c27dbb6ca98bf88812f4f63d74b5d9",
	"schema_version_5": "46397e2f5f2c82116786127e9f6a403e975b14d2ca7b652a48cd1ba843e6a27c",
}
