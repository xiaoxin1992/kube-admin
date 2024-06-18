create table cluster
(
    id          bigint auto_increment comment '自增ID'
        primary key,
    zone        varchar(200)                           not null comment '区域',
    host        text                                   not null comment '主机地址',
    token       text                                   not null comment '集群token',
    version     varchar(200) default ''                not null comment '集群版本',
    remark      text                                   not null comment '备注',
    create_time timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time timestamp    default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    constraint cluster_pk_2
        unique (zone)
);

create table users
(
    id           bigint auto_increment comment 'id'
        primary key,
    role         int          default 2                 not null comment '角色: 1 表示管理员, 2表示普通用户, 默认是2 普通用户',
    username     varchar(200)                           not null comment '用户名',
    display_name varchar(200)                           not null comment '显示用户名',
    password     varchar(255)                           not null comment '密码',
    email        varchar(200) default ''                not null comment '用户邮箱',
    phone        varchar(12)  default ''                not null comment '用户手机号',
    create_time  timestamp    default CURRENT_TIMESTAMP not null,
    update_time  timestamp    default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    constraint users_pk_2
        unique (username)
)
    comment '用户表';
