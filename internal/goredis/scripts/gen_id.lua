local prefix = KEYS[1]

local data_center_id = ARGV[1]

local worker_id = ARGV[2]

local last_timestamp_key = prefix.."_last_timestamp"

local sequence_key = prefix.."_sequence"

local last_timestamp = redis.call("GET", last_timestamp_key)

if not last_timestamp then
  last_timestamp = 0
end

local time = redis.call("TIME")

local sequence_bits = 12

local squence_mask = ~(-1 << sequence_bits)

local current_timestamp = time[1] * 1000 + math.floor(time[2] / 1000)

local sequence = redis.call("GET", sequence_key)

if not sequence then
  sequence = 0
end

local twepoch = 1288834974657

local worker_id_bits = 5

local data_center_id_bits = 5

local worker_id_shift = sequence_bits

local data_center_id_shift = worker_id_shift + worker_id_bits

local timestamp_left_shift = data_center_id_shift + data_center_id_bits

if current_timestamp == last_timestamp then
  sequence = (sequence + 1) & squence_mask
  if sequence == 0 then
    repeat
      time = redis.call("TIME")
      current_timestamp = time[1] * 1000 + math.floor(time[2] / 1000)
    until current_timestamp ~= last_timestamp
  end
elseif current_timestamp > last_timestamp then
  sequence = 0
else
  return { err = "Clock moved backwards. Refusing to generate id for "..(last_timestamp - current_timestamp).."ms" } 
end

redis.call("SET", last_timestamp_key, current_timestamp)
redis.call("SET", sequence_key, sequence)

return ((current_timestamp - twepoch) << timestamp_left_shift)
    | (data_center_id << data_center_id_shift)
    | (worker_id << worker_id_shift)
    | sequence

