DROP TABLE IF EXISTS "Upload";
CREATE TABLE "Upload" (
        "id" serial PRIMARY KEY,
        "path" varchar(255) NOT NULL,
        "name" varchar(255) DEFAULT NULL::character varying,
        "type" varchar(255) DEFAULT NULL::character varying
);
COMMENT ON COLUMN "Upload"."path" IS 'Путь к файлу';
COMMENT ON COLUMN "Upload"."name" IS 'Имя файла';
COMMENT ON COLUMN "Upload"."type" IS 'Content Type';
