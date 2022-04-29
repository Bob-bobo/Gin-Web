
DROP TABLE IF EXISTS `blog_tag`;
CREATE TABLE `blog_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签管理';

-------------
--- 用户表 ---
-------------

DROP TABLE IF EXISTS `user_tab`;

create table `user_tab`(
                     `id`    int auto_increment  primary key,
                     `phone`          varchar(255) not null,
                     `username`       varchar(255) not null,
                     `password`       varchar(255) not null,
                     `gender`         char(255)    not null,
                     `true_name`       varchar(255) null,
                     `birthday`       varchar(255) null,
                     `email`          varchar(255) null,
                     `personal_brief`  varchar(255) null,
                     `avatar_img_url`   text         not null,
                     `recently_landed` varchar(255) null
) engine=innodb charset = utf8;

-------------
--- 角色表 ---
-------------

DROP TABLE IF EXISTS `role_tab`;

create table `role_tab`(
                     `id` int(11) not null auto_increment,
                     `name` varchar(255) not null,
                     primary key(id)
)engine=innodb auto_increment=4 default charset=utf8;


-------------
--- 用户角色表 ---
-------------


CREATE TABLE `user_role_tab` (
                               `User_id` int(11) NOT NULL,
                               `Role_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-------------
--- 活动表 ---
-------------

CREATE TABLE `activity_tab`(
                             `id` int(11) not null auto_increment,
                             `act_name` varchar(255) not null,
                             `act_catego` int(11) not null,
                             `act_start` timestamp not null,
                             `act_end` timestamp not null,
                             `act_site` varchar(255) null,
                             `act_pers` int(4) null,
                             `act_detail` varchar(255) null,
                             Primary key(id)
)engine=innodb default charset=utf8;

-------------
--- 活动用户表 ---
-------------

CREATE TABLE `activity_user_tab`(
                                  `user_id` int(11) not null,
                                  `activity_id` int(11) not null
)engine=innodb default charset=utf8;

-------------
--- 分类表 ---
-------------

CREATE TABLE `category_tab`(
                             `id` int(11) not null auto_increment,
                             `category_name` varchar(255) not null,
                             primary key(id)
)engine=innodb default charset=utf8;

-------------
--- 评论表 ---
-------------

create table `comment_tab`(
                            `id` int(11) not null auto_increment,
                            `user_id` int(11) not null,
                            `activity_id` int(11) not null,
                            `comment_content` varchar(255) not null,
                            `publish_time` timestamp
                                primary key(id)
)engine=innodb default charset=utf8;
