-- +goose Up
-- +goose StatementBegin

-- Назначаем права для super_admin и admin
INSERT INTO role_permissions (role_id, permission_id) VALUES
     -- super_admin получает все права: user, role, permission
    ('5f43769e-bb58-4120-9ec8-9caaf9409ff3', '041a6cff-5d0d-4537-9f8a-2f58a2b3d429'),  -- user.create
    ('5f43769e-bb58-4120-9ec8-9caaf9409ff3', '1bc0e2b7-eab5-4a07-9a27-fbd6f2d0a570'),  -- user.delete
    ('5f43769e-bb58-4120-9ec8-9caaf9409ff3', 'ddc8ba1a-81ec-47c4-8e2c-c86818b40d4c'),  -- user.read
    ('5f43769e-bb58-4120-9ec8-9caaf9409ff3', '854869bf-4455-4c97-a19e-53a5a531dc77'),  -- user.list
    ('5f43769e-bb58-4120-9ec8-9caaf9409ff3', '147f5bc5-f071-47a3-b23b-b6d076f51676'),  -- user.update
    ('5f43769e-bb58-4120-9ec8-9caaf9409ff3', 'da3ca285-8001-4dca-8486-ad667bd5c4a8'),  -- user.restore

    ('5f43769e-bb58-4120-9ec8-9caaf9409ff3', 'a6feac36-3d67-45f8-b787-9e1e6c322096'),  -- role.create
    ('5f43769e-bb58-4120-9ec8-9caaf9409ff3', 'adaa977a-6573-4437-a15c-2b9511193022'),  -- role.delete
    ('5f43769e-bb58-4120-9ec8-9caaf9409ff3', '2e62e6de-1764-4d74-be50-d08be673bc66'),  -- role.read
    ('5f43769e-bb58-4120-9ec8-9caaf9409ff3', 'ff71d389-e0f7-4cba-9df8-0ea95fb18e1f'),  -- role.list
    ('5f43769e-bb58-4120-9ec8-9caaf9409ff3', 'd7fdfacd-55b5-4c9e-bb39-93240ad55546'),  -- role.update
    ('5f43769e-bb58-4120-9ec8-9caaf9409ff3', '11da4c13-8252-439f-a5d8-f85c167b2141'),  -- role.restore

    ('5f43769e-bb58-4120-9ec8-9caaf9409ff3', 'ca24487a-dfd6-4ca4-82b7-a86dbc0a8992'),  -- permission.create
    ('5f43769e-bb58-4120-9ec8-9caaf9409ff3', 'fc257e28-14db-4030-8bc6-741287aab3f3'),  -- permission.delete
    ('5f43769e-bb58-4120-9ec8-9caaf9409ff3', 'fc57a6b0-9d96-47da-b9df-7745394365f5'),  -- permission.read
    ('5f43769e-bb58-4120-9ec8-9caaf9409ff3', 'd1f0d48b-32bf-4087-9b50-443a80c0bc06'),  -- permission.list
    ('5f43769e-bb58-4120-9ec8-9caaf9409ff3', '2fbdf044-18e9-478b-b3e1-eb9f5670ca7c'),  -- permission.update
    ('5f43769e-bb58-4120-9ec8-9caaf9409ff3', 'c7d12711-6d51-4156-b53e-539c517c9c97'),  -- permission.restore

    -- admin получает ограниченные права: только на чтение и список прав
    ('9936aed6-680a-4f4b-a093-389a30be4a15', '041a6cff-5d0d-4537-9f8a-2f58a2b3d429'),  -- user.create
    ('9936aed6-680a-4f4b-a093-389a30be4a15', '1bc0e2b7-eab5-4a07-9a27-fbd6f2d0a570'),  -- user.delete
    ('9936aed6-680a-4f4b-a093-389a30be4a15', 'ddc8ba1a-81ec-47c4-8e2c-c86818b40d4c'),  -- user.read
    ('9936aed6-680a-4f4b-a093-389a30be4a15', '854869bf-4455-4c97-a19e-53a5a531dc77'),  -- user.list
    ('9936aed6-680a-4f4b-a093-389a30be4a15', '147f5bc5-f071-47a3-b23b-b6d076f51676'),  -- user.update
    ('9936aed6-680a-4f4b-a093-389a30be4a15', 'da3ca285-8001-4dca-8486-ad667bd5c4a8'),  -- user.restore

    ('9936aed6-680a-4f4b-a093-389a30be4a15', 'a6feac36-3d67-45f8-b787-9e1e6c322096'),  -- role.create
    ('9936aed6-680a-4f4b-a093-389a30be4a15', 'adaa977a-6573-4437-a15c-2b9511193022'),  -- role.delete
    ('9936aed6-680a-4f4b-a093-389a30be4a15', '2e62e6de-1764-4d74-be50-d08be673bc66'),  -- role.read
    ('9936aed6-680a-4f4b-a093-389a30be4a15', 'ff71d389-e0f7-4cba-9df8-0ea95fb18e1f'),  -- role.list
    ('9936aed6-680a-4f4b-a093-389a30be4a15', 'd7fdfacd-55b5-4c9e-bb39-93240ad55546'),  -- role.update
    ('9936aed6-680a-4f4b-a093-389a30be4a15', '11da4c13-8252-439f-a5d8-f85c167b2141'),  -- role.restore

    -- admin получает только возможность просматривать права (read, list), но не изменять их
    ('9936aed6-680a-4f4b-a093-389a30be4a15', 'ca24487a-dfd6-4ca4-82b7-a86dbc0a8992'),  -- permission.create
    ('9936aed6-680a-4f4b-a093-389a30be4a15', 'fc257e28-14db-4030-8bc6-741287aab3f3'),  -- permission.delete
    ('9936aed6-680a-4f4b-a093-389a30be4a15', 'fc57a6b0-9d96-47da-b9df-7745394365f5'),  -- permission.read
    ('9936aed6-680a-4f4b-a093-389a30be4a15', 'd1f0d48b-32bf-4087-9b50-443a80c0bc06'),  -- permission.list
    ('9936aed6-680a-4f4b-a093-389a30be4a15', '2fbdf044-18e9-478b-b3e1-eb9f5670ca7c'),  -- permission.update
    ('9936aed6-680a-4f4b-a093-389a30be4a15', 'c7d12711-6d51-4156-b53e-539c517c9c97'),  -- permission.restore
-- +goose StatementEnd
