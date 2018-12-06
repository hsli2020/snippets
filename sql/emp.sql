﻿-- 创建表DEPT
CREATE TABLE dept( /*部门表*/
    deptno MEDIUMINT   UNSIGNED  NOT NULL  DEFAULT 0,
    dname VARCHAR(20)  NOT NULL  DEFAULT "",
    loc VARCHAR(13) NOT NULL DEFAULT ""
) ENGINE=MyISAM DEFAULT CHARSET=utf8 ;

-- 创建表EMP雇员
CREATE TABLE emp (
    empno  MEDIUMINT UNSIGNED  NOT NULL  DEFAULT 0,
    ename VARCHAR(20) NOT NULL DEFAULT "",
    job VARCHAR(9) NOT NULL DEFAULT "",
    mgr MEDIUMINT UNSIGNED NOT NULL DEFAULT 0,
    hiredate DATE NOT NULL,
    sal DECIMAL(7,2)  NOT NULL,
    comm DECIMAL(7,2) NOT NULL,
    deptno MEDIUMINT UNSIGNED NOT NULL DEFAULT 0
) ENGINE=MyISAM DEFAULT CHARSET=utf8 ;

-- 工资级别表
CREATE TABLE salgrade (
    grade MEDIUMINT UNSIGNED NOT NULL DEFAULT 0,
    losal DECIMAL(17,2)  NOT NULL,
    hisal DECIMAL(17,2)  NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

INSERT INTO salgrade VALUES (1,700,1200);
INSERT INTO salgrade VALUES (2,1201,1400);
INSERT INTO salgrade VALUES (3,1401,2000);
INSERT INTO salgrade VALUES (4,2001,3000);
INSERT INTO salgrade VALUES (5,3001,9999);

-- 随机产生字符串

-- 定义一个新的命令结束符合
delimiter $$

-- 删除自定的函数
drop  function rand_string $$

-- 这里我创建了一个函数. 
create function rand_string(n INT)
returns varchar(255)
begin 
  declare chars_str varchar(100) default 'abcdefghijklmnopqrstuvwxyzABCDEFJHIJKLMNOPQRSTUVWXYZ';
  declare return_str varchar(255) default '';
  declare i int default 0;
  while i < n do 
    set return_str=concat(return_str,substring(chars_str,floor(1+rand()*52),1));
    set i = i + 1;
  end while;
  return return_str;
end $$

delimiter ;         select rand_string(6);

-- 随机产生部门编号

delimiter $$        drop function rand_num $$

-- 这里我们又自定了一个函数
create function rand_num( )
returns int(5)
begin 
  declare i int default 0;
  set i = floor(10+rand()*500);
  return i;
end $$

delimiter ;       select rand_num();
-------------------------------------------
-- 向emp表中插入记录(海量的数据)
delimiter $$      drop procedure insert_emp $$

create procedure insert_emp(in start int(10),in max_num int(10))
begin
  declare i int default 0; 
  set autocommit = 0;  
  repeat
    set i = i + 1;
    insert into emp values ((start+i),rand_string(6),'SALESMAN',001,curdate(),2000,400,rand_num());
  until i = max_num
  end repeat;
  commit;
end $$

-- 调用刚刚写好的函数, 1800000条记录,从100001号开始
delimiter ;      call insert_emp(100001,1800000);
---------------------------------------------------------------
-- 向dept表中插入记录
delimiter $$     drop procedure insert_dept $$

create procedure insert_dept(in start int(10),in max_num int(10))
begin
  declare i int default 0; 
  set autocommit = 0;  
  repeat
    set i = i + 1;
    insert into dept values ((start+i) ,rand_string(10),rand_string(8));
  until i = max_num
  end repeat;
  commit;
end $$

delimiter ;    call insert_dept(100,10);
-------------------------------------------------
-- 向salgrade 表插入数据

delimiter $$   drop procedure insert_salgrade $$

create procedure insert_salgrade(in start int(10),in max_num int(10))
begin
  declare i int default 0; 
  set autocommit = 0;
  ALTER TABLE emp DISABLE KEYS;  
  repeat
    set i = i + 1;
    insert into salgrade values ((start+i) ,(start+i),(start+i));
  until i = max_num
  end repeat;
  commit;
end $$
delimiter ;    call insert_salgrade(10000,1000000);
