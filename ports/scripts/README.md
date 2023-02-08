### Scripts

In this folder there are several initialization scripts that are setting up the postgres database. I have separated them such that:

- [`postgres](postgres/) contains the scripts creating the tables. It is not the responsibility of these scripts to manage databases, users etc. They should be mounted inside the container, and executed in the correct database by the correct user.
- [`initpostgres`](initpostgres/) is useful only when trying to run the database through the docker compose file inside this service. It contains a shell script that creates the database and user, and runs the initialization scripts as the new user. This is done similarly to the compose file for the entire project, except that it allows to only set up the database, not the entire project.

Note that postgres 15 is the latest version, but that changed the way how `GRANT ALL` works. Now, the new user must be provided privileges on the public schema as well. I went thus with postgres 14 since I'm more familiar with it.