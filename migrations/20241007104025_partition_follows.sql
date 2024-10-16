-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS public.posts;

CREATE TABLE public.follows (
  id bigserial primary key,
  follower_id uuid not null,
  followee_id uuid not null,
  "created_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL DEFAULT CURRENT_TIMESTAMP,
) partition by hash (follower_id);

create UNIQUE index if not exists tx_u_follow_follower_followee_id_idx on public.follows (follower_id, followee_id);

CREATE TABLE follows_part_1 PARTITION OF follows FOR VALUES WITH (MODULUS 4, REMAINDER 0);
CREATE TABLE follows_part_2 PARTITION OF follows FOR VALUES WITH (MODULUS 4, REMAINDER 1);
CREATE TABLE follows_part_3 PARTITION OF follows FOR VALUES WITH (MODULUS 4, REMAINDER 2);
CREATE TABLE follows_part_4 PARTITION OF follows FOR VALUES WITH (MODULUS 4, REMAINDER 3);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd