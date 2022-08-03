CREATE TABLE "meetings" (
  "id" SERIAL not null,
  "comment" text,
  "user_name" text,
  "in_time" text PRIMARY KEY not null,
  "in_meet" boolean not null default(false));


insert into "meetings"(id,comment,user_name,in_time,in_meet)
values(1,'Никем не занята','Vladvlk','11:00','false'),
      (2,'Никем не занята',' ','11:30','false'),
      (3,'Никем не занята',' ','12:00','false'),
      (4,'Никем не занята',' ','12:30','false'),
      (5,'Никем не занята',' ','13:00','false'),
      (6,'Никем не занята',' ','13:30','false'),
      (7,'Никем не занята',' ','14:00','false'),
      (8,'Никем не занята',' ','14:30','false'),
      (9,'Никем не занята',' ','15:00','false'),
      (10,'Никем не занята',' ','15:30','false'),
      (11,'frot','vladvlk','16:00','true'),
      (12,'back','senya','16:30','true'),
      (13,'git',' ','17:00','true'),
      (14,'check',' ','17:30','true'),
      (15,'prod',' ','18:00','true'),
      (16,'junior',' ','18:30','true'),
      (17,'senior',' ','19:00','true');