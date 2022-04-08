-- create admin user with password "admin1234"
insert into users(login, email, displayName, password)
  values ('admin', 'admin@example.com', 'Administrative User', '$2a$14$Nf0unAko3.VQcr/X.52.Bu6H1jN9Ir/6ps3hk1X4Sqp9fAyFlMvxe')
