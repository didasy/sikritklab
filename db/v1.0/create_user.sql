use sikritklab;

create user 'sikrit'@'localhost' identified by 'sikritpassword';
grant all privileges on posts.* to 'sikrit'@'localhost';
grant all privileges on threads.* to 'sikrit'@'localhost';