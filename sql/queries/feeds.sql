-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (   
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT feeds.name, url, users.name
FROM feeds
INNER JOIN users
    ON feeds.user_id = users.id;

-- name: FeedByURL :one
SELECT *
FROM feeds
WHERE url = $1;

-- name: CreateFeedFollow :one

WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES
    ($1,
    $2,
    $3,
    $4,
    $5)
    RETURNING *
)
SELECT inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
 FROM inserted_feed_follow
    INNER JOIN users
        ON inserted_feed_follow.user_id = users.id
    INNER JOIN feeds
        ON inserted_feed_follow.feed_id = feeds.id;

-- name: GetFeedFollowsForUser :many

SELECT 
feed_follows.*,
users.name AS user_name,
feeds.name AS feed_name
 FROM feed_follows
    INNER JOIN users 
        ON feed_follows.user_id = users.id
    INNER JOIN feeds
        ON feed_follows.feed_id = feeds.id
WHERE feed_follows.user_id = $1;

-- name: DeleteFeedFollow :exec

DELETE
FROM feed_follows
USING feeds, users
WHERE 
feed_follows.feed_id = feeds.id
AND feed_follows.user_id = users.id
AND users.name = $1 AND feeds.url = $2;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET 
last_fetched_at = NOW() ::TIMESTAMP,
updated_at = NOW() ::TIMESTAMP
WHERE feeds.id = $1;

-- name: GetNextFeedToFetch :one
SELECT *
FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;


