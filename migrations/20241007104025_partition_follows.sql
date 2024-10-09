-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS public.posts;

CREATE TABLE public.follows (
  id bigserial,
  follower_id uuid not null,
  followee_id uuid not null,
  "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  primary key (id, created_at)
) partition by range(created_at);

create index if not exists tx_follow_created_at_idx on public.follows (created_at);

create UNIQUE index if not exists tx_u_follow_follower_followee_id_idx on public.follows (follower_id, followee_id);

CREATE
OR REPLACE FUNCTION create_follow_partitions(start_date DATE, end_date DATE) RETURN VOID AS $ $ DECLARE curr_start DATE := start_date;

curr_end DATE;

BEGIN WHILE curr_start < end_date LOOP curr_end := curr_start + INTERVAL '7 days';

EXECUTE format(
  'CREATE TABLE IF NOT EXISTS follow_%s PARTITION OF public."follows" FOR VALUES FROM (%L) TO (%L)',
  to_char(curr_start, 'YYYYMMDD'),
  curr_start,
  curr_end
);

curr_start := curr_end;

END LOOP;

END;

$ $ LANGUAGE plpgsql;

SELECT
  create_follow_partitions('2024-10-08' :: DATE, '2025-01-01' :: DATE);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd