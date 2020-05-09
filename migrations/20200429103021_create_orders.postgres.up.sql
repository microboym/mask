CREATE TABLE "orders" (
    "phone" char(11) references users("phone") NOT NULL check("phone" ~ '^[0-9]+$'),
    PRIMARY KEY("phone"),
    "img" bytea NOT NULL,
    "created_at" timestamp NOT NULL default CURRENT_TIMESTAMP
);