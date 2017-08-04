require "file_utils"

path = ARGV[0]

dir_name = File.dirname(path)
file_name = File.basename(path)

FileUtils.mkdir_p(dir_name, mode: 0o755)
FileUtils.touch("#{dir_name}/#{file_name}")
