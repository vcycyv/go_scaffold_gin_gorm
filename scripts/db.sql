CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE {entity}
(
    id UUID DEFAULT uuid_generate_v4(),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    name text,
    CONSTRAINT {entity}_pkey PRIMARY KEY (id)
)
