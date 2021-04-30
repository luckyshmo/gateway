-- inclinometer data
data="01016077c3fec38cf100c3d1ab800000"

--[[ Перевод массива hex в массив чисел ]]--
function string.fromhex(str)
    return (str:gsub('..', function (cc)
        return string.char(tonumber(cc, 16))
    end))
end

-- define parsed data
parsed_data = {
	packet_type,	-- Тип пакета (DATA 1 or INFO 2) 
	time_stamp, 	-- Время измерения
	angle_x, 		-- Угол наклона в угловых минутах
	angle_y, 		-- Угол наклона в угловых минутах
	temperature, 	-- Температура, умноженная на 10
}

-- convert string of HEX values into byte array (raw string)
arr = string.fromhex(data)

parsed_data.packet_type = tonumber(arr:byte(1))
parsed_data.time_stamp = (tonumber(arr:byte(3)) << 24) + (tonumber(arr:byte(4)) << 16) + (tonumber(arr:byte(5)) << 8) + tonumber(arr:byte(6))
parsed_data.angle_x = (tonumber(arr:byte(7)) << 24) + (tonumber(arr:byte(8)) << 16) + (tonumber(arr:byte(9)) << 8) + tonumber(arr:byte(10))
parsed_data.angle_y = (tonumber(arr:byte(11)) << 24) + (tonumber(arr:byte(12)) << 16) + (tonumber(arr:byte(13)) << 8) + tonumber(arr:byte(14))
parsed_data.temperature = (tonumber(arr:byte(15)) << 8) + tonumber(arr:byte(16))

print(parsed_data.temperature)
