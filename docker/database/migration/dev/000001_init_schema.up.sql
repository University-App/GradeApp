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