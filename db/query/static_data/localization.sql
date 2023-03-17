-- name: GetLocalizations :many
SELECT * FROM localization;

-- name: GetLocalizationByEnumId :one
SELECT * FROM localization WHERE enum_id = ?;
