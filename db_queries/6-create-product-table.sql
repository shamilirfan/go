CREATE TABLE
    product (
        id SERIAL PRIMARY KEY,
        title varchar(100) NOT NULL,
        description TEXT,
        price DOUBLE PRECISION NOT NULL,
        img_url TEXT NOT NULL,
        created_at TIMESTAMP
        WITH
            TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP
        WITH
            TIME ZONE DEFAULT CURRENT_TIMESTAMP
    )