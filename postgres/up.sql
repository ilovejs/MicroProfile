DROP TABLE IF EXISTS profiles;

CREATE TABLE profiles (
  id       serial PRIMARY KEY,
  name     text not null,
  gender   boolean,
  dob      date,
  postcode numeric,
  phone_no text
);
