DELETE FROM `post` 
WHERE `status` = 'hidden' 
AND `updated_time` < DATE_SUB(NOW(), INTERVAL 7 DAY);

DELETE FROM `comment` 
WHERE `status` = 'hidden' 
AND `updated_at` < DATE_SUB(NOW(), INTERVAL 7 DAY);