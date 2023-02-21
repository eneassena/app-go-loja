create table products (
	`id` int not null auto_increment primary key,
	`name` varchar(100) not null,
	`type` varchar(100) not null,
	`count` int(11) not null,
	`price` decimal(18,2) not null
);