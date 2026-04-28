CREATE TABLE subcribtions (
    id SERIAL PRIMARY KEY,
    service_name VARCHAR(255) NOT NULL,
    user_id UUID NOT NULL,
    subcribtion_date_begining TIMESTAMPTZ DEFAULT NOW()
    subcribtion_date_ending TIMESTAMPTZ DEFAULT NOW()
);