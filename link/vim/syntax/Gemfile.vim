" Vim Syntax File
" Language:     Gemfile
" Maintainer:   Mustaqil Ali <git@mustaqila.li>
" Filenames:    Gemfile

" Adopt general Ruby syntax highlight first
runtime syntax/ruby.vim

syntax case match
syntax keyword gemfileKeywords        source gem gemspec git github path group platforms
syntax match gemfileGemOptions        /:require\|:group\|:groups\|:platforms\|:github\|:git\|:path/
syntax match gemfileGemspecOptions    /:path\|:name\|:development_group/
syntax match gemfileGitOptions        /:branch\|:tag\|:ref\|:submodules/

highlight def link gemfileKeywords Function
highlight def link gemfileGemOptions Label
highlight def link gemfileGemspecOptions Label
highlight def link gemfileGitOptions Label
