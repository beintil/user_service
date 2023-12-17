CREATE TABLE IF NOT EXISTS business_constants
(
    type text,
    label text,
    code int unique,
    info text
);

INSERT INTO business_constants (type, label, code, info) VALUES ('message_status', 'read_message', 1001, 'Status message read');
INSERT INTO business_constants (type, label, code, info) VALUES ('message_status', 'not_read_message', 1002, 'Status message not read');
