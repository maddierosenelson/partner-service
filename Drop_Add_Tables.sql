drop table keys cascade;
drop table groups cascade;
drop table partners cascade;
drop table groups_to_keys cascade;
drop table partner_mappings cascade;

CREATE TABLE keys (
    id serial primary key,
    name varchar
);

CREATE TABLE groups (
    id serial primary key,
    name varchar
);

CREATE TABLE partners (
    id serial primary key,
    name varchar,
    code varchar
);

CREATE TABLE groups_to_keys (
    id serial primary key,
    group_id int,
    key_id int,
    FOREIGN KEY(group_id) REFERENCES groups(id),
    FOREIGN KEY(key_id) REFERENCES keys(id)
);

CREATE TABLE partner_mappings (
    id serial primary key,
    partner_id int,
    key_id int,
    FOREIGN KEY(partner_id) REFERENCES keys(id),
    FOREIGN KEY(key_id) REFERENCES keys(id),
    value varchar
);

INSERT INTO keys (name) VALUES ('Currency');
INSERT INTO keys (name) VALUES ('Type of Payment');
INSERT INTO keys (name) VALUES ('860');
INSERT INTO keys (name) VALUES ('850');
INSERT INTO keys (name) VALUES ('Color');
INSERT INTO keys (name) VALUES ('Gender');
INSERT INTO keys (name) VALUES ('Sleeves');
INSERT INTO keys (name) VALUES ('ISAID');

INSERT INTO groups (name) VALUES ('EDI');
INSERT INTO groups (name) VALUES ('Style');
INSERT INTO groups (name) VALUES ('Money');

INSERT INTO partners (name, code) VALUES ('Kohls', 'KOH');
INSERT INTO partners (name, code) VALUES ('JC Penny', 'JCP');
INSERT INTO partners (name, code) VALUES ('Mustang', 'MUS');
INSERT INTO partners (name, code) VALUES ('Barrett', 'BAR');
INSERT INTO partners (name, code) VALUES ('Dicks', 'DIC');
INSERT INTO partners (name, code) VALUES ('Fanatics', 'FAN');

INSERT INTO groups_to_keys (group_id, key_id) VALUES (1, 3);
INSERT INTO groups_to_keys (group_id, key_id) VALUES (1, 4);
INSERT INTO groups_to_keys (group_id, key_id) VALUES (2, 5);
INSERT INTO groups_to_keys (group_id, key_id) VALUES (2, 6);
INSERT INTO groups_to_keys (group_id, key_id) VALUES (2, 7);
INSERT INTO groups_to_keys (group_id, key_id) VALUES (3, 1);
INSERT INTO groups_to_keys (group_id, key_id) VALUES (3, 2);

INSERT INTO partner_mappings (partner_id, key_id, value) VALUES (1, 1, 'USD');
INSERT INTO partner_mappings (partner_id, key_id, value) VALUES (1, 2, 'Credit');
INSERT INTO partner_mappings (partner_id, key_id, value) VALUES (2, 1, 'CAD');
INSERT INTO partner_mappings (partner_id, key_id, value) VALUES (2, 2, 'Cash');
INSERT INTO partner_mappings (partner_id, key_id, value) VALUES (2, 5, 'Red');
INSERT INTO partner_mappings (partner_id, key_id, value) VALUES (2, 6, 'Youth');
INSERT INTO partner_mappings (partner_id, key_id, value) VALUES (2, 7, 'Long');
INSERT INTO partner_mappings (partner_id, key_id, value) VALUES (3, 3, 'Sent');
INSERT INTO partner_mappings (partner_id, key_id, value) VALUES (3, 4, 'Received');
INSERT INTO partner_mappings (partner_id, key_id, value) VALUES (4, 3, 'Not Sent');
INSERT INTO partner_mappings (partner_id, key_id, value) VALUES (4, 3, 'Not Received');
INSERT INTO partner_mappings (partner_id, key_id, value) VALUES (4, 1, 'CAD');
INSERT INTO partner_mappings (partner_id, key_id, value) VALUES (3, 1, 'USD');
INSERT INTO partner_mappings (partner_id, key_id, value) VALUES (5, 1, 'USD');
INSERT INTO partner_mappings (partner_id, key_id, value) VALUES (5, 2, 'Debit');
INSERT INTO partner_mappings (partner_id, key_id, value) VALUES (5, 5, 'Black');
INSERT INTO partner_mappings (partner_id, key_id, value) VALUES (5, 6, 'Women');
INSERT INTO partner_mappings (partner_id, key_id, value) VALUES (5, 7, 'Short');
INSERT INTO partner_mappings (partner_id, key_id, value) VALUES (4, 8, 'BARRETT1142');
INSERT INTO partner_mappings (partner_id, key_id, value) VALUES (3, 8, 'MUSTANGDRINK');
INSERT INTO partner_mappings (partner_id, key_id, value) VALUES (6, 8, 'FANATICSWS');
