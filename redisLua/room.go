package redisLua

const (
	JOIN = `
		lacal users = KEYS[1]
		local roomId = KEYS[2]
		`
)
