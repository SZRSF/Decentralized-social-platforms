# CREATE TABLE `user`
# (
#     `user_id`       bigint(20)          NOT NULL COMMENT '用户ID',
#     `user_name`     varchar(20)         NOT NULL COMMENT '用户名',
#     `phone_num`     varchar(11)         NULL COMMENT '手机号码',
#     `emil`          varchar(64)         NULL COMMENT '邮箱',
#     `gender`        tinyint(4)          NULL COMMENT '性别',
#     `user_password` varchar(64)         NULL COMMENT '用户密码',
#     `head_img`      varchar(255)        NULL COMMENT '头像地址',
#     `works_count`   int(15)             NOT NULL DEFAULT 0 COMMENT '作品数量',
#     `follow_count`  int(15)             NOT NULL DEFAULT 0 COMMENT '关注数量',
#     `fans_count`  int(15)               NOT NULL DEFAULT 0 COMMENT '粉丝数量',
#     `like_count`    int(20)             NOT NULL DEFAULT 0 COMMENT '获赞数量',
#     `collect_count` int(20)             NOT NULL DEFAULT 0 COMMENT '收藏',
#     `joined_family` int(20)             NOT NULL DEFAULT 0 COMMENT '加入的家',
#     `browsing_history` int(20)          NOT NULL DEFAULT 0 COMMENT '浏览记录',
#     `invite_id`     int(11)             NULL COMMENT '邀请ID',
#     `time_stamp`    bigint              NULL COMMENT '时间戳',
#     `create_time`   timestamp           NULL COMMENT '账号创建时间',
#     PRIMARY KEY (`user_id`)
# ) ENGINE = InnoDB
#   DEFAULT CHARSET = utf8mb4
#   COLLATE = utf8mb4_general_ci
#     COMMENT='用户表';

# CREATE TABLE `family`
# (
#     `id`                  int(11)      NOT NULL COMMENT '表id',
#     `family_id`           int(11)      NULL COMMENT '家id',
#     `family_name`         varchar(128) NULL COMMENT '家的名称',
#     `family_founder_name` varchar(32)  NULL COMMENT '家的创建人',
#     `family_head_img`      varchar(100)        NULL COMMENT '家头像地址',
#     `family_parents_id`   int(64)      NULL COMMENT '家长',
#     `family_user_id`      int(64)      NULL COMMENT '加入家中成员',
#     `introduction`        varchar(255) NULL COMMENT '家介绍',
#     `create_time`         timestamp    NULL COMMENT '家创建时间',
#     PRIMARY KEY (`id`),
#     UNIQUE INDEX `idx_family_id` (`family_id`),
#     UNIQUE INDEX `idx_family_name` (`family_name`)
# ) ENGINE = InnoDB
#   DEFAULT CHARSET = utf8mb4
#   COLLATE = utf8mb4_general_ci
#   COMMENT='家详情表';;

# CREATE TABLE `post`
# (
#     `id`          bigint(20)                               NOT NULL AUTO_INCREMENT,
#     `post_id`     bigint(20)                               NOT NULL COMMENT '作品id',
#     `title`       varchar(128) COLLATE utf8mb4_general_ci  NOT NULL COMMENT '标题',
#     `content`     varchar(8192) COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
#     `author_id`   bigint(20)                               NOT NULL COMMENT '作者的用户id',
#     `family_id`   bigint(20)                               NULL COMMENT '所属家',
#     `status`      tinyint(4)                               NOT NULL DEFAULT 1 COMMENT '帖子状态',
#     `create_time` timestamp                                NULL     DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
#     `update_time` timestamp                                NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
#     PRIMARY KEY (`id`),
#     UNIQUE KEY `idx_post_id` (`post_id`),
#     KEY `idx_author_id` (`author_id`),
#     KEY `idx_community_id` (`family_id`)
# ) ENGINE = InnoDB
#   DEFAULT CHARSET = utf8mb4
#   COLLATE = utf8mb4_general_ci
#   COMMENT='作品表';;

# CREATE TABLE user_family
# (
#     `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
#     `user_id`     bigint(20)  NOT NULL COMMENT '用户ID',
#     `family_id`   int(11)  NOT NULL COMMENT '频道ID',
#     is_owner BOOLEAN NOT NULL DEFAULT false COMMENT '是否为拥有者',
#     created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
#     updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
#     FOREIGN KEY (user_id) REFERENCES user(user_id),
#     FOREIGN KEY (family_id) REFERENCES family(family_id),
#     UNIQUE KEY (user_id, family_id)
# ) ENGINE = InnoDB
#   DEFAULT CHARSET = utf8mb4
#   COLLATE = utf8mb4_general_ci
# COMMENT ='用户关注家中间表';

# CREATE TABLE `user_follow`
# (
#     `id`           bigint(20)    NOT NULL AUTO_INCREMENT COMMENT '关注ID',
#     `follower_id`  bigint(20)    NOT NULL COMMENT '关注者ID',
#     `following_id` bigint(20)    NOT NULL COMMENT '被关注者ID',
#     `status`       tinyint(1) NOT NULL COMMENT '关注状态，0为未关注，1为已关注，2为互相关注',
#     `created_at`   datetime   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '关注时间',
#     PRIMARY KEY (`id`),
#     UNIQUE KEY `uk_follower_following` (`follower_id`, `following_id`),
#     KEY `idx_follower_id` (`follower_id`),
#     KEY `idx_following_id` (`following_id`)
# ) ENGINE = InnoDB
#   DEFAULT CHARSET = utf8mb4
#   COLLATE = utf8mb4_general_ci
#   COMMENT ='用户关注表';

CREATE TABLE `user_collection`
(
    `id`         bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id`    bigint(20) NOT NULL,
    `post_id` bigint(20) NOT NULL,
    `created_at` datetime   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_article` (`user_id`, `post_id`),
    KEY `idx_user_id` (`user_id`),
    KEY `idx_post_id` (`post_id`),
    CONSTRAINT `fk_user_collection_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_user_collection_post_id` FOREIGN KEY (`post_id`) REFERENCES `post` (`post_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
    COMMENT ='用户收藏表';
