# Create databases with charset utf8mb4
CREATE DATABASE IF NOT EXISTS usr CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE IF NOT EXISTS img CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE IF NOT EXISTS mst CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

# Set default character set and collation (optional, often server-wide configured)
# SET GLOBAL character_set_server = 'utf8mb4';
# SET GLOBAL collation_server = 'utf8mb4_unicode_ci';

# Consider time zone adjustment based on server time zone (if needed)
# Assuming server is in UTC (GMT+0) and you want Jakarta time (UTC+7)
SET GLOBAL time_zone = '+7:00';
SET GLOBAL time_zone = 'ASIA/JAKARTA';
-- SET @@global.time_zone = '+00:00';
# This statement would only affect the current session

# You might want to set the time zone in your application code
# for each connection if needed for specific time zone handling.