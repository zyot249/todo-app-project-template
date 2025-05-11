DROP TABLE IF EXISTS todo;
create table todo
(
    id          serial NOT NULL,
    title       text NOT NULL ,
    description text NOT NULL ,
    created_at  time NOT NULL ,
    finished    bool default false
);

ALTER TABLE todo ADD CONSTRAINT todo_pkey PRIMARY KEY (id);