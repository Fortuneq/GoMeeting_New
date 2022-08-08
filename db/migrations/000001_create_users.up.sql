CREATE TABLE "meetings_1" (
  "id" SERIAL not null,
  "comment" text,
  "user_name" text,
  "user_chat_id" int,
  "priority" int,
  "in_time" text PRIMARY KEY not null,
  "in_meet" boolean not null default(false));


CREATE TABLE "meetings_2" (
  "id" SERIAL not null,
  "comment" text,
  "user_name" text,
  "user_chat_id" int,
  "priority" int,
  "in_time" text PRIMARY KEY not null,
  "in_meet" boolean not null default(false));


insert into "meetings_1"(id,comment,user_name,user_chat_id,priority,in_time,in_meet)
values(1,'KUVFKWiGEk','Tomi Gentry',654565,1,'11:00','true'),
      (2,'Заняла для обсуждения кофе','SophieVL',8765456,1,'11:30','true'),
      (3,'Заняла ','greoiner',765436,1,'12:00','true'),
      (4,'wmdRmdvOea','Sohail Gamble',98765,1,'12:30','true'),
      (5,' EbsBlWbIto','Jayden Pham',9876,1,'13:00','true'),
      (6,'LujgfwvOzF','Emily-Jane Philip',2347654,1,'13:30','true'),
      (7,'vpwUsrUWmR','Gaia Wood',123456,1,'14:00','true'),
      (8,'aZdwCKpRAR ','Alicia Peel',874592,1,'14:30','true'),
      (9,'svXeHfGoWH','Kerys Marsden',98471289,1,'15:00','true'),
      (10,'stRsmDccdF','Philip Baird',1241241,1,'15:30','true'),
      (11,'dOknjTupja','Cheryl Chen',94320324,1,'16:00','true'),
      (12,'Никем не занята','',0,1,'16:30','false'),
      (13,'Никем не занята','',0,1,'17:00','false'),
      (14,'Никем не занята','',0,1,'17:30','false'),
      (15,'Никем не занята','',0,1,'18:00','false'),
      (16,'Никем не занята','',0,1,'18:30','false'),
      (17,'Никем не занята','',0,1,'19:00','false');

insert into "meetings_2"(id,comment,user_name,user_chat_id,priority,in_time,in_meet)
values(1,'Кем то занята','vfghjk','1234567',1,'11:00','true'),
      (2,'Никем не занята',' ',0,1,'11:30','false'),
      (3,'Никем не занята',' ',0,1,'12:00','false'),
      (4,'Никем не занята',' ',0,1,'12:30','false'),
      (5,'Никем не занята',' ',0,1,'13:00','false'),
      (6,'Никем не занята',' ',0,1,'13:30','false'),
      (7,'Никем не занята',' ',0,1,'14:00','false'),
      (8,'Никем не занята',' ',0,1,'14:30','false'),
      (9,'Никем не занята',' ',0,1,'15:00','false'),
      (10,'Никем не занята',' ',0,1,'15:30','false'),
      (11,'Никем не занята',' ',0,1,'16:00','false'),
      (12,'Никем не занята',' ',0,1,'16:30','false'),
      (13,'Никем не занята',' ',0,1,'17:00','false'),
      (14,'Никем не занята',' ',0,1,'17:30','false'),
      (15,'Никем не занята',' ',0,1,'18:00','false'),
      (16,'Никем не занята',' ',0,1,'18:30','false'),
      (17,'Никем не занята',' ',0,1,'19:00','false');


