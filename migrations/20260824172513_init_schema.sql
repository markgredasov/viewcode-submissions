-- +goose Up
CREATE SCHEMA IF NOT EXISTS users;

CREATE TYPE users.submission_status AS ENUM ('PENDING', 'RUNNING', 'ACCEPTED', 'WRONG_ANSWER', 'TIMEOUT', 'INTERNAL');

CREATE TYPE users.available_languages AS ENUM ('go', 'py', 'java', 'cpp', 'js', 'ts', 'csharp', 'ko', 'sw', 'rust');

CREATE TABLE IF NOT EXISTS users.submissions (
    id                  UUID                          PRIMARY KEY,
    user_id             UUID                          NOT NULL,
    problem_id          UUID                          NOT NULL,
    tests_total         SMALLINT                      NOT NULL,
    tests_passed        SMALLINT                      NOT NULL,
    condition           users.submission_status       NOT NULL,
    lang                users.available_languages     NOT NULL,
    execution_time_ms   INT                           NOT NULL,
    created_at          TIMESTAMP                     NOT NULL          DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS users.submissions;

DROP TYPE IF EXISTS users.available_languages;

DROP TYPE IF EXISTS users.submission_status;

DROP SCHEMA IF EXISTS users;
