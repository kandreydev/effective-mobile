-- +goose Up
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS subscriptions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    service_name    TEXT NOT NULL,
    price           INT NOT NULL,
    user_id         UUID NOT NULL,
    start_date      DATE NOT NULL,
    end_date        DATE, --optional
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT price_nonnegative CHECK (price >= 0)
);

-- Keep updated_at fresh on UPDATE
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = now();
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd


CREATE TRIGGER trg_set_updated_at
BEFORE UPDATE ON subscriptions
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

-- +goose Down
DROP TRIGGER IF EXISTS trg_set_updated_at ON subscriptions;
DROP FUNCTION IF EXISTS set_updated_at();
DROP TABLE IF EXISTS subscriptions;
