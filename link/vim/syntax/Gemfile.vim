" Vim Syntax File
" Language:     Gemfile
" Maintainer:   Mustaqil Ali <git@mustaqila.li>
" Filenames:    Gemfile
" Last Change:  2012 Jan 21


" Adopt general Ruby syntax highlight first
runtime syntax/ruby.vim

syntax case match
syntax keyword gemfileKeywords        source gem gemspec git path group platforms
syntax match gemfileGemOptions        /:require\|:group\|:groups\|:platforms\|:git\|:path/
syntax match gemfileGemspecOptions    /:path\|:name\|:development_group/
syntax match gemfileGitOptions        /:branch\|:tag\|:ref\|:submodules/

highlight def link gemfileKeywords Function
highlight def link gemfileGemOptions Label
highlight def link gemfileGemspecOptions Label
highlight def link gemfileGitOptions Label
