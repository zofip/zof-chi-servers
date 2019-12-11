CREATE TABLE domains (
    id SERIAL PRIMARY KEY,
    servers_changed BOOL,
    ssl_grade STRING(100),
    previous_ssl_grade STRING(50),
    logo STRING(100),
    title STRING(100),
    is_down BOOL,
    host STRING(100)
);

CREATE TABLE servers (
    id SERIAL PRIMARY KEY,
    address STRING(100),
    ssl_grade STRING(100),
    country STRING(5),
    owner STRING(100),
    domain_id SERIAL REFERENCES domains(id) ON DELETE CASCADE,
    INDEX fk_ref_domains (domain_id ASC)
);

INSERT INTO domains (servers_changed, ssl_grade, previous_ssl_grade, logo, title, is_down, host)
VALUES (true, 'B', 'A+', 'https://uploads-ssl.webflow.com/5b559a554de48fbcb01fd277/5d2e09d5b78ea81d63876dc3_LogoBlanco.svg',
'Truora', false, 'truora.com');

INSERT INTO domains (servers_changed, ssl_grade, previous_ssl_grade, logo, title, is_down, host)
VALUES (true, 'A', 'B+', 'http://gymvirtual.com/wp-content/uploads/2015/02/logo-gym-virtual.png',
'GymVirtual', false, 'gymvirtual.com/');

INSERT INTO servers (address, ssl_grade, country, owner, domain_id)
VALUES ('server1', 'B', 'US', 'Amazon.com, Inc.', '510841046735945729');

INSERT INTO servers (address, ssl_grade, country, owner, domain_id)
VALUES ('server2', 'A+', 'US', 'Amazon.com, Inc.', '510841046735945729');

INSERT INTO servers (address, ssl_grade, country, owner, domain_id)
VALUES ('server3', 'A', 'US', 'Amazon.com, Inc.', '510841046735945729');

INSERT INTO servers (address, ssl_grade, country, owner, domain_id)
VALUES ('server1', 'AA', 'US', 'Azure.com', '510841635558948865');

INSERT INTO servers (address, ssl_grade, country, owner, domain_id)
VALUES ('server2', 'BB', 'US', 'Azure.com', '510841635558948865');


