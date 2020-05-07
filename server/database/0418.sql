CREATE TABLE  users (
  id int auto_increment primary key,
  email varchar(256) not null,
  hashed_password varchar(512) not null,
  `name` varchar(256) not null,
  created_at datetime default current_timestamp,
  updated_at datetime default current_timestamp on update current_timestamp,
  deleted_at datetime default null
);

CREATE TABLE  body_info (
  user_id int,
  `weight` decimal,
  `body_fat` decimal,
  created_at datetime default current_timestamp,
  updated_at datetime default current_timestamp on update current_timestamp,
  deleted_at datetime default null,

  foreign key(user_id) references users(id)
);

CREATE TABLE  crunch (
  user_id int,
  `rep` int,
  `set` int,
  created_at datetime default current_timestamp,
  updated_at datetime default current_timestamp on update current_timestamp,
  deleted_at datetime default null,

  foreign key(user_id) references users(id)
);

CREATE TABLE  dead_lift (
  user_id int,
  `weight` decimal,
  `rep` int,
  `set` int,
  created_at datetime default current_timestamp,
  updated_at datetime default current_timestamp on update current_timestamp,
  deleted_at datetime default null,

  foreign key(user_id) references users(id)
);

CREATE TABLE  chin_up (
  user_id int,
  `rep` int,
  `set` int,
  created_at datetime default current_timestamp,
  updated_at datetime default current_timestamp on update current_timestamp,
  deleted_at datetime default null,

  foreign key(user_id) references users(id)
);

CREATE TABLE  dumbell_curl (
  user_id int,
  `weight` decimal,
  `rep` int,
  `set` int,
  created_at datetime default current_timestamp,
  updated_at datetime default current_timestamp on update current_timestamp,
  deleted_at datetime default null,

  foreign key(user_id) references users(id)
);

CREATE TABLE  bench_press (
  user_id int,
  `weight` decimal,
  `rep` int,
  `set` int,
  created_at datetime default current_timestamp,
  updated_at datetime default current_timestamp on update current_timestamp,
  deleted_at datetime default null,

  foreign key(user_id) references users(id)
);

CREATE TABLE  hip_thrust (
  user_id int,
  `weight` decimal,
  `rep` int,
  `set` int,
  created_at datetime default current_timestamp,
  updated_at datetime default current_timestamp on update current_timestamp,
  deleted_at datetime default null,

  foreign key(user_id) references users(id)
);

CREATE TABLE  squat (
  user_id int,
  `weight` decimal,
  `rep` int,
  `set` int,
  created_at datetime default current_timestamp,
  updated_at datetime default current_timestamp on update current_timestamp,
  deleted_at datetime default null,

  foreign key(user_id) references users(id)
);

CREATE TABLE  shoulder_press (
  user_id int,
  `weight` decimal,
  `rep` int,
  `set` int,
  created_at datetime default current_timestamp,
  updated_at datetime default current_timestamp on update current_timestamp,
  deleted_at datetime default null,

  foreign key(user_id) references users(id)
);

CREATE TABLE  cable_press_down (
  user_id int,
  `weight` decimal,
  `rep` int,
  `set` int,
  created_at datetime default current_timestamp,
  updated_at datetime default current_timestamp on update current_timestamp,
  deleted_at datetime default null,

  foreign key(user_id) references users(id)
);