create table cats(
  id varchar primary key,
  name varchar,
  year_of_birth integer,
  picture_id varchar
);

create table feedings(
  id varchar primary key,
  cat_id varchar references cats(id),
  cat_food_id  varchar,
  finished_eating boolean,
  new_pack boolean,
  didnt_eat boolean,
  time timestamp not null,
  note varchar
);


create table incidents(
  id varchar primary key,
  cat_id varchar references cats(id),
  kind varchar not null,
  time timestamp not null,
  picture_id varchar
);
