-- 都道府県マスタ
CREATE TABLE prefectures (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 郵便番号マスタ
CREATE TABLE postal_codes (
    id SERIAL PRIMARY KEY,
    code TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 企業テーブル（外部キー付き）
CREATE TABLE companies (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    postal_code_id INT REFERENCES postal_codes(id),
    prefecture_id INT REFERENCES prefectures(id),
    address TEXT,
    contact_name TEXT,
    phone TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
