CREATE TABLE articles
(
    id integer PRIMARY KEY,
    user_id uuid,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE articles_tags
(
    id integer PRIMARY KEY,
    article_id integer,
    tag_id integer
);

CREATE TABLE tags
(
    id integer PRIMARY KEY,
    name char(10),
    img_url varchar(255)
);