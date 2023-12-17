create table if not exists person
(
    id serial not null constraint person_pk primary key,
    user_id      integer not null constraint user_id references "user" on update cascade on delete cascade,
    first_name   text    not null,
    last_name    text    not null,
    middle_name  text default null,
    age          integer default null,
    gender       text default null,
    birthday     date default null,
    phone_number varchar(15) default null,
    created_at   timestamp not null ,
    updated_at   timestamp default null
);

COMMENT ON TABLE person IS 'User information';
COMMENT ON COLUMN person.id IS 'Unique identifier for each person';
COMMENT ON COLUMN person.user_id IS 'Foreign key referencing the user associated with this person';
COMMENT ON COLUMN person.first_name IS 'First name of the person';
COMMENT ON COLUMN person.last_name IS 'Last name of the person';
COMMENT ON COLUMN person.middle_name IS 'Middle name of the person';
COMMENT ON COLUMN person.age IS 'Age of the person';
COMMENT ON COLUMN person.gender IS 'Gender of the person';
COMMENT ON COLUMN person.birthday IS 'Date of birth of the person';
COMMENT ON COLUMN person.phone_number IS 'Phone number of the person';
COMMENT ON COLUMN person.created_at IS 'Timestamp indicating when the record was created';
COMMENT ON COLUMN person.updated_at IS 'Timestamp indicating when the record was last updated';
