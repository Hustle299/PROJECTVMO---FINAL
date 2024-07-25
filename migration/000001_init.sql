CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS departments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(45) NOT NULL,
    address VARCHAR(100) NOT NULL,
    status SMALLINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON COLUMN departments.status IS '0: inactive\n1: active';

CREATE TABLE IF NOT EXISTS countries (
    id SERIAL PRIMARY KEY,
    name VARCHAR(45) NOT NULL,
    status SMALLINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    role_id INT DEFAULT NULL,
    department_id INT DEFAULT NULL,
    email VARCHAR(45) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    name VARCHAR(45) NOT NULL,
    surname VARCHAR(45) NOT NULL,
    gender VARCHAR(20) DEFAULT NULL,
    dob DATE DEFAULT NULL,
    mobile VARCHAR(15) DEFAULT NULL,
    country_id INT DEFAULT NULL,
    resident_country_id INT DEFAULT NULL,
    avatar VARCHAR(100) DEFAULT NULL,
    verification_status SMALLINT DEFAULT 0,
    status SMALLINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_users_roles FOREIGN KEY (role_id) REFERENCES roles (id),
    CONSTRAINT fk_users_departments FOREIGN KEY (department_id) REFERENCES departments (id),
    CONSTRAINT fk_users_countries FOREIGN KEY (country_id) REFERENCES countries (id),
    CONSTRAINT fk_users_resident_countries FOREIGN KEY (resident_country_id) REFERENCES countries (id)
);

COMMENT ON COLUMN users.verification_status IS '0: unverified\n1: verified';
COMMENT ON COLUMN users.status IS '0: inactive\n1: active';

CREATE TABLE IF NOT EXISTS requests (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    type VARCHAR(45) NOT NULL,
    status SMALLINT NOT NULL,
    reject_notes VARCHAR(255) DEFAULT NULL,
    verifier_id INT DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_requests_users FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT fk_requests_verifiers FOREIGN KEY (verifier_id) REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS user_identities (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    number VARCHAR(30) NOT NULL,
    type VARCHAR(45) NOT NULL,
    status SMALLINT NOT NULL,
    expiry_date DATE NOT NULL,
    place_issued VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user_identities_users FOREIGN KEY (user_id) REFERENCES users (id)
);
