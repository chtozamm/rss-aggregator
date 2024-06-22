-- +goose Up
ALTER TABLE users ADD COLUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT (
  -- Generate random bytes > cast them into binary > apply hash function > encode into hexadecimal
  encode(sha256(random()::text::bytea), 'hex') 
);

-- +goose Down
ALTER TABLE users DROP COLUMN api_key;