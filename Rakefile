require 'rake'

task :default => 'infect'

desc "Hook our dotfiles into system-standard positions."
task :infect => 'structure' do

   # The files we want to link the roots.
   Dir.glob('link/**').each do |linkable|
      file = linkable.split('/').last
      link(linkable, "#{ENV["HOME"]}/.#{file}")
   end

   # The files we want to link the leafs
   Dir.glob('specific/**/*').each do |linkable|
      file = linkable.split('/')
      file.delete_at 0
      file = "#{ENV["HOME"]}/.#{file.join('/')}"

      if !File.directory? linkable
         dir = File.dirname file
         if !Dir.exists? dir
            FileUtils.mkdir_p dir
         end

         link(linkable, file)
      end
   end

   # Link all of bin.
   Dir.glob('bin/*').each do |linkable|
      link(linkable, "#{ENV["HOME"]}/#{linkable}")
   end
end

desc "Build wanted directory structure."
task :structure do
   dirs = [ 'Projects', 'bin', 'tmp' ].map {|dir|
      "#{ENV["HOME"]}/#{dir}"
   }.keep_if {|dir| !File.exist? dir }

   FileUtils.mkdir dirs
end

def link file, target
   overwrite = false
   backup = false

   if File.exists?(target) || File.symlink?(target)
      puts "File already exists: #{target}, what do you want to do? [s]kip, [o]verwrite, [b]ackup"
      case STDIN.gets.chomp
      when 'o' then overwrite = true
      when 'b' then backup = true
      end

      # Overwrite
      FileUtils.rm_rf(target) if overwrite

      # Backup
      `mv "$HOME/.#{file}" "$HOME/tmp/#{file}.#{Time.now.to_i}.backup"` if backup
   end

   # Do the link...
   `ln -s "$PWD/#{file}" "#{target}"`
end
