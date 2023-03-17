-- name: GetSignatureWeapons :many
SELECT * FROM signature_weapon;

-- name: GetSignatureWeaponByEnumId :one
SELECT * FROM signature_weapon WHERE enum_id = ?;
