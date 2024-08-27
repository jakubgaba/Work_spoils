CREATE TABLE "test" (
    "id" INTEGER NOT NULL UNIQUE, -- unique ID test
    "name" TEXT NOT NULL,
    "location" TEXT NOT NULL,
    PRIMARY KEY ("id")
);

INSERT INTO "main"."test"(
    "id",
    "name",
    "location"
)
VALUES
    ('1','Test_1','Russia'),
    ('2','Test_2','Canada'),
    ('3','Test_3','Nigeria'),
    ('4','Test_4','Poland'),
    ('5','Test_5','Russia'),
    ('6','Test_6','Nigeria'),
    ('7','Test_7','Nigeria'),
    ('8','Test_8','UK');

