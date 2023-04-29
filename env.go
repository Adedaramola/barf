package barf

import "github.com/opensaucerer/barf/env"

/*
Env loads environment variables into the given EnvStruct from the give EnvPath.

The EnvPath is optional and defaults to ".env".

The struct must have a tag named "barfenv" with the following format: `barfenv:"key=YOUR_ENV_KEY;required=true"` or `barfenv:"key=YOUR_ENV_KEY;required=false"`
*/
var Env = env.Env