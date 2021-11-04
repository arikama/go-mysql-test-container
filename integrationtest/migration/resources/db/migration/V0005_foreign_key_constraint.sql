ALTER TABLE `child` ADD CONSTRAINT FOREIGN KEY (`parent_id`) REFERENCES `parent` (`id`);
