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
INSERT INTO keys (name) VALUES ('ISAID');
INSERT INTO keys (name) VALUES ('Qualifier');
INSERT INTO keys (name) VALUES ('DM_VENDOR_CODE');


INSERT INTO groups (name) VALUES ('EDI');
INSERT INTO groups (name) VALUES ('Money');

INSERT INTO partners (name, code) VALUES ('Mustang', 'MUS');
INSERT INTO partners (name, code) VALUES ('Barrett', 'BAR');
INSERT INTO partners (name, code) VALUES ('Fanatics', 'FAN');

INSERT INTO groups_to_keys (group_id, key_id) VALUES (1, 2);
INSERT INTO groups_to_keys (group_id, key_id) VALUES (1, 3);
INSERT INTO groups_to_keys (group_id, key_id) VALUES (2, 1);
INSERT INTO groups_to_keys (group_id, key_id) VALUES (1, 4);


INSERT INTO partner_mappings (partner_id, key_id, value) VALUES (2, 1, 'USD');
INSERT INTO partner_mappings (partner_id, key_id, value) VALUES (1, 1, 'CAD');
INSERT INTO partner_mappings (partner_id, key_id, value) VALUES (2, 2, 'BARRETT1142');
INSERT INTO partner_mappings (partner_id, key_id, value) VALUES (1, 2, 'MUSTANGDRINK');
INSERT INTO partner_mappings (partner_id, key_id, value) VALUES (3, 2, 'FANATICSWS');
INSERT INTO partner_mappings (partner_id, key_id, value) VALUES (1, 3, 'ZZ');
INSERT INTO partner_mappings (partner_id, key_id, value) VALUES (2, 3, 'ZZ');
INSERT INTO partner_mappings (partner_id, key_id, value) VALUES (3, 3, 'ZZ');
INSERT INTO partner_mappings (partner_id, key_id, value) VALUES (1, 4, 'MUS');
INSERT INTO partner_mappings (partner_id, key_id, value) VALUES (2, 4, 'BAR');
INSERT INTO partner_mappings (partner_id, key_id, value) VALUES (3, 4, 'FAN');
