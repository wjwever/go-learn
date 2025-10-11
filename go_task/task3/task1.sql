use task3;

drop table if exists students;

create table students (
    id int primary key auto_increment,
    name varchar(30) not null,
    age int not null,
    grade varchar(30) not null
);

insert into students values(1, '张三', 20, '三年级');

select * from students where age > 18;

update students set grade = '四年级' where name = '张三';

select * from students;

delete from students where age < 15;
