CREATE table users ( 
    phone char(11) primary key NOT NULL check(phone ~ '^[0-9]+$'), 
    passwd char(6) NOT NULL check(passwd ~ '^[0-9]+$') 
);
