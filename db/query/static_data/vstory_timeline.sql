-- name: GetVstoryTimelines :many
SELECT * FROM vstory_timeline;

-- name: GetVstoryTimelineByEnumId :one
SELECT * FROM vstory_timeline WHERE enum_id = ?;
