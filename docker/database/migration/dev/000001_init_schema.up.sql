CREATE TABLE "Grade" (
                         "id" SERIAL PRIMARY KEY,
                         "grade_number" int NOT NULL,
                         "course_name" varchar NOT NULL,
                         "student_id" int
);

CREATE TABLE "Unite" (
                         "id" SERIAL PRIMARY KEY,
                         "unite_name" varchar NOT NULL
);

CREATE TABLE "Course" (
                          "id" SERIAL PRIMARY KEY,
                          "course_name" varchar NOT NULL,
                          "unite_id" int
);

CREATE TABLE "Student" (
                           "id" SERIAL PRIMARY KEY,
                           "last_name" varchar NOT NULL,
                           "first_name" varchar NOT NULL
);

ALTER TABLE "Grade" ADD FOREIGN KEY ("student_id") REFERENCES "Student" ("id");

ALTER TABLE "Course" ADD FOREIGN KEY ("unite_id") REFERENCES "Unite" ("id");

CREATE TABLE "Student_Course" (
                                  "Student_id" int NOT NULL,
                                  "Course_id" int NOT NULL,
                                  PRIMARY KEY ("Student_id", "Course_id")
);

ALTER TABLE "Student_Course" ADD FOREIGN KEY ("Student_id") REFERENCES "Student" ("id");

ALTER TABLE "Student_Course" ADD FOREIGN KEY ("Course_id") REFERENCES "Course" ("id");

CREATE TABLE "CourseAverage" (
                           "id" SERIAL PRIMARY KEY,
                           "course_name" varchar NOT NULL,
                           "average" int NOT NULL
);

CREATE TABLE "StudentCourseAverage" (
                                        "id" SERIAL PRIMARY KEY,
                                        "course_name" varchar NOT NULL,
                                        "student_name" varchar NOT NULL,
                                        "average" int NOT NULL
);

----------------------------------------------------------------------------------------
/*
    SEEDING DEV DATABASE
 */

INSERT INTO "Student" (id, last_name, first_name) VALUES (1, 'Djek', 'Pm');
INSERT INTO "Student" (id, last_name, first_name) VALUES (2, 'Djek', 'test');
INSERT INTO "Student" (id, last_name, first_name) VALUES (3, 'Djek', 'test1');

INSERT INTO "Unite" (id, unite_name) VALUES (1, 'UE1');
INSERT INTO "Unite" (id, unite_name) VALUES (2, 'UE2');

INSERT INTO "Course" (id, course_name, unite_id) VALUES (1, 'Course1', 1);
INSERT INTO "Course" (id, course_name, unite_id) VALUES (2, 'Course3', 1);
INSERT INTO "Course" (id, course_name, unite_id) VALUES (3, 'Course2', 2);
INSERT INTO "Course" (id, course_name, unite_id) VALUES (4, 'Course4', 2);

INSERT INTO "Grade" (id, grade_number, course_name, student_id) VALUES (1, 10, 'Course1', 1);
INSERT INTO "Grade" (id, grade_number, course_name, student_id) VALUES (2, 15, 'Course1', 2);
INSERT INTO "Grade" (id, grade_number, course_name, student_id) VALUES (3, 20, 'Course1', 1);
INSERT INTO "Grade" (id, grade_number, course_name, student_id) VALUES (4, 13, 'Course1', 2);

INSERT INTO "Student_Course" ("Student_id", "Course_id") VALUES (1, 1);
INSERT INTO "Student_Course" ("Student_id", "Course_id") VALUES (1, 2);
INSERT INTO "Student_Course" ("Student_id", "Course_id") VALUES (1, 3);
INSERT INTO "Student_Course" ("Student_id", "Course_id") VALUES (1, 4);
INSERT INTO "Student_Course" ("Student_id", "Course_id") VALUES (2, 1);
INSERT INTO "Student_Course" ("Student_id", "Course_id") VALUES (2, 2);
INSERT INTO "Student_Course" ("Student_id", "Course_id") VALUES (2, 3);
INSERT INTO "Student_Course" ("Student_id", "Course_id") VALUES (2, 4);


ALTER SEQUENCE "Student_id_seq" RESTART WITH 4;
ALTER SEQUENCE "Unite_id_seq" RESTART WITH 3;
ALTER SEQUENCE "Course_id_seq" RESTART WITH 5;
ALTER SEQUENCE "Grade_id_seq" RESTART WITH 5;