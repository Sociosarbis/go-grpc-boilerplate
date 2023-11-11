local prefix = KEYS[1];

local list_key = prefix..'list'

local count = redis.call('LLEN', list_key)

if count == 0 then
  return { err = "no worker id available" }
end

local time = redis.call("TIME")

local current_timestamp = time[1] * 1000 + math.floor(time[2] / 1000)

local id = redis.call('LPOP', list_key)

local set_key = prefix..'set'

redis.call('ZADD', set_key, id, current_timestamp)

return id