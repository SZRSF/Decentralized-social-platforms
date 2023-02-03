CREATE TABLE `user`(
    `user_id`       bigint(20)             NOT NULL COMMENT '用户ID',
    `user_name`     varchar(20)         NOT NULL COMMENT '用户名',
    `phone_num`     varchar(11)         NOT NULL COMMENT '手机号码',
    `emil`          varchar(64)         NOT NULL COMMENT '邮箱',
    `gender`        tinyint(4) ZEROFILL NOT NULL COMMENT '性别',
    `user_password` varchar(32)         NOT NULL COMMENT '用户密码',
    `head_img`      varchar(100)        NULL COMMENT '头像地址',
    `invite_id`     int(11)             NULL COMMENT '邀请ID',
    `time_stamp`    bigint              NULL COMMENT '时间戳',
    `create_time`   timestamp           NULL COMMENT '账号创建时间',
    PRIMARY KEY (`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;