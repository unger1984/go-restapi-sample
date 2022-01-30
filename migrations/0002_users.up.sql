CREATE TABLE "User"
(
    "id"          serial       PRIMARY KEY ,
    "email"       varchar(255) not null unique,
    "password"    varchar(255) not null,
    "avatarId"    int4
);
COMMENT ON COLUMN "User"."email" IS 'Уникальный Email';
COMMENT ON COLUMN "User"."password" IS 'Пароль пользователя';
COMMENT ON COLUMN "User"."avatarId" IS 'Ссылка на uploads';
ALTER TABLE "User" ADD CONSTRAINT "User_avatarId_fkey" FOREIGN KEY ("avatarId") REFERENCES "Upload" ("id") ON DELETE SET NULL ON UPDATE NO ACTION;
