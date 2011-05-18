require 'rake'

task :default => 'infect'

desc "Hook our dotfiles into system-standard positions."
task :infect => 'structure' do
  linkables = Dir.glob('link/**')

  skip_all = false
  overwrite_all = false
  backup_all = false

  linkables.each do |file|
    overwrite = false
    backup = false

    target = "#{ENV["HOME"]}/.#{file}"

    if File.exists?(target) || File.symlink?(target)
      unless skip_all || overwrite_all || backup_all
        puts "File already exists: #{target}, what do you want to do? [s]kip, [S]kip all, [o]verwrite, [O]verwrite all, [b]ackup, [B]ackup all"
        case STDIN.gets.chomp
        when 'o' then overwrite = true
        when 'b' then backup = true
        when 'O' then overwrite_all = true
        when 'B' then backup_all = true
        when 'S' then skip_all = true
        end
      end

      # Overwrite
      FileUtils.rm_rf(target) if overwrite || overwrite_all

      # Backup
      `mv "$HOME/.#{file}" "$HOME/tmp/#{file}.#{Time.now.to_i}.backup"` if backup || backup_all
    end

    # Do the link...
    `ln -s "$PWD/#{linkable}" "#{target}"`
  end
end

desc "Build wanted directory structure."
task :structure do
   dirs = [ 'Projects', 'bin', 'tmp' ].map {|dir|
      "#{ENV["HOME"]}/#{dir}"
   }.keep_if {|dir| !File.exist? dir }

   FileUtils.mkdir dirs
end
