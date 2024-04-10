package system

func (s Storage) Migrations() error {
	var err error
	// Create profile table
	_, err = s.Conn.Exec(`
        create table if not exists profiles (
            id text primary key not null,
            created datetime not null default current_timestamp,
            updated datetime not null default current_timestamp,
            user_id text unique not null,
            active boolean not null,
            username text not null,
            about text not null,
            first_name text not null,
            last_name text not null,
            email text not null,
            country text not null,
            street_address text not null,
            city text not null,
            state text not null,
            zip text not null,
            email_notifications text not null,
            push_notification text not null,
            resume text not null,
            cover text not null,
            position text not null,
            skills text not null
        )`)
	if err != nil {
		return err
	}
	// Index user_id
	_, err = s.Conn.Exec(`create index if not exists idx_profiles_user_id on profiles (user_id)`)
	if err != nil {
		return err
	}

	// Create notes table
	_, err = s.Conn.Exec(`
        create table if not exists notes (
            id text primary key not null,
            created datetime not null default current_timestamp,
            updated datetime not null default current_timestamp,
            deleted datetime not null default '2400-01-01 00:00:00',
            user_id text not null,
            title text not null,
            content text not null
        )`)
	if err != nil {
		return err
	}
	// Create index for user_id
	_, err = s.Conn.Exec(`create index if not exists user_id on notes (user_id)`)
	if err != nil {
		return err
	}
	return nil
}
