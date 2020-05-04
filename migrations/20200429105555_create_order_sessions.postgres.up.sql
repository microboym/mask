CREATE TABLE "order_sessions" (
    "phone" char(11) references users("phone") NOT NULL check("phone" ~ '^[0-9]+$'),
    PRIMARY KEY("phone"),
    "captcha1" bytea references captchas("img_sha"),
    "captcha2" bytea references captchas("img_sha")
        check("captcha1" is not null OR "captcha2" is null),
    "created_at" timestamp NOT NULL default CURRENT_TIMESTAMP
);