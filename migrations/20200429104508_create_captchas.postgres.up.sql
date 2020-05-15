CREATE TABLE "captchas"(
    "img_sha" bytea NOT NULL PRIMARY KEY,
    "img" bytea NOT NULL,
    "equation" char(5), check("equation" ~ '^\d+[+-]\d+'),
);