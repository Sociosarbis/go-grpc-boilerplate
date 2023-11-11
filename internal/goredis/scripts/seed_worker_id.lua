local prefix = KEYS[1]

local capacity = ARGV[1]

local list_key = prefix..'list'

if redis.call('EXISTS', list_key) then
  return { err = "list already exists" }
else
  for i = 0,capacity do
    redis.call("LPUSH", i)
  end
end