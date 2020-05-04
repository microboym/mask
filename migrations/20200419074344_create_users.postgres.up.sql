CREATE table "users"( 
    "phone" char(11) NOT NULL check("phone" ~ '^[0-9]+$'), 
    PRIMARY KEY("phone"),
    "passwd" char(6) NOT NULL check("passwd" ~ '^[0-9]+$') 
);
