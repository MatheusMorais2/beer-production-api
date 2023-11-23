CREATE TYPE role AS ENUM ('creator', 'admin', 'technician', 'brewer');

CREATE TABLE brewery_invite (
    "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    "inviting_user_id" uuid,
    "invited_user_id" uuid,
    "brewery_id" uuid,
    "role" role DEFAULT 'brewer'
);

ALTER TABLE brewery_user DROP COLUMN role;
ALTER TABLE brewery_user ADD COLUMN role role DEFAULT 'brewer';