CREATE TABLE IF NOT EXISTS owners (
    id         SERIAL PRIMARY KEY,
    name       TEXT NOT NULL,
    birth_date DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS pets (
    id         SERIAL PRIMARY KEY,
    name       TEXT NOT NULL,
    birth_date DATE NOT NULL,
    breed      TEXT,
    color      TEXT,
    owner_id   INTEGER REFERENCES owners(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS pet_friends (
    pet_id    INTEGER REFERENCES pets(id) ON DELETE CASCADE,
    friend_id INTEGER REFERENCES pets(id) ON DELETE CASCADE,
    PRIMARY KEY (pet_id, friend_id),
    CHECK (pet_id != friend_id)
);