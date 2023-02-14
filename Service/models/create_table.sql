# CREATE TABLE `user`
# (
#     `user_id`       bigint(20)          NOT NULL COMMENT '用户ID',
#     `user_name`     varchar(20)         NOT NULL COMMENT '用户名',
#     `phone_num`     varchar(11)         NULL COMMENT '手机号码',
#     `emil`          varchar(64)         NULL COMMENT '邮箱',
#     `gender`        tinyint(4) ZEROFILL NULL COMMENT '性别',
#     `user_password` varchar(64)         NULL COMMENT '用户密码',
#     `head_img`      varchar(100)        NULL COMMENT '头像地址',
#     `invite_id`     int(11)             NULL COMMENT '邀请ID',
#     `time_stamp`    bigint              NULL COMMENT '时间戳',
#     `create_time`   timestamp           NULL COMMENT '账号创建时间',
#     PRIMARY KEY (`user_id`)
# ) ENGINE = InnoDB
#   DEFAULT CHARSET = utf8mb4
#   COLLATE = utf8mb4_general_ci;

# CREATE TABLE `family`
# (
#     `id`                  int(11)      NOT NULL COMMENT '表id',
#     `family_id`           int(11)      NULL COMMENT '家id',
#     `family_name`         varchar(128) NULL COMMENT '家的名称',
#     `family_founder_name` varchar(32)  NULL COMMENT '家的创建人',
#     `family_parents_id`   int(64)      NULL COMMENT '家长',
#     `family_user_id`      int(64)      NULL COMMENT '加入家中成员',
#     `introduction`        varchar(255) NULL COMMENT '家介绍',
#     `create_time`         timestamp    NULL COMMENT '家创建时间',
#     PRIMARY KEY (`id`),
#     UNIQUE INDEX `idx_family_id` (`family_id`),
#     UNIQUE INDEX `idx_family_name` (`family_name`)
# ) ENGINE = InnoDB
#   DEFAULT CHARSET = utf8mb4
#   COLLATE = utf8mb4_general_ci;

CREATE TABLE `post`
(
    `id`          bigint(20)                               NOT NULL AUTO_INCREMENT,
    `post_id`     bigint(20)                               NOT NULL COMMENT '作品id',
    `title`       varchar(128) COLLATE utf8mb4_general_ci  NOT NULL COMMENT '标题',
    `content`     varchar(8192) COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
    `author_id`   bigint(20)                               NOT NULL COMMENT '作者的用户id',
    `family_id`   bigint(20)                               NULL COMMENT '所属家',
    `status`      tinyint(4)                               NOT NULL DEFAULT 1 COMMENT '帖子状态',
    `create_time` timestamp                                NULL     DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp                                NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_post_id` (`post_id`),
    KEY `idx_author_id` (`author_id`),
    KEY `idx_community_id` (`family_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;