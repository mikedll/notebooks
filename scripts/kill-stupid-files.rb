#!/usr/bin/env ruby

require 'fileutils'

deleted = 0
dirs = ["."]
dirs.each do |dir|
  Dir.glob("#{dir}/**/*").each do |file|
    match = /\d{4}-\d{2}-\d{2}(-\d+)?/.match(file)
    if match
      puts "Removing #{file}"
      FileUtils.rm file 
      deleted += 1
    end
  end
end

puts "Deleted #{deleted} file(s)"
