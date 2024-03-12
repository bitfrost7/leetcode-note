Create table If Not Exists Person (id int, email varchar(255));
Truncate table Person;
insert into Person (id, email) values ('1', 'a@b.com');
insert into Person (id, email) values ('2', 'c@d.com');
insert into Person (id, email) values ('3', 'a@b.com');


# 表: Person
#
# +-------------+---------+
# | Column Name | Type    |
# +-------------+---------+
# | id          | int     |
# | email       | varchar |
# +-------------+---------+
# id 是该表的主键（具有唯一值的列）。
# 此表的每一行都包含一封电子邮件。电子邮件不包含大写字母。
#
#
# 编写解决方案来报告所有重复的电子邮件。 请注意，可以保证电子邮件字段不为 NULL。
#
# 以 任意顺序 返回结果表。

# select Email from
#     (
#         select Email, count(Email) as num
#         from Person
#         group by Email
#     ) as a
# where num > 1;

select email from person group by email having count(email) > 1 ;
