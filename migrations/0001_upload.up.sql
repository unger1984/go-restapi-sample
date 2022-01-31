DROP TABLE IF EXISTS "Upload";
CREATE TABLE "Upload" (
        "id" serial PRIMARY KEY,
        "path" varchar(255) NOT NULL,
        "name" varchar(255) DEFAULT NULL::character varying,
        "type" varchar(255) DEFAULT NULL::character varying,
        "created_at" timestamptz(6) NOT NULL,
        "updated_at" timestamptz(6) NOT NULL
);
COMMENT ON COLUMN "Upload"."path" IS 'Путь к файлу';
COMMENT ON COLUMN "Upload"."name" IS 'Имя файла';
COMMENT ON COLUMN "Upload"."type" IS 'Content Type';
