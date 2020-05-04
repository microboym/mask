CREATE TABLE "captchas"(
    "img_sha" bytea NOT NULL PRIMARY KEY,
    "img" bytea NOT NULL,
    "answer" integer check("answer" < 100)
);