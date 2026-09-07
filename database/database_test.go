package database

import "testing"

func TestConfigFromEnvUsesDefaults(t *testing.T) {
	for _, key := range []string{
		"MONGODB_URI",
		"MONGODB_DATABASE",
		"MONGODB_USERS_COLLECTION",
		"MONGODB_HOBBIES_COLLECTION",
	} {
		t.Setenv(key, "")
	}

	cfg := configFromEnv()

	if cfg.uri != "mongodb://localhost:27017" {
		t.Fatalf("unexpected default URI: %q", cfg.uri)
	}
	if cfg.database != "go-mongodb" {
		t.Fatalf("unexpected default database: %q", cfg.database)
	}
	if cfg.userCollection != "users" || cfg.hobbyCollection != "hobbies" {
		t.Fatal("unexpected default collection names")
	}
}

func TestConfigFromEnvUsesEnvironment(t *testing.T) {
	t.Setenv("MONGODB_URI", "mongodb://database.internal:27018")
	t.Setenv("MONGODB_DATABASE", "app_database")
	t.Setenv("MONGODB_USERS_COLLECTION", "app_users")
	t.Setenv("MONGODB_HOBBIES_COLLECTION", "app_hobbies")

	cfg := configFromEnv()

	if cfg.uri != "mongodb://database.internal:27018" {
		t.Fatalf("unexpected URI: %q", cfg.uri)
	}
	if cfg.database != "app_database" {
		t.Fatalf("unexpected database: %q", cfg.database)
	}
	if cfg.userCollection != "app_users" || cfg.hobbyCollection != "app_hobbies" {
		t.Fatal("collection names were not loaded from the environment")
	}
}
