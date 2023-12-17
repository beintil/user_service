create table if not exists "user"
(
    id         serial not null constraint id primary key,
    password   text not null,
    email      text unique,
    created_at timestamp not null,
    updated_at timestamp default null
);

COMMENT ON TABLE "user" IS 'User table is store with authorization data';
COMMENT ON COLUMN "user".id IS 'Unique identifier for each user';
COMMENT ON COLUMN "user".password IS 'Password user';
COMMENT ON COLUMN "user".email IS 'Email user';
COMMENT ON COLUMN "user".created_at IS 'Timestamp indicating when the record was created';
COMMENT ON COLUMN "user".updated_at IS 'Timestamp indicating when the record was last updated';