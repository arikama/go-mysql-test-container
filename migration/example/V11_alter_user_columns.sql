ALTER TABLE `user` ADD COLUMN `username` VARCHAR(64) NOT NULL UNIQUE AFTER `id`;
ALTER TABLE `user` ADD COLUMN `name` VARCHAR(64) NOT NULL UNIQUE AFTER `username`;
