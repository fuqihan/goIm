local REDIS_ROOM_USERS, REDIS_ROOM_USER_INFO, REDIS_USER_ROOMS = KEYS[1], KEYS[2], KEYS[3];
local RoomId, UserIds, now = ARGV[1], cmsgpack.unpack(ARGV[2]), ARGV[3]

for v, userId in ipairs(UserIds) do
    redis.call("SADD", string.format(REDIS_ROOM_USERS, RoomId), userId)
    redis.call("HSET", string.format(REDIS_ROOM_USER_INFO, RoomId, userId), "currentViewTime", now)
    redis.call("SADD", string.format(REDIS_USER_ROOMS, userId), RoomId)
end
