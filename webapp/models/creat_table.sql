create table user
(
    id          bigint(20) not null auto_increment,
    user_id     bigint(20) not null,
    user_name   varchar(64),
    password    varchar(64),
    email       varchar(64),
    gender      varchar(20),
    create_time timestamp  not null default current_timestamp comment '创建时间',
    update_time timestamp  not null default current_timestamp on update current_timestamp comment '修改时间',
    primary key (id)
);

create table community
(
    id             int(11)   not null auto_increment,
    community_id   int(10)   not null,
    community_name varchar(128),
    introduction   varchar(256),
    create_time    timestamp not null default current_timestamp comment '创建时间',
    update_time    timestamp not null default current_timestamp on update current_timestamp comment '修改时间',
    primary key (id)
);

insert into community(id, community_id, community_name, introduction)
values (1, 1, 'go', 'golang');
insert into community(id, community_id, community_name, introduction)
values (2, 2, 'leetcode', '刷题');
insert into community(id, community_id, community_name, introduction)
values (3, 3, 'cf', '火麒麟');
insert into community(id, community_id, community_name, introduction)
values (4, 4, '无畏契约', '战神');

create table post
(
    id           bigint(20)                               not null auto_increment,
    post_id      bigint(20)                               not null comment '帖子id',
    title        varchar(128) collate utf8mb4_general_ci  not null comment '标题',
    content      varchar(8192) collate utf8mb4_general_ci not null comment '内容',
    author_id    bigint(20)                               not null comment '作者的用户id',
    community_id bigint(20)                               not null comment '所属社区',
    status       tinyint(4)                               not null default '1' comment '帖子状态',
    create_time  timestamp                                not null default current_timestamp comment '创建时间',
    update_time  timestamp                                not null default current_timestamp on update current_timestamp comment '修改时间',
    primary key (id),
    unique key idx_post_id (post_id)
);

create table vote
(
    id      bigint(20) not null auto_increment,
    post_id bigint(20) not null comment '帖子id',
    user_id bigint(20) not null comment '用户id',
    primary key (id)
);

create table comment
(
    id      bigint(20)    not null auto_increment,
    post_id bigint(20)    not null comment '帖子id',
    user_id bigint(20)    not null comment '用户id',
    content varchar(8192) not null comment '评论内容',
    primary key (id)
);