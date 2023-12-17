CREATE TABLE IF NOT EXISTS chat (
    id SERIAL NOT NULL CONSTRAINT chat_id PRIMARY KEY,
    owner_id INT NOT NULL REFERENCES "user"(id) on update cascade on delete cascade,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP default null
);

CREATE TABLE IF NOT EXISTS chat_member (
    id SERIAL NOT NULL CONSTRAINT chat_member_id PRIMARY KEY,
    chat_id INTEGER NOT NULL REFERENCES chat(id) on update cascade on delete cascade,
    user_id INTEGER NOT NULL REFERENCES "user"(id) on update cascade on delete cascade
);