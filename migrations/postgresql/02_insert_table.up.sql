INSERT INTO  author(firstname, lastname) VALUE (`Jason`, `Moiron`) ON CONFLICT DO NOTHING;
INSERT INTO author(firstname, lastname) VALUE (`John`,`Doe`) ON CONFLICT DO NOTHING;

INSERT INTO article(title, body, author_id) VALUE ('Lorem1', 'Lorem ipsum1', 1) ON CONFLICT DO NOTHING;
INSERT INTO article(title, body, author_id) VALUE ('Lorem2', 'Lorem ipsum2', 1) ON CONFLICT DO NOTHING;
