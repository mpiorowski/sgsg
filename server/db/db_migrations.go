package db

func Migrations() error {
	var err error
	// Create the users table
	_, err = Db.Exec(`
        create table if not exists users (
            id text primary key not null,
            created timestamp not null default current_timestamp,
            updated timestamp not null default current_timestamp,
            deleted timestamp,
            email text unique not null,
            role int not null,
            sub text unique not null,
            avatar text not null default '',
            subscription_id text not null default '',
            subscription_end timestamp
        )`)
	if err != nil {
		return err
	}

	// Create tokens table
	_, err = Db.Exec(`
        create table if not exists tokens (
            id text primary key not null,
            created timestamp not null default current_timestamp,
            updated timestamp not null default current_timestamp,
            deleted timestamp,
            user_id text not null,
            provider text not null,
            access_token text not null,
            refresh_token text not null,
            token_type text not null,
            expires timestamp not null
        )`)
	if err != nil {
		return err
	}

	// Create profile table
	_, err = Db.Exec(`
        create table if not exists profiles (
            id text primary key not null,
            created timestamp not null default current_timestamp,
            updated timestamp not null default current_timestamp,
            deleted timestamp,
            user_id text not null,
            username text not null,
            about text not null,
            resume_id text not null,
            cover_url text not null
        )`)
	if err != nil {
		return err
	}
	// Create index for user_id
	_, err = Db.Exec(`create index if not exists user_id on profiles (user_id)`)
	if err != nil {
		return err
	}

	// Create notes table
	_, err = Db.Exec(`
        create table if not exists notes (
            id text primary key not null,
            created timestamp not null default current_timestamp,
            updated timestamp not null default current_timestamp,
            deleted timestamp,
            user_id text not null,
            title text not null,
            content text not null
        )`)
	if err != nil {
		return err
	}
	// Create index for user_id
	_, err = Db.Exec(`create index if not exists user_id on notes (user_id)`)
	if err != nil {
		return err
	}

	return nil
}
