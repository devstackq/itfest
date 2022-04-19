--versionDocs - src doc ?
CREATE TABLE bimbo_company
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE bimbo_departament
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE bimbo_position
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE bimbo_role
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE bimbo_user
(
    id BIGSERIAL PRIMARY KEY,
    full_name VARCHAR(80) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    company_id INTEGER NOT NULL REFERENCES bimbo_company(id) ON DELETE CASCADE ON UPDATE CASCADE,
    departament_id INTEGER NOT NULL REFERENCES bimbo_departament(id) ON DELETE CASCADE ON UPDATE CASCADE,
    position_id INTEGER NOT NULL REFERENCES bimbo_position(id) ON DELETE CASCADE ON UPDATE CASCADE,
    role_id INTEGER NOT NULL REFERENCES bimbo_role(id) ON DELETE CASCADE ON UPDATE CASCADE --def 1 worker
);

CREATE TABLE bimbo_document
(
    id BIGSERIAL PRIMARY KEY,
    path VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    author_id INTEGER NOT NULL REFERENCES bimbo_user(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE bimbo_type_access
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE bimbo_access
(
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    document_id INTEGER NOT NULL REFERENCES bimbo_document(id) ON DELETE CASCADE ON UPDATE CASCADE,
    access_type_id INTEGER NOT NULL REFERENCES bimbo_type_access(id) ON DELETE CASCADE ON UPDATE CASCADE,
    to_whom_id INTEGER NOT NULL REFERENCES bimbo_user(id) ON DELETE CASCADE ON UPDATE CASCADE
);


CREATE TABLE bimbo_type_event
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);
CREATE TABLE bimbo_history_event
(
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    document_id INTEGER NOT NULL REFERENCES bimbo_document(id) ON DELETE CASCADE ON UPDATE CASCADE,
    type_event_id INTEGER NOT NULL REFERENCES bimbo_type_event(id) ON DELETE CASCADE ON UPDATE CASCADE,
    author_id INTEGER NOT NULL REFERENCES bimbo_user(id) ON DELETE CASCADE ON UPDATE CASCADE,
    to_whom_id INTEGER NOT NULL REFERENCES bimbo_user(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- CREATE TABLE bimbo_template_category
-- (
--     id BIGSERIAL PRIMARY KEY,
--     name VARCHAR(50) NOT NULL UNIQUE,
--     template_id INTEGER NOT NULL REFERENCES bimbo_template(id) ON DELETE CASCADE ON UPDATE CASCADE
-- );

CREATE TABLE bimbo_template
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    body TEXT
);

-- Компания?  : amount : float32 : 3
CREATE TABLE bimbo_choice
(
    id BIGSERIAL PRIMARY KEY,
    jsonKey VARCHAR(50) NOT NULL,
    value VARCHAR(10000) NOT NULL,
    title VARCHAR(500) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    htmlTagKey VARCHAR(100) NOT NULL,
    -- htmlTagValue VARCHAR(100) NOT NULL,
    template_id INTEGER NOT NULL REFERENCES bimbo_template(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- backend return client - json {map[string]string}

-- CREATE TABLE bimbo_tepmlate_choice_index 
-- (
--     id BIGSERIAL PRIMARY KEY,
--     name VARCHAR(50) NOT NULL UNIQUE,
--     template_id INTEGER NOT NULL REFERENCES bimbo_template(id) ON DELETE CASCADE ON UPDATE CASCADE,
--     choice_id INTEGER NOT NULL REFERENCES bimbo_choice(id) ON DELETE CASCADE ON UPDATE CASCADE
-- )

--pending/done/reject
CREATE TABLE bimbo_sign_status
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);
--doc22 -> user12 -> user 55
CREATE TABLE bimbo_sign
(
    id BIGSERIAL PRIMARY KEY,
    from_who INTEGER NOT NULL REFERENCES bimbo_user(id) ON DELETE CASCADE ON UPDATE CASCADE,
    to_whom_id INTEGER NOT NULL REFERENCES bimbo_user(id) ON DELETE CASCADE ON UPDATE CASCADE,
    status_id INTEGER NOT NULL REFERENCES bimbo_sign_status(id) ON DELETE CASCADE ON UPDATE CASCADE,
    document_id INTEGER NOT NULL REFERENCES bimbo_document(id) ON DELETE CASCADE ON UPDATE CASCADE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
