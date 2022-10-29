CREATE TABLE IF NOT EXISTS "author" (
    "id" SERIAL PRIMARY KEY, 
    "firstname" varchar(255) NOT NULL,
    "lastname" varchar(255) NOT NULL,
    "create_at" TIMESTAMP DEFAULT (Now()),
     "update_at" TIMESTAMP DEFAULT (Now())
)

CREATE TABLE IF NOT EXISTS "article"(
    "id" SERIAL PRIMARY KEY, 
    "title" varchar(255) NOT NULL,
    "body" TEXT,
    "author_id" INT,

    "create_at" TIMESTAMP DEFAULT (Now()),
    "update_at" TIMESTAMP DEFAULT (Now()),
    CONSTRAINT fk_author FOREIGN KEY(author_id) REFERENCES author(id)
);