CREATE TABLE IF NOT EXISTS post
(
    id serial not null constraint post_pk primary key,
    group_id INT NOT NULL,
    FOREIGN KEY (group_id) REFERENCES "group"(id) on update cascade on delete cascade
);