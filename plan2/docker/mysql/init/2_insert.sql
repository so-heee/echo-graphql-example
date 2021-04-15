USE sample;

create table todos (
  id         int auto_increment,
  userId     int,
  text       text,
  created_at timestamp not null default current_timestamp,
  updated_at timestamp not null default current_timestamp on update current_timestamp,
  primary key(id)
);

insert into `todos`(
    `id`,
    `userId`,
    `text`,
    `created_at`,
    `updated_at`
)
values(
    1,
    1,
    'aaaa',
    '2020-01-01 00:00:00',
    '2020-01-02 00:00:00'
),
(
    2,
    2,
    'bbbb',
    '2020-02-01 00:00:00',
    '2020-02-02 00:00:00'
),
(
    3,
    3,
    'cccc',
    '2020-03-01 00:00:00',
    '2020-03-02 00:00:00'
),
(
    4,
    4,
    'dddd',
    '2020-04-01 00:00:00',
    '2020-04-02 00:00:00'
),
(
    5,
    5,
    'eeee',
    '2020-05-01 00:00:00',
    '2020-05-02 00:00:00'
),
(
    6,
    6,
    'ffff',
    '2020-06-01 00:00:00',
    '2020-06-02 00:00:00'
),
(
    7,
    7,
    'gggg',
    '2020-07-01 00:00:00',
    '2020-07-02 00:00:00'
)
;