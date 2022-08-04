CREATE TABLE "meetings" (
  "id" SERIAL not null,
  "comment" text,
  "user_name" text,
  "in_time" text PRIMARY KEY not null,
  "in_meet" boolean not null default(false));


insert into "meetings"(id,comment,user_name,in_time,in_meet)
values(1,'Никем не занята','Vladvlk','11:00','false'),
      (2,'Заняла для обсуждения кофе','SophieVL','11:30','true'),
      (3,'Заняла ','greoiner','12:00','true'),
      (4,'Никем не занята',' ','12:30','false'),
      (5,'Никем не занята',' ','13:00','false'),
      (6,'Никем не занята',' ','13:30','false'),
      (7,'Никем не занята',' ','14:00','false'),
      (8,'Никем не занята',' ','14:30','false'),
      (9,'Никем не занята',' ','15:00','false'),
      (10,'Никем не занята',' ','15:30','false'),
      (11,'Никем не занята',' ','16:00','false'),
      (12,'Никем не занята',' ','16:30','false'),
      (13,'Никем не занята',' ','17:00','false'),
      (14,'Никем не занята',' ','17:30','false'),
      (15,'Никем не занята',' ','18:00','false'),
      (16,'Никем не занята',' ','18:30','false'),
      (17,'Никем не занята',' ','19:00','false');