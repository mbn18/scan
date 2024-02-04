INSERT INTO service (id, name)
values ('b1d13374-e8f0-4bd4-b156-1b5c85852c60', 'service A'),
       ('5981bc0c-4e6b-4cb4-bae3-b004c00fc4a8', 'service B'),
       ('74e1c073-1494-47e4-b113-db6b166fa36d', 'service C');

INSERT INTO resource (urn, name, type, data)
VALUES ('urn:okta:users:123', 'jack', 'IDP_User', '{"mfa": true}'),
       ('urn:github:users:875', 'jack', 'Service_User', '{}'),
       ('urn:meow:users:456', 'miki', 'IDP_User', '{"mfa": true, "ganesh": "elephant"}'),
       ('urn:github:users:876', 'miki', 'Service_User', '{"shiva": "hands"}'),
       ('urn:meow:users:457', 'Tahash', 'Service_User', '{"mfa": true, "ganesh": "elephant"}');

INSERT INTO service_resource (service_id, resource_urn)
VALUES ('b1d13374-e8f0-4bd4-b156-1b5c85852c60', 'urn:okta:users:123'),
       ('b1d13374-e8f0-4bd4-b156-1b5c85852c60', 'urn:github:users:875'),
       ('5981bc0c-4e6b-4cb4-bae3-b004c00fc4a8', 'urn:meow:users:456'),
       ('5981bc0c-4e6b-4cb4-bae3-b004c00fc4a8', 'urn:github:users:876'),
       ('74e1c073-1494-47e4-b113-db6b166fa36d', 'urn:meow:users:457');

CREATE MATERIALIZED VIEW service_user AS
SELECT t1.urn, t1.name, t1.type, jsonb_build_object('idp', split_part(t2.urn, ':', 2)) || t1.data || t2.data AS data
FROM resource t1 JOIN resource t2 ON(t1.name=t2.name)
WHERE t1.type = 'Service_User' AND t2.type = 'IDP_User';
