Postgres
--------

This package provides functionality to interact with a postgres database. It also implements the `store.Store` interface, which describes the possible interactions with a database provided by the service.

Database Interaction
--------------------

I've used `gorm` as the object-relational mapper of this database. However, I do want to note that I usually tend to prefer configuring the entire database schema from within the database itself via sql, rather than from the object relational mapper. I belive there are benefits and drawbacks for both approaches, but the reasons I prefer to go with this approach generally are:

- I'm familiar with Postgres, so it's just faster to write sql and set up initialization and migration scripts, as well as backups;
- I don't have to learn the ORM in-depth, I can just benefit from minimal functionality such as transactions, thread-safety, connectivity etc;
- Less coupling, database is more independent of the application (and also language-agnostic) and can be worked on separately by a dedicated database team. This can be especially useful if the application is extremely data-heavy.

I find that one of the main drawbacks is that it's harder to handle errors, because errors might be issued from the database driver and not from the ORM itself. This can be annoying at times, but for me is overshadowed by the benefits above.

Basically, my go-to is, the less I have to configure in the ORM, the better. In particular, I've never felt like using gorm's [auto migration feature](https://gorm.io/docs/migration.html). I do recognise that this is probably a bit extreme, and a combination of using the ORM features and pre-configured database would be the sweet spot, but I never felt the need yet of using those additional ORM features.