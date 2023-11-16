local prefix = KEYS[1]

local capacity = ARGV[1]

local list_key = prefix..'_list'

if redis.call('EXISTS', list_key) == 1 then
  return { err = "list already exists" }
else
  for i = 0,capacity-1 do
    redis.call("LPUSH", list_key, i)
  end
end
return 0