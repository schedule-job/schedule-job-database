CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE jobs
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    job_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    name text NOT NULL,
    description text NOT NULL,
    author text NOT NULL,
    members text[],
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    PRIMARY KEY (id),
    UNIQUE (id)
);

CREATE TABLE actions
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    type text NOT NULL,
    name text NOT NULL,
    payload json NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    job_id uuid NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (id)
);

CREATE TABLE triggers
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    name text NOT NULL,
    payload json NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    job_id uuid NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (id)
);

CREATE TABLE request_logs
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    job_id uuid NOT NULL,
    status text NOT NULL DEFAULT 'progress',
    request_url text NOT NULL,
    request_method text NOT NULL,
    request_headers json,
    request_body text DEFAULT '',
    response_headers json,
    response_body text DEFAULT '',
    response_status_code integer,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    PRIMARY KEY (id),
    UNIQUE (id)
);

CREATE TABLE authorizations
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    name text NOT NULL,
    payload json,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    PRIMARY KEY (id),
    UNIQUE (id)
);

CREATE TABLE urls
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    category text NOT NULL,
    url text NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    PRIMARY KEY (id),
    UNIQUE (id)
) 