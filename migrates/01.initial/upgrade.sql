BEGIN;

CREATE TABLE aether.user
(
    username        VARCHAR PRIMARY KEY      NOT NULL UNIQUE,
    password        VARCHAR                  NOT NULL,
    created_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

/* password = josue */
INSERT INTO aether.user (username, password) VALUES('josue', '$2a$14$0XemxjHRP2ix/sjlVi9Y9OUMe/l7BHqYkxt8A3y1fOa0/nwUBdFUS');

GRANT USAGE ON ALL SEQUENCES IN SCHEMA aether TO aether_api;
GRANT SELECT ON aether.user TO aether_api;

COMMIT;