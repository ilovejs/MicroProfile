DROP TABLE IF EXISTS profiles;

CREATE TABLE profiles (
  id       serial PRIMARY KEY,
  name     text not null,
  gender   boolean,
  dob      date,
  postcode numeric,
  phone_no text
);
INSERT INTO public.profiles (id, name, gender, dob, postcode, phone_no) VALUES (70, 'Michael Zhuang', true, '2019-10-28', 1111, '425569579');
INSERT INTO public.profiles (id, name, gender, dob, postcode, phone_no) VALUES (71, 'Michael Zhong', true, '2019-10-28', 1112, '425569579');
INSERT INTO public.profiles (id, name, gender, dob, postcode, phone_no) VALUES (72, 'Michael Ho', true, '2019-10-28', 1112, '425569579');
INSERT INTO public.profiles (id, name, gender, dob, postcode, phone_no) VALUES (73, 'Michael Wong', true, '2019-10-28', 1112, '425569579');
