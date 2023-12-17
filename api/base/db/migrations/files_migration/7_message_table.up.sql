CREATE TABLE IF NOT EXISTS message (
    id  serial not null constraint message_pk primary key,
    sender_id INT NOT NULL,
    receiver_id INT default null,
    group_id INT default null,
    post_id INT default null,
    content TEXT NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status_id INT NOT NULL,
    is_deleted BOOLEAN DEFAULT false,
    chat_id INT NOT NULL ,
    created_at   timestamp not null,
    updated_at   timestamp default null,
    FOREIGN KEY (sender_id) REFERENCES "user"(id) on update cascade on delete cascade,
    FOREIGN KEY (receiver_id) REFERENCES "user"(id) on update cascade on delete cascade,
    FOREIGN KEY (group_id) REFERENCES "group"(id) on update cascade on delete cascade,
    FOREIGN KEY (post_id) REFERENCES post(id) on update cascade on delete cascade,
    FOREIGN KEY (chat_id) REFERENCES chat(id) on update cascade on delete cascade,
    FOREIGN KEY (status_id) REFERENCES business_constants(code) on update cascade on delete cascade
);

COMMENT ON TABLE message IS 'Table for storing messages';

COMMENT ON COLUMN message.id IS 'Unique identifier for each message';
COMMENT ON COLUMN message.sender_id IS 'Identifier of the sender';
COMMENT ON COLUMN message.receiver_id IS 'Identifier of the receiver (can be NULL)';
COMMENT ON COLUMN message.group_id IS 'Identifier of the group (can be NULL)';
COMMENT ON COLUMN message.post_id IS 'Identifier of the post (can be NULL)';
COMMENT ON COLUMN message.content IS 'Text content of the message';
COMMENT ON COLUMN message.timestamp IS 'Timestamp indicating when the message was sent';
COMMENT ON COLUMN message.status_id IS 'Status of the message';
COMMENT ON COLUMN message.is_deleted IS 'Flag indicating whether the message is deleted';
COMMENT ON COLUMN message.chat_id IS 'Identifier of the chat (can be NULL)';
