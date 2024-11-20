-- Table: mst_role
CREATE TABLE mst_role (
    id_role VARCHAR PRIMARY KEY,
    role_name VARCHAR NOT NULL
);

-- Table: mst_user
CREATE TABLE mst_user (
    id_user VARCHAR PRIMARY KEY,
    name VARCHAR NOT NULL,
    email VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    phone_number VARCHAR NOT NULL,
    role_id VARCHAR NOT NULL,
    CONSTRAINT fk_role
        FOREIGN KEY (role_id)
        REFERENCES mst_role (id_role)
);


INSERT INTO mst_role (id_role, role_name) 
VALUES
    ('ROLE001', 'Ketua Umum'),
    ('ROLE002', 'Sekretaris Umum'),
    ('ROLE003', 'Bendahara Umum'),
    ('ROLE004', 'Departement Keorganisasian'),
    ('ROLE005', 'Departement Pembelajaran'),
    ('ROLE006', 'Departement Humas');