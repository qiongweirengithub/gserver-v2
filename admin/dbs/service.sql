CREATE TABLE `g_service` (
  `id` int(10) primary key,
  `name` varchar(25) NOT NULL,
  `status` tinyint(3) NOT NULL DEFAULT '0',
  `host` varchar(25) NOT NULL,
  `port` varchar(25) NOT NULL,
  `docker_image` varchar(25) NOT NULL,
  `container_id` varchar(25) NOT NULL,
  `service_type` varchar(25) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

alter TABLE g_service modify id int(10) auto_increment;

