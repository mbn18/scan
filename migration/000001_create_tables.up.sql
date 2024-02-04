CREATE TABLE service
(
    id   UUID default gen_random_uuid() PRIMARY KEY, -- should we allow default? probably not.
    name TEXT UNIQUE not null -- is it unique, can there be two services with same name?
);

CREATE TABLE resource
(
    urn          TEXT PRIMARY KEY,
    type         TEXT NOT NULL, -- Would consider to put types on external table
    name         TEXT,
    data         JSONB,
    generated_at timestamp
);

CREATE TABLE service_resource
(
    service_id   UUID REFERENCES service (id) NOT NULL,
    resource_urn TEXT REFERENCES resource (urn)
);

-- Full --------------------------

-- we need trigger or constraint that check that there is a ref to resource
CREATE TABLE scan_full
(
    id          UUID      default gen_random_uuid() PRIMARY KEY,
    executed_at timestamp default NOW()
);

CREATE TABLE scan_full_resource
(
    scan_full_id UUID REFERENCES scan_full (id) not null,
    resource_urn TEXT REFERENCES resource (urn) not null
);

CREATE TABLE scan_full_service
(
    scan_full_id UUID REFERENCES scan_full (id) not null,
    service_id   UUID references service (id)   not null
);

-- Partial --------------------------

CREATE TABLE scan_partial
(
    id          UUID      default gen_random_uuid() PRIMARY KEY,
    service_id  UUID references service (id) not null,
    executed_at timestamp default NOW()
);

CREATE TABLE scan_partial_resource
(
    scan_partial_id UUID REFERENCES scan_partial (id) not null,
    resource_urn    TEXT REFERENCES resource (urn)    not null
);
