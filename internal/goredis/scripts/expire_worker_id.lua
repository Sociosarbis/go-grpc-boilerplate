local prefix = KEYS[1]
local time_to_live = ARGV[1]

local set_key = prefix..'_set'
local time = redis.call("TIME")
local current_timestamp = time[1] * 1000 + math.floor(time[2] / 1000)

local expired_max = current_timestamp - time_to_live - 1
local ids = redis.call('ZRANGEBYSCORE', set_key, 0, expired_max)

if #ids ~= 0 then
  local list_key = prefix..'_list'
  redis.call('ZREMRANGEBYSCORE', 0, expired_max)
  redis.call('LPUSH', list_key, table.unpack(ids))
end
